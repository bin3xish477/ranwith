package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	red    = "\u001b[91m"
	green  = "\u001b[32m"
	blue   = "\u001b[94m"
	yellow = "\u001b[33m"
	end    = "\u001b[0m"
	bold   = "\u001b[1m"
	underL = "\u001b[4m"

	procDir = "/proc"
)

type Proc struct {
	Name    string
	Pid     int
	CmdLine string
}

func (p *Proc) Pprint() {
	fmt.Printf(
		"[%s%sName%s=%s, %s%sPid%s=%d, %s%sCmdLine%s='%s']\n",
		red, underL, end, p.Name, green, underL, end, p.Pid, blue, underL, end, p.CmdLine,
	)
}

var (
	procList []*Proc
)

func isProcess(f string) bool {
	if _, err := strconv.Atoi(f); err == nil {
		return true
	}
	return false
}

func newProc(f string) *Proc {
	d, _ := os.ReadFile(fmt.Sprintf("%s/%s/cmdline", procDir, f))
	cmd := string(d)
	
	d, _ = os.ReadFile(fmt.Sprintf("%s/%s/comm", procDir, f))
	name := strings.TrimRight(string(d), "\n")
	
	pid, _ := strconv.Atoi(f)

	return &Proc{
		Name:    name,
		Pid:     pid,
		CmdLine: cmd,
	}
}

func createProcessList() {
	procFiles, err := ioutil.ReadDir(procDir)
	if err != nil {
		fmt.Printf("not able to list files in %s", procDir)
		return
	}
	for _, file := range procFiles {
		fileName := file.Name()
		if isProcess(fileName) {
			proc := newProc(fileName)
			procList = append(procList, proc)
		}
	}
}

func main() {
	createProcessList()
	for _, proc := range procList {
		proc.Pprint()
	}
}
