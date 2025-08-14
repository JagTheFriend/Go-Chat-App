package types

type Message struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	AuthorId  string `json:"authorId"`
	ChannelId string `json:"channelId"`
}
