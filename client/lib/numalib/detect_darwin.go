//go:build darwin

package numalib

import (
	"github.com/hashicorp/nomad/client/lib/idset"
	"github.com/shoenig/go-m1cpu"
)

func PlatformScanners() []SystemScanner {
	return []SystemScanner{
		new(MacOS),
	}
}

type MacOS struct{}

func (m *MacOS) ScanSystem(top *Topology) {
	switch m1cpu.IsAppleSilicon() {
	case true:
		m.scanAppleSilicon(top)
	case false:
		m.scanLegacyX86(top)
	}
}

func (m *MacOS) scanAppleSilicon(top *Topology) {
	const (
		nodeID   = NodeID(0)
		socketID = SocketID(0)
		maxSpeed = KHz(0)
	)

	// all apple hardware is non-numa; just assume it is so
	top.NodeIDs = idset.Empty[NodeID]()
	top.NodeIDs.Insert(nodeID)

	pCoreCount := m1cpu.PCoreCount()
	pCoreSpeed := KHz(m1cpu.PCoreHz() / 1024)

	nthCore := CoreID(0)

	for i := 0; i < pCoreCount; i++ {
		top.insert(nodeID, socketID, nthCore, performance, maxSpeed, pCoreSpeed)
		nthCore++
	}

	eCoreCount := m1cpu.ECoreCount()
	eCoreSpeed := KHz(m1cpu.ECoreHz() / 1024)

	for i := 0; i < eCoreCount; i++ {
		top.insert(nodeID, socketID, nthCore, efficiency, maxSpeed, eCoreSpeed)
		nthCore++
	}
}

func (m *MacOS) scanLegacyX86(top *Topology) {
	// hello
}