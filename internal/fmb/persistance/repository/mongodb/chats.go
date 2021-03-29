package mongodb

import (
	"crypto/tls"
	"net"

	"github.com/Haski007/fav-music-bot/internal/fmb/config"
	"github.com/Haski007/fav-music-bot/internal/fmb/persistance/model"
	"github.com/Haski007/fav-music-bot/internal/fmb/persistance/repository"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/sirupsen/logrus"
)

var session *mgo.Session

type ChatRepository struct {
	ChatsColl *mgo.Collection
}

func (r *ChatRepository) InitChatsConn(cfg config.Mongo) {

	dialInfo := mgo.DialInfo{
		Addrs:    cfg.Addrs,
		Username: cfg.UserName,
		Password: cfg.Password,
	}
	tlsConfig := &tls.Config{}
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig) // add TLS config
		return conn, err
	}

	var err error
	session, err = mgo.DialWithInfo(&dialInfo)
	if err != nil {
		logrus.Fatalf("[mgo DialWithInfo] dialInfo: %+v | err: %s", dialInfo, err)
	}

	if err = session.Ping(); err != nil {
		logrus.Fatalf("[mgo Ping] dialInfo: %+v | err: %s", dialInfo, err)
	}

	r.ChatsColl = session.DB(cfg.DBName).C("chats")
}

// ---> CHATS

func (r *ChatRepository) GetAllChats(chats *[]model.Chat) {
	if err := r.ChatsColl.Find(bson.M{}).All(chats); err != nil {
		logrus.Errorf("[GetAllChats] err: %s", err)
		return
	}
}

func (r *ChatRepository) SaveNewChat(chat *model.Chat) error {
	if r.ChatExists(chat.ID) {
		return repository.ErrChatAlreadyExists
	}

	return r.ChatsColl.Insert(chat)
}

func (r *ChatRepository) RemoveChat(chatID int64) error {

	if !r.ChatExists(chatID) {
		return repository.ErrChatDoesNotExist
	}

	return r.ChatsColl.RemoveId(chatID)
}

func (r *ChatRepository) ChatExists(id int64) bool {
	count, _ := r.ChatsColl.FindId(id).Count()
	if count != 0 {
		return true
	}
	return false
}

// ---> Music

func (r *ChatRepository) PushPostedMusic(chatID int64, musicID string) error {
	if !r.ChatExists(chatID) {
		return repository.ErrChatDoesNotExist
	}

	findQuery := bson.M{
		"_id": chatID,
	}
	updateQuery := bson.M{
		"$push": bson.M{
			"posted_music": musicID,
		},
	}

	err := r.ChatsColl.Update(findQuery, updateQuery)
	return err
}

func (r *ChatRepository) PostedMusicExists(chatID int64, musicID string) bool {
	query := bson.M{
		"_id":          chatID,
		"posted_music": musicID,
	}
	count, _ := r.ChatsColl.Find(query).Count()
	if count != 0 {
		return true
	}
	return false
}
