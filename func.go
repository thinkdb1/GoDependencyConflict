package GoDependencyConflict

import (
	"context"
	"fmt"
)

func Func1(ctx context.Context, c1 string) string {
	return c1
}
func Func2(ctx context.Context, c1 int) string {
	return fmt.Sprint(c1)
}
