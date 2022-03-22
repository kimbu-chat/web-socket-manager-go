package services

import (
	"context"
	"fmt"

	"github.com/kimbu-chat/web-socket-manager-go/internal/apiproto"
	"github.com/kimbu-chat/web-socket-manager-go/internal/config"
)

func BroadcastData(userIds []int64, data []byte) error {
	client := config.GetGRPCClient()

	request := apiproto.BroadcastRequest{
		Channels: convertIntArrayToString(userIds),
		Data:     data,
	}

	response, err := client.Broadcast(context.Background(), &request)
	if err != nil {
		return err
	}

	if response.Error != nil {
		fmt.Printf("PublishToUsers error, code: %v. Message: %v\n", response.Error.Code, response.Error.Message)
	}

	return nil
}

func convertIntArrayToString(numbers []int64) []string {
	length := len(numbers)

	result := make([]string, length)

	for i := 0; i < length; i++ {
		result[i] = fmt.Sprintf("#%v", fmt.Sprint(numbers[i]))
	}

	return result
}