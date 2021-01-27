package util

import (
	"bufio"
	"net"
	"printsServer/config"
	"strings"
	"time"
)

func GetConn() (net.Conn, error) {
	conn, err := net.DialTimeout("tcp", config.PrinterAddr+":"+config.PrinterPort, 1 * time.Second)
	
	return conn, err
}
func GetPageCount() (string, error) {
	var command = "@PJL INFO PAGECOUNT"
	return RunSingleCommand(command)
}

func GetStatus() (string, error) {
	var command = "@PJL INFO STATUS"
	return RunSingleCommand(command)
}

func Gettoner() (string, error) {
	var command = "@PJL INFO TONERCOUNT5"
	return RunSingleCommand(command)
}

func RunCommand(command string) (string, error){
	conn, err := net.DialTimeout("tcp", config.PrinterAddr+":"+config.PrinterPort, 2 * time.Second)
	if err != nil{
		return "", err
	}
	_, err = write(conn, command)
	if err != nil{
		return "", err
	}
	status, err := read(conn)
	if err != nil{
		return "", err
	}
	_ = conn.Close()
	return strings.Replace(status, command, "", 1), err
}

func RunSingleCommand(command string) (string, error) {
	conn, err := net.DialTimeout("tcp", config.PrinterAddr+":"+config.PrinterPort, 2 * time.Second)
	if err != nil{
		return "", err
	}
	_, err = write(conn, "\x1b%-12345X "+command+"\r\n")
	if err != nil{
		return "", err
	}
	status, err := read(conn)
	if err != nil{
		return "", err
	}
	_ = conn.Close()
	return strings.Replace(status, command, "", 1), err
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
