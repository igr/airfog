package main

import (
	"foursquare.com/airfog/cmd/actions"
)

func main() {
	ctx := actions.AirflowApi()

	runCli(ctx)

	//dagRun, res, err := api.DAGRunApi.GetDagRun(ctx, dagId, *dagRunId).Execute()
	//if err != nil {
	//	exit(res)
	//}
}
