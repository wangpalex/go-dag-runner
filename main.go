package main

import (
	"dag_runner/dag"
	"dag_runner/types"
	"fmt"
	"time"
)

func main() {
	reqCtx := &types.ReqContext{}
	runner := dag.NewTestDag(reqCtx)
	start := time.Now()
	runner.Run()
	runner.Wait()
	fmt.Printf("Spent %f secs", time.Since(start).Seconds())
}
