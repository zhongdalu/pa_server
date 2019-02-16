//@Auth:zdl
package gateway_server

import (
	"bufio"
	"com.deer.pa_server/protocol"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/g/util/gconv"
	"gitee.com/sky_big/workmod/myLog"
	"net"
	"strings"
)

func Run() {
	netListen, err := net.Listen("tcp", ":9988")
	if err!=nil {
		myLog.Error(err)
	}
	defer netListen.Close()
	fmt.Println("Waiting for clients")
	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}
		fmt.Println(conn.RemoteAddr().String(), " tcp connect success")
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	//声明一个临时缓冲区，用来存储被截断的数据
	tmpBuffer := make([]byte, 0)
	//存放不同socket和fac_port 数据
	mutilabMap:=make(map[string][]byte)
	buffer := make([]byte, 0)//1024
	for {
		str,err:=reader.ReadString('\n')
		if err==nil{
			// 去掉字符串尾部的回车
			str = strings.TrimSpace(str)
			fmt.Println(str)
			data:=protocol.StandardStc{}
			err = json.Unmarshal([]byte(str), &data)
			if err!=nil {
				myLog.Error(err)
				continue
			}
			if data.Details.Data=="" {
				continue
			}
			buffer,err=hex.DecodeString(data.Details.Data)
			if err !=nil{
				myLog.Error(err)
				continue
			}
			_,ok :=mutilabMap[gconv.String(data.Details.FacPort)+"_"+gconv.String(data.Details.Socket)]
			if !ok{
				mutilabMap[gconv.String(data.Details.FacPort)+"_"+gconv.String(data.Details.Socket)]=make([]byte,0)
			}
			tmpBuffer=mutilabMap[gconv.String(data.Details.FacPort)+"_"+gconv.String(data.Details.Socket)]
			mutilabMap[gconv.String(data.Details.FacPort)+"_"+gconv.String(data.Details.Socket)] = protocol.Unpack(append(tmpBuffer, buffer...),gconv.String(data.Details.FacPort))//buffer[:n]...
		}else{
			myLog.Error("connection error: ",err)
			//	_ = conn.Close()
		}
	}
}
