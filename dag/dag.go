package dag

import (
	"dag_runner/types"
	"fmt"
	"sync"
)

type Task interface {
	Run(*types.ReqContext) error
	Deps() []string
}

type Runner struct {
	nodes map[string]*Node
	wg    *sync.WaitGroup
}

func NewRunner(reqCtx *types.ReqContext, tasks map[string]Task) *Runner {
	// Init DAG nodes
	nodeMap := make(map[string]*Node, len(tasks))
	for name, task := range tasks {
		nodeMap[name] = &Node{
			Name:    name,
			Task:    task,
			reqCtx:  reqCtx,
			deps:    task.Deps(),
			wg:      new(sync.WaitGroup),
			doneWGs: make([]*sync.WaitGroup, 0),
		}
	}
	// Construct dependencies
	outCnt := make(map[string]int, len(nodeMap))
	for _, node := range nodeMap {
		for _, dep := range node.deps {
			node.wg.Add(1)
			nodeMap[dep].doneWGs = append(nodeMap[dep].doneWGs, node.wg)
			outCnt[dep]++
		}
	}
	// Construct runner
	runner := &Runner{
		nodes: nodeMap,
		wg:    new(sync.WaitGroup),
	}
	for name, node := range nodeMap {
		if outCnt[name] == 0 { // Is last level node
			runner.wg.Add(1)
			node.doneWGs = append(node.doneWGs, runner.wg)
		}
	}

	return runner
}

func (d *Runner) Run() {
	for _, node := range d.nodes {
		go node.run()
	}
}

func (d *Runner) Wait() {
	d.wg.Wait()
}

type Node struct {
	Name   string
	Task   Task
	Error  error
	reqCtx *types.ReqContext

	deps    []string
	wg      *sync.WaitGroup
	doneWGs []*sync.WaitGroup
}

func (t *Node) run() {
	defer t.done()
	t.wait()
	defer t.recover()
	t.Error = t.Task.Run(t.reqCtx)
}

func (t *Node) wait() {
	if t.wg != nil {
		t.wg.Wait()
	}
}

func (t *Node) done() {
	for _, wg := range t.doneWGs {
		wg.Done()
	}
}

func (t *Node) recover() {
	if e := recover(); e != nil {
		t.Error = fmt.Errorf("[%s] panic: %s", t.Name, e)
	}
}
