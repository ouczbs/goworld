package proto

import (
	"github.com/ouczbs/goworld/engine/netutil"
	protolib "google.golang.org/protobuf/proto"
)
//length with 2 bytes ,not exit on down
func appendReservedText(packet * netutil.Packet){
	packet.AppendUint16(1)
}
func SendPackageWithPbMessage(gwc *GoWorldConnection,pb protolib.Message)error{
	packet := gwc.NewPacket()
	appendReservedText(packet)
	out , err := protolib.Marshal(pb)
	if err != nil{
		return err
	}
	packet.AppendBytes(out)
	return gwc.SendPacketRelease(packet)
}