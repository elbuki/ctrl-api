package handler

type Handler struct {
	api *API
}

func NewHandler(a *API) *Handler {
	return &Handler{a}
}
