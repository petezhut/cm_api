package main

// TODO: Get the URL generation sorted.  This is ridiculous.

import (
	"fmt"

	"github.com/petezhut/cm_api/go/api/cluster"
	"github.com/petezhut/cm_api/go/logging"
)

func main() {
	C := cluster.New("avatar-jmcfarland-1.gce.cloudera.com")
	//C.Update()
	logging.INFO(fmt.Sprintf("%s", C))
	logging.INFO(fmt.Sprintf("%s", C.GetAPIVersion()))
	logging.DEBUG(fmt.Sprintf("%s", C.GetHosts()))
}
