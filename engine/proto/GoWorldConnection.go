package proto

import (
	"context"
	"github.com/xiaonanln/pktconn"
	"net"

	"github.com/xiaonanln/go-xnsyncutil/xnsyncutil"
	"github.com/xiaonanln/goworld/engine/common"
	"github.com/xiaonanln/goworld/engine/gwlog"
	"github.com/xiaonanln/goworld/engine/netutil"
	"github.com/xiaonanln/goworld/engine/netutil/compress"
)

// GoWorldConnection is the network protocol implementation of GoWorld components (dispatcher, gate, game)
type GoWorldConnection struct {
	packetConn   *pktconn.PacketConn
	closed       xnsyncutil.AtomicBool
	autoFlushing bool
}

// NewGoWorldConnection creates a GoWorldConnection using network connection
func NewGoWorldConnection(conn netutil.Connection, compressConnection bool, compressFormat string) *GoWorldConnection {
	var compressor compress.Compressor
	if compressConnection {
		compressor = compress.NewCompressor(compressFormat)
	}

	_ = compressor // FIXME: compressor not supported in pktconn
	return &GoWorldConnection{
		packetConn: pktconn.NewPacketConn(context.Background(), conn),
	}
}

// SendSetGameID sends MT_SET_GAME_ID message
func (gwc *GoWorldConnection) SendSetGameID(id uint16, isReconnect bool, isRestore bool, isBanBootEntity bool,
	eids []common.EntityID) error {
	packet := netutil.NewPacket()
	packet.WriteUint16(MT_SET_GAME_ID)
	packet.WriteUint16(id)
	packet.WriteBool(isReconnect)
	packet.WriteBool(isRestore)
	packet.WriteBool(isBanBootEntity)
	// put all entity IDs to the packet

	packet.WriteUint32(uint32(len(eids)))
	for _, eid := range eids {
		packet.WriteEntityID(eid)
	}
	return gwc.SendPacketRelease(packet)
}

// SendSetGateID sends MT_SET_GATE_ID message
func (gwc *GoWorldConnection) SendSetGateID(id uint16) error {
	packet := netutil.NewPacket()
	packet.WriteUint16(MT_SET_GATE_ID)
	packet.WriteUint16(id)
	return gwc.SendPacketRelease(packet)
}

// SendNotifyCreateEntity sends MT_NOTIFY_CREATE_ENTITY message
func (gwc *GoWorldConnection) SendNotifyCreateEntity(id common.EntityID) error {
	packet := netutil.NewPacket()
	packet.WriteUint16(MT_NOTIFY_CREATE_ENTITY)
	packet.WriteEntityID(id)
	return gwc.SendPacketRelease(packet)
}

// SendNotifyDestroyEntity sends MT_NOTIFY_DESTROY_ENTITY message
func (gwc *GoWorldConnection) SendNotifyDestroyEntity(id common.EntityID) error {
	packet := netutil.NewPacket()
	packet.WriteUint16(MT_NOTIFY_DESTROY_ENTITY)
	packet.WriteEntityID(id)
	return gwc.SendPacketRelease(packet)
}

// SendNotifyClientConnected sends MT_NOTIFY_CLIENT_CONNECTED message
func (gwc *GoWorldConnection) SendNotifyClientConnected(id common.ClientID, bootEid common.EntityID) error {
	packet := netutil.NewPacket()
	packet.WriteUint16(MT_NOTIFY_CLIENT_CONNECTED)
	packet.WriteClientID(id)
	packet.WriteEntityID(bootEid)
	return gwc.SendPacketRelease(packet)
}

// SendNotifyClientDisconnected sends MT_NOTIFY_CLIENT_DISCONNECTED message
func (gwc *GoWorldConnection) SendNotifyClientDisconnected(id common.ClientID, ownerEntityID common.EntityID) error {
	packet := netutil.NewPacket()
	packet.WriteUint16(MT_NOTIFY_CLIENT_DISCONNECTED)
	packet.WriteEntityID(ownerEntityID)
	packet.WriteClientID(id)
	return gwc.SendPacketRelease(packet)
}

// SendCreateEntitySomewhere sends MT_CREATE_ENTITY_SOMEWHERE message
func (gwc *GoWorldConnection) SendCreateEntitySomewhere(gameid uint16, entityid common.EntityID, typeName string, data map[string]interface{}) error {
	packet := netutil.NewPacket()
	packet.WriteUint16(MT_CREATE_ENTITY_SOMEWHERE)
	packet.WriteUint16(gameid)
	packet.WriteEntityID(entityid)
	packet.WriteVarStr(typeName)
	packet.WriteData(data)
	return gwc.SendPacketRelease(packet)
}

// SendLoadEntitySomewhere sends MT_LOAD_ENTITY_SOMEWHERE message
func (gwc *GoWorldConnection) SendLoadEntitySomewhere(typeName string, entityID common.EntityID, gameid uint16) error {
	packet := netutil.NewPacket()
	packet.WriteUint16(MT_LOAD_ENTITY_SOMEWHERE)
	packet.WriteUint16(gameid)
	packet.WriteEntityID(entityID)
	packet.WriteVarStr(typeName)
	return gwc.SendPacketRelease(packet)
}

// SendSrvdisRegister
func (gwc *GoWorldConnection) SendSrvdisRegister(srvid string, info string, force bool) error {
	packet := netutil.NewPacket()
	packet.WriteUint16(MT_SRVDIS_REGISTER)
	packet.WriteVarStr(srvid)
	packet.WriteVarStr(info)
	packet.WriteBool(force)
	return gwc.SendPacketRelease(packet)
}

// SendCallEntityMethod sends MT_CALL_ENTITY_METHOD message
func (gwc *GoWorldConnection) SendCallEntityMethod(id common.EntityID, method string, args []interface{}) error {
	packet := netutil.NewPacket()
	packet.WriteUint16(MT_CALL_ENTITY_METHOD)
	packet.WriteEntityID(id)
	packet.WriteVarStr(method)
	packet.WriteArgs(args)
	return gwc.SendPacketRelease(packet)
}

// SendCallEntityMethodFromClient sends MT_CALL_ENTITY_METHOD_FROM_CLIENT message
func (gwc *GoWorldConnection) SendCallEntityMethodFromClient(id common.EntityID, method string, args []interface{}) error {
	packet := netutil.NewPacket()
	packet.WriteUint16(MT_CALL_ENTITY_METHOD_FROM_CLIENT)
	packet.WriteEntityID(id)
	packet.WriteVarStr(method)
	packet.WriteArgs(args)
	return gwc.SendPacketRelease(packet)
}

// SendCreateEntityOnClient sends MT_CREATE_ENTITY_ON_CLIENT message
func (gwc *GoWorldConnection) SendCreateEntityOnClient(gameid uint16, clientid common.ClientID, typeName string, entityid common.EntityID,
	isPlayer bool, clientData map[string]interface{}, x, y, z float32, yaw float32) error {
	packet := netutil.NewPacket()
	packet.WriteUint16(MT_CREATE_ENTITY_ON_CLIENT)
	packet.WriteUint16(gameid)
	packet.WriteClientID(clientid)
	packet.WriteBool(isPlayer)
	packet.WriteEntityID(entityid)
	packet.WriteVarStr(typeName)
	packet.WriteFloat32(x)
	packet.WriteFloat32(y)
	packet.WriteFloat32(z)
	packet.WriteFloat32(yaw)
	packet.WriteData(clientData)
	return gwc.SendPacketRelease(packet)
}

// SendSyncPositionYawFromClient sends MT_SYNC_POSITION_YAW_FROM_CLIENT message
func (gwc *GoWorldConnection) SendSyncPositionYawFromClient(entityID common.EntityID, x, y, z float32, yaw float32) error {
	packet := netutil.NewPacket()
	packet.WriteUint16(MT_SYNC_POSITION_YAW_FROM_CLIENT)
	packet.WriteEntityID(entityID)
	packet.WriteFloat32(x)
	packet.WriteFloat32(y)
	packet.WriteFloat32(z)
	packet.WriteFloat32(yaw)
	return gwc.SendPacketRelease(packet)
}

//func (gwc *GoWorldConnection) SendSetClientClientID(clientid common.ClientID) error {
//	packet := netutil.NewPacket()
//	packet.WriteUint16(MT_SET_CLIENT_CLIENTID)
//	packet.WriteClientID(clientid)
//	return gwc.SendPacketRelease(packet)
//}

func (gwc *GoWorldConnection) SetHeartbeatFromClient() error {
	packet := netutil.NewPacket()
	packet.WriteUint16(MT_HEARTBEAT_FROM_CLIENT)
	return gwc.SendPacketRelease(packet)

}

// SendDestroyEntityOnClient sends MT_DESTROY_ENTITY_ON_CLIENT message
func (gwc *GoWorldConnection) SendDestroyEntityOnClient(gateid uint16, clientid common.ClientID, typeName string, entityid common.EntityID) error {
	packet := netutil.NewPacket()
	packet.WriteUint16(MT_DESTROY_ENTITY_ON_CLIENT)
	packet.WriteUint16(gateid)
	packet.WriteClientID(clientid)
	packet.WriteVarStr(typeName)
	packet.WriteEntityID(entityid)
	return gwc.SendPacketRelease(packet)
}

// SendNotifyMapAttrChangeOnClient sends MT_NOTIFY_MAP_ATTR_CHANGE_ON_CLIENT message
func (gwc *GoWorldConnection) SendNotifyMapAttrChangeOnClient(gateid uint16, clientid common.ClientID, entityid common.EntityID, path []interface{}, key string, val interface{}) error {
	packet := netutil.NewPacket()
	packet.WriteUint16(MT_NOTIFY_MAP_ATTR_CHANGE_ON_CLIENT)
	packet.WriteUint16(gateid)
	packet.WriteClientID(clientid)
	packet.WriteEntityID(entityid)
	packet.WriteData(path)
	packet.WriteVarStr(key)
	packet.WriteData(val)
	return gwc.SendPacketRelease(packet)
}

// SendNotifyMapAttrDelOnClient sends MT_NOTIFY_MAP_ATTR_DEL_ON_CLIENT message
func (gwc *GoWorldConnection) SendNotifyMapAttrDelOnClient(gateid uint16, clientid common.ClientID, entityid common.EntityID, path []interface{}, key string) error {
	packet := netutil.NewPacket()
	packet.WriteUint16(MT_NOTIFY_MAP_ATTR_DEL_ON_CLIENT)
	packet.WriteUint16(gateid)
	packet.WriteClientID(clientid)
	packet.WriteEntityID(entityid)
	packet.WriteData(path)
	packet.WriteVarStr(key)
	return gwc.SendPacketRelease(packet)
}

// SendNotifyMapAttrClearOnClient sends MT_NOTIFY_MAP_ATTR_CLEAR_ON_CLIENT message
func (gwc *GoWorldConnection) SendNotifyMapAttrClearOnClient(gateid uint16, clientid common.ClientID, entityid common.EntityID, path []interface{}) error {
	packet := netutil.NewPacket()
	packet.WriteUint16(MT_NOTIFY_MAP_ATTR_CLEAR_ON_CLIENT)
	packet.WriteUint16(gateid)
	packet.WriteClientID(clientid)
	packet.WriteEntityID(entityid)
	packet.WriteData(path)
	return gwc.SendPacketRelease(packet)
}

// SendNotifyListAttrChangeOnClient sends MT_NOTIFY_LIST_ATTR_CHANGE_ON_CLIENT message
func (gwc *GoWorldConnection) SendNotifyListAttrChangeOnClient(gateid uint16, clientid common.ClientID, entityid common.EntityID, path []interface{}, index uint32, val interface{}) error {
	packet := netutil.NewPacket()
	packet.WriteUint16(MT_NOTIFY_LIST_ATTR_CHANGE_ON_CLIENT)
	packet.WriteUint16(gateid)
	packet.WriteClientID(clientid)
	packet.WriteEntityID(entityid)
	packet.WriteData(path)
	packet.WriteUint32(index)
	packet.WriteData(val)
	return gwc.SendPacketRelease(packet)
}

// SendNotifyListAttrPopOnClient sends MT_NOTIFY_LIST_ATTR_POP_ON_CLIENT message
func (gwc *GoWorldConnection) SendNotifyListAttrPopOnClient(gateid uint16, clientid common.ClientID, entityid common.EntityID, path []interface{}) error {
	packet := netutil.NewPacket()
	packet.WriteUint16(MT_NOTIFY_LIST_ATTR_POP_ON_CLIENT)
	packet.WriteUint16(gateid)
	packet.WriteClientID(clientid)
	packet.WriteEntityID(entityid)
	packet.WriteData(path)
	return gwc.SendPacketRelease(packet)
}

// SendNotifyListAttrAppendOnClient sends MT_NOTIFY_LIST_ATTR_APPEND_ON_CLIENT message
func (gwc *GoWorldConnection) SendNotifyListAttrAppendOnClient(gateid uint16, clientid common.ClientID, entityid common.EntityID, path []interface{}, val interface{}) error {
	packet := netutil.NewPacket()
	packet.WriteUint16(MT_NOTIFY_LIST_ATTR_APPEND_ON_CLIENT)
	packet.WriteUint16(gateid)
	packet.WriteClientID(clientid)
	packet.WriteEntityID(entityid)
	packet.WriteData(path)
	packet.WriteData(val)
	return gwc.SendPacketRelease(packet)
}

// SendCallEntityMethodOnClient sends MT_CALL_ENTITY_METHOD_ON_CLIENT message
func (gwc *GoWorldConnection) SendCallEntityMethodOnClient(gateid uint16, clientid common.ClientID, entityID common.EntityID, method string, args []interface{}) (err error) {
	packet := netutil.NewPacket()
	packet.WriteUint16(MT_CALL_ENTITY_METHOD_ON_CLIENT)
	packet.WriteUint16(gateid)
	packet.WriteClientID(clientid)
	packet.WriteEntityID(entityID)
	packet.WriteVarStr(method)
	packet.WriteArgs(args)
	return gwc.SendPacketRelease(packet)
}

// SendSetClientFilterProp sends MT_SET_CLIENTPROXY_FILTER_PROP message
func (gwc *GoWorldConnection) SendSetClientFilterProp(gateid uint16, clientid common.ClientID, key, val string) (err error) {
	packet := netutil.NewPacket()
	packet.WriteUint16(MT_SET_CLIENTPROXY_FILTER_PROP)
	packet.WriteUint16(gateid)
	packet.WriteClientID(clientid)
	packet.WriteVarStr(key)
	packet.WriteVarStr(val)
	return gwc.SendPacketRelease(packet)
}

// SendClearClientFilterProp sends MT_CLEAR_CLIENTPROXY_FILTER_PROPS message
func (gwc *GoWorldConnection) SendClearClientFilterProp(gateid uint16, clientid common.ClientID) (err error) {
	packet := netutil.NewPacket()
	packet.WriteUint16(MT_CLEAR_CLIENTPROXY_FILTER_PROPS)
	packet.WriteUint16(gateid)
	packet.WriteClientID(clientid)
	return gwc.SendPacketRelease(packet)
}

// SendCallFilterClientProxies sends MT_CALL_FILTERED_CLIENTS message
func AllocCallFilterClientProxiesPacket(op FilterClientsOpType, key, val string, method string, args []interface{}) netutil.GWPacket {
	packet := netutil.NewPacket()
	packet.WriteUint16(MT_CALL_FILTERED_CLIENTS)
	packet.WriteOneByte(byte(op))
	packet.WriteVarStr(key)
	packet.WriteVarStr(val)
	packet.WriteVarStr(method)
	packet.WriteArgs(args)
	return packet
}

func AllocCallNilSpacesPacket(exceptGameID uint16, method string, args []interface{}) netutil.GWPacket {
	// construct one packet for multiple sending
	packet := netutil.NewPacket()
	packet.WriteUint16(MT_CALL_NIL_SPACES)
	packet.WriteUint16(exceptGameID)
	packet.WriteVarStr(method)
	packet.WriteArgs(args)
	return packet
}

func AllocGameLBCInfoPacket(lbcinfo GameLBCInfo) netutil.GWPacket {
	packet := netutil.NewPacket()
	packet.WriteUint16(MT_GAME_LBC_INFO)
	packet.WriteData(lbcinfo)
	return packet
}

// SendQuerySpaceGameIDForMigrate sends MT_QUERY_SPACE_GAMEID_FOR_MIGRATE message
func (gwc *GoWorldConnection) SendQuerySpaceGameIDForMigrate(spaceid common.EntityID, entityid common.EntityID) error {
	packet := netutil.NewPacket()
	packet.WriteUint16(MT_QUERY_SPACE_GAMEID_FOR_MIGRATE)
	packet.WriteEntityID(spaceid)
	packet.WriteEntityID(entityid)
	return gwc.SendPacketRelease(packet)
}

// SendMigrateRequest sends MT_MIGRATE_REQUEST message
func (gwc *GoWorldConnection) SendMigrateRequest(entityID common.EntityID, spaceID common.EntityID, spaceGameID uint16) error {
	packet := netutil.NewPacket()
	packet.WriteUint16(MT_MIGRATE_REQUEST)
	packet.WriteEntityID(entityID)
	packet.WriteEntityID(spaceID)
	packet.WriteUint16(spaceGameID)
	return gwc.SendPacketRelease(packet)
}

// SendCancelMigrate sends MT_CANCEL_MIGRATE message to dispatcher to unblock the entity
func (gwc *GoWorldConnection) SendCancelMigrate(entityid common.EntityID) error {
	packet := netutil.NewPacket()
	packet.WriteUint16(MT_CANCEL_MIGRATE)
	packet.WriteEntityID(entityid)
	return gwc.SendPacketRelease(packet)
}

// SendRealMigrate sends MT_REAL_MIGRATE message
func (gwc *GoWorldConnection) SendRealMigrate(eid common.EntityID, targetGame uint16, data []byte) error {
	packet := netutil.NewPacket()
	packet.WriteUint16(MT_REAL_MIGRATE)
	packet.WriteEntityID(eid)
	packet.WriteUint16(targetGame)
	packet.WriteVarBytes(data)
	return gwc.SendPacketRelease(packet)
}

//// SendStartFreezeGame sends MT_START_FREEZE_GAME message
//func (gwc *GoWorldConnection) SendStartFreezeGame(gameid uint16) error {
//	packet := netutil.NewPacket()
//	packet.WriteUint16(MT_START_FREEZE_GAME)
//	return gwc.SendPacketRelease(packet)
//}

func AllocStartFreezeGamePacket() netutil.GWPacket {
	packet := netutil.NewPacket()
	packet.WriteUint16(MT_START_FREEZE_GAME)
	return packet
}

func MakeNotifyGameConnectedPacket(gameid uint16) netutil.GWPacket {
	pkt := netutil.NewPacket()
	pkt.WriteUint16(MT_NOTIFY_GAME_CONNECTED)
	pkt.WriteUint16(gameid)
	return pkt
}

func MakeNotifyGameDisconnectedPacket(gameid uint16) netutil.GWPacket {
	pkt := netutil.NewPacket()
	pkt.WriteUint16(MT_NOTIFY_GAME_DISCONNECTED)
	pkt.WriteUint16(gameid)
	return pkt
}

func MakeNotifyDeploymentReadyPacket() netutil.GWPacket {
	pkt := netutil.NewPacket()
	pkt.WriteUint16(MT_NOTIFY_DEPLOYMENT_READY)
	return pkt
}

func (gwc *GoWorldConnection) SendSetGameIDAck(dispid uint16, isDeploymentReady bool, connectedGameIDs []uint16, rejectEntities []common.EntityID, srvdisRegisterMap map[string]string) error {
	pkt := netutil.NewPacket()
	pkt.WriteUint16(MT_SET_GAME_ID_ACK)
	pkt.WriteUint16(dispid)

	pkt.WriteBool(isDeploymentReady)

	pkt.WriteUint16(uint16(len(connectedGameIDs)))
	for _, gameid := range connectedGameIDs {
		pkt.WriteUint16(gameid)
	}
	// put rejected entity IDs to the packet
	pkt.WriteUint32(uint32(len(rejectEntities)))
	for _, eid := range rejectEntities {
		pkt.WriteEntityID(eid)
	}
	// put all services to the packet
	pkt.WriteMapStringString(srvdisRegisterMap)
	return gwc.SendPacketRelease(pkt)
}

// SendPacket send a packet to remote
func (gwc *GoWorldConnection) SendPacket(packet netutil.GWPacket) error {
	return gwc.packetConn.Send(packet.Packet)
}

// SendPacketRelease send a packet to remote and then release the packet
func (gwc *GoWorldConnection) SendPacketRelease(packet netutil.GWPacket) error {
	err := gwc.packetConn.Send(packet.Packet)
	packet.Release()
	return err
}

// Recv receives the next packet and retrive the message type
func (gwc *GoWorldConnection) Recv(recvChan chan *pktconn.Packet) error {
	err := gwc.packetConn.Recv(recvChan)
	gwlog.Errorf("GoWorldConnection:  Recv error: %+v", err)
	return err
}

// Close this connection
func (gwc *GoWorldConnection) Close() error {
	gwc.closed.Store(true)
	return gwc.packetConn.Close()
}

// IsClosed returns if the connection is closed
func (gwc *GoWorldConnection) IsClosed() bool {
	return gwc.closed.Load()
}

// RemoteAddr returns the remote address
func (gwc *GoWorldConnection) RemoteAddr() net.Addr {
	return gwc.packetConn.RemoteAddr()
}

// LocalAddr returns the local address
func (gwc *GoWorldConnection) LocalAddr() net.Addr {
	return gwc.packetConn.LocalAddr()
}

func (gwc *GoWorldConnection) String() string {
	return gwc.packetConn.String()
}
