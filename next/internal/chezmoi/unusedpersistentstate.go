package chezmoi

// An unusedPersistentState is a PersistentState that is never used.
type unusedPersistentState struct{}

func newUnusedPersistentState() unusedPersistentState                          { return unusedPersistentState{} }
func (unusedPersistentState) CopyTo(PersistentState) error                     { panic("CopyTo") }
func (unusedPersistentState) Get([]byte, []byte) ([]byte, error)               { panic("Get") }
func (unusedPersistentState) Delete([]byte, []byte) error                      { panic("Delete") }
func (unusedPersistentState) ForEach([]byte, func([]byte, []byte) error) error { panic("ForEach") }
func (unusedPersistentState) Set([]byte, []byte, []byte) error                 { panic("Set") }
