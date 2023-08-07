package indices

import "github.com/rahularcota/address-book/dtos"

type IIndex interface {
	Search(req *dtos.SearchIndexRequest) ([]int, error)
	AddEntity(req *dtos.AddToIndexRequest) error
}
