package main

import (
	"GoLearn/save_the_world_with_go/14_protocolbuffers/pb/example_05/user"
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func main() {
    info := anypb.Any{Value: []byte(`John rules`), TypeUrl: "urltype"}
    
	userA := user.User{UserId: "John", Email: "john@gmail.com", Info: []*anypb.Any{&info}}
    
    fmt.Println("To encode:", userA.String())

    encoded, err := proto.Marshal(&userA)
    if err != nil {
        panic(err)
    }
    recovered := user.User{}
    err = proto.Unmarshal(encoded, &recovered)
    fmt.Println("Recovered:", recovered.String())
}

