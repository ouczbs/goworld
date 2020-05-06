package up

import (
	protolib "github.com/gogo/protobuf/proto"
	"github.com/ouczbs/goworld/engine/proto/pb"
)

func SendSetGameID(gwc * proto.GoWorldConnection ,data * pb.SET_GAME_ID) error {
	packet := gwc.PacketConn.NewPacket()
	//packet.AppendUint16(&pb.SET_GAME_ID)
	// put all entity IDs to the packet
	packet.AppendUint32(protolib.Marshal(data))
	return gwc.SendPacketRelease(packet)
}