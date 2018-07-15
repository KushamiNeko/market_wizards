package config

import (
	"os"
	"strings"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////

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

	NamespaceIBD         = "--ibd-checkup--"
	NamespaceMarketSmith = "--marketsmith--"
	NamespaceTransaction = "--transaction--"
	NamespaceAdmin       = "--admin--"

	NamespaceWatchList    = "--watch-list--"
	NamespacePostAnalysis = "--post-analysis--"
	NamespaceChartsStudy  = "--charts-study--"
	NamespaceExperience   = "--experience--"

	CollectionUser = "--user--"

	//KindUser      = "--user--"

	//StorageNamespaceIBDs   = "IBDCheckups"
	//StorageNamespaceCharts = "Charts"

	ImageQuality = 100

	MongoURL = "mongodb://localhost:27017"

	//UIDSecret = "426592B26163B15AD5AB64838EA4B7A91C57BC9AB33A92CD24F92295EB49DFE06F38D6EB1350512CCE023CE73FAB2228008A28C3122323AE6F36D35841058B73"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	PriceInterval         = 50
	IntervalFormat        = "%v ~ %v"
	PriceIntervalFormat   = "$%v ~ $%v"
	PercentIntervalFormat = "%v%% ~ %v%%"

	StyleFormat = "color: %v; opacity: %v;"

	WinnerColor = "#03A9F4"
	LoserColor  = "#F44336"

	WinnerOpacity = 0.5
	LoserOpacity  = 0.5
)

//////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	//WatchListPosition []float64 = []float64{0.05, 0.08, 0.1, 0.12, 0.13, 0.15, 0.2, 0.25, 0.3, 0.35}
	WatchListPosition []float64 = []float64{0.05, 0.075, 0.1, 0.125, 0.15, 0.2, 0.25, 0.3}
)

//////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	NullValue = "N/A"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	//WebFrontendOrigin string
	//AppEngineOrigin string

	//ProjectID     string
	//ProjectBucket string

	UIDSecret string

	TLSConnection bool
)

//////////////////////////////////////////////////////////////////////////////////////////////////////

func Init() {
	//WebFrontendOrigin = os.Getenv("GCP_WEB_FRONTEND_ORIGIN")
	//if WebFrontendOrigin == "" {
	//panic("NO ENV VARIABLE for GCP_WEB_FRONTEND_ORIGIN")
	//}

	//AppEngineOrigin = os.Getenv("GCP_APP_ENGINE_ORIGIN")
	//if AppEngineOrigin == "" {
	//panic("NO ENV VARIABLE for GCP_APP_ENGINE_ORIGIN")
	//}

	//ProjectID = os.Getenv("GCP_PROJECT_ID")
	//if ProjectID == "" {
	//panic("NO ENV VARIABLE for GCP_PROJECT_ID")
	//}

	//ProjectBucket = os.Getenv("GCP_STORAGE_BUCKET")
	//if ProjectBucket == "" {
	//panic("NO ENV VARIABLE for GCP_STORAGE_BUCKET")
	//}

	UIDSecret = os.Getenv("GCP_UID_SECRET")
	if UIDSecret == "" {
		panic("NO ENV VARIABLE for GCP_UID_SECRET")
	}

	TLSConnection = strings.Contains(os.Getenv("GCP_TLS_CONNECTION"), "true")
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
