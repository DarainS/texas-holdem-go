package main

import (
	//"./src/crawer"
	"github.com/henrylee2cn/pholcus/exec"
	_ "./src/crawer"
)

func main() {
	//crawer.BuildRequestTest()
	exec.DefaultRun("web")
}
