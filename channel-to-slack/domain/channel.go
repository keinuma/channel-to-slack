package domain

import (
	"encoding/json"
	"fmt"
)

const (
	EventTypeMessage  = "message"
	EventTypeUserChat = "userChat"
)

const (
	ChatTypeGroup    = "group"
	ChatTypeUserChat = "userChat"
)

const (
	PersonTypeManager = "manager"
	PersonTypeUser    = "user"
	PersonTypeBot     = "bot"
)

const (
	ManagerEventColor = "#DAA038"
	UserEventColor    = "#2EB886"
)

type Channel struct {
	Event  string        `json:"event"`
	Type   string        `json:"type"`
	Entity ChannelEntity `json:"entity"`
	Refers ChannelRefers `json:"refers"`
}

type ChannelEntity struct {
	PlainText  string `json:"plainText"`
	ChatType   string `json:"chatType"`
	ChannelId  string `json:"channelId"`
	ChatId     string `json:"chatId"`
	PersonType string `json:"personType"`
}

type ChannelRefers struct {
	Manager ChannelRefersManager `json:"manager"`
	User    ChannelRefersUser    `json:"user"`
}

type ChannelRefersManager struct {
	Name string `json:"name"`
}

type ChannelRefersUser struct {
	Name string `json:"name"`
}

func (c *Channel) ToSlack() *Slack {
	if c.Type == EventTypeUserChat || c.Entity.ChatType == ChatTypeGroup || c.Entity.PersonType == PersonTypeBot {
		return nil
	}

	var title string
	var color string
	link := fmt.Sprintf("https://desk.channel.io/#/channels/%s/user_chats/%s", c.Entity.ChannelId, c.Entity.ChatId)

	if c.Entity.PersonType == PersonTypeManager {
		title = fmt.Sprintf("%sが[%s](%s)に返信しました", c.Refers.Manager.Name, c.Refers.User.Name, link)
		color = ManagerEventColor
	} else {
		title = fmt.Sprintf("[%s](%s)からメッセージが来ました", c.Refers.User.Name, link)
		color = UserEventColor
	}
	return NewSlack(title, color, c.Entity.PlainText)
}

func NewChannel(body string) (*Channel, error) {
	var channel *Channel

	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &channel)
	if err != nil {
		return nil, err
	}

	return channel, nil
}
