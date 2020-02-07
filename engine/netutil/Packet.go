package netutil

import (
	"github.com/xiaonanln/goworld/engine/common"
	"github.com/xiaonanln/goworld/engine/gwlog"
	"github.com/xiaonanln/pktconn"
)

type GWPacket struct {
	*pktconn.Packet
}

func (packet GWPacket) WriteEntityID(id common.EntityID) {
	packet.WriteBytes([]byte(id))
}

// ReadEntityID reads one EntityID from the beginning of unread  payload
func (packet GWPacket) ReadEntityID() common.EntityID {
	return common.EntityID(packet.ReadBytes(common.ENTITYID_LENGTH))
}

func (packet GWPacket) WriteClientID(id common.ClientID) {
	packet.WriteBytes([]byte(id))
}

// ReadClientID reads one ClientID from the beginning of unread  payload
func (packet GWPacket) ReadClientID() common.ClientID {
	return common.ClientID(packet.ReadBytes(common.CLIENTID_LENGTH))
}

func (packet GWPacket) WriteVarStr(s string) {
	packet.WriteVarBytesI([]byte(s))
}

// ReadVarStr reads a varsize string from the beginning of unread  payload
func (packet GWPacket) ReadVarStr() string {
	b := packet.ReadVarBytes()
	return string(b)
}

func (packet GWPacket) WriteData(msg interface{}) {
	dataBytes, err := MSG_PACKER.PackMsg(msg, nil)
	if err != nil {
		gwlog.Panic(err)
	}

	packet.WriteVarBytes(dataBytes)
}

func (packet GWPacket) WriteVarBytes(b []byte) {
	packet.WriteVarBytesI(b)
}

// ReadVarBytes reads a varsize slice of bytes from the beginning of unread  payload
func (packet GWPacket) ReadVarBytes() []byte {
	blen := packet.ReadUint32()
	return packet.ReadBytes(int(blen))
}

func (packet GWPacket) WriteArgs(args []interface{}) {
	argCount := uint16(len(args))
	packet.WriteUint16(argCount)

	for _, arg := range args {
		packet.WriteData(arg)
	}
}

func (packet GWPacket) WriteMapStringString(m map[string]string) {
	packet.WriteUint32(uint32(len(m)))
	for k, v := range m {
		packet.WriteVarStr(k)
		packet.WriteVarStr(v)
	}
}

func NewPacket() GWPacket {
	return GWPacket{Packet: pktconn.NewPacket()}
}
