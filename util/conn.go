package util

import (
	"net"
	"printsServer/config"
	"strconv"
	"time"
)

func GetConn() net.Conn {
	conn, err := net.Dial("tcp", config.PrinterAddr + ":" + config.PrinterPort)
	if err != nil {
		WriteError(strconv.FormatInt(int64(time.Millisecond), 10) + " " + err.Error())
		print("error " + err.Error() )
		// panic(err)
	}
	return conn
}