package flags

import "flag"

// Debug shows if debugging messages are eabled or disabled by the user's choice at CLI
var Debug bool
var LessDebug bool

func init() {
	flag.BoolVar(&Debug, "debug", false, "enable debugging messages")
	flag.BoolVar(&LessDebug, "no-debug-verbose", false, "disable verbose debugging (calling location)")
}
