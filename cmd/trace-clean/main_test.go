package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

const (
	checkMark    = "\u2713"
	crossMark    = "\u2717"
	testDataPath = "./testdata"
)

var noCheckOutFlag = flag.Bool("nocheckout", false, "leave modified files in place")

func TestMain(t *testing.T) {
	os.Args = append(os.Args, "-r", testDataPath)
	main()
	t.Logf("Test cmd on test data.")
	err := filepath.Walk(testDataPath, makeWalker(t))
	if err != nil {
		t.Fatalf("Error processing testdata: %s.", err)
	}
	if *noCheckOutFlag {
		return
	}
	err = cleanTestData(testDataPath)
	if err != nil {
		t.Fatalf("Error cleaning testdata: %s.", err)
	}
}

func makeWalker(t *testing.T) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || isGolden(path) {
			return nil
		}
		ok, err := compareGoldenFile(path)
		if err != nil {
			return err
		}
		if !ok {
			t.Errorf("\t %s '%s' doesn't match golden file.", crossMark, path)
		} else {
			t.Logf("\t %s '%s' matches golden file.", checkMark, path)
		}
		return nil
	}
}

func compareGoldenFile(path string) (bool, error) {
	modified, err := ioutil.ReadFile(path)
	if err != nil {
		return false, err
	}
	golden, err := ioutil.ReadFile(path + ".golden")
	if err != nil {
		return false, err
	}
	return bytes.Equal(modified, golden), nil
}

func isGolden(filePath string) bool {
	return strings.HasSuffix(filePath, ".golden")
}

func cleanTestData(path string) error {
	cmd := exec.Command("git", "checkout", "--", path)
	return cmd.Run()
}
