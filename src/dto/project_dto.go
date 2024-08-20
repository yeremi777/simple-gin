package dto

type Project struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type StoreProjectRequest struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type UpdateProjectRequest struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}
