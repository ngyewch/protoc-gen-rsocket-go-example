package client

import (
	"context"
	"fmt"
	"github.com/ngyewch/protoc-gen-rsocket-go-example/pb"
	pb1 "github.com/ngyewch/protoc-gen-rsocket-go-example/pb/request/pb"
	"github.com/ngyewch/protoc-gen-rsocket-go/runtime"
	"github.com/rsocket/rsocket-go"
)

type Client struct {
}

func New() *Client {
	return &Client{}
}

func (c *Client) Start(ctx context.Context) error {
	ctx1, cancel := context.WithCancelCause(ctx)
	client, err := rsocket.Connect().
		OnClose(func(err error) {
			cancel(fmt.Errorf("connection closed: %w", err))
		}).
		Acceptor(func(ctx context.Context, socket rsocket.RSocket) rsocket.RSocket {
			return rsocket.NewAbstractSocket(
				rsocket.RequestResponse(runtime.RSocketServerRequestResponseHandler(runtime.NewServers(
					pb.NewTestListenerServer(3, c),
				))),
			)
		}).
		Transport(rsocket.TCPClient().
			SetAddr("127.0.0.1:7878").
			Build(),
		).
		Start(ctx1)
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

	<-ctx1.Done()
	return ctx1.Err()
}

func (c *Client) OnEvent(ctx context.Context, event *pb.Event) (*pb.Void, error) {
	fmt.Printf("<< %s\n", event.Id)
	return &pb.Void{}, nil
}
