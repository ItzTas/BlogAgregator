package main

import (
	"time"

	"github.com/ItsTas/BlogAgregator/internal/database"
	"github.com/google/uuid"
)

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Url         string    `json:"url"`
	Description string    `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	FeedID      uuid.UUID `json:"feed_id"`
	ImageUrl    string    `json:"image_url"`
}

func databasePostToPost(post database.Post) Post {
	var imgURL string
	if post.ImageUrl.Valid {
		imgURL = post.ImageUrl.String
	}
	p := Post{
		ID:          post.ID,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		Title:       post.Title,
		Url:         post.Url,
		Description: post.Description,
		PublishedAt: post.PublishedAt,
		FeedID:      post.FeedID,
		ImageUrl:    imgURL,
	}
	return p
}
