package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type job struct {
	name   string
	shell  string
	stdout string
	stderr string
}

type lsf struct {
	job
	pid string
}

type local struct {
	job
	pid string
}

func NewLSF(name, shell, stdout, stderr string) lsf {
	jb1 := job{name: name, shell: shell, stdout: stdout, stderr: stderr}
	j := lsf{jb1, ""}
	return j
}

func NewLocal() local {
	var j local
	return j
}

func ExtractJobId(m string) []string {
	slice_jobs_line := strings.Split(strings.Trim(m, "\n"), "\n")
	var slice_jobs []string
	for _, value := range slice_jobs_line[1:] {
		jobid := strings.Fields(value)[0]
		slice_jobs = append(slice_jobs, jobid)
	}
	return slice_jobs
}

func (m lsf) GetJobId() []string {
	cmd := exec.Command("bjobs")
	buf, err := cmd.Output()
	if err != nil {
		log.Println(err)
	}
	slice_jobs := ExtractJobId(string(buf))
	return slice_jobs
}

func (j local) GetJobId() []string {
	cmd := exec.Command("ps", "x")
	buf, err := cmd.Output()
	if err != nil {
		log.Println(err)
	}
	slice_jobs := ExtractJobId(string(buf))
	return slice_jobs
}

//func (p, job) getStatus(tp string) bool{
//	if tp == "lsf" {

func main() {
	a := NewLSF("test", "test.sh", "out", "err")
	slice_tmp := a.GetJobId()
	fmt.Println(strings.Join(slice_tmp, "\n"))
	fmt.Println(a.name)
}
