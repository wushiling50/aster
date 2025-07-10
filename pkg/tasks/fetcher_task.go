package tasks

import (
	"encoding/json"
	"strconv"

	"github.com/hibiken/asynq"
	"github.com/wushiling50/aster/pkg/constants"
)

type FetchPayload struct {
	Type        int    `json:"type"`
	Id          int64  `json:"id"`
	UpdateAfter string `json:"updateAfter"`
	SearchLimit int64  `json:"searchLimit"`
}

func GetNewFetcherTaskKey(fetchType int, id int64) string {
	return constants.FetcherTaskName + constants.TaskSeparator + strconv.Itoa((fetchType)) +
		constants.TaskSeparator + strconv.Itoa(int(id))
}

func NewFetcherTask(fetchType int, id int64, updateAfter string, searchLimit int64) (*asynq.Task, string, error) {
	payload, err := json.Marshal(FetchPayload{
		Type:        fetchType,
		Id:          id,
		UpdateAfter: updateAfter,
		SearchLimit: searchLimit,
	})
	if err != nil {
		return nil, "", err
	}
	return asynq.NewTask(constants.FetcherTaskName, payload), GetNewFetcherTaskKey(fetchType, id), nil
}
