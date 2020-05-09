package proto

import (
	"github.com/ouczbs/goworld/engine/proto/pb"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"strings"
)

var pbMessageHandles map[uint32]func(* protoreflect.ProtoMessage , uint32)
var pbMessageTypes map[uint32]protoreflect.MessageType

func RegisterDownHandle(cmd pb.CommandList ,handle func(* protoreflect.ProtoMessage, uint32)){
	pbMessageHandles[uint32(cmd)] = handle
}
func unRegisterDownHandle(cmd pb.CommandList){
	pbMessageHandles[uint32(cmd)] = nil
}
func newPbMessage(cmd pb.CommandList) (protoreflect.ProtoMessage , error){
	messageType := pbMessageTypes[uint32(cmd)]
	if messageType != nil {
		return messageType.New().Interface() ,nil
	}
	pbName := pb.CommandList_name[int32(cmd)]
	pbName = strings.Replace(pbName, "MT_" , "pb." , 1)
	messageType,err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(pbName))
	if err != nil {
		return nil , err
	}
	pbMessageTypes[uint32(cmd)] = messageType
	return messageType.New().Interface() ,err
}
func handlePbMessage(){

}