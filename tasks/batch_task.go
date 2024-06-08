package tasks

import (
	"dag_runner/types"
)

type Runnable interface {
	Run(*types.ReqContext) error
}

type BatchTask struct {
	subTasks []Runnable
	deps     []string
}

func NewSerialTask(subTasks []Runnable, deps []string) *BatchTask {
	return &BatchTask{subTasks: subTasks, deps: deps}
}

func (t *BatchTask) Run(rc *types.ReqContext) error {
	for _, task := range t.subTasks {
		err := task.Run(rc)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *BatchTask) Deps() []string {
	return t.deps
}

func (t *BatchTask) AddSubTask(r ...Runnable) {
	t.subTasks = append(t.subTasks, r...)
}

func (t *BatchTask) AddDependency(d ...string) {
	t.deps = append(t.deps, d...)
}
