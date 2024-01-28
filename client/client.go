package client

import (
	"context"
	"fmt"
	"github.com/ngyewch/protoc-gen-rsocket-go-example/pb"
	pb1 "github.com/ngyewch/protoc-gen-rsocket-go-example/pb/request/pb"
	"github.com/rsocket/rsocket-go"
)

type Client struct {
}

func New() *Client {
	return &Client{}
}

func (c *Client) Start(ctx context.Context) error {
	client, err := rsocket.Connect().
		Transport(rsocket.TCPClient().
			SetAddr("127.0.0.1:7878").
			Build(),
		).
		Start(ctx)
	if err != nil {
		return err
	}
	defer func(client rsocket.Client) {
		_ = client.Close()
	}(client)

	async := true

	if async {
		testServiceClientAsync := pb.NewRSocketTestServiceClientAsync(0, client)
		rsp, err := testServiceClientAsync.Test(ctx, &pb1.TestRequest{
			Text: "Hello, world!",
		}).Collect()
		if err != nil {
			return err
		}
		fmt.Printf("%+v\n", rsp)

		testService2ClientAsync := pb.NewRSocketTestService2ClientAsync(1, client)
		rsp, err = testService2ClientAsync.Test(ctx, &pb1.TestRequest{
			Text: "Hello, world!",
		}).Collect()
		if err != nil {
			return err
		}
		fmt.Printf("%+v\n", rsp)
	} else {
		testServiceClient := pb.NewRSocketTestServiceClient(0, client)
		rsp, err := testServiceClient.Test(ctx, &pb1.TestRequest{
			Text: "Hello, world!",
		})
		if err != nil {
			return err
		}
		fmt.Printf("%+v\n", rsp)

		testService2Client := pb.NewRSocketTestService2Client(1, client)
		rsp, err = testService2Client.Test(ctx, &pb1.TestRequest{
			Text: "Hello, world!",
		})
		if err != nil {
			return err
		}
		fmt.Printf("%+v\n", rsp)
	}

	return nil
}
