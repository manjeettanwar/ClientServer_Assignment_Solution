## Programming Assignment
Along with this README we have provided a registration.proto file.  This is a gRPC protobuf definition os a server you are to implement.   You will also be required to implement a simple client demonstrating the functions of the server.

gRPC supports many programming language, https://grpc.io/docs/languages/,  you may choose to use any language that you wish and are most comfortable with.  WE have provide a sample/starter client in Go, which you can extend or write your own in a different programming language.

You should not need to modify the proto file for this assignment, you will need to generate the protobuf files for the language you will be using using protoc.

### Service Registration
The server is a very simple ServiceRegistration server.  Some services(s) can connect to the server and registers themselves as a service of some type.  Client applications can request the registered services and use this information to connect to the service.

For this exercise you may keep all registration in memory in a data structure if your choosing.  You do not need to worry about persisting the data in this exercise.

For registration you can use the provide Go sample, you should register (and/or remove) several services and then have a client application that retrieves the set of registered services.


### Requirements
1. The service types SVC_1, SVC_2, SVC_3 are considered singleton services.  Meaning at any tines there can be only a single service registered for each of these types
2. The service types MSVC_1 and MSVC_2 can have any number of these services registered.
3. The combination of service name and type are should be unique with in the service.  For example there could be 2 services registered with name foo as long as they are of different types.
4. If any of the above conditions are violated the client should receive an error response from the server
5. Multiple clients should be able to connect and interact with the server.
