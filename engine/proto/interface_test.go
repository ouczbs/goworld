package proto

import (
	"fmt"
	"github.com/ouczbs/goworld/engine/proto/pb"
	protolib "google.golang.org/protobuf/proto"
	"reflect"
	"testing"
)
//go get github.com/gogo/protobuf@master
func myfunc(v interface{}){
	res_type := reflect.TypeOf(v)
	val := reflect.ValueOf(v)
	fmt.Println(res_type) //string
	fmt.Println(v)
	fmt.Println(val)
}
func myfunc1(pb1 protolib.Message){
	fmt.Println(pb1 ,"myfuc1") //string
	pb2 := pb1.(*pb.SET_GATE_ID)
	pb2.GateId = 3
}
func myfunc2(pb1 * pb.SET_GATE_ID){
	fmt.Println(&pb1,"myfuc2") //string
	fmt.Println(pb1,"myfuc2") //string
	pb1.GateId = 1
}
func mymessage(bytes []byte, pb protolib.Message){
	protolib.Unmarshal(bytes , pb)
}
func call1(){
	pb1 := &pb.SET_GATE_ID{GateId: 2}
	fmt.Println(&pb1,"call 1") //string
	myfunc1(pb1)
	myfunc2(pb1)
	fmt.Println(pb1,"call 2") //string
}
func call2(){
	pb1 := &pb.SET_GATE_ID{GateId: 1}
	fmt.Println(&pb1,"call 1") //string
	out , err := protolib.Marshal(pb1)
	fmt.Println(out,err)
	pb2 := &pb.SET_GATE_ID{}
	mymessage(out,pb2)
	//protolib.Unmarshal(out , pb2)
	fmt.Println(&pb2,"call 2")
	fmt.Println(pb2.GateId,"call 2")
}
func wrapmessage(bytes []byte) * pb.WrapMessage{
	message := &pb.WrapMessage{}
	protolib.Unmarshal(nil , message)
	return message
}
func message(bytes []byte , pb protolib.Message){
	protolib.Unmarshal(bytes , pb)
}
func TestInterface(t *testing.T)  {
	pb1 := &pb.SET_GATE_ID{GateId: 1}
	out , err := protolib.Marshal(pb1)
	fmt.Println(out ,err)
	pb2 := &pb.WrapMessage{
		Cmd:pb.CommandList_MT_SET_GATE_ID,
	}
	fmt.Println(pb2.Content)
	protolib.Unmarshal(pb2.Content , pb1)
	pb2.Content = out
	bytes, err := protolib.Marshal(pb2)
	fmt.Println(bytes ,err)
	pb3 := wrapmessage(bytes)
	fmt.Println(pb3)
	//fmt.Println(int(pb3.Cmd) , pb3.Cmd)
}
