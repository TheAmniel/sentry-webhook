package models

type SentryPayload struct {
	Action string             `json:"action"`
	Data   SentryPayloadData  `json:"data"`
	Actor  SentryPayloadActor `json:"actor"`
}

// TODO: Add issue???
type SentryPayloadData struct {
	Error SentryPayloadDataError `json:"error"`
}

type SentryPayloadDataError struct {
	Title     string      `json:"title,omniempty"`
	WebURL    string      `json:"web_url"`
	Platform  string      `json:"platform"`
	Message   string      `json:"message,omniempty"`
	Project   int64       `json:"project"`
	Tags      [][2]string `json:"tags,omniempty"`
	Exception interface{} `json:"exception,omniempty"`
}

type SentryPayloadActor struct {
	Type string `json:"type"`
	Name string `json:"name"`
	ID   int    `json:"id"`
}
