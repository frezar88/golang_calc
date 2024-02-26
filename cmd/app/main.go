package main

import (
	"bufio"
	"calculator/internal/calculator"
	"calculator/internal/engine"
	"calculator/internal/model"
	"calculator/internal/parser"
	"calculator/internal/printer"
	"fmt"
	"os"
)

func main() {
	const Done = "done\n"
	const NewLine = '\n'

	fmt.Fprintln(os.Stdout, "Привет. Введите выражение для вычисления. Когда закончите напишите: done")
	fmt.Fprintln(os.Stdout, "Возможные выражения: - , +")
	fmt.Fprintln(os.Stdout, "Маска выражения: <число1><пробел><действие><пробел><число2>")
	fmt.Fprintln(os.Stdout, "Пример выражения: 2 + 2")

	var userInputs []model.UserInput

	for {
		in := bufio.NewReader(os.Stdin)
		str, err := in.ReadString(NewLine)
		if err != nil {
			fmt.Fprintln(os.Stderr, "ошибка ввода", err)
			continue
		}
		if str == Done {
			break
		}

		var n1, n2 int
		var action string

		n, err := fmt.Sscanf(str, "%d %s %d", &n1, &action, &n2)
		if n != 3 || err != nil {
			fmt.Fprintln(os.Stderr, "ошибка ввода", err)
			continue
		}

		ui := model.UserInput{
			N1:     n1,
			N2:     n2,
			Action: action,
		}
		userInputs = append(userInputs, ui)
	}

	if len(userInputs) == 0 {
		fmt.Fprintln(os.Stderr, "Нет выражений - нет вычислений")
		return
	}
	calcParser := parser.NewParser(
		userInputs,
		model.ActionMinus,
		model.ActionPlus,
		model.ActionMultiple,
		model.ActionDivision,
	)
	calcEngine := engine.NewEngine(0)
	calcPrinter := printer.NewPrinter()

	appCalculator := calculator.NewCalculator(*calcParser, *calcEngine, *calcPrinter)
	appCalculator.Run(userInputs)

	fmt.Fprint(os.Stdout, "exit")

}
