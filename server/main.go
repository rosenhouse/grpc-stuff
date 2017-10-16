package main

import (
	"log"
	"net"
	"time"

	"github.com/rosenhouse/grpc-stuff/policy"
	"google.golang.org/grpc"
)

type internalAPIServer struct{}

func (s *internalAPIServer) WatchPolicies(filter *policy.PoliciesFilter, outStream policy.InternalAPI_WatchPoliciesServer) error {
	log.Printf("got filter: %v\n", filter)
	outStream.Send(testData["bulkUpdate0"])
	time.Sleep(1 * time.Second)

	outStream.Send(testData["eventUpdate1"])
	time.Sleep(1 * time.Second)
	outStream.Send(testData["eventUpdate2"])
	time.Sleep(1 * time.Second)
	outStream.Send(testData["bulkUpdate3"])
	return nil
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:5000")
	if err != nil {
		log.Fatalf("listen: %s", err)
	}

	apiServer := &internalAPIServer{}

	grpcServer := grpc.NewServer()
	policy.RegisterInternalAPIServer(grpcServer, apiServer)
	grpcServer.Serve(listener)
}

var testData = map[string]*policy.PoliciesUpdate{
	"bulkUpdate0": &policy.PoliciesUpdate{
		Method: policy.PoliciesUpdate_REPLACE,
		Policy: []*policy.Policy{
			{
				Source:      &policy.App{Id: "some-source0"},
				Destination: &policy.App{Id: "some-destination0"},
				Protocol:    "tcp",
				PortRange:   "8001",
			},
			{
				Source:      &policy.App{Id: "some-source1"},
				Destination: &policy.App{Id: "some-destination1"},
				Protocol:    "tcp",
				PortRange:   "8002",
			},
			{
				Source:      &policy.App{Id: "some-source3"},
				Destination: &policy.App{Id: "some-destination3"},
				Protocol:    "udp",
				PortRange:   "8003-8030",
			},
		},
	},
	"eventUpdate1": &policy.PoliciesUpdate{
		Method: policy.PoliciesUpdate_UNION,
		Policy: []*policy.Policy{
			{
				Source:      &policy.App{Id: "some-source4"},
				Destination: &policy.App{Id: "some-destination4"},
				Protocol:    "tcp",
				PortRange:   "9004,9044",
			},
			{
				Source:      &policy.App{Id: "some-source5"},
				Destination: &policy.App{Id: "some-destination5"},
				Protocol:    "tcp",
				PortRange:   "8005",
			},
		},
	},
	"eventUpdate2": &policy.PoliciesUpdate{
		Method: policy.PoliciesUpdate_UNION,
		Policy: []*policy.Policy{
			{
				Source:      &policy.App{Id: "some-source6"},
				Destination: &policy.App{Id: "some-destination6"},
				Protocol:    "tcp",
				PortRange:   "9006,9066",
			},
			{
				Source:      &policy.App{Id: "some-source5"},
				Destination: &policy.App{Id: "some-destination5"},
				Protocol:    "tcp",
				PortRange:   "8005",
			},
		},
	},
	"bulkUpdate3": &policy.PoliciesUpdate{
		Method: policy.PoliciesUpdate_REPLACE,
		Policy: []*policy.Policy{
			{
				Source:      &policy.App{Id: "some-source0"},
				Destination: &policy.App{Id: "some-destination0"},
				Protocol:    "tcp",
				PortRange:   "8001",
			},
			{
				Source:      &policy.App{Id: "some-source1"},
				Destination: &policy.App{Id: "some-destination1"},
				Protocol:    "tcp",
				PortRange:   "8002",
			},
			{
				Source:      &policy.App{Id: "some-source4"},
				Destination: &policy.App{Id: "some-destination4"},
				Protocol:    "tcp",
				PortRange:   "9004,9044",
			},
			{
				Source:      &policy.App{Id: "some-source6"},
				Destination: &policy.App{Id: "some-destination6"},
				Protocol:    "tcp",
				PortRange:   "9006,9066",
			},
		},
	},
}
