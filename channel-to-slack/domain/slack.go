package domain

type Slack struct {
	Attachments []SlackAttachment `json:"attachments"`
}

type SlackAttachment struct {
	Color  string                  `json:"color"`
	Blocks []SlackAttachmentsBlock `json:"blocks"`
}

type SlackAttachmentsBlock struct {
	Type string                    `json:"type"`
	Text SlackAttachmentsBlockText `json:"text"`
}

type SlackAttachmentsBlockText struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func NewSlack(title string, color string, text string) *Slack {
	titleBlock := SlackAttachmentsBlock{
		Type: "section",
		Text: SlackAttachmentsBlockText{Type: "mkdown", Text: title},
	}
	textBlock := SlackAttachmentsBlock{
		Type: "section",
		Text: SlackAttachmentsBlockText{Type: "mkdown", Text: text},
	}
	blocks := []SlackAttachmentsBlock{titleBlock, textBlock}

	attachments := SlackAttachment{
		Color:  color,
		Blocks: blocks,
	}
	return &Slack{Attachments: []SlackAttachment{attachments}}
}
