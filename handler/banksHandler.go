package handler

type BanksHandler interface {
}

type banksHandler struct {
}

func NewBanksHandler() banksHandler {
	return banksHandler{}
}
