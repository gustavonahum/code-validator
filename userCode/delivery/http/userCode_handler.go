package http

type UserCodeHandler struct {
	UCUsecase userCode.Usecase
}

func NewUserCodeHandler(us userCode.Usecase) {
	handler := &UserCodeHandler{
		UCUsecase: us,
	}
	
}