package deadlock

import "errors"

// ErrNotReady indicates that the GC session is unavailable or was interrupted.
var ErrNotReady = errors.New("the Deadlock game coordinator session is not ready")
