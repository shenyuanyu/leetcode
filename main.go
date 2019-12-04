package main

import (
	"flag"
	"fmt"
	"strings"
)

var replacerStr = flag.String("r", `[,{,],}`, "set the replacer string, split by comma")

func main() {
	flag.Parse()

	if !strings.Contains(*replacerStr, ",") {
		return
	}

	s := ""
	if len(flag.Args()) > 0 {
		s = flag.Arg(0)
	}

	replacer := strings.NewReplacer(strings.Split(*replacerStr, ",")...)

	fmt.Println(replacer.Replace(s))
}
