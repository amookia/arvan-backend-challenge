package service

type Middleware interface {
	UserQuotaRequest(string) bool
	UserQuotaTraffic(string) bool
}
