package domain

type ChatRequest struct {
	Message string  `json:"message"`
}

type ChatResponse struct {

}

type ChatUsecase interface {
	SendMessage(uid string, msg string) error
}