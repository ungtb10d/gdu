package storage

import (
	"github.com/dgraph-io/badger/v3"
	"github.com/ungtb10d/gdu/v5/pkg/fs"
	"github.com/ungtb10d/gdu/v5/pkg/storage/itempb"
	"google.golang.org/protobuf/proto"
)

// StoreDir saves item info into badger DB
func StoreDir(dir fs.Item) {
	options := badger.DefaultOptions("/tmp/badger")
	options.Logger = nil
	db, err := badger.Open(options)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Update(func(txn *badger.Txn) error {
		dirpb := &itempb.Dir{
			Item: &itempb.Item{
				Path: dir.GetPath(),
			},
		}

		data, _ := proto.Marshal(dirpb)
		txn.Set([]byte(dir.GetPath()), data)
		return nil
	})
}
