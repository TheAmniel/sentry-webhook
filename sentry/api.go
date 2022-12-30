package sentry

type Sentry struct {
	Action string `json:"action"`
	Data   *Data  `json:"data"`
	Actor  *Actor `json:"actor"`
}

// TODO: Add issue???
type Data struct {
	Error *DataError `json:"error"`
}

type DataError struct {
	EventID     string             `json:"event_id"`
	Environment string             `json:"environment"`
	Title       string             `json:"title"`
	WebURL      string             `json:"web_url"`
	URL         string             `json:"url"`
	Platform    string             `json:"platform"`
	Message     string             `json:"message"`
	Level       string             `json:"level"`
	Project     int64              `json:"project"`
	Timestamp   float64            `json:"timestamp"`
	Tags        [][2]string        `json:"tags,omitempty"`
	Contexts    *DataErrorContexts `json:"contexts"`
	Exception   interface{}        `json:"exception,omitempty"`
}

type DataErrorContexts struct {
	Browser *DataErrorContextsBrowser `json:"browser"`
	OS      *DataErrorContextsBrowser `json:"os"`
}

type DataErrorContextsBrowser struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Version string `json:"version"`
}

type Actor struct {
	Type string `json:"type"`
	Name string `json:"name"`
	ID   int    `json:"id"`
}
