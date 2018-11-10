package config

import (
	"image/color"
	"math"
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

	StatisticBase = 20
	StatisticSMA  = 5
)

//////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	ChartMultiplier = 0.65

	ChartWidth  = 1920.0 * ChartMultiplier
	ChartHeight = 1080.0 * ChartMultiplier

	ChartFontSizeL = 25.0 * ChartMultiplier
	ChartFontSizeM = 25.0 * ChartMultiplier
	ChartFontSizeS = 25.0 * ChartMultiplier

	ChartPointRadius = 9.0 * ChartMultiplier
	ChartBarWidth    = 20.0 * ChartMultiplier
	ChartLineWidth   = 2.0 * ChartMultiplier

	ChartFont = "Helvetica"
	//ChartFont          = "Helvetica-Oblique"
	ChartLabelRotation = -90.0 * math.Pi / 180.0

	ChartLabelPaddingX       = 15.0 * ChartMultiplier
	ChartLegendPaddingYRatio = 1.15

	ChartBarPaddingXRatio = 1.035

	ChartDataUrlFormat = "data:image/png;base64,%s"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	WinnerRGBA = color.NRGBA{R: 3, G: 169, B: 255, A: 127}
	LoserRGBA  = color.NRGBA{R: 255, G: 50, B: 70, A: 127}

	SMARGBA  = color.NRGBA{R: 0, G: 150, B: 136, A: 127}
	InfoRGBA = color.NRGBA{R: 255, G: 111, B: 0, A: 127}
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
