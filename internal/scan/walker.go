package scan

import (
	"context"
	"os"
	"path/filepath"
)

func WalkFiles(ctx context.Context, root string, outCh chan<- string) error {
    defer close(outCh)

    return filepath.WalkDir(root, func(p string, d os.DirEntry, err error) error {
        if err != nil { // unreadable dir/file
            return nil // keep walking
        }
        if !d.Type().IsRegular() {
            return nil
        }
        select {
        case <-ctx.Done():
            return ctx.Err()
        case outCh <- p:
        }
        return nil
    })
}
