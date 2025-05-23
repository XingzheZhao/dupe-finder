package scan

import (
	"context"
	"crypto/sha256"
	"io"
	"os"
	"sync"
)

// FindDuplicates spins workers that hash paths from fileCh and returns a map
// hash -> []FileInfo (only entries with >1 file are kept)
func FindDuplicates(
    ctx context.Context,
    fileCh <-chan string, // read-only
    workers int,
) (Result, error) {
    // pair is what each worker emits.
    type pair struct {
        Hash string
        File FileInfo
        Err error
    }

    out := make(chan pair)

    // -------------------------------------------------------
    // Worker function 
    // -------------------------------------------------------
    var wg sync.WaitGroup
    hashWorker := func() {
        defer wg.Done()

        // re-use one buffer per worker to avoid allocations
        buf := make([]byte, 32*1024)

        for p := range fileCh {
            h := sha256.New()

            f, err := os.Open(p)
            if err != nil {
                out <- pair{Err: err}
                continue
            }
            f, err := f.Seek(0, io.SeekEnd)
            _, _ = f.Seek(0, io.SeekStart)

        }
    }

}
