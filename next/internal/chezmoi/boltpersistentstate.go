package chezmoi

import (
	"path/filepath"

	vfs "github.com/twpayne/go-vfs"
	"go.etcd.io/bbolt"
)

// A BoltPersistentState is a state persisted with bolt.
type BoltPersistentState struct {
	db *bbolt.DB
}

// NewBoltPersistentState returns a new BoltPersistentState. options is mutated.
//
//nolint:interfacer
func NewBoltPersistentState(fs vfs.FS, path string, options *bbolt.Options) (*BoltPersistentState, error) {
	if err := vfs.MkdirAll(fs, filepath.Dir(path), 0o777); err != nil {
		return nil, err
	}
	if options == nil {
		options = &bbolt.Options{}
	}
	options.OpenFile = fs.OpenFile
	db, err := bbolt.Open(path, 0o600, options)
	if err != nil {
		return nil, err
	}
	return &BoltPersistentState{
		db: db,
	}, nil
}

// Close closes b.
func (b *BoltPersistentState) Close() error {
	return b.db.Close()
}

// CopyTo copies b to p.
func (b *BoltPersistentState) CopyTo(p PersistentState) error {
	return b.db.View(func(tx *bbolt.Tx) error {
		return tx.ForEach(func(bucket []byte, b *bbolt.Bucket) error {
			return b.ForEach(func(key, value []byte) error {
				return p.Set(bucket, key, value)
			})
		})
	})
}

// Delete deletes the value associate with key in bucket. If bucket or key does
// not exist then Delete does nothing.
func (b *BoltPersistentState) Delete(bucket, key []byte) error {
	return b.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucket)
		if b == nil {
			return nil
		}
		return b.Delete(key)
	})
}

// Get returns the value associated with key in bucket.
func (b *BoltPersistentState) Get(bucket, key []byte) ([]byte, error) {
	var value []byte
	if err := b.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucket)
		if b == nil {
			return nil
		}
		if v := b.Get(key); v != nil {
			value = make([]byte, len(v))
			copy(value, v)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return value, nil
}

// ForEach calls fn for each key, value pair in bucket.
func (b *BoltPersistentState) ForEach(bucket []byte, fn func(k, v []byte) error) error {
	return b.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucket)
		if b == nil {
			return nil
		}
		return b.ForEach(fn)
	})
}

// Set sets the value associated with key in bucket. bucket will be created if
// it does not already exist.
func (b *BoltPersistentState) Set(bucket, key, value []byte) error {
	return b.db.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			return err
		}
		return b.Put(key, value)
	})
}
