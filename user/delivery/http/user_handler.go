package http

import "github.com/gustavonahum/code-validator/user"

type UserHandler struct {
	UUsecase user.Usecase
}

func NewUserHandler(us user.Usecase) {
	handler := &UserHandler{
		UUsecase: us,
	}

}
