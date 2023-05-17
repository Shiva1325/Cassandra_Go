package service

import (
	"net"
	"sync"
)

type HostInfo struct {
	mu             sync.RWMutex
	peer           []int32
	connectAddress net.IP
	port           int
	hostId         int32
	tokens         []string
}

func (h *HostInfo) HostID() int32 {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.hostId
}
