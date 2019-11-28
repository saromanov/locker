package lock

// Locker defines main interface for locks
type Locker interface {
	Lock() bool
	Unlock()
}
