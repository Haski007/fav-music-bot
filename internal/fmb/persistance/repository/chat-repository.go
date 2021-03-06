package repository

import "errors"

var (
	ErrChatAlreadyExists = errors.New("CHAT_ALREADY_EXISTS")
	ErrChatDoesNotExist  = errors.New("CHAT_DOES_NOT_EXIST")

	ErrUserAlreadyExists = errors.New("USER_ALREADY_EXISTS")
	ErrUserDoesNotExist  = errors.New("USER_DOES_NOT_EXIST")

	ErrPlaylistAlreadyExists = errors.New("PLAYLIST_ALREADY_EXISTS")
)

type ChatRepository interface {
	SaveNewChat()
}
