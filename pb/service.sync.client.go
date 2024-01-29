package pb

import (
	"context"
	pb "github.com/ngyewch/protoc-gen-rsocket-go-example/pb/request/pb"
	pb1 "github.com/ngyewch/protoc-gen-rsocket-go-example/pb/response/pb"
	runtime "github.com/ngyewch/protoc-gen-rsocket-go/runtime"
	rsocketgo "github.com/rsocket/rsocket-go"
	proto "google.golang.org/protobuf/proto"
)

type TestServiceClient struct {
	selector uint64
	handler  runtime.ClientRequestResponseHandler
}

func NewTestServiceClient(selector uint64, handler runtime.ClientRequestResponseHandler) *TestServiceClient {
	return &TestServiceClient{
		handler:  handler,
		selector: selector,
	}
}
func NewRSocketTestServiceClient(selector uint64, rs rsocketgo.RSocket) *TestServiceClient {
	return &TestServiceClient{
		handler:  runtime.RSocketClientRequestResponseHandler(rs),
		selector: selector,
	}
}
func (c *TestServiceClient) Test(ctx context.Context, req *pb.TestRequest) (*pb1.TestResponse, error) {
	rspBytes, err := runtime.HandleClientRequestResponse(ctx, c.selector, "Test", req, c.handler)
	if err != nil {
		return nil, err
	}
	var rsp pb1.TestResponse
	err = proto.Unmarshal(rspBytes, &rsp)
	if err != nil {
		return nil, err
	}
	return &rsp, nil
}

type TestService2Client struct {
	selector uint64
	handler  runtime.ClientRequestResponseHandler
}

func NewTestService2Client(selector uint64, handler runtime.ClientRequestResponseHandler) *TestService2Client {
	return &TestService2Client{
		handler:  handler,
		selector: selector,
	}
}
func NewRSocketTestService2Client(selector uint64, rs rsocketgo.RSocket) *TestService2Client {
	return &TestService2Client{
		handler:  runtime.RSocketClientRequestResponseHandler(rs),
		selector: selector,
	}
}
func (c *TestService2Client) Test(ctx context.Context, req *pb.TestRequest) (*pb1.TestResponse, error) {
	rspBytes, err := runtime.HandleClientRequestResponse(ctx, c.selector, "Test", req, c.handler)
	if err != nil {
		return nil, err
	}
	var rsp pb1.TestResponse
	err = proto.Unmarshal(rspBytes, &rsp)
	if err != nil {
		return nil, err
	}
	return &rsp, nil
}

type TestListenerClient struct {
	selector uint64
	handler  runtime.ClientRequestResponseHandler
}

func NewTestListenerClient(selector uint64, handler runtime.ClientRequestResponseHandler) *TestListenerClient {
	return &TestListenerClient{
		handler:  handler,
		selector: selector,
	}
}
func NewRSocketTestListenerClient(selector uint64, rs rsocketgo.RSocket) *TestListenerClient {
	return &TestListenerClient{
		handler:  runtime.RSocketClientRequestResponseHandler(rs),
		selector: selector,
	}
}
func (c *TestListenerClient) OnEvent(ctx context.Context, req *Event) (*Void, error) {
	rspBytes, err := runtime.HandleClientRequestResponse(ctx, c.selector, "OnEvent", req, c.handler)
	if err != nil {
		return nil, err
	}
	var rsp Void
	err = proto.Unmarshal(rspBytes, &rsp)
	if err != nil {
		return nil, err
	}
	return &rsp, nil
}
