package order

import (
	"fmt"

	"github.com/sanprasirt/go-grpc-api-gateway/pkg/config"
	"github.com/sanprasirt/go-grpc-api-gateway/pkg/order/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.OrderServiceClient
}

func InitServiceClient(c *config.Config) pb.OrderServiceClient {
	// using WihtInsecure() because no SSL running
	cc, err := grpc.Dial(c.OrderSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Cloud not connect:", err)
	}

	return pb.NewOrderServiceClient(cc)
}