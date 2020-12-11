package chezmoi

var (
	bucketTombstone = make(map[string][]byte)
	valueTombstone  = make([]byte, 0)
)

// A dryRunPersistentState wraps a PersistentState and drops all writes but
// records that they occurred.
//
// FIXME mock writes (e.g. writes should not affect the underlying
// PersistentState but subsequent reads should return as if the write occurred).
type dryRunPersistentState struct {
	s        PersistentState
	diff     map[string]map[string][]byte
	modified bool
}

// newDryRunPersistentState returns a new dryRunPersistentState that wraps s.
func newDryRunPersistentState(s PersistentState) *dryRunPersistentState {
	return &dryRunPersistentState{
		s:    s,
		diff: make(map[string]map[string][]byte),
	}
}

// Get implements PersistentState.Get.
func (s *dryRunPersistentState) Get(bucket, key []byte) ([]byte, error) {
	return s.s.Get(bucket, key)
}

// Delete implements PersistentState.Delete.
func (s *dryRunPersistentState) Delete(bucket, key []byte) error {
	s.modified = true
	return nil
}

// ForEach implements PersistentState.ForEach.
func (s *dryRunPersistentState) ForEach(bucket []byte, fn func(k, v []byte) error) error {
	return s.s.ForEach(bucket, fn)
}

// Set implements PersistentState.Set.
func (s *dryRunPersistentState) Set(bucket, key, value []byte) error {
	s.modified = true
	// FIXME do we need to remember that the value has been set?
	return nil
}
