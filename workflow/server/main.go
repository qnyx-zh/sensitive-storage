package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"sync"
	"time"
)

var mutex = &sync.Mutex{}

func main() {
	http.HandleFunc("/upgrade", func(writer http.ResponseWriter, request *http.Request) {
		mutex.Lock()
		cmd := exec.Command("chmod", "a+x", "startup.sh")
		_ = cmd.Run()
		f, _ := os.Create("upgrade.log")
		f.WriteString(fmt.Sprintf("新版本更新：%v\n", time.Now().Unix()))
		defer func() {
			f.Close()
			mutex.Unlock()
		}()
		cmd = exec.Command("sh", "./startup.sh")
		_ = cmd.Run()
		time.Sleep(time.Second * 10)
		return
	})
	err := http.ListenAndServe(":10012", nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
