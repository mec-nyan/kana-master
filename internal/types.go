package internal

type Action = uint

type Name = string

type Description = string

type ActionItem struct {
	Name
	Action
	Description
}

type ActionMenu []ActionItem

type Mode = uint

type ModeItem struct {
	Name
	Mode
	Description
}

type ModeMenu []ModeItem
