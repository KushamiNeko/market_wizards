package headerutils

//////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"net/http"
	"strings"
	"time"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	DefaultMaxAllowAge     = 24
	DefaultMaxAllowAgeUnit = time.Hour
)

//////////////////////////////////////////////////////////////////////////////////////////////////////

func CrossOriginOptions(w http.ResponseWriter) {

	headers := strings.Join([]string{
		"Origin",
		"Authorization",
		"Content-Type",
		"X-Endpoint-API-UserInfo",
		"X-API-Key",
	}, ",")

	w.Header().Add("Access-Control-Allow-Headers", headers)

	//w.Header().Add("Access-Control-Allow-Origin", config.WebFrontendOrigin)

	w.Header().Add("Access-Control-Allow-Credentials", "true")

	w.Header().Add("Vary", "Origin")

	maxAllowAge := DefaultMaxAllowAge * DefaultMaxAllowAgeUnit
	w.Header().Add("Access-Control-Allow-Max-Age", string(int(maxAllowAge.Seconds())))
}

////////////////////////////////////////////////////////////////////////////////////////////////////////

//func AccessControlAllowFronted(w http.ResponseWriter) {
//w.Header().Add("Access-Control-Allow-Origin", config.WebFrontendOrigin)
//}

//////////////////////////////////////////////////////////////////////////////////////////////////////

func AccessControlAllowCredentials(w http.ResponseWriter) {
	w.Header().Add("Access-Control-Allow-Credentials", "true")
}

//////////////////////////////////////////////////////////////////////////////////////////////////////

func AccessControlVaryOrigin(w http.ResponseWriter) {
	w.Header().Add("Vary", "Origin")
}

//////////////////////////////////////////////////////////////////////////////////////////////////////

func AccessControlAllowOrigin(w http.ResponseWriter, origin string) {
	w.Header().Add("Access-Control-Allow-Origin", origin)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////

func AccessControl(w http.ResponseWriter) {
	//w.Header().Add("Access-Control-Allow-Origin", config.WebFrontendOrigin)
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Vary", "Origin")
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
