//客户端发送封包
package main

import (
	"bufio"
	"fmt"
	"github.com/gogf/gf/g/os/glog"
	"gitee.com/sky_big/workmod"
	"gitee.com/sky_big/workmod/myLog"
	"net"
	"os"
	"time"
)


func sender(conn net.Conn) {
	//{"SendDt":"2018-11-05T05:52:55.270","SourceModule":"GC_Server","SourceIP":"192.168.2.254","DestinationMoudule":"PA_Server","DestinationIP":"223.104.254.141","DestinationPort":59787,"ProtocolCode":"1","Details":{"socket":17930,"Type":"1","fac_port":6060,"ReturnCode":"800","Data":""}}

	/*	for i := 0; i < 1000; i++ {
			words := "{\"Id\":1,\"Name\":\"golang\",\"Message\":\"message\"}"
			data,err:=protocol.PacketString(words)
			if err!=nil{
				glog.Error(err)
				return
			}
			glog.Debug("data=>",data)
			conn.Write(data)
		}
		fmt.Println("send over")*/

	file,err:=os.Open("List 20181105.Log")
	if err!=nil{
		glog.Error(err)
		return
	}
	reder:=bufio.NewReader(file)
	count:=0
	fmt.Println("reder.size()==============",reder.Size())
	for  {
		line,_,err:=reder.ReadLine()
		if err!=nil {
			myLog.Error(err)
			break
		}
		if len(line)>0 {
			count+=1
			fmt.Println("count=======",count)
			myLog.Println(count,"   => ",string(line))
			_,err=conn.Write(append(line,[]byte("\n")...))
		}

		/*var strc =define.StandardStc{}
		err=json.Unmarshal(line, &strc)
		if err!=nil {
			glog.Error(err)
			continue
		}
		if strc.Details.Data != "" {
			count+=1
		//	glog.Println(count ,"   send  Data===============",strc.Details.Data)
			myLog.Println(count,"    ",strc.Details.Data)
			da,err:=hex.DecodeString(strc.Details.Data)
			if err!=nil{
				glog.Error(err)
			}else{
				_,err=conn.Write(da)
			}
			//break
		}*/

		if err!=nil{
			glog.Error(err)
		}
	}
	fmt.Println("count=======",count)

}

func main() {
	//binary.BigEndian.Uint16()
	workmod.Init()
	glog.SetPath("./send_log")
	glog.SetLevel(glog.LEVEL_ALL&^glog.LEVEL_ERRO)
	server := "127.0.0.1:9988"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	defer conn.Close()
	fmt.Println("connect success")
	go sender(conn)
	for {
		time.Sleep(1 * 1e9)
	}
}