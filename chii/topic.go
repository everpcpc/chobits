package chii

// Topic ...
type Topic struct {
	ID        int    `json:"id,omitemtpy"`
	URL       string `json:"url,omitemtpy"`
	Title     string `json:"title,omitemtpy"`
	MainID    string `json:"main_id,omitemtpy"`
	Timestamp int    `json:"timestamp,omitemtpy"`
	Lastpost  int    `json:"lastpost,omitemtpy"`
	Replies   int    `json:"replies,omitemtpy"`
	User      User   `json:"user,omitemtpy"`
}
