package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func TerminateVim(path string, info os.FileInfo, err error) error {
	var proc []int
	if strings.Count(path, "/") == 3 {
		if strings.Contains(path, "/status") {
			pid, err := strconv.Atoi(path[6:strings.LastIndex(path, "/")])
			if err != nil {
				return err
			}
			f, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			name := string(f[6:bytes.IndexByte(f, '\n')])
			if name == "vim" {
				log.Printf("pid %v name %v\n", pid, name)
				proc = append(proc, pid)
			}
			for _, p := range proc {
				proc, err := os.FindProcess(p)
				if err != nil {
					return err
				}
				proc.Kill()
			}
			return nil
		}
	}
	return nil
}

func main() {
	err := filepath.Walk("/proc", TerminateVim)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Killed vim\n")
}
