package main

import (
	"bytes"
	"fmt"
	"time"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"istio.io/file-envoy/adsc"
)

const (
	typePrefix = "type.googleapis.com/envoy.api.v2."

	// Constants used for XDS

	// ClusterType is used for cluster discovery. Typically first request received
	clusterType = typePrefix + "Cluster"
	// EndpointType is used for EDS and ADS endpoint discovery. Typically second request.
	endpointType = typePrefix + "ClusterLoadAssignment"
	// ListenerType is sent after clusters and endpoints.
	listenerType = typePrefix + "Listener"
	// RouteType is sent after listeners.
	routeType = typePrefix + "RouteConfiguration"
)

func main() {
	grpc := fmt.Sprintf("localhost:15010")
	adsc, err := adsc.Dial(grpc, "", &adsc.Config{
		IP:        "10.11.0.1",
		Namespace: "envoy",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Waiting for updates")
	adsc.Watch()
	for {
		select {
		case r := <-adsc.RawUpdates:
			fmt.Printf("Got update: %v\n", r.TypeUrl)
		case r := <-adsc.Updates:
			fmt.Printf("Got small update: %v\n", r)
		case <-time.After(time.Second * 10):
			fmt.Println("Done")
			return
		}
	}
}

func MarshallJson(w proto.Message) ([]byte, error) {
	buffer := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{}).Marshal(buffer, w)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
