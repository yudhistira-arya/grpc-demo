package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"sync"
	"testing"
	"time"
	meteorite "yudhistira.dev/demo/meteorite/grpc/client/yudhistira.dev/demo/meteorite/grpc/api"
)

var result *meteorite.MeteoriteLandingList

func Benchmark1GrpcRequest(b *testing.B) {
	var conn = CreateConnection()
	defer conn.Close()

	client := meteorite.NewMeteoriteLandingsServiceClient(conn)

	var response *meteorite.MeteoriteLandingList
	var err error
	for n := 0; n < b.N; n++ {
		ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
		defer cancel()

		response, err  = client.GetMeteorite(ctx, &empty.Empty{})
		handleError(err)
	}
	result = response
}

func Benchmark10GrpcRequest(b *testing.B) {
	var conn = CreateConnection()
	defer conn.Close()

	concurrentRequest := 10

	client := meteorite.NewMeteoriteLandingsServiceClient(conn)
	var response *meteorite.MeteoriteLandingList
	var err error
	for n := 0; n < b.N; n++ {
		var wg sync.WaitGroup
		wg.Add(concurrentRequest)
		for c := 0; c < concurrentRequest; c++ {
			go func() {
				ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
				defer cancel()

				defer wg.Done()
				response, err  = client.GetMeteorite(ctx, &empty.Empty{})
				handleError(err)
			}()
		}
		wg.Wait()
	}
	result = response
}

func Benchmark100GrpcRequest(b *testing.B) {
	var conn1 = CreateConnection()
	defer conn1.Close()

	concurrentRequest := 100
	client1 := meteorite.NewMeteoriteLandingsServiceClient(conn1)

	var response *meteorite.MeteoriteLandingList
	var err error
	for n := 0; n < b.N; n++ {
		var wg sync.WaitGroup
		wg.Add(concurrentRequest)
		for c := 0; c < concurrentRequest; c++ {
			go func() {
				ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
				defer cancel()
				defer wg.Done()

				response, err  = client1.GetMeteorite(ctx, &empty.Empty{})
				handleError(err)
			}()
		}
		wg.Wait()
	}
	result = response
}

