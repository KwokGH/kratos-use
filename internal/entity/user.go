package entity

type UserDTO struct {
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt int64 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// 用户名称
	Name string `json:"name,omitempty"`
	// 用户账号
	Account string `json:"account,omitempty"`
	// 手机号
	Mobile string `json:"mobile,omitempty"`
	// 密码
	Password string `json:"password,omitempty"`
	// 加盐
	PasswordSalt string `json:"password_salt,omitempty"`
}

type RegisterInput struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type LoginInput struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}
type LoginOutput struct {
	Token string `json:"token"`
}
