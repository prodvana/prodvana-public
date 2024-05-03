package cmdutil

import (
	"context"
	"io"
	"time"

	"github.com/gosuri/uilive"
)

func RenderWatch(ctx context.Context, render func(ctx context.Context, writer io.Writer) (bool, error)) error {
	writer := uilive.New()
	// uilive.Writer automatically flushes every millisecond.
	// So, whatever is in the buffer at flush will be written to screen, overwriting whatever was previously written.
	// If a render iteration writes multiple lines with at least 1ms between them (which can happen randomly for a lot of lines),
	// partial output of a render will be sent to screen.
	//
	// Instead, make each iteration of render write to a buffer and then flush that buffer to the screen.
	// Disable polling by setting the refresh interval to a large value.
	writer.RefreshInterval = 5000 * time.Hour
	writer.Start()
	defer writer.Stop()
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		stop, err := render(ctx, writer)
		if err != nil {
			return err
		}
		writer.Flush()
		if stop {
			return nil
		}
		select {
		case <-ticker.C:
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
