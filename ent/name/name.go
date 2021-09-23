package name

type Name struct {
	NameString string `json:"nameString,omitempty"`
	Canonical  string `json:"canonical,omitempty"`
	Authorship string `json:"authorship,omitempty"`
	Year       string `json:"year,omitempty"`
}
