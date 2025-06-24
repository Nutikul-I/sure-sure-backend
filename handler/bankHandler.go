package handler

type BankHandler interface {
}

type bankHandler struct {
}

func NewBankHandler() bankHandler {
	return bankHandler{}
}
