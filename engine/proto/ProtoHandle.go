package proto

import "github.com/ouczbs/goworld/engine/proto/pb"

type DownHandles struct {
	handles map[uint16]func()
}
var moduleDownHandles = &DownHandles{}

func RegisterDownHandle(cmd pb.CommandList ,handle func()){
	moduleDownHandles.handles[uint16(cmd)] = handle
}
func unRegisterDownHandle(cmd pb.CommandList){
	moduleDownHandles.handles[uint16(cmd)] = nil
}