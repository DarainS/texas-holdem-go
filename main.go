package main

import (
	_ "./src/crawler2"
	"github.com/henrylee2cn/pholcus/exec"
)


func main() {
	exec.DefaultRun("web")
}
