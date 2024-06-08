package tasks

import (
	"dag_runner/types"
	"fmt"
	"time"
)

type PrintATask struct{}

func NewPrintATask() PrintATask {
	return PrintATask{}
}

func (t PrintATask) Run(reqCtx *types.ReqContext) error {
	fmt.Println("A")
	time.Sleep(1 * time.Second)
	return nil
}

func (t PrintATask) Deps() []string {
	return []string{}
}
