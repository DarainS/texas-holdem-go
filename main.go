package main

import (
	_ "./src/crawler2"
	"github.com/henrylee2cn/pholcus/exec"
)

func main() {
	//crawer.BuildRequestTest()
	exec.DefaultRun("web")
}
