package repository

type Repository struct {
	Url      string `json:"html_url"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
}
