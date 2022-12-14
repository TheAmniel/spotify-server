package socket

import "encoding/json"

const (
	// Default Opcode when receiving core events [RECEIVE ONLY]
	SocketDispatch = 0

	// Sends this when clients initially connect [RECEIVE ONLY]
	SocketHello = 1

	// This is what the client sends when receiving opcode 1 [SEND ONLY]
	SocketInitialize = 2

	// Clients should send Opcode 3 [SEND/RECEIVE]
	SocketHeartbeat = 3

	// Sends this when clients sends heartbeat [RECEIVE ONLY]
	SocketHeartbeatACK = 4

	// TODO: implement
	// Sends this when server request reconnect [RECEIVE ONLY]
	SocketReconnect = 5

	// This is what the client sends with session_id when receiving opcode 5  [SEND ONLY]
	SocketResume = 6
)

const (
	CloseInvalidOpcode        = 4001
	CloseInvalidMessage       = 4002
	CloseNotAuthenticated     = 4003
	CloseByServerRequest      = 4004
	CloseAlreadyAuthenticated = 4005
)

type JSON map[string]any

type SocketMessage struct {
	OP int    `json:"op"`
	T  string `json:"t,omitempty"`
	D  any    `json:"d,omitempty"`
}

func (sm *SocketMessage) ToBytes() []byte {
	if bytes, err := json.Marshal(sm); err != nil {
		return nil
	} else {
		return bytes
	}
}
