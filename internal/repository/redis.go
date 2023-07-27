package repository

type Redis interface {
	UserRemaining(string) int
}
