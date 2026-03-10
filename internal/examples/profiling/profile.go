package profiling

import (
	"net/http"
	_ "net/http/pprof"

	log "github.com/sirupsen/logrus"
)

func GenerateReport() {

	log.Info(http.ListenAndServe("localhost:6060", nil))

}
