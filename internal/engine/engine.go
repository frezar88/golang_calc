package engine

import (
	"calculator/internal/model"
	"errors"
	"time"
)

var ErrorUnSupportAction = errors.New("unknown action \n")
var ErrorZeroValue = errors.New("на ноль делить нельзя \n")

type Engine struct {
	delay int
}

func NewEngine(delay int) *Engine {
	return &Engine{delay: delay}
}

func (e *Engine) Calculate(operations []*model.Operation) {
	for _, operation := range operations {
		time.Sleep(time.Duration(e.delay) * time.Second)
		var result int
		var resultError error
		switch operation.Action {
		case model.ActionMinus:
			result = operation.Number1 - operation.Number2
		case model.ActionPlus:
			result = operation.Number1 + operation.Number2
		case model.ActionMultiple:
			result = operation.Number1 * operation.Number2
		case model.ActionDivision:
			if operation.Number2 == 0 {
				resultError = ErrorZeroValue
				break
			}
			result = operation.Number1 / operation.Number2
		default:
			resultError = ErrorUnSupportAction
		}
		operation.Result = struct {
			Data  int
			Error error
		}{
			Data: result, Error: resultError,
		}
	}
}
