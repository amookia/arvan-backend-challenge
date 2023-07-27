package service

type Middleware interface {
	IsUserLimited(string) bool
}
