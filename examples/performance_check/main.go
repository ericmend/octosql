package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/cube2222/octosql/datasources/json"
	"github.com/cube2222/octosql/execution"
	"github.com/cube2222/octosql/execution/aggregates"
	"github.com/cube2222/octosql/execution/nodes"
)

func main() {
	start := time.Now()

	// entries := make([]memory.Entry, 2000000)
	// for i := range entries {
	// 	entries[i] = memory.Entry{Record: execution.NewRecord([]octosql.Value{octosql.NewInt(rand.Intn(4)), octosql.NewString(getRandomGroupName()), octosql.NewInt(rand.Intn(100)), octosql.NewInt(rand.Intn(50))}, false)}
	// }

	var plan execution.Node

	// plan = &memory.Datasource{
	// 	Entries: entries,
	// }
	plan = &json.Datasource{
		Path:   "goals_big.json",
		Fields: []string{"time", "team"},
	}

	// TODO: Add map and filter in between.

	plan = nodes.NewGroupBy(
		[]func() nodes.Aggregate{aggregates.NewCountPrototype()},
		[]execution.Expression{execution.NewVariable(0, 0)},
		[]execution.Expression{execution.NewVariable(0, 1)},
		plan,
		execution.NewCountingTriggerPrototype(100000),
	)

	if err := plan.Run(
		execution.ExecutionContext{
			Context:         context.Background(),
			VariableContext: &execution.VariableContext{},
		},
		func(ctx execution.ProduceContext, record execution.Record) error {
			log.Printf("%+v", record)
			return nil
		},
		func(ctx execution.ProduceContext, msg execution.MetadataMessage) error {
			return nil
		},
	); err != nil {
		log.Fatal(err)
	}
	log.Println(time.Since(start))
}

func getRandomGroupName() string {
	switch rand.Intn(10) {
	case 0, 1, 2, 3:
		return "test1"
	case 4, 5:
		return "test2"
	case 6, 7, 8:
		return "test3"
	case 9:
		return "test4"
	}
	return "bad"
}
