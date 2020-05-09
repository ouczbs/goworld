package proto

import (
	"github.com/ouczbs/goworld/engine/proto/pb"
	protolib "google.golang.org/protobuf/proto"
)

func RecvPbWrapMessage(bytes []byte) * pb.WrapMessage{
	message := &pb.WrapMessage{}
	protolib.Unmarshal(bytes , message)
	return message
}
func RecvPbMessage(bytes []byte , pb protolib.Message){
	protolib.Unmarshal(bytes , pb)
}
