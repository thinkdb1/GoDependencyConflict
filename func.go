package GoDependencyConflict

import "context"

func f1(ctx context.Context, c1 string) string {
	return c1
}
