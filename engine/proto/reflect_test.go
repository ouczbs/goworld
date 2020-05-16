package proto

import (
	"fmt"
	"github.com/ouczbs/goworld/engine/proto/pb"
	protolib "google.golang.org/protobuf/proto"
	"strings"

	//"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"reflect"
	"testing"
)
func TestReflect2(t *testing.T) {
	//MessageType
	//protoregistry.GlobalTypes.FindMessageByName
	messagetowrap()
}
func messagetowrap() error{
	message,err := protoregistry.GlobalTypes.FindMessageByName("pb.SET_GATE_ID")
	if err != nil {
		return err
	}
	//ark := message.New().(*pb.SET_GATE_ID)
	//fmt.Println(ark)
	inte := message.New().Interface()
	reltype := reflect.ValueOf(inte)
	ack,ok := inte.(*pb.SET_GATE_ID)
	if !ok {
		fmt.Println("error")
		return err
	}
	fmt.Println(ack ,ok)
	ack.GateId =2
	buf,err := protolib.Marshal(ack)
	test := protolib.Unmarshal(buf , inte)
	res := pb.CommandList_name[2]
	res = strings.Replace(res, "MT_" , "pb." , 1)
	pb2,ok2 := inte.(*pb.SET_GATE_ID)
	fmt.Println(pb2 , ok2)
	fmt.Println(ack ,	res , test, inte , reflect.TypeOf(inte))
	fmt.Println(reltype.Elem(),reltype.Kind(),reltype.String())
	fmt.Println(message)
	return nil
}