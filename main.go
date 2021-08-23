package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/vpiyush/getir-go-app/apis"
	"github.com/vpiyush/getir-go-app/common"
	"net/http"
	"os"
)

//init sets up logging
func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	// register endpoints
	log.Debug("Registering api endpoint /api/v1/records")
	http.HandleFunc("/api/v1/records", apis.GetRecords)
	log.Debug("Registering api endpoint /api/v1/pair")
	http.HandleFunc("/api/v1/pair", apis.HandlePair)
	serveruri := common.Cfg.Server.Host + ":" + common.Cfg.Server.Port
	log.Debug("listening on ", serveruri)
	// bind and serve http
	http.ListenAndServe(serveruri, nil)
}
