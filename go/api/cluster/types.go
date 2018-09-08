package cluster

import (
	"net/url"

	"github.com/petezhut/cm_api/go/api/auth"
	"github.com/petezhut/cm_api/go/api/manager"
)

// Host - Object for Host-related information
type Host struct {
	HostID   string
	HostName string
}

type ClusterUrls struct {
	APIURL               *url.URL
	ClusterURL           *url.URL
	ClusterHostsURL      *url.URL
	ClusterVersionURL    *url.URL
	ClusterApiVersionURL *url.URL
}

type Cluster struct {
	ClusterName string
	Admin       *auth.Auth
	Manager     *manager.Manager
	Hosts       []*Host
	APIVersion  string
	Version     string
	URLS        *ClusterUrls
}
