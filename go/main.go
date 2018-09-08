package main

import (
	"fmt"

	"github.com/petezhut/cm_api/go/api/cluster"
	"github.com/petezhut/cm_api/go/logging"
)

func main() {
	logging.DEBUG("1")
	C := cluster.NewCluster()
	logging.DEBUG("2")
	C.Update()
	logging.DEBUG("3")
	logging.INFO(fmt.Sprintf("%s", C))
	logging.DEBUG(C.GetClusterHosts())
}
