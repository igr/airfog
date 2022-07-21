package actions

import (
	"fmt"
	"foursquare.com/airfog/cmd/util"
	"foursquare.com/airfog/internal/airflow"
	"time"
)

func PostRunDag(ctx ApiCtx, dagId string) airflow.DAGRun {
	currentTime := time.Now()
	dagRunId := fmt.Sprintf("%s-run-%s", dagId, currentTime.Format("20060102150405-Z07"))
	dagRun := airflow.NewDAGRun()
	dagRun.DagRunId.Set(&dagRunId)

	createdDagRun, res, err := ctx.Api.DAGRunApi.
		PostDagRun(ctx.ApiContext, dagId).
		DAGRun(*dagRun).
		Execute()

	if err != nil {
		util.Exit(res)
	}

	return createdDagRun
}
