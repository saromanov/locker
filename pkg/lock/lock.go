package lock

// Locker defines main interface for locks
type Locker interface {
	Lock() bool
	Unlock()
}

// Config provides definition for common config
// for all type of locks
type Config struct {
	Address string
}
