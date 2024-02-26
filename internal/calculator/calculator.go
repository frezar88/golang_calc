package calculator

import (
	"calculator/internal/engine"
	"calculator/internal/model"
	"calculator/internal/parser"
	"calculator/internal/printer"
)

type Calculator struct {
	parser  parser.Parser
	engine  engine.Engine
	printer printer.Printer
}

func NewCalculator(
	calcParser parser.Parser,
	calcEngine engine.Engine,
	calcPrinter printer.Printer,
) *Calculator {
	o := new(Calculator)
	o.parser = calcParser
	o.engine = calcEngine
	o.printer = calcPrinter
	return o
}

func (c Calculator) Run(rawData []model.UserInput) {
	operations := c.parser.PrepareData()
	c.engine.Calculate(operations)
	c.printer.Print(operations)
}
