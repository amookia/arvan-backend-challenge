package repository

type Redis interface {
	UserRemaining(string, int) int
	Exists(string) bool
	UserMonthlyUsageUpdate(string, int64)
	UserMonthlyUsage(string) int64
}
