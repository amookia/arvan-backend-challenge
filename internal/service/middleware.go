package service

type Middleware interface {
	IsUserLimited(string) bool
	IsUserLimitedCapacity(string, int64) bool
}
