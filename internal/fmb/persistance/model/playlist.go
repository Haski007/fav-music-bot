package model

type Playlist struct {
	ID    string `json:"id" bson:"id"`
	Title string `json:"title" bson:"title"`
}
