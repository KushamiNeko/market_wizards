package config

//////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"os"
	"strings"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	CacheStoreKeyLens = 9

	KeyLengthMin     = 16
	KeyLengthShort   = 32
	KeyLengthDefault = 64
	KeyLengthStrong  = 128
	KeyLengthMax     = 256

	//KeyLengthMin     = 32
	//KeyLengthShort   = 64
	//KeyLengthDefault = 128
	//KeyLengthStrong  = 256
	//KeyLengthMax     = 512

	//NamespaceIBD         = "--IBD-Checkup--"
	NamespaceTransaction = "--Transaction--"

	StorageNamespaceIBDs   = "IBDCheckups"
	StorageNamespaceCharts = "Charts"

	NamespaceUser = "--User--"
	KindUser      = "--user--"

	ImageQuality = 100
)

//////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	//WebFrontendOrigin string
	AppEngineOrigin string

	ProjectID     string
	ProjectBucket string

	UIDSecret string

	TLSConnection bool
)

//////////////////////////////////////////////////////////////////////////////////////////////////////

func Init() {
	//WebFrontendOrigin = os.Getenv("GCP_WEB_FRONTEND_ORIGIN")
	//if WebFrontendOrigin == "" {
	//panic("NO ENV VARIABLE for GCP_WEB_FRONTEND_ORIGIN")
	//}

	AppEngineOrigin = os.Getenv("GCP_APP_ENGINE_ORIGIN")
	if AppEngineOrigin == "" {
		panic("NO ENV VARIABLE for GCP_APP_ENGINE_ORIGIN")
	}

	ProjectID = os.Getenv("GCP_PROJECT_ID")
	if ProjectID == "" {
		panic("NO ENV VARIABLE for GCP_PROJECT_ID")
	}

	ProjectBucket = os.Getenv("GCP_STORAGE_BUCKET")
	if ProjectBucket == "" {
		panic("NO ENV VARIABLE for GCP_STORAGE_BUCKET")
	}

	UIDSecret = os.Getenv("GCP_UID_SECRET")
	if UIDSecret == "" {
		panic("NO ENV VARIABLE for GCP_UID_SECRET")
	}

	TLSConnection = strings.Contains(os.Getenv("GCP_TLS_CONNECTION"), "true")
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
