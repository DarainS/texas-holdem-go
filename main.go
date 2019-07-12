package main

import (
	_ "./src/crawler2"
	_ "./src/crawler"
	"github.com/henrylee2cn/pholcus/exec"
)


func main() {
	exec.DefaultRun("web")
}
