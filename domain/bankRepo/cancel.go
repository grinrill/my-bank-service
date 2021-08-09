package bankRepo

// CancelFunc Should be returned by initializers
// to close connections
type CancelFunc func() error