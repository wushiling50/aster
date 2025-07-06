package tasks

import (
	"github.com/hibiken/asynq"
	"github.com/wushiling50/aster/pkg/constants"
)

func FetcherTaskPusher(c *asynq.Client, fetchType int, id int64, updateAfter string, searchLimit int64) (err error) {
	var (
		task   *asynq.Task
		taskId string
	)

	if task, taskId, err = NewFetcherTask(fetchType, id, updateAfter, searchLimit); err != nil {
		return
	}

	_, err = c.Enqueue(
		task,
		asynq.TaskID(taskId),
		asynq.Queue(constants.FetcherTaskQueue),
		asynq.MaxRetry(constants.FetchMaxRetry),
	)

	if err != nil {
		return
	}

	return
}
