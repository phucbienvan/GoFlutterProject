package req

type ReqSignup struct {
	FullName string `validate:"required"`
	Email    string `validate:"required"`
	Password string `validate:"required"`
}
