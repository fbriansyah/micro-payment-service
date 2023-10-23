package port

import "context"

type EventEmitterPort interface {
	// Push event to event broker with custom severity
	Push(ctx context.Context, event string, severity string) error
}
