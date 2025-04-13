package antipatterns

import "net/http"

type ID int
type Message int
type User int
type Users int
type Conversation int
type Chat int

/* плохо */
type API interface {
	GetMessage(ID) Message
	SendMessage(Message) ID
	DeleteMessage(ID) error
	MarkAsRead(ID)
	EditMessage(ID) error
	CreateChat(string) ID
	EditChat(string) error
	GetChatUsers(ID) []Users
	ConnectChat(ID) error
	AddMessageToChat(Message, Chat)
	GetHistory(User) []Message
	GetConversations(User) []Conversation
	// etc...
}

/* хорошо */
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
