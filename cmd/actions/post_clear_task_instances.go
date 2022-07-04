package actions

import (
	"foursquare.com/airfog/cmd/util"
	"foursquare.com/airfog/internal/airflow"
)

func PostClearTaskInstances(ctx ApiCtx, dagId string, taskIds []string) {
	includeParentDag := true
	dryRun := false

	_, res, err := ctx.Api.DAGApi.PostClearTaskInstances(ctx.ApiContext, dagId).ClearTaskInstance(airflow.ClearTaskInstance{
		DryRun:           &dryRun,
		TaskIds:          &taskIds,
		IncludeParentdag: &includeParentDag,
	}).Execute()
	if err != nil {
		util.Exit(res)
	}
}
