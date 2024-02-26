package model

type UserInput struct {
	N1, N2 int
	Action string
}

type Operation struct {
	Number1 int
	Number2 int
	Action  Action
	Result  struct {
		Data  int
		Error error
	}
}

func NewOperation(n1, n2 int, op Action) *Operation {
	return &Operation{
		Number1: n1,
		Number2: n2,
		Action:  op,
	}
}

type Action string

const (
	ActionPlus     Action = "+"
	ActionMinus    Action = "-"
	ActionMultiple Action = "*"
	ActionDivision Action = "/"
)

type Actions []Action

func (actions Actions) IsAction(rawAction string) bool {
	for _, action := range actions {
		if action == Action(rawAction) {
			return true
		}
	}
	return false
}
