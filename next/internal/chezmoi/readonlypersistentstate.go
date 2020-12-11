package chezmoi

import "os"

// A readOnlyPersistentState wraps a PeristentState but returns an error on any
// write.
type readOnlyPersistentState struct {
	s PersistentState
}

func newReadOnlyPersistentState(s PersistentState) PersistentState {
	return &readOnlyPersistentState{
		s: s,
	}
}

// CopyTo implements PersistentState.CopyTo.
func (s *readOnlyPersistentState) CopyTo(p PersistentState) error {
	return s.s.CopyTo(p)
}

// Get implements PersistentState.Get.
func (s *readOnlyPersistentState) Get(bucket, key []byte) ([]byte, error) {
	return s.s.Get(bucket, key)
}

// Delete implements PersistentState.Delete.
func (s *readOnlyPersistentState) Delete(bucket, key []byte) error {
	return os.ErrPermission
}

// ForEach implements PersistentState.ForEach.
func (s *readOnlyPersistentState) ForEach(bucket []byte, fn func(k, v []byte) error) error {
	return s.s.ForEach(bucket, fn)
}

// Set implements PersistentState.Set.
func (s *readOnlyPersistentState) Set(bucket, key, value []byte) error {
	return os.ErrPermission
}
