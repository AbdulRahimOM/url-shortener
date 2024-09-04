package domain

type ErrorRes struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}
type SuccessRes struct {
	Status       bool   `json:"status"`
	Message      string `json:"message"`
	ShortenedUrl string `json:"shortened_url"`
}
type UrlRequest struct {
    LongUrl string `json:"longUrl"`
}