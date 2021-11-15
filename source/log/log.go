package log

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

var fileDir string

// Init def .
func Init() error {
	fileDir = strings.TrimSuffix(os.Getenv("LOG_FILEDIR"), "/")
	if fileDir == "" {
		fileDir = "/alird/logs"
	}
	if _, err := os.Stat(fileDir); err != nil {
		os.MkdirAll(fileDir, 0x755)
	}
	var setlog = func() {
		f, _ := os.OpenFile(fileDir+"/"+time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
		log.SetOutput(io.MultiWriter(f, os.Stdout))
	}
	setlog()
	log.SetFlags(log.Ldate | log.LstdFlags)
	go startTimer(setlog)
	go checkFileTime(5)
	return nil
}

// 零点定时器
func startTimer(f func()) {
	go func() {
		var now, next time.Time
		var t *time.Timer
		for {
			now = time.Now()
			// 计算下一个零点
			next = now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
			t = time.NewTimer(next.Sub(now))
			<-t.C
			f()
		}
	}()
}

// 删除过期文件
func checkFileTime(saveDay int) {
	var expiretime time.Time
	var err error
	var listfile []os.FileInfo
	var file os.FileInfo
	for {
		//遍历文件夹下所有的文件
		listfile, err = ioutil.ReadDir(fileDir)
		if err == nil {
			for _, file = range listfile {
				//筛选.log文件
				if path.Ext(file.Name()) == ".log" {
					expiretime = file.ModTime().AddDate(0, 0, saveDay)
					//过期文件 删除
					if time.Now().After(expiretime) {
						os.Remove(fileDir + "/" + file.Name())
					}
				}
			}
		}
		time.Sleep(1 * 24 * time.Hour) //每天清理一次
	}
}
