package types

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	ChannelId string `json:"channelId"`
}
