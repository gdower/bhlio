package reference

type Reference struct {
	RefString string `json:"refString,omitempty"`
	Authors   string `json:"authors,omitempty"`
	Journal   string `json:"journal,omitempty"`
	Volume    string `json:"volume,omitempty"`
	Pages     string `json:"pages,omitempty"`
	Year      string `json:"year,omitempty"`
}
