package main

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/ItsTas/BlogAgregator/internal/auth"
	"github.com/ItsTas/BlogAgregator/internal/client"
	"github.com/ItsTas/BlogAgregator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func parsePubDate(pubDateToParse string) (time.Time, error) {
	const (
		layoutRFC3339 = time.RFC3339
		layoutRSS1    = "Mon, 02 Jan 2006 15:04:05 -0700"
		layoutRSS2    = "Mon, 02 Jan 2006 15:04:05 MST"
	)

	if pubDate, err := time.Parse(layoutRFC3339, pubDateToParse); err == nil {
		return pubDate, nil
	}

	if pubDate, err := time.Parse(layoutRSS1, pubDateToParse); err == nil {
		return pubDate, nil
	}

	if pubDate, err := time.Parse(layoutRSS2, pubDateToParse); err == nil {
		return pubDate, nil
	}

	return time.Time{}, fmt.Errorf("unable to parse date: %v", pubDateToParse)
}

func (cfg *apiConfig) fetchFromFeed(feed database.Feed) (*client.RSS, error) {
	data, err := cfg.Client.FetchXML(feed.Url)
	if err != nil {
		return nil, err
	}

	params := database.MarkFeedsFetchedParams{
		ID: feed.ID,
		LastFetchedAt: sql.NullTime{
			Time:  time.Now().UTC(),
			Valid: true,
		},
		UpdatedAt: time.Now().UTC(),
	}

	_, err = cfg.DB.MarkFeedsFetched(context.TODO(), params)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (cfg *apiConfig) scraper(interval time.Duration, n int) {
	timer := time.NewTicker(interval)
	for range timer.C {
		feeds, err := cfg.DB.GetNextFeedsToFetch(context.TODO(), int32(n))
		if err != nil {
			fmt.Printf("\nfailed to fetch: %v\n", err)
			continue
		}

		var wg sync.WaitGroup
		for _, feed := range feeds {
			wg.Add(1)
			go cfg.saveFeedToDataBase(feed, &wg)
		}
		wg.Wait()
	}
}

func (cfg *apiConfig) saveFeedToDataBase(feed database.Feed, wg *sync.WaitGroup) {
	defer wg.Done()

	feedData, err := cfg.fetchFromFeed(feed)
	if err != nil {
		fmt.Printf("\nFailed to fetch feed: %v\n error: %v\n", feed, err)
		return
	}
	for _, item := range feedData.Channel.Items {
		id := uuid.New()
		pubDate, err := parsePubDate(item.PubDate)
		if err != nil {
			fmt.Printf("\nCouldn't parse date: %v\n", err.Error())
			continue
		}

		imageURL := auth.ExtractImageUrl(item)
		postParams := database.CreatePostParams{
			ID:          id,
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: pubDate,
			FeedID:      feed.ID,
			ImageUrl:    convertStringToSQLNULLString(imageURL),
		}
		_, err = cfg.DB.CreatePost(context.TODO(), postParams)
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok {
				if pqErr.Code == "23505" {
					continue
				}
			}
			fmt.Printf("\n Failed to create post: %v\nWith Params: %v\n", postParams, err.Error())
		}
	}
}

func convertStringToSQLNULLString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  s != "",
	}
}
