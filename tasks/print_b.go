package tasks

import (
	"dag_runner/types"
	"fmt"
	"time"
)

type PrintBTask struct{}

func NewPrintBTask() PrintBTask {
	return PrintBTask{}
}

func (t PrintBTask) Run(reqCtx *types.ReqContext) error {
	fmt.Println("B")
	time.Sleep(1 * time.Second)
	return nil
}

func (t PrintBTask) Deps() []string {
	return []string{"print_a"}
}
