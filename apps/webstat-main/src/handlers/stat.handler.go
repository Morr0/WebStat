package handlers

import (
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/morr0/webstat/apps/webstat-main/src/models"
	"github.com/morr0/webstat/apps/webstat-main/src/utilities"
)

type CollectStatRequest struct {
	Url string
}

func CollectStat(request CollectStatRequest, awsSession session.Session) int {
	timestamp := time.Now().UnixMicro()

	stat := models.StatModel{
		Url:       request.Url,
		CreatedAt: timestamp,
	}

	client := dynamodb.New(&awsSession)

	dynamoObj, err := dynamodbattribute.MarshalMap(stat)
	if err != nil {
		return http.StatusBadRequest
	}

	tableName := utilities.GetStatsTableName()
	reqInput := &dynamodb.PutItemInput{
		Item:      dynamoObj,
		TableName: &tableName,
	}
	_, putErr := client.PutItem(reqInput)
	if putErr != nil {
		return http.StatusInternalServerError
	}

	return http.StatusOK

}
