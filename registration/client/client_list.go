package main

import (
	"context"
	"flag"
	"fmt"
	pb "google.golang.org/grpc/examples/registration/client/src/registry"

	"google.golang.org/grpc"
)

var serviceName = flag.String(
	"serviceName",
	"Trial2",
	"name if the service to be registered",
)

var serviceType = flag.String(
	"serviceType",
	"MSVC_2",
	"type of the service being registered.",
)

var serviceHost = flag.String(
	"serviceHost",
	"localhost",
	"hostname or ip of the service being registered.",
)

var servicePort = flag.Int(
	"servicePort",
	111,
	"port of service being registered.",
)

var grpcServerHost = flag.String(
	"grpcServerHost",
	"localhost",
	"hostname or ip of gRPC service to use for registration.",
)

var grpcServerPort = flag.Int(
	"grpcServerPort",
	10080,
	"port of gRPC service to use for registration.",
)

func main() {
	flag.Parse()

	if len(*serviceName) == 0 {
		fmt.Printf("Service name is missing\n")
		return
	}

	if len(*serviceType) == 0 {
		fmt.Printf("Service type is missing\n")
		return
	}

	if len(*serviceType) == 0 {
		fmt.Printf("Service type is missing\n")
		return
	}

	//serviceTypeVal, ok := pb.ServiceType_value[*serviceType]
	//if !ok {
		fmt.Printf("Specified Service type %s,is invalid\n", *serviceType)
	//}

	if *servicePort <= 0 {
		fmt.Printf("Service port is missing\n")
		return
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", *grpcServerHost, *grpcServerPort), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	client := pb.NewServiceRegistrationClient(conn)

	req := &pb.ListServicesRequest{
		}

	resp, err := client.ListServices(context.Background(), req)
	if err != nil {
		//fmt.Printf("Service Registration failed : %v ",err)
		fmt.Printf("Service List Failed : [%v]\n\n",err)
	} else{
		fmt.Printf("ServiceList is shared successfully \n\n")
		for i := 0; i < len(resp.Servicea); i++ {
			fmt.Printf("[%s][%s]\n",resp.Servicea[i].Name,resp.Servicea[i].Type.String())
		}
	}

}
