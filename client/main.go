package main

import (
	"context"
	"io"
	"log"

	"github.com/rosenhouse/grpc-stuff/policy"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dial: %s", err)
	}
	defer conn.Close()
	client := policy.NewInternalAPIClient(conn)
	policiesFilter := &policy.PoliciesFilter{
		App: []*policy.App{
			{Id: "foo"},
			{Id: "bar"},
		},
	}
	updateStream, err := client.WatchPolicies(context.Background(), policiesFilter)
	if err != nil {
		log.Fatalf("watch: %s", err)
	}
	log.Println("streaming...")
	for {
		update, err := updateStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("recv: %s", err)
		}
		log.Println(update)
	}
	log.Println("done")
}
