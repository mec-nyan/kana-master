package internal

type Action = uint
type Description = string

type MenuItem struct {
	Action
	Description
}

type Menu []MenuItem
