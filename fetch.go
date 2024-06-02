package main

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/ItsTas/BlogAgregator/internal/client"
	"github.com/ItsTas/BlogAgregator/internal/database"
)

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

func (cfg *apiConfig) fetcherWorker(interval time.Duration, n int) {
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
			go func(feed database.Feed) {
				defer wg.Done()

				feedData, err := cfg.fetchFromFeed(feed)
				if err != nil {
					fmt.Printf("\nFailed to fetch feed: %v\n error: %v\n", feed, err)
					return
				}
				for _, item := range feedData.Channel.Items {
					fmt.Println(item.Title)
				}
			}(feed)
		}
		wg.Wait()
	}
}
