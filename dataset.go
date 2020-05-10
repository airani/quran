package quran

import "os"

type datasetFile string

const (
	datasetQuranSimple           datasetFile = "quran-simple-min.xml"
	datasetTranslateFaFooladvand             = "fa.fooladvand.xml"
)

var datasetPath string

func init() {
	datasetPath = os.Getenv("PATH_DATASET")
	if datasetPath == "" {
		datasetPath = "./dataset/"
	}
}
