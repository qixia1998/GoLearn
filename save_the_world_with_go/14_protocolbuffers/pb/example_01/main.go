package main

import (
	"fmt"
	"GoLearn/save_the_world_with_go/14_protocolbuffers/pb/example_01/user"
	"google.golang.org/protobuf/proto"
)

func main() {
	u := user.User{Email: "john@gmail.com", UserId: "John"}

	fmt.Println("To encode:", u.String())
	encoded, err := proto.Marshal(&u)
	if err != nil {
		panic(err)
	}

	v := user.User{}
	err = proto.Unmarshal(encoded, &v)
	if err != nil {
		panic(err)
	}
	fmt.Println("Recovered:", v.String())
}