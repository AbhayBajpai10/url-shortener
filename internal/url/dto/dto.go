package dto

type CreateURLRequest struct {
	OriginalURL string `json:"original_url" binding:"required,url"`
}
type CreateURLResponse struct {
	ID          int    `json:"id"`
	OriginalURL string `json:"original_url"`
	ShortCode   string `json:"short_code"`
	ShortURL    string `json:"short_url"`
}
type URLResponse struct {
	ID          int    `json:"id"`
	OriginalURL string `json:"original_url"`
	ShortCode string `json:"short_code"`
	Clicks int `json:"clicks"`
	ShortURL    string `json:"short_url"`
}
