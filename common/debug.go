package common

import (
	"fmt"
	"runtime"

	"time"

	"github.com/cg-/space-a/flags"
)

// Debug prints debugging messages if enabled
func Debug(s string) {
	if flags.Debug {
		t := time.Now()
		timeStr := fmt.Sprintf("%d:%d.%d", t.Hour(), t.Minute(), t.Second())
		fileLoc := ""
		if !flags.LessDebug {
			_, file, no, ok := runtime.Caller(1)
			if ok {
				fileLoc = fmt.Sprintf("{%s@%d}", file, no)
			}
		}
		fmt.Printf("%s%s debug: %s\n", timeStr, fileLoc, s)
	}
}
