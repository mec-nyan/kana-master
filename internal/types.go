package internal

type Action = string
type Description = string

type MenuItem struct {
	Action
	Description
}
type Menu []MenuItem
