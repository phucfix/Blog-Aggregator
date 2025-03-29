package main

import (
    "context"
    "net/http"
    "io"
    "encoding/xml"
    "html"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
    req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
    if err != nil {
        return nil, err
    }
    req.Header.Set("User-Agent", "gator")

    client := &http.Client{
		Timeout: 10 * time.Second,
	}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    data, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var rssFeedResp RSSFeed;
    if err := xml.Unmarshal(data, &rssFeedResp); err != nil {
        return nil, err
    }

    // Decode escaped HTML entities (like &ldquo;) of the rssFeedResp
	rssFeedResp.Channel.Title = html.UnescapeString(rssFeedResp.Channel.Title)
	rssFeedResp.Channel.Description = html.UnescapeString(rssFeedResp.Channel.Description)
	for i, item := range rssFeedResp.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		rssFeedResp.Channel.Item[i] = item
	}

    return &rssFeedResp, nil
}
