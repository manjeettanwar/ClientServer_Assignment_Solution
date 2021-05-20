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
	"Trial3",
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

	serviceTypeVal, ok := pb.ServiceType_value[*serviceType]
	if !ok {
		fmt.Printf("Specified Service type %s,is invalid\n", *serviceType)
		return
	}

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

	req := &pb.RegisterServiceRequest{
		Service: &pb.Service{
			Name: *serviceName,
			Type: pb.ServiceType(serviceTypeVal),
			Host: *serviceHost,
			Port: int32(*servicePort),
		}}

	resp, err := client.RegisterService(context.Background(), req)
	if err != nil {
		//fmt.Printf("Service Registration failed : %v ",err)
		fmt.Printf("Service Registration failed: [%v] \n",err)
//		fmt.Println(err)
	}else{
		fmt.Printf("Service with name %s and type %s registered sucessfully, id:%v\n", *serviceName, *serviceType, resp.Id)
	}

}
