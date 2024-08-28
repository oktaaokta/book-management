package handler

type Handler struct {
	uc usecaseInterface
}

func New(uc usecaseInterface) *Handler {
	return &Handler{
		uc: uc,
	}
}
