package discord

import "time"

type Webhook struct {
	Username  string   `json:"username,omitempty"`
	AvatarURL string   `json:"avatar_url,omitempty"`
	Content   string   `json:"content"`
	Embeds    []*Embed `json:"embeds,omitempty"`
}

type Embed struct {
	Title       string        `json:"title,omitempty"`
	Type        string        `json:"type"`
	URL         string        `json:"url,omitempty"`
	Description string        `json:"description,omitempty"`
	Color       int           `json:"color,omitempty"`
	Timestamp   time.Time     `json:"timestamp,omitempty"`
	Fields      []*EmbedField `json:"fields,omitempty"`
	Footer      *EmbedFooter  `json:"footer,omitempty"`
}

type EmbedFooter struct {
	Text         string `json:"text"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}
