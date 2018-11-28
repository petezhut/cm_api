package cluster

import (
	"encoding/json"
	"fmt"
	"github.com/petezhut/cm_api/go/api"
	"github.com/petezhut/cm_api/go/api/auth"
	"github.com/petezhut/cm_api/go/api/manager"
	"github.com/petezhut/cm_api/go/logging"
	"net/url"
)

func (cl *ClusterObj) GetClusterVersionURL() *url.URL {
	if cl.URLS.ClusterVersionURL.String() == "" {
		logging.DEBUG("Setting ClusterVersionURL")
		realVersionURL, _ := url.Parse(fmt.Sprintf(ClusterURL, cl.Manager.Hostname, cl.Manager.Port, cl.APIVersion))
		logging.DEBUG(fmt.Sprintf("Found: %s", realVersionURL))
		cl.URLS.ClusterVersionURL = realVersionURL
		cl.URLS.ClusterURL = realVersionURL
	}
	logging.DEBUG("Returning ClusterVersionURL")
	return cl.URLS.ClusterVersionURL

}

// NewClusterURLs - This creates a container object of cluster-related URLs
func NewClusterObjURLs() *ClusterUrls {
	urls := new(ClusterUrls)
	emptyUrl, _ := url.Parse("")
	urls.APIURL = emptyUrl
	urls.ClusterURL = emptyUrl
	urls.ClusterHostsURL = emptyUrl
	urls.ClusterVersionURL = emptyUrl
	urls.ClusterApiVersionURL = emptyUrl
	return urls
}

// EmptyURL - This is to create a new ClusterObj Object
func New(CmHost string) *ClusterObj {
// func NewClusterObj(CmHost string) *ClusterObj {
	newClusterObj := new(ClusterObj)
	newClusterObj.Admin = auth.NewAuth()
	newClusterObj.Manager = manager.NewManager(CmHost)
	newClusterObj.URLS = NewClusterObjURLs()
	newClusterObj.Update()
	return newClusterObj
}

// This produces the string representation of the ClusterUrls object
func (cu *ClusterUrls) String() string {
	return fmt.Sprintf("ClusterURLs(APIURL: %s, ClusterURL: %s, ClusterHostsURL: %s, ClusterVersionURL: %s, ClusterApiVersionURL: %s)",
	cu.APIURL, cu.ClusterURL, cu.ClusterHostsURL, cu.ClusterVersionURL, cu.ClusterApiVersionURL)
}

// This produces the string representation of the ClusterObj object
func (cl *ClusterObj) String() string {
	return fmt.Sprintf("Cluster(ClusterName: %s, Admin: %s, Manager: %s, APIVersion: %s, Version: %s, URLS: %s)",
		cl.ClusterName, cl.Admin, cl.Manager, cl.APIVersion, cl.Version, cl.URLS)
}

// Update the ClusterObj object to have some core information
func (cl *ClusterObj) Update() {
	logging.DEBUG("Starting Update")
	cl.setAPIVersion()
	logging.DEBUG("Finished Setting APIVersion")
	var clusterList api.ClusterList
	_ = json.Unmarshal(api.Get(cl.Admin, cl.GetClusterVersionURL()), &clusterList)
	logging.DEBUG(fmt.Sprintf("%s", clusterList))
	cl.Version = clusterList.Items[0].FullVersion
	cl.ClusterName = clusterList.Items[0].Name
}

func (cl *ClusterObj) GetClusterAPIVersionURL() *url.URL {
	if cl.URLS.ClusterVersionURL.String() == "" {
		logging.DEBUG("Setting ClusterApiVersionURL")
		cl.URLS.ClusterApiVersionURL = cl.GetClusterVersionURL()
	}
	logging.DEBUG("Returning ClusterApiVersionURL")
	return cl.URLS.ClusterApiVersionURL
}

// setAPIVersion - an internal method to query the API version in use on the cluster
func (cl *ClusterObj) setAPIVersion() {
	logging.DEBUG("Staring APIVersion")
	cl.APIVersion = "v19"
	logging.DEBUG(fmt.Sprintf("ClusterAPIVersionURL = %s", cl.GetClusterAPIVersionURL()))
	// cl.APIVersion = string(api.Get(cl.Admin, cl.GetClusterAPIVersionURL()))
}

func (cl *ClusterObj) GetAPIVersion() string {
	return cl.APIVersion
}

func (cl *ClusterObj) getClusterHostsURL() *url.URL {
	if cl.URLS.ClusterHostsURL.String() == "" {
		logging.DEBUG("Setting ClusterHostsURL")
		clusterHostURL, _ := url.Parse(fmt.Sprintf("%s/%s/hosts", cl.URLS.ClusterURL, cl.ClusterName))
		logging.DEBUG(fmt.Sprintf("ClusterHostURL = %s", clusterHostURL))
		cl.URLS.ClusterHostsURL = clusterHostURL
	}
	logging.DEBUG("Returning ClusterHostsURL")
	return cl.URLS.ClusterHostsURL
}

func (cl *ClusterObj) GetHosts() *api.HostRefList {
	var clusterHosts *api.HostRefList // []*api.Host
	_ = json.Unmarshal(api.Get(cl.Admin, cl.getClusterHostsURL()), &clusterHosts)
	return clusterHosts
}
