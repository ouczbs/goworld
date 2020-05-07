package proto

import (
	"fmt"

	"google.golang.org/protobuf/reflect/protoregistry"
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {
	//MessageType
	//protoregistry.GlobalTypes.FindMessageByName
	message,err := protoregistry.GlobalTypes.FindMessageByName("pb.SET_GATE_ID")
	if err != nil {
		return
	}
	pb2 := message.New()
	//eface := *(*emptyInterface)(unsafe.Pointer(&i))
	// t :=toType
	/*
		mi := &MessageInfo{
			Desc:          legacyLoadMessageDesc(t, name),
			GoReflectType: t,
		}
	*/
	pb3 := pb2.Interface()
	//pbn := pb3.(* pb.SET_GATE_ID)
	pb4 := pb3.ProtoReflect()
	pb4.Interface()
	//pb5 := pb4.(* protolib.Message)
	a := reflect.TypeOf(pb4)
	fmt.Println(pb2 ,pb3 ,pb4 ,a)
}