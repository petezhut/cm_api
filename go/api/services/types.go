package services

import "net/url"

// Service this is the General-use structure for the services endpoint
type Service struct {
	ServiceName      string
	ServiceType      string
	ServiceURL       *url.URL
	RoleInstancesURL *url.URL
}
