package service

import (
	"bufio"
	"io"
	"log"
	"os/exec"
	"runtime/debug"
	"sync"
)

// 读取流信息
func readLog(wg *sync.WaitGroup, out chan string, reader io.ReadCloser) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r, string(debug.Stack()))
		}
	}()
	defer wg.Done()
	r := bufio.NewReader(rea der)
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF || err != nil {
			return
		}
		out <- string(line)
	}
}

// 执行shell
func RunCommand(out chan string, name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	if err := cmd.Start(); err != nil {
		return err
	}
	wg := sync.WaitGroup{}
	defer wg.Wait()
	wg.Add(2)
	go readLog(&wg, out, stdout)
	go readLog(&wg, out, stderr)
	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}