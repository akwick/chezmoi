package chezmoi

// A nullPersistentState is an empty PersistentState that returns the zero value
// for all reads and silently consumes all writes.
type nullPersistentState struct{}

func (nullPersistentState) CopyTo(s PersistentState) error                          { return nil }
func (nullPersistentState) Get(bucket, key []byte) ([]byte, error)                  { return nil, nil }
func (nullPersistentState) Delete(bucket, key []byte) error                         { return nil }
func (nullPersistentState) ForEach(bucket []byte, fn func(k, v []byte) error) error { return nil }
func (nullPersistentState) OpenOrCreate() error                                     { return nil }
func (nullPersistentState) Set(bucket, key, value []byte) error                     { return nil }
