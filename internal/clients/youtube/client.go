package youtube

import (
	"context"
	"fmt"
	"net/http"
)

type client struct {
	url         string
	accessToken string
	httpClient  *http.Client
}

type Client interface {
	SearchTrack(ctx context.Context, trackName string) (response *http.Response, err error)
}

func NewClientYouTube(url, accessToken string, httpClient *http.Client) *client {
	return &client{url: url, accessToken: accessToken, httpClient: httpClient}
}

func (c *client) SearchTrack(ctx context.Context, trackName string) (response *http.Response, err error) {
	params := map[string]string{
		"part":       "snippet",
		"maxResults": "1",
		"q":          trackName,
		"type":       "video",
		"key":        c.accessToken,
	}
	uri := fmt.Sprintf("%s/search", c.url)


	request, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	query := request.URL.Query()
	for k, v := range params {
		query.Set(k, v)
	}
	request.URL.RawQuery = query.Encode()


	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	return c.httpClient.Do(request)
}
