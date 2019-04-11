package yourUserEntity

type UserEntityController struct {
	UserEntityRepository UserEntityRepository
}

type UserEntityRepository struct {
}

type UserEntity struct {
	Username string `json:"username"    validate:"required,username-req"`
	Email    string `json:"email"       validate:"required,email"`
	Password string `json:"password"    validate:"required,password-req"`
	Enable   bool   `json:"enable"		  validate:"required"`
}

type UserEntities []UserEntity

type UserEntityVerify struct {
	Username string `json:"username"    validate:"required,username-req"`
	Password string `json:"password"    validate:"required,password-req"`
}

type EmailVerify struct {
	Email string `json:"email"       validate:"required,email"`
}
