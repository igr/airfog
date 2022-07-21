package actions

import (
	"foursquare.com/airfog/cmd/util"
	"foursquare.com/airfog/internal/airflow"
)

func GetAllDags(ctx ApiCtx) []airflow.DAG {
	var allDags []airflow.DAG
	off := 0
	limit := 100
	total := 0
	for {
		dagsCollection, res, err := ctx.Api.DAGApi.
			GetDags(ctx.ApiContext).
			Offset(int32(off)).
			Limit(int32(limit)).
			Execute()

		if err != nil {
			util.Exit(res)
		}

		dags := *dagsCollection.Dags
		totalInstances := *dagsCollection.TotalEntries

		for _, dag := range dags {
			allDags = append(allDags, dag)
		}

		off += len(dags)
		total += len(dags)

		if total >= int(totalInstances) {
			break
		}
	}
	return allDags
}
