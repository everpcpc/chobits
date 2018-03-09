package chii

// Character ...
type Character struct {
	ID       int      `json:"id"`
	URL      string   `json:"url"`
	Name     string   `json:"name"`
	NameCN   string   `json:"name_cn,omitempty"`
	RoleName string   `json:"role_name,omitempty"`
	Images   Images   `json:"images,omitempty"`
	Comment  int      `json:"comments,omitempty"`
	Collects int      `json:"collects,omitempty"`
	Info     Info     `json:"info,omitempty"`
	Actors   []Person `json:"actors,omitempty"`
}
