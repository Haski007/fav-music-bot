package api

import (
	"google.golang.org/api/youtube/v3"
)

func (srv *YoutubeService) GetLikedIDs(limit int64, playlistID string) ([]*youtube.PlaylistItem, error) {
	call := srv.SRV.PlaylistItems.List([]string{"snippet"})

	var resp *youtube.PlaylistItemListResponse
	var pageToken string
	var err error
	for {
		call.PlaylistId(playlistID)
		call.MaxResults(limit)
		call.PageToken(pageToken)
		resp, err = call.Do()
		if err != nil {
			return nil, err
		}
		pageToken = resp.NextPageToken
		if pageToken == "" {
			break
		}
	}

	return resp.Items, nil
}
