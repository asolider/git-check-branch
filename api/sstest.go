package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

func copyAndCapture(msg chan []byte, r io.Reader) ([]byte, error) {
	var out []byte
	buf := make([]byte, 1024, 1024)
	for {
		n, err := r.Read(buf[:])
		if n > 0 {
			d := buf[:n]
			out = append(out, d...)
			msg <- out
			//_, err := w.Write(d)
			if err != nil {
				return out, err
			}
		}
		if err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				err = nil
			}
			return out, err
		}
	}
	// never reached
	panic(true)
	return nil, nil
}

func test1(msg chan []byte) {
	cmd := exec.Command("wget", "https://dl.google.com/go/go1.11.2.linux-amd64.tar.gz")
	var stdout, stderr []byte
	var errStdout, errStderr error
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	cmd.Start()

	go func() {
		stdout, errStdout = copyAndCapture(msg, stdoutIn)
		//msg <- stdout
	}()

	go func() {
		stderr, errStderr = copyAndCapture(msg, stderrIn)
		msg <- stderr
	}()

	err := cmd.Wait()
	if err != nil {
		log.Printf("cmd.Run() failed with %s\n", err)
	}
	if errStdout != nil || errStderr != nil {
		log.Printf("failed to capture stdout or stderr\n")
	}
	fmt.Print("读取完毕")
}
