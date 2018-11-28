package manager

import (
	"fmt"

	"github.com/petezhut/cm_api/go/api"
)

func NewManager(hostname string) *Manager {
	return &Manager{
		Hostname: hostname,
		Port: api.DefaultAPIPort}
}

func (cm *Manager) String() string {
	return fmt.Sprintf("Manager(Hostname: %s, Port: %d)", cm.Hostname, cm.Port)
}
