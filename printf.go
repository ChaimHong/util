package util

import (
	"log"
)

func Printfn(args ...interface{}) {
	v := ""
	for _, a := range args {
		switch a.(type) {
		case string:
			v += "%s "
		case int8, int16, int64:
			v += "%d "
		default:
			v += "%#v "
		}
	}

	v += "\n"

	log.Printf(v, args...)
}

func DebugPrintf(name string, args ...interface{}) {
	v := ""
	args = append([]interface{}{name}, args...)
	for _, a := range args {
		switch a.(type) {
		case string:
			v += "%s "
		case int8, int16, int64:
			v += "%d "
		default:
			v += "%#v "
		}
	}

	v += "\n"

	log.Printf(v, args...)
}
