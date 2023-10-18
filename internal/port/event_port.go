package port

import "context"

type EventEmitter interface {
	// Push event to event broker with custom severity
	Push(ctx context.Context, event string, severity string) error
}
