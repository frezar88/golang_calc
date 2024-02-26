package printer

import (
	"calculator/internal/model"
	"fmt"
	"os"
)

type Printer struct {
}

func NewPrinter() *Printer {
	return &Printer{}
}

func (p Printer) Print(operations []*model.Operation) {
	for _, o := range operations {
		if o.Result.Error != nil {
			fmt.Fprint(os.Stdout, "Ошибка вычисления: ", o.Result.Error)
			continue
		}
		fmt.Fprintf(os.Stdout, "%d %s %d = %d\n", o.Number1, o.Action, o.Number2, o.Result.Data)

	}
}
