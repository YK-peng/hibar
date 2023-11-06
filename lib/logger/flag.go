package logger

import (
	"flag"
	"hibar/lib/buildinfo"
)

func LogAllFlags() {
	glogger.Infof("build version: %s", buildinfo.Version)
	glogger.Infof("command-line flags")
	flag.Visit(func(f *flag.Flag) {
		glogger.Infof("  -%s=%q", f.Name, f.Value)
	})
}
