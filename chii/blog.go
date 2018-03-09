package chii

// Blog ...
type Blog struct {
	ID        int    `json:"id,omitemtpy"`
	URL       string `json:"url,omitemtpy"`
	Title     string `json:"title,omitemtpy"`
	Summary   string `json:"summary,omitemtpy"`
	Image     string `json:"image,omitemtpy"`
	Replies   int    `json:"replies,omitemtpy"`
	Timestamp int    `json:"timestamp,omitemtpy"`
	Dateline  string `json:"dateline,omitemtpy"`
	User      User   `json:"user,omitemtpy"`
}
