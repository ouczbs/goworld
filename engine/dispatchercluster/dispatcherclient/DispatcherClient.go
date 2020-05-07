package dispatcherclient

import (
	"github.com/ouczbs/goworld/engine/proto/pb"
	"github.com/xiaonanln/netconnutil"
	"net"

	"github.com/ouczbs/goworld/engine/consts"
	"github.com/ouczbs/goworld/engine/gwlog"
	"github.com/ouczbs/goworld/engine/proto"
)

type DispatcherClientType int

const (
	GameDispatcherClientType DispatcherClientType = 1 + iota
	GateDispatcherClientType
)

// DispatcherClient is a client connection to the dispatcher
type DispatcherClient struct {
	*proto.GoWorldConnection
	dctype        DispatcherClientType
	isReconnect   bool
	isRestoreGame bool
}

func newDispatcherClient(dctype DispatcherClientType, conn net.Conn, isReconnect bool, isRestoreGame bool) *DispatcherClient {
	conn = netconnutil.NewNoTempErrorConn(conn)
	gwc := proto.NewGoWorldConnection(netconnutil.NewBufferedConn(conn, consts.BUFFERED_READ_BUFFSIZE, consts.BUFFERED_WRITE_BUFFSIZE))
	if dctype != GameDispatcherClientType && dctype != GateDispatcherClientType {
		gwlog.Fatalf("invalid dispatcher client type: %v", dctype)
	}

	dc := &DispatcherClient{
		GoWorldConnection: gwc,
		dctype:            dctype,
		isReconnect:       isReconnect,
		isRestoreGame:     isRestoreGame,
	}
	dc.SetAutoFlush(consts.DISPATCHER_CLIENT_FLUSH_INTERVAL)
	return dc
}

// Close the dispatcher client
func (dc *DispatcherClient) Close() error {
	return dc.GoWorldConnection.Close()
}
func (dc *DispatcherClient) SendSetGameID(dcm * DispatcherConnMgr) error{
	data := &pb.SET_GAME_ID{
		GameId:uint32(dcm.gid),
		IsReconnect:dcm.isReconnect,
		IsRestore:dcm.isRestoreGame,
		IsBanBootEntity:dcm.isBanBootEntity,
	}
	eids := dcm.delegate.GetEntityIDsForDispatcher(dcm.dispid)
	for _, eid := range  eids{
		data.EntityIdList = append(data.EntityIdList, string(eid))
	}
	return proto.SendPackageWithPbMessage( dc.GoWorldConnection , data)
}
func (dc *DispatcherClient) SendSetGateID(id uint16) error {
	data := &pb.SET_GATE_ID{GateId: uint32(id)}
	return proto.SendPackageWithPbMessage(dc.GoWorldConnection, data)
}