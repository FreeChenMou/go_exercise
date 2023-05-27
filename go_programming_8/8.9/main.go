// The du4 command computes the disk usage of the files in a directory.
package main

// The du4 variant includes cancellation:
// it terminates quickly when the user hits return.

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var record int

var delay time.Duration = 50 * time.Millisecond //定时输出和RTT

//!-1
//exercise8.9  不管抢占和顺序，目的随要求修改
func main() {
	// Determine the initial directories.
	roots := os.Args[1:]
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse each root of the file tree in parallel.
	for _, root := range roots {
		var n sync.WaitGroup
		n.Add(1)
		fileSizes := make(chan int64)
		go walkDir(root, &n, fileSizes)
		go count(root, &n, fileSizes)
	}
	for record < len(os.Args[1:]) {

	}
}

func count(root string, n *sync.WaitGroup, fileSizes chan int64) {
	var nfiles, nbytes int64
	// Print the results periodically.
	tick := time.Tick(delay)
	go func() {
		for {
			select {
			case size, ok := <-fileSizes:
				// ...
				//!-3
				if !ok {
					break // fileSizes was closed
				}

				nfiles++
				nbytes += size
			case <-tick:
				fmt.Printf("%s file count and sizes: ", root)
				printDiskUsage(nfiles, nbytes)

				if _, ok := <-fileSizes; !ok {
					break
				}
			}
		}
	}()
	n.Wait()
	time.Sleep(2 * delay)
	close(fileSizes)
	record++
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.2f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
//!+4
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()

	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
		//!+4
	}

}

//!-4

var sema = make(chan struct{}, 20) // concurrency-limiting counting semaphore

// dirents returns the entries of directory dir.
//!+5
func dirents(dir string) []os.FileInfo {

	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token

	// ...read directory...
	//!-5

	f, err := os.Open(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	defer f.Close()

	entries, err := f.Readdir(0) // 0 => no limit; read all entries
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		// Don't return: Readdir may return partial results.
	}
	return entries
}
