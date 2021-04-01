package api

import "google.golang.org/api/youtube/v3"

func (srv *YoutubeService) GetPlaylist(id string) ([]*youtube.Playlist, error) {
	call := srv.SRV.Playlists.List([]string{"snippet"})

	call.Id(id)

	resp, err := call.Do()
	if err != nil {
		return nil, err
	}

	return resp.Items, nil
}
