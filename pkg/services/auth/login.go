package auth

type Login interface {
	Login(email, password string) bool
}

type LoginInformation struct {
	email string
	pass  string
}

func (l *LoginInformation) Login(email, password string) bool {
	return l.email == email && l.pass == password
}

func StaticLogin(email, password string) Login {
	return &LoginInformation{email, password}
}
