package model

type Chat struct {
	ID           int64     `json:"_id" bson:"_id,omitempty"`
	Name         string    `json:"name" bson:"name"`
	Playlist     *Playlist `json:"playlist" bson:"playlist"`
	PostedVideos []string  `json:"posted_videos" bson:"posted_videos"`
}

type Publisher struct {
	Username string `json:"username" bson:"username"`
	SecureID string `json:"secure_id" bson:"secure_id"`
}

func NewPublisher(username, secID string) *Publisher {
	return &Publisher{
		Username: username,
		SecureID: secID,
	}
}

func NewChat(id int64, name string) *Chat {
	return &Chat{
		ID:           id,
		Name:         name,
		PostedVideos: nil,
	}
}
