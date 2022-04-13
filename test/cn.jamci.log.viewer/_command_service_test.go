package cn_jamci_log_viewer

import "fmt"
import _ "cn.jamci.log.viewer/src/cn.jamci.log.viewer/service"

// 测试
func main() {
	out := make(chan string)
	defer close(out)
	go func() {
		for {
			str, ok := <-out
			if !ok {
				break
			}
			fmt.Println(str)
		}
	}()
	args := []string{"-c", "ping www.baidu.com"}
	if err := RunCommand(out, "bash", args...); err != nil {
		return
	}
}
