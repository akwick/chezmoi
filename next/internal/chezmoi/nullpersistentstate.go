package chezmoi

// A NullPersistentState is an empty PersistentState that returns the zero value
// for all reads and silently consumes all writes.
type NullPersistentState struct{}

// NewNullPersistentState returns a new NullPersistentState.
func NewNullPersistentState() NullPersistentState {
	return NullPersistentState{}
}

// CopyTo implements PersistentState.CopyTo.
func (NullPersistentState) CopyTo(s PersistentState) error {
	return nil
}

// Get implements PersistentState.Get.
func (NullPersistentState) Get(bucket, key []byte) ([]byte, error) {
	return nil, nil
}

// Delete implements PersistentState.Delete.
func (NullPersistentState) Delete(bucket, key []byte) error {
	return nil
}

// ForEach implements PersistentState.ForEach.
func (NullPersistentState) ForEach(bucket []byte, fn func(k, v []byte) error) error {
	return nil
}

// Set implements PersistentState.Set.
func (NullPersistentState) Set(bucket, key, value []byte) error {
	return nil
}
