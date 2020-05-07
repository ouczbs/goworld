package proto

import (
	protolib "github.com/gogo/protobuf/proto"
	"github.com/ouczbs/goworld/engine/proto/pb"
)

func RecvPbWrapMessage(bytes []byte) * pb.WrapMessage{
	message := &pb.WrapMessage{}
	protolib.Unmarshal(bytes , message)
	return message
}
func RecvPbMessage(bytes []byte , pb protolib.Message){
	protolib.Unmarshal(bytes , pb)
}
