package tasks

import (
	"dag_runner/types"
	"fmt"
)

func NewBatchPrintEFTask() *BatchTask {
	t := new(BatchTask)
	t.AddSubTask(new(PrintETask), new(PrintFTask))
	t.AddDependency("print_c", "print_d")
	return t
}

type PrintETask struct{}

func (t PrintETask) Run(reqCtx *types.ReqContext) error {
	fmt.Println("E")
	return nil
}

type PrintFTask struct{}

func (t PrintFTask) Run(reqCtx *types.ReqContext) error {
	fmt.Println("F")
	return nil
}
