package tasks

import (
	"dag_runner/types"
	"fmt"
	"time"
)

type PrintDTask struct{}

func NewPrintDTask() PrintDTask {
	return PrintDTask{}
}

func (t PrintDTask) Run(reqCtx *types.ReqContext) error {
	fmt.Println("D")
	time.Sleep(1 * time.Second)
	return nil
}

func (t PrintDTask) Deps() []string {
	return []string{"print_a"}
}
