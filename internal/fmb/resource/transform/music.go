package transform

import (
	"fmt"

	"github.com/Haski007/fav-music-bot/internal/fmb/persistance/model"
	"google.golang.org/api/youtube/v3"
)

func DecodeYoutubeVideos(items []*youtube.PlaylistItem) (music []*model.Music) {
	music = make([]*model.Music, len(items))
	for i, item := range items {
		music[i] = &model.Music{
			ID:     item.Id,
			Title:  item.Snippet.Title,
			Author: item.Snippet.VideoOwnerChannelTitle,
			Image:  item.Snippet.Thumbnails.Maxres.Url,
			URL:    fmt.Sprintf("https://www.youtube.com/watch?v=%s", item.Snippet.ResourceId.VideoId),
		}
	}
	return music
}
