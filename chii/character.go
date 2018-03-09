package chii

// Character ...
type Character struct {
	ID       int      `json:"id"`
	URL      string   `json:"url"`
	Name     string   `json:"name"`
	NameCN   string   `json:"name_cn,omitempty"`
	RoleName string   `json:"role_name,omitempty"`
	Images   Images   `json:"images,omitemtpy"`
	Comment  int      `json:"comments,omitemtpy"`
	Collects int      `json:"collects,omitemtpy"`
	Info     Info     `json:"info,omitemtpy"`
	Actors   []Person `json:"actors,omitemtpy"`
}

// CharacterAlias ...
type CharacterAlias struct {
	JP     string `json:"jp,omitempty"`
	Romaji string `json:"romaji,omitempty"`
	ZH     string `json:"zh,omitempty"`
	Kana   string `json:"kana,omitempty"`
}
