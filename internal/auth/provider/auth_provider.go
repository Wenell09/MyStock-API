package provider

type AuthProvider interface {
	Login(email, password string) (string, error)
}
