package dag

import (
	"dag_runner/tasks"
	"dag_runner/types"
)

var testTasks = map[string]Task{
	"print_a": tasks.NewPrintATask(),
	"print_b": tasks.NewPrintBTask(),
	"print_c": tasks.NewPrintCTask(),
	"print_d": tasks.NewPrintDTask(),
}

func NewTestDag(reqCtx *types.ReqContext) *Runner {
	return NewRunner(reqCtx, testTasks)
}
