package main

import (
	"foursquare.com/airfog/cmd/actions"
	"foursquare.com/airfog/cmd/util"
	"strconv"
)

type CleartiCmd struct {
	Dag      string `required:"true" help:"DagID"`
	DagRunId string `optional:"true" help:"DagRunID. If omitted, the latest run will be used."`
}

func (cmd *CleartiCmd) Run(apiCtx actions.ApiCtx, globals *Globals) error {
	util.PrintValue("DagID:", cmd.Dag)
	if cmd.DagRunId == "" {
		util.PrintInfo("DagRunID not specified, using the latest.")
		cmd.DagRunId = *actions.GetLatestDagRun(apiCtx, cmd.Dag)
		if &cmd.DagRunId == nil {
			util.Error("DagRunId not found")
		}
	}
	util.PrintValue("DagRunID:", cmd.DagRunId)

	failedTaskIds := actions.GetTaskInstances(apiCtx, cmd.Dag, cmd.DagRunId)

	if len(failedTaskIds) == 0 {
		util.Print("Nothing to clear.")
		return nil
	}

	util.Print("Clearing " + strconv.Itoa(len(failedTaskIds)) + " task instances.")

	actions.PostClearTaskInstances(apiCtx, cmd.Dag, failedTaskIds)

	return nil
}
