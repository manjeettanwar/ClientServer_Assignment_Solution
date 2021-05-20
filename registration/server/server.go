/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"log"
	"net"
	"fmt"
//	"math/rand"
	"errors"
	//status "google.golang.org/grpc/status"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/registration/client/src/registry"
)

const (
	port = "0.0.0.0:10080"
)

type ServiceNode struct {
	value *pb.Service
	next  *ServiceNode
}

type RegisterServiceGrpcImpl struct {
	pb.UnimplementedServiceRegistrationServer
	m map[int32]*ServiceNode
	servicecount int
}


//NewRegisterServiceGrpcImpl returns the pointer to the implementation.
func NewRegisterServiceGrpcImpl() *RegisterServiceGrpcImpl {
	return &RegisterServiceGrpcImpl{
		m: map[int32]*ServiceNode{
			0: nil,
			1: nil,
			2: nil,
			10: nil,
			11: nil,
		},
	}
}


//RegisterService

func (server *RegisterServiceGrpcImpl)  RegisterService(ctx context.Context,in *pb.RegisterServiceRequest) (out *pb.RegisterServiceResponse,err error) {
	if _, found := pb.ServiceType_value[in.Service.Type.String()]; !found {
		fmt.Printf("Specified Service type %s,is invalid\n",in.Service.Type.String() )
	}
	fmt.Printf("RegisterService : Register Service Request :ServiceType[%s],  ServiceName[%s]\n", in.Service.Type, in.Service.Name)

	ServiceTypeNumber32 := int32(in.Service.Type.Number())
	newNode := ServiceNode{}
	newNode.value = in.Service
	newNode.next = nil
	if (ServiceTypeNumber32 < 3 ){
		if (server.m[ServiceTypeNumber32] == nil){
			server.m[ServiceTypeNumber32] = &newNode
			server.servicecount++
		} else{
			//newNode = nil
			out = &pb.RegisterServiceResponse{
				Id: out.GetId(),
			}
			fmt.Printf("RegisterService : Can't ad more than 1 service for [%s]\n\n",in.Service.Type)
//			err=fmt.Errorf("%v", "Service Already Registered")
			return out, errors.New("Service Already Registered")
		}
	} else {
		if (server.m[ServiceTypeNumber32] == nil){
			server.m[ServiceTypeNumber32] = &newNode
			server.servicecount++
		}else {
			var temp = server.m[ServiceTypeNumber32]
			var pre = temp
			for temp != nil {
				if (temp.value.Name == newNode.value.Name) {
					//newNode = nil
					out = &pb.RegisterServiceResponse{
						Id: out.GetId(),
					}
//					fmt.Println("Already found name for MSVC_1 or MSVC_2 ")
					fmt.Printf("RegisterService : Service Registeration failed :[%s]:[%s] : Already Exists\n\n ", in.Service.Type, in.Service.Name)
//					err=fmt.Errorf("%v", "Service Already Registered")
					return out, errors.New("Service Already Registered")
				}
				pre=temp
				temp=temp.next
			}
			pre.next = &newNode
			server.servicecount++
		}
	}
	fmt.Printf("RegisterService : Service Registered Successfully :[%s]:[%s]\n\n",in.Service.Type, in.Service.Name)
	out = &pb.RegisterServiceResponse{
		Id: out.GetId(),
	}

	return out, nil
}


//RemoveService

func (server *RegisterServiceGrpcImpl) RemoveService(ctx context.Context,in *pb.RemoveServiceRequest) (out *pb.RemoveServiceResponse,err error) {
	//return nil, status.Errorf(codes.Unimplemented, "method RemoveService not implemented")

	if _, found := pb.ServiceType_value[in.Type.String()]; !found {
		fmt.Printf("Remove Specified Service type %s,is invalid \n",in.Type.String() )
	}
	fmt.Printf("RemoveServiceRequest to remove : ServiceType[%s],  ServiceName[%s] \n", in.Type, in.Name)

	ServiceTypeNumber32 := int32(in.Type.Number())
	if (ServiceTypeNumber32 < 3 ){
		if (server.m[ServiceTypeNumber32] == nil){
		} else{
			server.m[ServiceTypeNumber32] = nil
			fmt.Printf("RemoveService: Service Removed [%s]:[%s]   \n\n", in.Type, in.Name)
			server.servicecount--
			out = &pb.RemoveServiceResponse{
			}
			return out , nil
		}
	} else {
		if (server.m[ServiceTypeNumber32] == nil){
		}else {
			var temp = server.m[ServiceTypeNumber32]
			var pre = temp
			for temp != nil {
				if (temp.value.Name == in.Name ) {
					if( pre == temp ){
						server.m[ServiceTypeNumber32]=nil
						temp = nil
						pre = nil
					}else{
						pre.next = temp.next
						temp = nil
					}
					fmt.Printf("RemoveService: Service Removed [%s]:[%s] \n\n", in.Type, in.Name)
					server.servicecount--
//					err=fmt.Errorf("%v", "Service Already Registered")
					out = &pb.RemoveServiceResponse{
					}
					return out , nil
				}
				pre=temp
				temp=temp.next
			}
		}
	}
	out = &pb.RemoveServiceResponse{
	}
	fmt.Printf("RemoveService: Requested service doesnt exists [%s]:[%s] \n\n", in.Type,  in.Name)
	err=errors.New("Requested Service doesnt exist")

	return out, err
}

// ListService

func (server *RegisterServiceGrpcImpl) ListServices(ctx context.Context,in *pb.ListServicesRequest) (out *pb.ListServicesResponse,err error) {
	//        return nil, status.Errorf(codes.Unimplemented, "method ListServices not implemented")
	serviceslist := make([]*pb.Service,server.servicecount)
	fmt.Println("Total no. of registered services  : ",server.servicecount)
	var i int=0
	for k := range server.m {
//		fmt.Printf("key[%s] value[%s]\n", k, m[k])
		if (server.m[k] != nil){
			var temp = server.m[k]
			for temp != nil {
				serviceslist[i] = temp.value
				i++
				temp=temp.next
			}
		}

	}

	out = &pb.ListServicesResponse{
		Servicea: serviceslist,
	}

        fmt.Printf("ListService: Shared service list \n\n")
        //err=errors.New("Requested Service doesnt exist")
        err=nil

        return out, err
}




//main

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}else { fmt.Println("Started listening on port : ",port)}
	s := grpc.NewServer()
	registerServiceImpl := NewRegisterServiceGrpcImpl()
	pb.RegisterServiceRegistrationServer(s, registerServiceImpl)
//	pb.RegisterServiceRegistrationServer(s, &RegisterServiceGrpcImpl{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}else { fmt.Println(" Success to serve")}
}
