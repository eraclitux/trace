package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

var recursiveFlag = flag.Bool("r", false, "process subdirs")
var verboseFlag = flag.Bool("v", false, "verbose output")
var dryRunFlag = flag.Bool("d", false, "do not make any modification, just print filenames to process")

func init() {
	log.SetFlags(0)
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		// FIXME read from stdin
		log.Fatalln("specify at least a file or a path with -r")
	}
	wg := sync.WaitGroup{}
	jobsChan := make(chan string)
	errChan := make(chan struct{}, 1)
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			for filePath := range jobsChan {
				if isGo(filePath) {
					err := clean(filePath)
					if err != nil {
						log.Printf("error cleaning file '%s': %s\n", filePath, err)
						select {
						case errChan <- struct{}{}:
						default:
							// error condition already signaled,
							// move along
						}
					}
				}
			}
			wg.Done()
		}()
	}
	if *recursiveFlag {
		// recursively walk filesystem
		// with path as root
		basePath := args[0]
		err := filepath.Walk(
			basePath,
			func(path string, info os.FileInfo, err error) error {
				jobsChan <- path
				return err
			},
		)
		if err != nil {
			log.Fatalln("traversing filesystem:", err)
		}
	} else {
		// arguments are filenames
		for _, path := range args {
			jobsChan <- path
		}
	}
	close(jobsChan)
	wg.Wait()
	select {
	case <-errChan:
		os.Exit(1)
	default:
		// no errors
	}
}

func clean(filePath string) error {
	if *dryRunFlag {
		fmt.Printf("'%s' to be processed\n", filePath)
		return nil
	}
	info, err := os.Stat(filePath)
	if err != nil {
		return err
	}
	dd, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	lines := bytes.Split(dd, []byte("\n"))
	// FIXME find a better way like in place
	// file modification. Read gofmt.
	cleanedLines := make([][]byte, 0, len(lines))
	cleaned := false
	for _, l := range lines {
		if bytes.Contains(l, []byte("trace.Pr")) || bytes.Contains(l, []byte("eraclitux/trace")) {
			cleaned = true
			continue
		}
		cleanedLines = append(cleanedLines, l)
	}
	dd = bytes.Join(cleanedLines, []byte("\n"))
	err = ioutil.WriteFile(filePath, dd, info.Mode().Perm())
	if err != nil {
		return err
	}
	if *verboseFlag && cleaned {
		fmt.Printf("'%s' cleaned\n", filePath)
	}
	return nil
}

func isGo(filePath string) bool {
	return strings.HasSuffix(filePath, ".go")
}
