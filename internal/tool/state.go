package tool

var (
	SaveFolder string = ".fdlr/"
)

type State struct {
	URL            string
	DownloadRanges []DownloadRange
}

type DownloadRange struct {
	URL       string
	Path      string
	RangeFrom int64
	RangeTo   int64
}
