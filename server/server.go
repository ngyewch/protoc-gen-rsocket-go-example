package server

import (
	"context"
	"fmt"
	"github.com/ngyewch/protoc-gen-rsocket-go-example/pb"
	pb1 "github.com/ngyewch/protoc-gen-rsocket-go-example/pb/request/pb"
	pb2 "github.com/ngyewch/protoc-gen-rsocket-go-example/pb/response/pb"
	"github.com/ngyewch/protoc-gen-rsocket-go/runtime"
	"github.com/rsocket/rsocket-go"
	"github.com/rsocket/rsocket-go/payload"
	"strings"
	"time"
)

type Server struct {
}

func New() *Server {
	return &Server{}
}

func (s *Server) Start(ctx context.Context) error {
	err := rsocket.Receive().
		Acceptor(func(ctx context.Context, setup payload.SetupPayload, sendingSocket rsocket.CloseableRSocket) (rsocket.RSocket, error) {
			ctx1, cancel := context.WithCancelCause(ctx)
			sendingSocket.OnClose(func(err error) {
				fmt.Println("**** client disconnected")
				cancel(err)
			})
			c := pb.NewRSocketTestListenerClient(3, sendingSocket)
			go func() {
				for {
					_, err := c.OnEvent(ctx1, &pb.Event{
						Id: time.Now().String(),
					})
					if err != nil {
						fmt.Printf("err: %v\n", err)
					}
					select {
					case <-time.After(5 * time.Second):
						break
					case <-ctx1.Done():
						return
					}
				}
			}()
			return rsocket.NewAbstractSocket(
				rsocket.RequestResponse(runtime.RSocketServerRequestResponseHandler(
					pb.NewTestServiceServer(0, newTestService1()),
					pb.NewTestServiceServer(1, newTestService2()),
				)),
			), nil
		}).
		Transport(rsocket.TCPServer().SetAddr(":7878").Build()).
		Serve(ctx)
	if err != nil {
		return err
	}

	return nil
}

// ----

type TestService1 struct {
}

func newTestService1() *TestService1 {
	return &TestService1{}
}

func (s *TestService1) Test(ctx context.Context, req *pb1.TestRequest) (*pb2.TestResponse, error) {
	return &pb2.TestResponse{
		Text: strings.ToUpper(req.Text),
	}, nil
}

// ----

type TestService2 struct {
}

func newTestService2() *TestService2 {
	return &TestService2{}
}

func (s *TestService2) Test(ctx context.Context, req *pb1.TestRequest) (*pb2.TestResponse, error) {
	return &pb2.TestResponse{
		Text: strings.ToLower(req.Text),
	}, nil
}
