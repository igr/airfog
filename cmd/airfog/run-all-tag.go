package main

import (
	"foursquare.com/airfog/cmd/actions"
	"foursquare.com/airfog/cmd/util"
	"strconv"
)

type RunAllTagCmd struct {
	Tag string `required:"true" help:"Tag name"`
}

func (cmd *RunAllTagCmd) Run(apiCtx actions.ApiCtx, globals *Globals) error {
	util.PrintValue("Tag:", cmd.Tag)

	dags := actions.GetAllDags(apiCtx)

	if len(dags) == 0 {
		util.Print("Nothing to run.")
		return nil
	}

	counter := 0

	for _, dag := range dags {
		for _, tag := range dag.Tags {
			if *tag.Name == cmd.Tag {
				util.PrintValue("DAG:", *dag.DagId)
				actions.PatchPauseDag(apiCtx, dag, false)
				actions.PostRunDag(apiCtx, *dag.DagId)
				counter++
			}
		}
	}

	util.Print("Ran " + strconv.Itoa(counter) + " DAGs.")

	return nil
}
