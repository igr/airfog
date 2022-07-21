package actions

import (
	"foursquare.com/airfog/cmd/util"
	"foursquare.com/airfog/internal/airflow"
)

func PatchPauseDag(ctx ApiCtx, dag airflow.DAG, pause bool) airflow.DAG {
	if dag.IsPaused.IsSet() {
		if *dag.IsPaused.Get() == pause {
			return dag
		}
	}
	dag.IsPaused.Set(&pause)

	requestDag := airflow.DAG{}
	requestDag.IsPaused.Set(&pause)

	newDag, res, err := ctx.Api.DAGApi.
		PatchDag(ctx.ApiContext, *dag.DagId).
		DAG(requestDag).
		UpdateMask([]string{"is_paused"}).
		Execute()

	if err != nil {
		util.Exit(res)
	}

	return newDag
}
