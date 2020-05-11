package quran

import "os"

type DatasetFile string

const (
	DatasetQuranSimple           DatasetFile = "quran-simple-min.xml"
	DatasetTranslateFaFooladvand             = "fa.fooladvand.xml"
)

var datasetPath string

func init() {
	datasetPath = os.Getenv("PATH_DATASET")
	if datasetPath == "" {
		datasetPath = "./dataset/"
	}
}
