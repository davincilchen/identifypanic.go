// You can edit this code!
// Click here and start typing.
package main

import "fmt"
import "runtime"
import "strings"

func identifyPanic() string {
	var name, file string
	var line int
	var pc [16]uintptr
	
	n := runtime.Callers(3, pc[:])
	for _, pc := range pc[:n] {
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}
		file, line = fn.FileLine(pc)
		name = fn.Name()
		if !strings.HasPrefix(name, "runtime.") {
			break
		}
	}
	
	switch {
	case name != "":
		return fmt.Sprintf("%v:%v", name, line)
	case file != "":
		return fmt.Sprintf("%v:%v", file, line)
	}
	
	return fmt.Sprintf("pc:%x", pc)
}

func recoverPanic() {
	r := recover()
	if r == nil {
		return
	}
	fmt.Println(identifyPanic())
}

func createPanic() {
	var s *string
	fmt.Println(*s)
}

func main() {
	defer recoverPanic()
	createPanic()
}
