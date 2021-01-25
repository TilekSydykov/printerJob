package util

import (
	"bufio"
	"net"
	"printsServer/config"
	"strconv"
	"strings"
	"time"
)

func GetConn() net.Conn {
	conn, err := net.Dial("tcp", config.PrinterAddr+":"+config.PrinterPort)
	if err != nil {
		WriteError(strconv.FormatInt(int64(time.Millisecond), 10) + " " + err.Error())
		print("error " + err.Error())
		// panic(err)
	}

	return conn
}

func GetPageCount() string {
	var command = "@PJL INFO PAGECOUNT"
	return RunSingleCommand(command)
}

func GetStatus() string {
	var command = "@PJL INFO STATUS"
	return RunSingleCommand(command)
}

func RunSingleCommand(command string) string {
	conn, err := net.Dial("tcp", config.PrinterAddr+":"+config.PrinterPort)
	HandleError(err)
	_, _ = write(conn, "\x1b%-12345X "+command+"\r\n")
	status, err := read(conn)
	if err != nil {
		println(err)
	}
	_ = conn.Close()
	return strings.Replace(status, command, "", 1)
}

func write(conn net.Conn, content string) (int, error) {
	writer := bufio.NewWriter(conn)
	number, err := writer.WriteString(content)
	if err == nil {
		err = writer.Flush()
	}
	return number, err
}

func read(conn net.Conn) (string, error) {
	reader := bufio.NewReader(conn)
	return reader.ReadString('')
}
