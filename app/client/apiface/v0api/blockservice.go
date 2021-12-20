package v0api

import (
	"context"
	"io"

	"github.com/ipfs/go-cid"
	ipld "github.com/ipfs/go-ipld-format"
)

type IDagService interface {
	// Rule[perm:read]
	DAGGetNode(ctx context.Context, ref string) (interface{}, error)
	// Rule[perm:read]
	DAGGetFileSize(ctx context.Context, c cid.Cid) (uint64, error)
	// Rule[perm:read]
	DAGCat(ctx context.Context, c cid.Cid) (io.Reader, error)
	// Rule[perm:write]
	DAGImportData(ctx context.Context, data io.Reader) (ipld.Node, error)
}
