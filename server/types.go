package server

type Sorter interface {
	Sort(paths [][]string) ([]string, error)
}

type Server struct {
	sorter Sorter
}

type status string

const (
	StatusSuccess = "success"
	StatusFailure = "failure"
)

type RequestSort struct {
	Paths [][]string `json:"paths"`
}

type ResponseSort struct {
	Path    []string `json:"path,omitempty"`
	Message *string  `json:"message,omitempty"`
	Status  status   `json:"status"`
}
