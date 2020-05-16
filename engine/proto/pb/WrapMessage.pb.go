// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: WrapMessage.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CommandList int32

const (
	// MT_SET_GAME_ID is a message type for game
	CommandList_MT_SYSTEM_SERVICE_MSG_TYPE_START CommandList = 0
	CommandList_MT_SET_GAME_ID                   CommandList = 1
	// MT_SET_GATE_ID is a message type for gate
	CommandList_MT_SET_GATE_ID CommandList = 2
	// MT_NOTIFY_CREATE_ENTITY is a message type for creating entities
	CommandList_MT_NOTIFY_CREATE_ENTITY CommandList = 3
	// MT_NOTIFY_DESTROY_ENTITY is a message type for destroying entities
	CommandList_MT_NOTIFY_DESTROY_ENTITY CommandList = 4
	// MT_KVREG_REGISTER is a message type for declaring services
	CommandList_MT_KVREG_REGISTER CommandList = 5
	// MT_UNDECLARE_SERVICE is a message type for undeclaring services
	CommandList_MT_UNDECLARE_SERVICE CommandList = 6
	// MT_CREATE_ENTITY_SOMEWHERE is a message type for creating entities
	CommandList_MT_CREATE_ENTITY_SOMEWHERE CommandList = 7
	// MT_LOAD_ENTITY_SOMEWHERE is a message type loading entities
	CommandList_MT_LOAD_ENTITY_SOMEWHERE CommandList = 8
	// MT_NOTIFY_CLIENT_CONNECTED is a message type for clients
	CommandList_MT_NOTIFY_CLIENT_CONNECTED CommandList = 9
	// MT_NOTIFY_CLIENT_DISCONNECTED is a message type for clients
	CommandList_MT_NOTIFY_CLIENT_DISCONNECTED CommandList = 10
	// MT_SYNC_POSITION_YAW_FROM_CLIENT is a message type for clients
	CommandList_MT_SYNC_POSITION_YAW_FROM_CLIENT CommandList = 11
	// MT_NOTIFY_ALL_GAMES_CONNECTED is a message type to notify all games connected
	CommandList_MT_NOTIFY_ALL_GAMES_CONNECTED CommandList = 12
	// NOT USED ANYMORE
	// MT_NOTIFY_GATE_DISCONNECTED is a message type to notify gate disconnected
	CommandList_MT_NOTIFY_GATE_DISCONNECTED CommandList = 13
	// MT_START_FREEZE_GAME is a message type for hot swapping
	CommandList_MT_START_FREEZE_GAME CommandList = 14
	// MT_START_FREEZE_GAME_ACK is a message type for hot swapping
	CommandList_MT_START_FREEZE_GAME_ACK CommandList = 15
	// Message types for migrating
	// MT_MIGRATE_REQUEST is a message type for entity migrations
	CommandList_MT_MIGRATE_REQUEST CommandList = 16
	// MT_REAL_MIGRATE is a message type for entity migrations
	CommandList_MT_REAL_MIGRATE CommandList = 17
	// MT_QUERY_SPACE_GAMEID_FOR_MIGRATE is a message type for entity migrations
	CommandList_MT_QUERY_SPACE_GAMEID_FOR_MIGRATE CommandList = 18
	CommandList_MT_CANCEL_MIGRATE                 CommandList = 19
	// MT_CALL_NIL_SPACES message is used to call nil spaces on all games
	CommandList_MT_CALL_NIL_SPACES CommandList = 20
	// MT_SET_GAME_ID_ACK is sent by dispatcher to game to ACK MT_SET_GAME_ID message
	CommandList_MT_SET_GAME_ID_ACK CommandList = 21
	// MT_NOTIFY_GAME_CONNECTED is sent by dispatcher to game to notify new game connected
	CommandList_MT_NOTIFY_GAME_CONNECTED    CommandList = 22
	CommandList_MT_NOTIFY_GAME_DISCONNECTED CommandList = 23
	CommandList_MT_NOTIFY_DEPLOYMENT_READY  CommandList = 24
	// MT_GAME_LBC_INFO contains game load balacing info
	CommandList_MT_GAME_LBC_INFO CommandList = 25
	// MT_GATE_SERVICE_MSG_TYPE_START is the first message types that should be handled by GateService
	CommandList_MT_GATE_SERVICE_MSG_TYPE_START CommandList = 1000
	// MT_REDIRECT_TO_GATEPROXY_MSG_TYPE_START is the first message type that should be redirected to client proxy
	CommandList_MT_REDIRECT_TO_GATEPROXY_MSG_TYPE_START CommandList = 1001
	// MT_CREATE_ENTITY_ON_CLIENT message type
	CommandList_MT_CREATE_ENTITY_ON_CLIENT CommandList = 1002
	// MT_DESTROY_ENTITY_ON_CLIENT message type
	CommandList_MT_DESTROY_ENTITY_ON_CLIENT CommandList = 1003
	// MT_NOTIFY_MAP_ATTR_CHANGE_ON_CLIENT message type
	CommandList_MT_NOTIFY_MAP_ATTR_CHANGE_ON_CLIENT CommandList = 1004
	// MT_NOTIFY_MAP_ATTR_DEL_ON_CLIENT message type
	CommandList_MT_NOTIFY_MAP_ATTR_DEL_ON_CLIENT CommandList = 1005
	// MT_NOTIFY_LIST_ATTR_CHANGE_ON_CLIENT message type
	CommandList_MT_NOTIFY_LIST_ATTR_CHANGE_ON_CLIENT CommandList = 1006
	// MT_NOTIFY_LIST_ATTR_POP_ON_CLIENT message type
	CommandList_MT_NOTIFY_LIST_ATTR_POP_ON_CLIENT CommandList = 1007
	// MT_NOTIFY_LIST_ATTR_APPEND_ON_CLIENT message type
	CommandList_MT_NOTIFY_LIST_ATTR_APPEND_ON_CLIENT CommandList = 1008
	// MT_SET_CLIENTPROXY_FILTER_PROP message type
	CommandList_MT_SET_CLIENTPROXY_FILTER_PROP CommandList = 1009
	// MT_CLEAR_CLIENTPROXY_FILTER_PROPS message type
	CommandList_MT_CLEAR_CLIENTPROXY_FILTER_PROPS CommandList = 1010
	// MT_NOTIFY_MAP_ATTR_CLEAR_ON_CLIENT message type
	CommandList_MT_NOTIFY_MAP_ATTR_CLEAR_ON_CLIENT CommandList = 1011
	// MT_REDIRECT_TO_GATEPROXY_MSG_TYPE_STOP message type
	CommandList_MT_REDIRECT_TO_GATEPROXY_MSG_TYPE_STOP CommandList = 1499
	// MT_CALL_FILTERED_CLIENTS message type: messages to be processed by GateService from Dispatcher, but not redirected to clients
	CommandList_MT_CALL_FILTERED_CLIENTS CommandList = 1501
	// MT_SYNC_POSITION_YAW_ON_CLIENTS message type
	CommandList_MT_SYNC_POSITION_YAW_ON_CLIENTS CommandList = 1502
	// MT_GATE_SERVICE_MSG_TYPE_STOP message type
	CommandList_MT_GATE_SERVICE_MSG_TYPE_STOP CommandList = 1999
	// MT_SET_CLIENT_CLIENTID message is sent to client to set its clientid
	CommandList_MT_SET_CLIENT_CLIENTID               CommandList = 2001
	CommandList_MT_UDP_SYNC_CONN_NOTIFY_CLIENTID     CommandList = 2002
	CommandList_MT_UDP_SYNC_CONN_NOTIFY_CLIENTID_ACK CommandList = 2003
	// MT_HEARTBEAT_FROM_CLIENT is sent by client to notify the gate server that the client is alive
	CommandList_MT_HEARTBEAT_FROM_CLIENT CommandList = 2004
)

// Enum value maps for CommandList.
var (
	CommandList_name = map[int32]string{
		0:    "MT_SYSTEM_SERVICE_MSG_TYPE_START",
		1:    "MT_SET_GAME_ID",
		2:    "MT_SET_GATE_ID",
		3:    "MT_NOTIFY_CREATE_ENTITY",
		4:    "MT_NOTIFY_DESTROY_ENTITY",
		5:    "MT_KVREG_REGISTER",
		6:    "MT_UNDECLARE_SERVICE",
		7:    "MT_CREATE_ENTITY_SOMEWHERE",
		8:    "MT_LOAD_ENTITY_SOMEWHERE",
		9:    "MT_NOTIFY_CLIENT_CONNECTED",
		10:   "MT_NOTIFY_CLIENT_DISCONNECTED",
		11:   "MT_SYNC_POSITION_YAW_FROM_CLIENT",
		12:   "MT_NOTIFY_ALL_GAMES_CONNECTED",
		13:   "MT_NOTIFY_GATE_DISCONNECTED",
		14:   "MT_START_FREEZE_GAME",
		15:   "MT_START_FREEZE_GAME_ACK",
		16:   "MT_MIGRATE_REQUEST",
		17:   "MT_REAL_MIGRATE",
		18:   "MT_QUERY_SPACE_GAMEID_FOR_MIGRATE",
		19:   "MT_CANCEL_MIGRATE",
		20:   "MT_CALL_NIL_SPACES",
		21:   "MT_SET_GAME_ID_ACK",
		22:   "MT_NOTIFY_GAME_CONNECTED",
		23:   "MT_NOTIFY_GAME_DISCONNECTED",
		24:   "MT_NOTIFY_DEPLOYMENT_READY",
		25:   "MT_GAME_LBC_INFO",
		1000: "MT_GATE_SERVICE_MSG_TYPE_START",
		1001: "MT_REDIRECT_TO_GATEPROXY_MSG_TYPE_START",
		1002: "MT_CREATE_ENTITY_ON_CLIENT",
		1003: "MT_DESTROY_ENTITY_ON_CLIENT",
		1004: "MT_NOTIFY_MAP_ATTR_CHANGE_ON_CLIENT",
		1005: "MT_NOTIFY_MAP_ATTR_DEL_ON_CLIENT",
		1006: "MT_NOTIFY_LIST_ATTR_CHANGE_ON_CLIENT",
		1007: "MT_NOTIFY_LIST_ATTR_POP_ON_CLIENT",
		1008: "MT_NOTIFY_LIST_ATTR_APPEND_ON_CLIENT",
		1009: "MT_SET_CLIENTPROXY_FILTER_PROP",
		1010: "MT_CLEAR_CLIENTPROXY_FILTER_PROPS",
		1011: "MT_NOTIFY_MAP_ATTR_CLEAR_ON_CLIENT",
		1499: "MT_REDIRECT_TO_GATEPROXY_MSG_TYPE_STOP",
		1501: "MT_CALL_FILTERED_CLIENTS",
		1502: "MT_SYNC_POSITION_YAW_ON_CLIENTS",
		1999: "MT_GATE_SERVICE_MSG_TYPE_STOP",
		2001: "MT_SET_CLIENT_CLIENTID",
		2002: "MT_UDP_SYNC_CONN_NOTIFY_CLIENTID",
		2003: "MT_UDP_SYNC_CONN_NOTIFY_CLIENTID_ACK",
		2004: "MT_HEARTBEAT_FROM_CLIENT",
	}
	CommandList_value = map[string]int32{
		"MT_SYSTEM_SERVICE_MSG_TYPE_START":        0,
		"MT_SET_GAME_ID":                          1,
		"MT_SET_GATE_ID":                          2,
		"MT_NOTIFY_CREATE_ENTITY":                 3,
		"MT_NOTIFY_DESTROY_ENTITY":                4,
		"MT_KVREG_REGISTER":                       5,
		"MT_UNDECLARE_SERVICE":                    6,
		"MT_CREATE_ENTITY_SOMEWHERE":              7,
		"MT_LOAD_ENTITY_SOMEWHERE":                8,
		"MT_NOTIFY_CLIENT_CONNECTED":              9,
		"MT_NOTIFY_CLIENT_DISCONNECTED":           10,
		"MT_SYNC_POSITION_YAW_FROM_CLIENT":        11,
		"MT_NOTIFY_ALL_GAMES_CONNECTED":           12,
		"MT_NOTIFY_GATE_DISCONNECTED":             13,
		"MT_START_FREEZE_GAME":                    14,
		"MT_START_FREEZE_GAME_ACK":                15,
		"MT_MIGRATE_REQUEST":                      16,
		"MT_REAL_MIGRATE":                         17,
		"MT_QUERY_SPACE_GAMEID_FOR_MIGRATE":       18,
		"MT_CANCEL_MIGRATE":                       19,
		"MT_CALL_NIL_SPACES":                      20,
		"MT_SET_GAME_ID_ACK":                      21,
		"MT_NOTIFY_GAME_CONNECTED":                22,
		"MT_NOTIFY_GAME_DISCONNECTED":             23,
		"MT_NOTIFY_DEPLOYMENT_READY":              24,
		"MT_GAME_LBC_INFO":                        25,
		"MT_GATE_SERVICE_MSG_TYPE_START":          1000,
		"MT_REDIRECT_TO_GATEPROXY_MSG_TYPE_START": 1001,
		"MT_CREATE_ENTITY_ON_CLIENT":              1002,
		"MT_DESTROY_ENTITY_ON_CLIENT":             1003,
		"MT_NOTIFY_MAP_ATTR_CHANGE_ON_CLIENT":     1004,
		"MT_NOTIFY_MAP_ATTR_DEL_ON_CLIENT":        1005,
		"MT_NOTIFY_LIST_ATTR_CHANGE_ON_CLIENT":    1006,
		"MT_NOTIFY_LIST_ATTR_POP_ON_CLIENT":       1007,
		"MT_NOTIFY_LIST_ATTR_APPEND_ON_CLIENT":    1008,
		"MT_SET_CLIENTPROXY_FILTER_PROP":          1009,
		"MT_CLEAR_CLIENTPROXY_FILTER_PROPS":       1010,
		"MT_NOTIFY_MAP_ATTR_CLEAR_ON_CLIENT":      1011,
		"MT_REDIRECT_TO_GATEPROXY_MSG_TYPE_STOP":  1499,
		"MT_CALL_FILTERED_CLIENTS":                1501,
		"MT_SYNC_POSITION_YAW_ON_CLIENTS":         1502,
		"MT_GATE_SERVICE_MSG_TYPE_STOP":           1999,
		"MT_SET_CLIENT_CLIENTID":                  2001,
		"MT_UDP_SYNC_CONN_NOTIFY_CLIENTID":        2002,
		"MT_UDP_SYNC_CONN_NOTIFY_CLIENTID_ACK":    2003,
		"MT_HEARTBEAT_FROM_CLIENT":                2004,
	}
)

func (x CommandList) Enum() *CommandList {
	p := new(CommandList)
	*p = x
	return p
}

func (x CommandList) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommandList) Descriptor() protoreflect.EnumDescriptor {
	return file_WrapMessage_proto_enumTypes[0].Descriptor()
}

func (CommandList) Type() protoreflect.EnumType {
	return &file_WrapMessage_proto_enumTypes[0]
}

func (x CommandList) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommandList.Descriptor instead.
func (CommandList) EnumDescriptor() ([]byte, []int) {
	return file_WrapMessage_proto_rawDescGZIP(), []int{0}
}

type WrapMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmd      CommandList `protobuf:"varint,1,opt,name=cmd,proto3,enum=pb.CommandList" json:"cmd,omitempty"` // 协议编号
	Request  uint32      `protobuf:"varint,2,opt,name=request,proto3" json:"request,omitempty"`             //请求ID
	Response uint32      `protobuf:"varint,3,opt,name=response,proto3" json:"response,omitempty"`           //回复ID
	Content  []byte      `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`              // 协议二进制
	Code     uint32      `protobuf:"varint,5,opt,name=code,proto3" json:"code,omitempty"`                   // 错误码
}

func (x *WrapMessage) Reset() {
	*x = WrapMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WrapMessage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WrapMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WrapMessage) ProtoMessage() {}

func (x *WrapMessage) ProtoReflect() protoreflect.Message {
	mi := &file_WrapMessage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WrapMessage.ProtoReflect.Descriptor instead.
func (*WrapMessage) Descriptor() ([]byte, []int) {
	return file_WrapMessage_proto_rawDescGZIP(), []int{0}
}

func (x *WrapMessage) GetCmd() CommandList {
	if x != nil {
		return x.Cmd
	}
	return CommandList_MT_SYSTEM_SERVICE_MSG_TYPE_START
}

func (x *WrapMessage) GetRequest() uint32 {
	if x != nil {
		return x.Request
	}
	return 0
}

func (x *WrapMessage) GetResponse() uint32 {
	if x != nil {
		return x.Response
	}
	return 0
}

func (x *WrapMessage) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *WrapMessage) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_WrapMessage_proto protoreflect.FileDescriptor

var file_WrapMessage_proto_rawDesc = []byte{
	0x0a, 0x11, 0x57, 0x72, 0x61, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x94, 0x01, 0x0a, 0x0b, 0x57, 0x72, 0x61, 0x70,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x2a, 0xf9,
	0x0b, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24,
	0x0a, 0x20, 0x4d, 0x54, 0x5f, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x5f, 0x53, 0x45, 0x52, 0x56,
	0x49, 0x43, 0x45, 0x5f, 0x4d, 0x53, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x41,
	0x52, 0x54, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x47,
	0x41, 0x4d, 0x45, 0x5f, 0x49, 0x44, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x54, 0x5f, 0x53,
	0x45, 0x54, 0x5f, 0x47, 0x41, 0x54, 0x45, 0x5f, 0x49, 0x44, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17,
	0x4d, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45,
	0x5f, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x10, 0x03, 0x12, 0x1c, 0x0a, 0x18, 0x4d, 0x54, 0x5f,
	0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x44, 0x45, 0x53, 0x54, 0x52, 0x4f, 0x59, 0x5f, 0x45,
	0x4e, 0x54, 0x49, 0x54, 0x59, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x4d, 0x54, 0x5f, 0x4b, 0x56,
	0x52, 0x45, 0x47, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x10, 0x05, 0x12, 0x18,
	0x0a, 0x14, 0x4d, 0x54, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x43, 0x4c, 0x41, 0x52, 0x45, 0x5f, 0x53,
	0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x06, 0x12, 0x1e, 0x0a, 0x1a, 0x4d, 0x54, 0x5f, 0x43,
	0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x53, 0x4f, 0x4d,
	0x45, 0x57, 0x48, 0x45, 0x52, 0x45, 0x10, 0x07, 0x12, 0x1c, 0x0a, 0x18, 0x4d, 0x54, 0x5f, 0x4c,
	0x4f, 0x41, 0x44, 0x5f, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x53, 0x4f, 0x4d, 0x45, 0x57,
	0x48, 0x45, 0x52, 0x45, 0x10, 0x08, 0x12, 0x1e, 0x0a, 0x1a, 0x4d, 0x54, 0x5f, 0x4e, 0x4f, 0x54,
	0x49, 0x46, 0x59, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45,
	0x43, 0x54, 0x45, 0x44, 0x10, 0x09, 0x12, 0x21, 0x0a, 0x1d, 0x4d, 0x54, 0x5f, 0x4e, 0x4f, 0x54,
	0x49, 0x46, 0x59, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x44, 0x49, 0x53, 0x43, 0x4f,
	0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x0a, 0x12, 0x24, 0x0a, 0x20, 0x4d, 0x54, 0x5f,
	0x53, 0x59, 0x4e, 0x43, 0x5f, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x59, 0x41,
	0x57, 0x5f, 0x46, 0x52, 0x4f, 0x4d, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x0b, 0x12,
	0x21, 0x0a, 0x1d, 0x4d, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x41, 0x4c, 0x4c,
	0x5f, 0x47, 0x41, 0x4d, 0x45, 0x53, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45, 0x44,
	0x10, 0x0c, 0x12, 0x1f, 0x0a, 0x1b, 0x4d, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f,
	0x47, 0x41, 0x54, 0x45, 0x5f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45,
	0x44, 0x10, 0x0d, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f,
	0x46, 0x52, 0x45, 0x45, 0x5a, 0x45, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x10, 0x0e, 0x12, 0x1c, 0x0a,
	0x18, 0x4d, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x46, 0x52, 0x45, 0x45, 0x5a, 0x45,
	0x5f, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x41, 0x43, 0x4b, 0x10, 0x0f, 0x12, 0x16, 0x0a, 0x12, 0x4d,
	0x54, 0x5f, 0x4d, 0x49, 0x47, 0x52, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53,
	0x54, 0x10, 0x10, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x54, 0x5f, 0x52, 0x45, 0x41, 0x4c, 0x5f, 0x4d,
	0x49, 0x47, 0x52, 0x41, 0x54, 0x45, 0x10, 0x11, 0x12, 0x25, 0x0a, 0x21, 0x4d, 0x54, 0x5f, 0x51,
	0x55, 0x45, 0x52, 0x59, 0x5f, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x49,
	0x44, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x4d, 0x49, 0x47, 0x52, 0x41, 0x54, 0x45, 0x10, 0x12, 0x12,
	0x15, 0x0a, 0x11, 0x4d, 0x54, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x5f, 0x4d, 0x49, 0x47,
	0x52, 0x41, 0x54, 0x45, 0x10, 0x13, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x54, 0x5f, 0x43, 0x41, 0x4c,
	0x4c, 0x5f, 0x4e, 0x49, 0x4c, 0x5f, 0x53, 0x50, 0x41, 0x43, 0x45, 0x53, 0x10, 0x14, 0x12, 0x16,
	0x0a, 0x12, 0x4d, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x49, 0x44,
	0x5f, 0x41, 0x43, 0x4b, 0x10, 0x15, 0x12, 0x1c, 0x0a, 0x18, 0x4d, 0x54, 0x5f, 0x4e, 0x4f, 0x54,
	0x49, 0x46, 0x59, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54,
	0x45, 0x44, 0x10, 0x16, 0x12, 0x1f, 0x0a, 0x1b, 0x4d, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x49, 0x46,
	0x59, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43,
	0x54, 0x45, 0x44, 0x10, 0x17, 0x12, 0x1e, 0x0a, 0x1a, 0x4d, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x49,
	0x46, 0x59, 0x5f, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45,
	0x41, 0x44, 0x59, 0x10, 0x18, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x54, 0x5f, 0x47, 0x41, 0x4d, 0x45,
	0x5f, 0x4c, 0x42, 0x43, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x19, 0x12, 0x23, 0x0a, 0x1e, 0x4d,
	0x54, 0x5f, 0x47, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x4d,
	0x53, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0xe8, 0x07,
	0x12, 0x2c, 0x0a, 0x27, 0x4d, 0x54, 0x5f, 0x52, 0x45, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x5f,
	0x54, 0x4f, 0x5f, 0x47, 0x41, 0x54, 0x45, 0x50, 0x52, 0x4f, 0x58, 0x59, 0x5f, 0x4d, 0x53, 0x47,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0xe9, 0x07, 0x12, 0x1f,
	0x0a, 0x1a, 0x4d, 0x54, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x45, 0x4e, 0x54, 0x49,
	0x54, 0x59, 0x5f, 0x4f, 0x4e, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0xea, 0x07, 0x12,
	0x20, 0x0a, 0x1b, 0x4d, 0x54, 0x5f, 0x44, 0x45, 0x53, 0x54, 0x52, 0x4f, 0x59, 0x5f, 0x45, 0x4e,
	0x54, 0x49, 0x54, 0x59, 0x5f, 0x4f, 0x4e, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0xeb,
	0x07, 0x12, 0x28, 0x0a, 0x23, 0x4d, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x4d,
	0x41, 0x50, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x4f,
	0x4e, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0xec, 0x07, 0x12, 0x25, 0x0a, 0x20, 0x4d,
	0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x4d, 0x41, 0x50, 0x5f, 0x41, 0x54, 0x54,
	0x52, 0x5f, 0x44, 0x45, 0x4c, 0x5f, 0x4f, 0x4e, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10,
	0xed, 0x07, 0x12, 0x29, 0x0a, 0x24, 0x4d, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f,
	0x4c, 0x49, 0x53, 0x54, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45,
	0x5f, 0x4f, 0x4e, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0xee, 0x07, 0x12, 0x26, 0x0a,
	0x21, 0x4d, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x5f,
	0x41, 0x54, 0x54, 0x52, 0x5f, 0x50, 0x4f, 0x50, 0x5f, 0x4f, 0x4e, 0x5f, 0x43, 0x4c, 0x49, 0x45,
	0x4e, 0x54, 0x10, 0xef, 0x07, 0x12, 0x29, 0x0a, 0x24, 0x4d, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x49,
	0x46, 0x59, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x5f, 0x41, 0x50, 0x50,
	0x45, 0x4e, 0x44, 0x5f, 0x4f, 0x4e, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0xf0, 0x07,
	0x12, 0x23, 0x0a, 0x1e, 0x4d, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e,
	0x54, 0x50, 0x52, 0x4f, 0x58, 0x59, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x50, 0x52,
	0x4f, 0x50, 0x10, 0xf1, 0x07, 0x12, 0x26, 0x0a, 0x21, 0x4d, 0x54, 0x5f, 0x43, 0x4c, 0x45, 0x41,
	0x52, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x50, 0x52, 0x4f, 0x58, 0x59, 0x5f, 0x46, 0x49,
	0x4c, 0x54, 0x45, 0x52, 0x5f, 0x50, 0x52, 0x4f, 0x50, 0x53, 0x10, 0xf2, 0x07, 0x12, 0x27, 0x0a,
	0x22, 0x4d, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x4d, 0x41, 0x50, 0x5f, 0x41,
	0x54, 0x54, 0x52, 0x5f, 0x43, 0x4c, 0x45, 0x41, 0x52, 0x5f, 0x4f, 0x4e, 0x5f, 0x43, 0x4c, 0x49,
	0x45, 0x4e, 0x54, 0x10, 0xf3, 0x07, 0x12, 0x2b, 0x0a, 0x26, 0x4d, 0x54, 0x5f, 0x52, 0x45, 0x44,
	0x49, 0x52, 0x45, 0x43, 0x54, 0x5f, 0x54, 0x4f, 0x5f, 0x47, 0x41, 0x54, 0x45, 0x50, 0x52, 0x4f,
	0x58, 0x59, 0x5f, 0x4d, 0x53, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x4f, 0x50,
	0x10, 0xdb, 0x0b, 0x12, 0x1d, 0x0a, 0x18, 0x4d, 0x54, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x46,
	0x49, 0x4c, 0x54, 0x45, 0x52, 0x45, 0x44, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x53, 0x10,
	0xdd, 0x0b, 0x12, 0x24, 0x0a, 0x1f, 0x4d, 0x54, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x50, 0x4f,
	0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x59, 0x41, 0x57, 0x5f, 0x4f, 0x4e, 0x5f, 0x43, 0x4c,
	0x49, 0x45, 0x4e, 0x54, 0x53, 0x10, 0xde, 0x0b, 0x12, 0x22, 0x0a, 0x1d, 0x4d, 0x54, 0x5f, 0x47,
	0x41, 0x54, 0x45, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x4d, 0x53, 0x47, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x10, 0xcf, 0x0f, 0x12, 0x1b, 0x0a, 0x16,
	0x4d, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x4c,
	0x49, 0x45, 0x4e, 0x54, 0x49, 0x44, 0x10, 0xd1, 0x0f, 0x12, 0x25, 0x0a, 0x20, 0x4d, 0x54, 0x5f,
	0x55, 0x44, 0x50, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x5f, 0x4e, 0x4f,
	0x54, 0x49, 0x46, 0x59, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x49, 0x44, 0x10, 0xd2, 0x0f,
	0x12, 0x29, 0x0a, 0x24, 0x4d, 0x54, 0x5f, 0x55, 0x44, 0x50, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x5f,
	0x43, 0x4f, 0x4e, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x43, 0x4c, 0x49, 0x45,
	0x4e, 0x54, 0x49, 0x44, 0x5f, 0x41, 0x43, 0x4b, 0x10, 0xd3, 0x0f, 0x12, 0x1d, 0x0a, 0x18, 0x4d,
	0x54, 0x5f, 0x48, 0x45, 0x41, 0x52, 0x54, 0x42, 0x45, 0x41, 0x54, 0x5f, 0x46, 0x52, 0x4f, 0x4d,
	0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0xd4, 0x0f, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_WrapMessage_proto_rawDescOnce sync.Once
	file_WrapMessage_proto_rawDescData = file_WrapMessage_proto_rawDesc
)

func file_WrapMessage_proto_rawDescGZIP() []byte {
	file_WrapMessage_proto_rawDescOnce.Do(func() {
		file_WrapMessage_proto_rawDescData = protoimpl.X.CompressGZIP(file_WrapMessage_proto_rawDescData)
	})
	return file_WrapMessage_proto_rawDescData
}

var file_WrapMessage_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_WrapMessage_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_WrapMessage_proto_goTypes = []interface{}{
	(CommandList)(0),    // 0: pb.CommandList
	(*WrapMessage)(nil), // 1: pb.WrapMessage
}
var file_WrapMessage_proto_depIdxs = []int32{
	0, // 0: pb.WrapMessage.cmd:type_name -> pb.CommandList
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_WrapMessage_proto_init() }
func file_WrapMessage_proto_init() {
	if File_WrapMessage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_WrapMessage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WrapMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_WrapMessage_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_WrapMessage_proto_goTypes,
		DependencyIndexes: file_WrapMessage_proto_depIdxs,
		EnumInfos:         file_WrapMessage_proto_enumTypes,
		MessageInfos:      file_WrapMessage_proto_msgTypes,
	}.Build()
	File_WrapMessage_proto = out.File
	file_WrapMessage_proto_rawDesc = nil
	file_WrapMessage_proto_goTypes = nil
	file_WrapMessage_proto_depIdxs = nil
}