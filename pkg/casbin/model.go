package casbin

type ApiResource interface {
	GetPath() string
	GetMethod() string
}
