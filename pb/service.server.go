package pb

import (
	"context"
	"fmt"
	pb "github.com/ngyewch/protoc-gen-rsocket-go-example/pb/request/pb"
	runtime "github.com/ngyewch/protoc-gen-rsocket-go/runtime"
	proto "google.golang.org/protobuf/proto"
)

type TestServiceServer struct {
	selector uint64
	service  TestService
}

func NewTestServiceServer(selector uint64, service TestService) *TestServiceServer {
	return &TestServiceServer{
		selector: selector,
		service:  service,
	}
}
func (s *TestServiceServer) Selector() uint64 {
	return s.selector
}
func (s *TestServiceServer) HandleRequestResponse(ctx context.Context, reqWrapper *runtime.RequestWrapper) (proto.Message, error) {
	if reqWrapper.Selector != s.selector {
		return nil, runtime.ErrorSelectorMismatch
	}
	switch reqWrapper.MethodName {
	case "Test":
		var req pb.TestRequest
		err := proto.Unmarshal(reqWrapper.Payload, &req)
		if err != nil {
			return nil, err
		}
		return s.service.Test(ctx, &req)
	default:
		return nil, fmt.Errorf("unknown method: %s", reqWrapper.MethodName)
	}
}

type TestService2Server struct {
	selector uint64
	service  TestService2
}

func NewTestService2Server(selector uint64, service TestService2) *TestService2Server {
	return &TestService2Server{
		selector: selector,
		service:  service,
	}
}
func (s *TestService2Server) Selector() uint64 {
	return s.selector
}
func (s *TestService2Server) HandleRequestResponse(ctx context.Context, reqWrapper *runtime.RequestWrapper) (proto.Message, error) {
	if reqWrapper.Selector != s.selector {
		return nil, runtime.ErrorSelectorMismatch
	}
	switch reqWrapper.MethodName {
	case "Test":
		var req pb.TestRequest
		err := proto.Unmarshal(reqWrapper.Payload, &req)
		if err != nil {
			return nil, err
		}
		return s.service.Test(ctx, &req)
	default:
		return nil, fmt.Errorf("unknown method: %s", reqWrapper.MethodName)
	}
}

type TestListenerServer struct {
	selector uint64
	service  TestListener
}

func NewTestListenerServer(selector uint64, service TestListener) *TestListenerServer {
	return &TestListenerServer{
		selector: selector,
		service:  service,
	}
}
func (s *TestListenerServer) Selector() uint64 {
	return s.selector
}
func (s *TestListenerServer) HandleRequestResponse(ctx context.Context, reqWrapper *runtime.RequestWrapper) (proto.Message, error) {
	if reqWrapper.Selector != s.selector {
		return nil, runtime.ErrorSelectorMismatch
	}
	switch reqWrapper.MethodName {
	case "OnEvent":
		var req Event
		err := proto.Unmarshal(reqWrapper.Payload, &req)
		if err != nil {
			return nil, err
		}
		return s.service.OnEvent(ctx, &req)
	default:
		return nil, fmt.Errorf("unknown method: %s", reqWrapper.MethodName)
	}
}
