package urls

import (
	"fmt"
	"net/url"

	"github.com/petezhut/cm_api/go/api/cluster"
	"github.com/petezhut/cm_api/go/logging"
)

func (thisCluster *cluster.Cluster) GetClusterVersionURL() *url.URL {
	if thisCluster.URLS.ClusterVersionURL.String() == "" {
		logging.DEBUG("Setting ClusterVersionURL")
		versionURL, _ := url.Parse(fmt.Sprintf(CLUSTER_URL, thisCluster.Manager.Hostname, thisCluster.Manager.Port, thisCluster.ApiVersion))
		thisCluster.URLS.ClusterVersionURL = versionURL
	}
	logging.DEBUG("Returning ClusterVersionURL")
	return thisCluster.URLS.ClusterVersionURL

}

func (thisCluster *Cluster) getClusterURL() *url.URL {
	if thisCluster.URLS.ClusterURL.String() == "" {
		logging.DEBUG("Setting ClusterURL")
		clusterURL, _ := url.Parse(fmt.Sprintf("%s/%s", thisCluster.getClusterVersionURL(), thisCluster.ClusterName))
		thisCluster.URLS.ClusterURL = clusterURL
	}
	logging.DEBUG("Returning ClusterURL")
	return thisCluster.URLS.ClusterURL
}

func (thisCluster *Cluster) getClusterHostsURL() *url.URL {
	if thisCluster.URLS.ClusterHostsURL.String() == "" {
		logging.DEBUG("Setting ClusterHostsURL")
		clusterHostURL, _ := url.Parse(fmt.Sprintf("%s/hosts", thisCluster.getClusterURL()))
		thisCluster.URLS.ClusterHostsURL = clusterHostURL
	}
	logging.DEBUG("Returning ClusterHostsURL")
	return thisCluster.URLS.ClusterHostsURL
}

func (thisCluster *Cluster) getClusterAPIVersionURL() *url.URL {
	if thisCluster.URLS.ClusterVersionURL.String() == "" {
		logging.DEBUG("Setting ClusterApiVersionURL")
		apiVersionURL, _ := url.Parse(fmt.Sprintf(VERSION_URL, thisCluster.Manager.Hostname, thisCluster.Manager.Port))
		thisCluster.URLS.ClusterApiVersionURL = apiVersionURL
	}
	logging.DEBUG("Returning ClusterApiVersionURL")
	return thisCluster.URLS.ClusterApiVersionURL
}
