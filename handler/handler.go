package handler

type Handler struct {
	// auth
	Auth AuthMiddle

	// usecase
	// UserUsecase UserUsecase
}

func NewHandler(auth AuthMiddle) *Handler {
	handler := new(Handler)
	// handler.UserUsecase = user
	handler.Auth = auth
	return handler
}
