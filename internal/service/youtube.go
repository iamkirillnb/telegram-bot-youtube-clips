package service

import (
	"context"
	"encoding/json"
	"fmt"
	"telegram_bot/internal/clients/youtube"
	"telegram_bot/internal/entities"
	"telegram_bot/pkg/logging"
)

type service struct {
	client youtube.Client
	logger *logging.Logger
}

func NewYoutubeService(client youtube.Client, logger *logging.Logger) YoutubeService {
	return &service{client: client, logger: logger}
}


type YoutubeService interface {
	FindTrackByName(ctx context.Context, trackName string) (string, error)
}

func (s *service) FindTrackByName(ctx context.Context, trackName string) (string, error) {
	response, err := s.client.SearchTrack(ctx, trackName)
	if err != nil {
		return "", err
	}


	var data entities.RestResponse
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return "", err
	}

	link := fmt.Sprintf("https://www.youtube.com/watch?v=%s", data.Items[0].Id.VideoId)
	return link, nil
}