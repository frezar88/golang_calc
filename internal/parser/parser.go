package parser

import (
	"calculator/internal/model"
	"fmt"
	"os"
)

type Parser struct {
	rowData []model.UserInput
	actions model.Actions
}

func NewParser(raw []model.UserInput, actions ...model.Action) *Parser {
	return &Parser{
		rowData: raw,
		actions: actions,
	}
}

func (p *Parser) PrepareData() []*model.Operation {
	var operations []*model.Operation
	for _, ui := range p.rowData {
		ok := p.actions.IsAction(ui.Action)
		if !ok {
			fmt.Fprintf(os.Stderr, "оператор %s не поддерживается\n", ui.Action)
			continue
		}

		operation := model.NewOperation(ui.N1, ui.N2, model.Action(ui.Action))
		operations = append(operations, operation)

	}
	return operations
}
