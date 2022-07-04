package actions

import "foursquare.com/airfog/cmd/util"

func GetTaskInstances(ctx ApiCtx, dagId string, dagRunId string) []string {
	var failedTaskIds []string
	off := 0
	limit := 100
	total := 0
	for {
		tasks, res, err := ctx.Api.TaskInstanceApi.
			GetTaskInstances(ctx.ApiContext, dagId, dagRunId).
			Offset(int32(off)).
			Limit(int32(limit)).
			Execute()

		if err != nil {
			util.Exit(res)
		}

		taskInstances := *tasks.TaskInstances
		totalInstances := *tasks.TotalEntries
		for _, task := range taskInstances {
			if task.State != nil && *task.State != "success" {
				failedTaskIds = append(failedTaskIds, *task.TaskId)
			}
		}

		off += len(taskInstances)
		total += len(taskInstances)

		if total >= int(totalInstances) {
			break
		}
	}
	return failedTaskIds
}
