package chii

// Person ...
type Person struct {
	ID       int      `json:"id"`
	URL      string   `json:"url"`
	Name     string   `json:"name"`
	NameCN   string   `json:"name_cn,omitemtpy"`
	RoleName string   `json:"role_name,omitemtpy"`
	Images   Images   `json:"images,omitemtpy"`
	Comment  int      `json:"comments,omitemtpy"`
	Collects int      `json:"collects,omitemtpy"`
	Info     Info     `json:"info,omitemtpy"`
	Jobs     []string `json:"jobs,omitemtpy"`
}
