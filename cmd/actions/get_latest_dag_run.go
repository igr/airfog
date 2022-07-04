package actions

import (
	"foursquare.com/airfog/cmd/util"
)

func GetLatestDagRun(ctx ApiCtx, dagId string) *string {
	dagRunsBatch, res, err := ctx.Api.DAGRunApi.GetDagRuns(ctx.ApiContext, dagId).
		Limit(10).
		Offset(0).
		OrderBy("-execution_date").
		Execute()
	if err != nil {
		util.Exit(res)
	}
	dagRuns := dagRunsBatch.GetDagRuns()
	if len(dagRuns) == 0 {
		return nil
	}
	return dagRuns[0].DagRunId.Get()
}
