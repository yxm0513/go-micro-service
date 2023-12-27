package topic_test

import (
    "fmt"
    "github.com/go-kit/kit/log"
    "github.com/opentracing/opentracing-go"
    client "github.com/yxm0513/go-micro-service/client/topic"
    p_topic "github.com/yxm0513/go-micro-service/proto/topic"
    "github.com/yxm0513/go-micro-service/service/topic/lib"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "net"
    "testing"
    "time"
)

func runTopicServer(addr string) *grpc.Server {
	service := lib.NewTopicService()
	ctx := context.Background()

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	srv := lib.MakeGRPCServer(ctx, service, opentracing.NoopTracer{}, log.NewNopLogger())
	s := grpc.NewServer()
	p_topic.RegisterTopicServer(s, srv)

	go func() {
		s.Serve(ln)
	}()
	time.Sleep(time.Second)
	return s
}

func TestNewTopicClient(t *testing.T) {
	s := runTopicServer(":8003")
	defer s.GracefulStop()
	conn, err := grpc.Dial(":8003", grpc.WithInsecure(), grpc.WithTimeout(time.Second))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	service := client.NewTopicClient(conn, opentracing.NoopTracer{}, log.NewNopLogger())
	req := &p_topic.GetTopicRequest{
		TopicId: 123,
	}
	resp, err := service.GetTopic(context.Background(), req)
	if err != nil {
		fmt.Println(resp, err)
	}
}
