package services

import (
	"context"
	"fmt"
	"strconv"

	"github.com/kimbu-chat/web-socket-manager-go/internal/apiproto"
	"github.com/kimbu-chat/web-socket-manager-go/internal/config"
)

func BroadcastData(userIds []int, data []byte) error {
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

func convertIntArrayToString(numbers []int) []string {
	length := len(numbers)

	result := make([]string, length)

	for i := 0; i < length; i++ {
		result[i] = strconv.Itoa(numbers[i])
	}

	return result
}
