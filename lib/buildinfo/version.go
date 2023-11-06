package buildinfo

import (
	"flag"
	"fmt"
	"os"
)

var version = flag.Bool("version", false, "show version")

// Version 通过-ldflags '-X'设置版本号
var Version string

// Init 必须在flag.Parse之后调用
func Init() {
	if *version {
		printVersion()
		os.Exit(0)
	}
}

func init() {
	oldUsage := flag.Usage
	flag.Usage = func() {
		printVersion()
		oldUsage()
	}
}

func printVersion() {
	fmt.Fprintf(flag.CommandLine.Output(), "%s\n", Version)
}
