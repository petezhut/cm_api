package manager

import (
	"fmt"

	"github.com/petezhut/cm_api/go/api"
)

func NewManager() *Manager {
	return &Manager{
		Hostname: "gojira-jmcfarland-1.gce.cloudera.com",
		Port: api.DefaultAPIPort}
}

func (cm *Manager) String() string {
	return fmt.Sprintf("Manager(Hostname: %s, Port: %d)", cm.Hostname, cm.Port)
}
