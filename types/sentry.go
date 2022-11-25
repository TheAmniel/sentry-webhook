package types

type Sentry struct {
	Action string       `json:"action"`
	Data   *SentryData  `json:"data"`
	Actor  *SentryActor `json:"actor"`
}

// TODO: Add issue???
type SentryData struct {
	Error SentryDataError `json:"error"`
}

type SentryDataError struct {
	Title     string      `json:"title,omitempty"`
	WebURL    string      `json:"web_url"`
	Platform  string      `json:"platform"`
	Message   string      `json:"message,omitempty"`
	Project   int64       `json:"project"`
	Tags      [][2]string `json:"tags,omitempty"`
	Exception interface{} `json:"exception,omitempty"`
}

type SentryActor struct {
	Type string `json:"type"`
	Name string `json:"name"`
	ID   int    `json:"id"`
}
