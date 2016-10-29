package search

import (
	"github.com/almighty/almighty-core/app"
	"golang.org/x/net/context"
)

// Repository encapsulates searching of woritems,users,etc
type Repository interface {
	SearchFullText(ctx context.Context, searchStr string, start *int, length *int) ([]*app.WorkItem, uint64, error)
	Validate(ctx context.Context, searchStr string) error
}