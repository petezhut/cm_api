package cluster

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/petezhut/cm_api/go/logging"
	"github.com/petezhut/cm_api/go/api"
	"github.com/petezhut/cm_api/go/api/auth"
	"github.com/petezhut/cm_api/go/api/manager"
	"github.com/petezhut/cm_api/go/api/urls"
)

// NewClusterURLs - This creates a container object of cluster-related URLs
func (thisCluster *Cluster) NewClusterURLs() *ClusterUrls{
	urls := new(ClusterUrls)
	emptyUrl, _ := url.Parse("")
	urls.APIURL = emptyUrl
	urls.ClusterURL = emptyUrl
	urls.ClusterHostsURL = emptyUrl
	urls.ClusterVersionURL = thisCluster.getClusterVersionURL()
	urls.ClusterApiVersionURL = emptyUrl
	return urls

}

// NewCluster - This is to create a new Cluster Object
func NewCluster() *Cluster {
	newCluster := new(Cluster)
	newCluster.Admin = auth.NewAuth()
	newCluster.Manager = manager.NewManager()
	newCluster.URLS = NewClusterURLs()
	return newCluster
}

// This produces the string representation of the ClusterURL object
func (clusterURLs *ClusterUrls) String() string {
	return fmt.Sprintf("URLS(APIURL: %s, ClusterURL: %s, ClusterHostsURL: %s, ClusterVersionURL: %s, ClusterAPIVersionURL: %s)",
		clusterURLs.APIURL, clusterURLs.ClusterURL, clusterURLs.ClusterHostsURL, clusterURLs.ClusterVersionURL, clusterURLs.ClusterApiVersionURL)
}

// This produces the string representation of the Cluster object
func (thisCluster *Cluster) String() string {
	return fmt.Sprintf("Cluster(ClusterName: %s, Admin: %s, Manager: %s, APIVersion: %s, Version: %s, URLS: %s)",
		thisCluster.ClusterName, thisCluster.Admin, thisCluster.Manager, thisCluster.APIVersion, thisCluster.Version, thisCluster.URLS)
}

// Update the Cluster object to have some core information
func (thisCluster *Cluster) Update() {
	logging.DEBUG("Starting Update")
	thisCluster.setAPIVersion()
	logging.DEBUG("Finished Setting APIVersion")
	var clusterList api.ClusterList
	json.Unmarshal(api.Get(thisCluster.Admin, thisCluster.URLS.ClusterVersionURL), &clusterList)
	thisCluster.Version = clusterList.Items[0].FullVersion
	thisCluster.ClusterName = clusterList.Items[0].Name
}

// setAPIVersion - an internal method to query the API version in use on the clutser
func (thisCluster *Cluster) setAPIVersion() {
	logging.DEBUG("Staring APIVersion")
	logging.DEBUG(fmt.Sprintf("ClusterAPIVersionURL = %s", thisCluster.URLS.ClusterApiVersionURL))
	logging.DEBUG(fmt.Sprintf("ClusterAPIVersionURL = %s", thisCluster.ge
	thisCluster.APIVersion = string(api.Get(thisCluster.Admin, thisCluster.URLS.ClusterApiVersionURL))
}

// GetClusterHosts - Collect the list of hostnames in the cluster
func (thisCluster *Cluster) GetClusterHosts() string {
	return string(api.Get(thisCluster.Admin, thisCluster.URLS.ClusterHostsURL))
}
