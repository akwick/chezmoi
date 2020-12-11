package chezmoi

// A PersistentState is a persistent state.
type PersistentState interface {
	CopyTo(s PersistentState) error
	Get(bucket, key []byte) ([]byte, error)
	Delete(bucket, key []byte) error
	ForEach(bucket []byte, fn func(k, v []byte) error) error
	Set(bucket, key, value []byte) error
}
