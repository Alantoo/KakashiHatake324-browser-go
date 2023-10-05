package browsergo

import (
	"log"
	"runtime/debug"
)

func CatchUnhandledError(position string) {
	if a := recover(); a != nil {
		stack := debug.Stack()
		log.Println(string(stack))
	}
}
