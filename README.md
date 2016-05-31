#Project#
Microservice samples

##API Interface##
Use protoc and protoc-gen-go to generate 

Go-micro uses code generation to provide client stub methods to reduce boiler plate code much like gRPC.

~~~~ 
go get github.com/micro/protobuf/{proto,protoc-gen-go}
protoc --go_out=plugins=micro:. greeter.proto 
~~~~

##Credit##
Asim Aslam & micro.mu

