package util

import (
	"bufio"
	"net"
	"printsServer/config"
	"printsServer/filesystem"
	"strings"
	"time"
)

type TrayNum int

const (
	FIRST = iota
	SECOND
)

func GetConn() (net.Conn, error) {

	conn, err := net.DialTimeout("tcp", config.PrinterAddr+":"+config.PrinterPort, 1*time.Second)

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

func RunCommand(command string) (string, error) {
	conn, err := net.DialTimeout("tcp", config.PrinterAddr+":"+config.PrinterPort, 2*time.Second)
	if err != nil {
		return "", err
	}
	_, err = write(conn, command)
	if err != nil {
		return "", err
	}
	status, err := read(conn)
	if err != nil {
		return "", err
	}
	_ = conn.Close()
	return strings.Replace(status, command, "", 1), err
}

func PrintDoc(docPath string, num TrayNum, duplexEnabled bool) error {
	conn, err := net.DialTimeout("tcp", config.PrinterAddr+":"+config.PrinterPort, 2*time.Second)
	if err != nil {
		return err
	}
	PDFBin, err := filesystem.RetrieveROM(docPath)
	if err != nil {
		return err
	}

	command := "\x1b%-12345X @PJL\r\n" +
		"@PJL JOB NAME = \"printPDF\" DISPLAY = \"Printing \"\r\n" +
		"@PJL SET MEDIASOURCE = TRAY3\r\n" +
		"@PJL ENTER LANGUAGE = PDF\r\n" +
		string(PDFBin) +
		"\x1b%-12345X @PJL\r\n" +
		"@PJL RESET" +
		"@PJL EOJ NAME = \"printPDF\"" +
		"\x1b%-12345X"
	_, err = write(conn, command)
	if err != nil {
		return err
	}
	return nil
}

func RunSingleCommand(command string) (string, error) {
	conn, err := net.DialTimeout("tcp", config.PrinterAddr+":"+config.PrinterPort, 2*time.Second)
	if err != nil {
		return "", err
	}
	_, err = write(conn, "\x1b%-12345X "+command+"\r\n")
	if err != nil {
		return "", err
	}
	status, err := read(conn)
	if err != nil {
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
