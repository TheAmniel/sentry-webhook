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
	Title     string      `json:"title,omitempty"`
	WebURL    string      `json:"web_url"`
	Platform  string      `json:"platform"`
	Message   string      `json:"message,omitempty"`
	Project   int64       `json:"project"`
	Tags      [][2]string `json:"tags,omitempty"`
	Exception interface{} `json:"exception,omitempty"`
}

type SentryPayloadActor struct {
	Type string `json:"type"`
	Name string `json:"name"`
	ID   int    `json:"id"`
}
