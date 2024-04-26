package models

type AccountType uint8

const (
	AccountTypeUnauthorizedUser AccountType = iota + 1
	AccountTypeUser
	AccountTypeAdmin
)

type Profile struct {
	ID          uint64
	Name        string
	AccountType AccountType
	CreatedAt   int64
	UpdatedAt   int64
}
