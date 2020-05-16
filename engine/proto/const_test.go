package proto

import (
	"fmt"
	"testing"
)
const (
	// MT_GATE_SERVICE_MSG_TYPE_START is the first message types that should be handled by GateService
	ak1 = 1000 + iota
	// MT_REDIRECT_TO_GATEPROXY_MSG_TYPE_START is the first message type that should be redirected to client proxy
	ak2
	// MT_CREATE_ENTITY_ON_CLIENT message type
	ak3
	// MT_DESTROY_ENTITY_ON_CLIENT message type
	ak4
	ak5 = 1499
)
func TestConst(t *testing.T) {
	fmt.Println("2")
	fmt.Println(ak1 , ak2 , ak3)
}
