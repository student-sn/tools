package logs

import (
	"fmt"
	"github.com/student-sn/tools/stopper"
	"os"
	"time"
)

type Logs struct {
	name    string
	file    *os.File
	stopper *stopper.Stopper
}

func Init(name string, stopper *stopper.Stopper) *Logs {
	path := ""
	if os.Getenv("ST_TOOLS_TEST") == "1" {
		if envPath := os.Getenv("ST_TOOLS_TEST_LOGS_PATH"); envPath != "" {
			path = envPath + "/"
		} else {
			path = "/tmp/st_tests/"
		}
	}
	if _, err := os.ReadDir(path + "logs"); os.IsNotExist(err) {
		err := os.Mkdir(path+"logs", os.ModePerm)
		if err != nil {
			panic(err)
		}
		if _, err := os.Open(path + "logs/" + name + ".log"); os.IsNotExist(err) {
			file, err := os.Create(path + "logs/" + name + ".log")
			if err != nil {
				panic(err)
			}
			return &Logs{
				name:    name,
				file:    file,
				stopper: stopper,
			}
		} else {
			file, err := os.Open(path + "logs/" + name + ".log")
			if err != nil {
				panic(err)
			}
			return &Logs{
				name:    name,
				file:    file,
				stopper: stopper,
			}
		}
	}
	return nil
}

func (l *Logs) Fatal(message any) {
	_, err := l.file.Write([]byte(fmt.Sprintf("[FATAL][%s][%s]: %v\n", l.name, time.Now().Format(time.RFC822Z), message)))
	if err != nil {
		return
	}
	fmt.Printf("\033[1;91m[FATAL][%s][%s]\033[0m: %v\n", l.name, time.Now().Format(time.RFC822Z), message)
	fmt.Printf("\033[1;91m[FATAL][%s][%s]\033[0m: %v\n", l.name, time.Now().Format(time.RFC822Z), "Stop system")
	if os.Getenv("ST_TOOLS_TEST") != "1" {
		l.stopper.Stop(-1, message)
	}
}

func (l *Logs) Log(message any) {
	_, err := l.file.Write([]byte(fmt.Sprintf("[LOG][%s][%s]: %v\n", l.name, time.Now().Format(time.RFC822Z), message)))
	if err != nil {
		return
	}
	fmt.Printf("\033[1;95m[LOG][%s][%s]\033[0m: %v\n", l.name, time.Now().Format(time.RFC822Z), message)
}

func (l *Logs) Warn(message string) {
	_, err := l.file.Write([]byte(fmt.Sprintf("[WARN][%s][%s]: %s\n", l.name, time.Now().Format(time.RFC822Z), message)))
	if err != nil {
		return
	}
	fmt.Printf("\033[1;93m[WARN][%s][%s]\033[0m: %s\n", l.name, time.Now().Format(time.RFC822Z), message)
}

func (l *Logs) Debug(message any) {
	//if os.Getenv("ST_DEBUG") == "1" {
	_, err := l.file.Write([]byte(fmt.Sprintf("[DEBUG][%s][%s]: %v\n", l.name, time.Now().Format(time.RFC822Z), message)))
	if err != nil {
		return
	}
	fmt.Printf("\033[1;94m[DEBUG][%s][%s]\033[0m: %v\n", l.name, time.Now().Format(time.RFC822Z), message)
	//}
}
