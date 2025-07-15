package tasks

import (
	"errors"

	"github.com/hibiken/asynq"
	"github.com/wushiling50/aster/pkg/constants"
	"github.com/wushiling50/aster/pkg/errno"
)

func FetcherTaskPusher(c *asynq.Client, fetchType int, id int64, updateAfter string, searchLimit int64) (err error) {
	var (
		task   *asynq.Task
		taskId string
	)

	if task, taskId, err = NewFetcherTask(fetchType, id, updateAfter, searchLimit); err != nil {
		err = errno.InternalAsynqError.WithError(err)
		return
	}

	_, err = c.Enqueue(
		task,
		asynq.TaskID(taskId),
		asynq.Queue(constants.FetcherTaskQueue),
		asynq.MaxRetry(constants.FetchMaxRetry),
	)

	if err != nil {
		if errors.Is(err, asynq.ErrTaskIDConflict) {
			err = nil
		} else {
			err = errno.InternalAsynqError.WithError(err)
			return
		}
	}

	return
}
