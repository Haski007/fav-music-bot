package api

import (
	"google.golang.org/api/youtube/v3"
)

func (srv *YoutubeService) GetLikedIDs(limit int64) ([]*youtube.Video, error) {
	call := srv.SRV.Videos.List([]string{"snippet"})

	call.MyRating("like")
	call.MaxResults(limit)
	resp, err := call.Do()
	if err != nil {
		return nil, err
	}

	return resp.Items, nil
}
