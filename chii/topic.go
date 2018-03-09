package chii

// Topic ...
type Topic struct {
	ID        int    `json:"id,omitempty"`
	URL       string `json:"url,omitempty"`
	Title     string `json:"title,omitempty"`
	MainID    string `json:"main_id,omitempty"`
	Timestamp int    `json:"timestamp,omitempty"`
	Lastpost  int    `json:"lastpost,omitempty"`
	Replies   int    `json:"replies,omitempty"`
	User      User   `json:"user,omitempty"`
}
