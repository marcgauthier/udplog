package udplog

import (
	"encoding/json"
	"fmt"
	"net"
	"runtime"
	"strings"
	"time"
)

type Msg struct {
	App   string
	Time  int64
	Msg   string
	Level int
}

var debug bool = false

var host, application string

func Init(Host, Applicationname string, _debug bool) {
	host = Host
	application = Applicationname
	debug = _debug
}

func send(msg string, level int) {
	m := Msg{App: application, Time: time.Now().UnixNano(), Msg: msg, Level: level}
	buf, _ := json.Marshal(m)

	conn, err := net.Dial("udp", host)
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	fmt.Fprintf(conn, string(buf))
}

// Info example:
//
// Info("timezone %s", timezone)
func Info(msg string, vars ...interface{}) {
	send(fmt.Sprintf(strings.Join([]string{"[INFO ]", msg}, " "), vars...), 0)
}

// Fatal example:
//
// Fatal(errors.New("db timezone must be UTC"))
func Fatal(err error) {
	pc, fn, line, _ := runtime.Caller(1)
	// Include function name if debugging
	if debug {
		send(fmt.Sprintf("[FATAL] %s [%s:%s:%d]", err, runtime.FuncForPC(pc).Name(), fn, line), 4)
	} else {
		send(fmt.Sprintf("[FATAL] %s [%s:%d]", err, fn, line), 4)
	}
}

// Error example:
//
// Error(errors.Errorf("Invalid timezone %s", timezone))
func Error(err error) {
	pc, fn, line, _ := runtime.Caller(1)
	// Include function name if debugging
	if debug {
		send(fmt.Sprintf("[ERROR] %s [%s:%s:%d]", err, runtime.FuncForPC(pc).Name(), fn, line), 3)
	} else {
		send(fmt.Sprintf("[ERROR] %s [%s:%d]", err, fn, line), 3)
	}
}
Footer
© 2022 GitHub, Inc.
Footer navigation
Terms
Privacy
Security
Status
Docs
Contact GitHub
Pricing
API
Training
Blog
About
