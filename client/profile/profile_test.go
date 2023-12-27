package profile_test

import (
    "fmt"
    "github.com/go-kit/kit/log"
    "github.com/opentracing/opentracing-go"
    client "github.com/yxm0513/go-micro-service/client/profile"
    p_profile "github.com/yxm0513/go-micro-service/proto/profile"
    "github.com/yxm0513/go-micro-service/service/profile/lib"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "net"
    "testing"
    "time"
)

func runProfileServer(addr string) *grpc.Server {
	service := lib.NewProfileService()
	ctx := context.Background()

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	srv := lib.MakeGRPCServer(ctx, service, opentracing.NoopTracer{}, log.NewNopLogger())
	s := grpc.NewServer()
	p_profile.RegisterProfileServer(s, srv)

	go func() {
		s.Serve(ln)
	}()
	time.Sleep(time.Second)
	return s
}

func TestNewProfileClient(t *testing.T) {
	s := runProfileServer(":8002")
	defer s.GracefulStop()
	conn, err := grpc.Dial(":8002", grpc.WithInsecure(), grpc.WithTimeout(time.Second))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	service := client.NewProfileClient(conn, opentracing.NoopTracer{}, log.NewNopLogger())
	req := &p_profile.GetProfileRequest{
		UserId: 123,
	}
	resp, err := service.GetProfile(context.Background(), req)
	if err != nil {
		fmt.Println(resp, err)
	}
}
