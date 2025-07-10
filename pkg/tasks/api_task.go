package tasks

import (
	"encoding/json"
	"strconv"

	"github.com/hibiken/asynq"
	"github.com/wushiling50/aster/pkg/constants"
)

type APIPayload struct {
	Type   int    `json:"type"`
	Id     int64  `json:"id"`
	TaskId string `json:"taskId"`
}

func GetNewAPITaskKey(fetchType int, id int64, reqId string) string {
	return constants.APITaskName + constants.TaskSeparator + strconv.Itoa(fetchType) +
		constants.TaskSeparator + strconv.Itoa(int(id)) + constants.TaskSeparator + reqId
}

func NewAPITask(fetchType int, id int64, reqId string) (*asynq.Task, string, error) {
	taskId := GetNewAPITaskKey(fetchType, id, reqId)

	payload, err := json.Marshal(APIPayload{
		Type:   fetchType,
		Id:     id,
		TaskId: taskId,
	})
	if err != nil {
		return nil, "", err
	}

	return asynq.NewTask(constants.APITaskName, payload), taskId, nil
}
