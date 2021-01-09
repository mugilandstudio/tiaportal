package main

import (
	"fmt"

	pb "github.com/mugilandstudio/tiaportal/proto"
)

func main() {
	req := &pb.EchoRequest{}
	fmt.Println(req)
}
