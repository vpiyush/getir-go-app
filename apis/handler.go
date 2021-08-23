// Packave api provides functions which are used to server the http Endpoints
package apis

import (
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/vpiyush/getir-go-app/daos"
	"github.com/vpiyush/getir-go-app/models"
	"github.com/vpiyush/getir-go-app/services"
	"net/http"
	"os"
)

// ErrorRespnse
type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

// RecordResponse to be sent back on GetRecords request
type RecordResponse struct {
	Code    int             `json:"code"`
	Msg     string          `json:"msg"`
	Records []models.Record `json:"records"`
}

//Request represnting GetRecords input structure, includes validation
// and json tags
type Request struct {
	StartDate string `validate:"required" json:"startdate"`
	EndDate   string `validate:"required" json:"enddate"`
	MinCount  int    `validate:"required" json:"mincount"`
	MaxCount  int    `validate:"required" json:"maxcount"`
}

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{})
}

// buildErrorResponse builds error based on error message and code
func buildErrorResponse(response http.ResponseWriter, err error, code int) {
	var errRsp = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   code,
	}
	response.WriteHeader(code)
	json.NewEncoder(response).Encode(errRsp)
}

// GetRecords Rest Endpoint to fetches records from database
func GetRecords(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	log.Debug("Entering GetRecords")
	var req Request
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		buildErrorResponse(response, err, http.StatusBadRequest)
		return
	}
	sDate, eDate, err := validateRequest(&req)
	if err != nil {
		log.Debug("Validation failed for GetRecords , Error: ", err)
		buildErrorResponse(response, err, http.StatusBadRequest)
		return
	}
	log.Debug("Request ", req)
	s := services.NewRecordService(daos.NewRecordDAO())
	records, err := s.Find(sDate, eDate, req.MinCount, req.MaxCount)
	var res RecordResponse
	if err != nil {
		res.Code = -1
		res.Msg = "Records not found"
		response.WriteHeader(http.StatusInternalServerError)
	} else {
		res.Code = 0
		res.Msg = "Success"
	}
	res.Records = records
	log.Debug("Exiting GetRecords, Response ", res)
	json.NewEncoder(response).Encode(res)
}

// InsertPairs creates a key-value entry in in-memory database
func InsertPair(response http.ResponseWriter, request *http.Request) {
	log.Debug("Entering InsertPair ")
	response.Header().Set("content-type", "application/json")
	var pair models.Pair
	if err := json.NewDecoder(request.Body).Decode(&pair); err != nil {
		buildErrorResponse(response, err, http.StatusBadRequest)
		return
	}
	if err := validatePair(&pair); err != nil {
		buildErrorResponse(response, err, http.StatusBadRequest)
		return
	}
	s := services.NewPairService(daos.NewPairDAO())
	res, err := s.Insert(pair.Key, pair.Value)
	if err != nil {
		log.Debug("InsertPair failed, Error: ", err)
		buildErrorResponse(response, err, http.StatusForbidden)
	} else {
		log.Debug("Exiting InsertPair, Response: ", res)
		json.NewEncoder(response).Encode(res)
	}
}

// GetPair fetches a key-value pair from in-memory database
func GetPair(response http.ResponseWriter, request *http.Request) {
	log.Debug("Entering GetPair")
	response.Header().Set("content-type", "application/json")
	var pair models.Pair
	pair.Key = request.FormValue("key")
	log.Debug("Recieved key ", pair.Key)
	s := services.NewPairService(daos.NewPairDAO())
	val, ok := s.Get(pair.Key)
	if !ok {
		err := errors.New("Key not Found")
		buildErrorResponse(response, err, http.StatusNotFound)
	} else {
		pair.Value = val
		log.Debug("Exiting GetPair, Res: ", pair)
		json.NewEncoder(response).Encode(pair)
	}
}

//Handle Pair common REST endpoint for GetPair and InsertPair
func HandlePair(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		GetPair(response, request)
	case "POST":
		InsertPair(response, request)
	default:
		log.Debug("Method not handled, Method: ", request.Method)
		err := errors.New("EndPoint not found")
		buildErrorResponse(response, err, http.StatusBadRequest)
	}
}
