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
	Url string `json:"url"`
}

func CollectStat(request CollectStatRequest, awsSession session.Session) int {
	timestamp := time.Now().UnixMicro()

	stat := models.StatModel{
		Timestamp: timestamp,
		Url:       request.Url,
	}

	client := dynamodb.New(&awsSession)

	dynamoObj, err := dynamodbattribute.MarshalMap(stat)
	if err != nil {
		return http.StatusBadRequest
	}

	tableName := utilities.GetStatsTableName()
	println("Table name:", tableName)
	reqInput := &dynamodb.PutItemInput{
		Item:      dynamoObj,
		TableName: &tableName,
	}
	_, putErr := client.PutItem(reqInput)
	if putErr != nil {
		println("Error Here", putErr.Error())
		return http.StatusInternalServerError
	}

	return http.StatusOK

}
