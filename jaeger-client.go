package main

import (
	"context"
	"flag"
	"io"
	"time"

	"github.com/golang/glog"
	"github.com/lucasponce/jaeger-proto-client/model"
	"google.golang.org/grpc"
)

func main() {
	flag.Set("logtostderr", "true")
	flag.Parse()

	jaegerAddress := "localhost:16685"

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock())
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(jaegerAddress, opts...)
	if err != nil {
		glog.Fatalf("[%s] failed to open: %v", jaegerAddress, err)
	}
	defer conn.Close()

	ctx := context.Background()
	grpcClient := model.NewQueryServiceClient(conn)
	ctx, cancel := context.WithTimeout(ctx, 4*time.Second)
	defer cancel()

	sTraceId := "2400171a1ac3b1cffde1fd36fd401ad0"
	traceId, err := model.TraceIDFromString(sTraceId)
	if err != nil {
		glog.Errorf("[%s] failed to parse TraceId: %v", sTraceId, err)
	}
	bTraceId := make([]byte, 16)
	_, err = traceId.MarshalTo(bTraceId)
	if err != nil {
		glog.Errorf("[%s] failed to marshall TraceId: %v", sTraceId, err)
	}
	stream, err := grpcClient.GetTrace(ctx, &model.GetTraceRequest{
		TraceId: bTraceId,
	})
	if err != nil {
		glog.Errorf("[%s] failed to FindTraces: %v", jaegerAddress, err)
	}

	for received, err := stream.Recv(); err != io.EOF; received, err = stream.Recv() {
		if err != nil {
			glog.Errorf("[%s] failed to process span: %v", jaegerAddress, err)
			break
		}
		for i := range received.Spans {
			span := received.Spans[i]
			traceId := model.TraceID{}
			traceId.Unmarshal(span.TraceId)
			spanId, err := model.SpanIDFromBytes(span.SpanId)
			if err != nil {
				glog.Errorf("[%s] failed to process spanId: %v", jaegerAddress, err)
				break
			}
			glog.Infof("TraceId: [%s]", traceId)
			glog.Infof("SpanId: [%s]", spanId)
		}

	}
}
