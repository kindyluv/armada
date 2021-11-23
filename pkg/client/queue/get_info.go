package queue

import (
	"context"
	"fmt"
	"time"

	"github.com/G-Research/armada/pkg/api"
	"github.com/G-Research/armada/pkg/client"
)

type GetInfoAPI func(string) (*api.QueueInfo, error)

func GetInfo(getConnectionDetails client.ConnectionDetails) GetInfoAPI {
	return func(queueName string) (*api.QueueInfo, error) {
		conn, err := client.CreateApiConnection(getConnectionDetails())
		if err != nil {
			return nil, fmt.Errorf("failed to connect to api because %s", err)
		}
		defer conn.Close()

		client := api.NewSubmitClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		queueInfo, e := client.GetQueueInfo(ctx, &api.QueueInfoRequest{Name: queueName})
		if e != nil {
			return nil, fmt.Errorf("get queue info request failed: %s", err)
		}

		return queueInfo, nil
	}
}
