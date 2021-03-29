package transform

import (
	"fmt"

	"github.com/Haski007/fav-music-bot/internal/fmb/persistance/model"
	"google.golang.org/api/youtube/v3"
)

func DecodeYoutubeVideos(videos []*youtube.Video) (music []*model.Music) {
	music = make([]*model.Music, len(videos))
	for i, video := range videos {
		music[i] = &model.Music{
			ID:     video.Id,
			Title:  video.Snippet.Title,
			Author: video.Snippet.ChannelTitle,
			Image:  video.Snippet.Thumbnails.Maxres.Url,
			URL:    fmt.Sprintf("https://www.youtube.com/watch?v=%s", video.Id),
		}
	}
	return music
}
