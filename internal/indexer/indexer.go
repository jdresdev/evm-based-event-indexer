package indexer

import "context"

type Indexer interface {
	Run(ctx context.Context) error
}
