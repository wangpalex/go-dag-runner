package tasks

import (
	"dag_runner/types"
	"fmt"
	"time"
)

type PrintCTask struct{}

func NewPrintCTask() PrintCTask {
	return PrintCTask{}
}

func (t PrintCTask) Run(reqCtx *types.ReqContext) error {
	fmt.Println("C")
	time.Sleep(1 * time.Second)
	return nil
}

func (t PrintCTask) Deps() []string {
	return []string{"print_b"}
}
