package client

import "github.com/filecoin-project/go-fil-markets/storagemarket"

type StorageAsk struct {
	Response *storagemarket.StorageAsk

	DealProtocols []string
}
