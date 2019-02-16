//@Auth:zdl
package huaLong

import (
	"com.deer.pa_server/util"
	"encoding/hex"
	"fmt"
	"github.com/gogf/gf/g/os/grpool"
	"gitee.com/sky_big/workmod/myLog"
	"time"
)

type ModelStc struct {
	name           string
	readerChannel  chan []byte
	HeaderBegin    byte
	HeaderEnd      byte
	HeaderLength   int
	SaveDataLength int
	CsBegin        int
	//协议尾
	TailEnd byte
}

func NewSub(na string) *ModelStc {
	modelStc := &ModelStc{
		name:           na,
		readerChannel:  make(chan []byte, 100),
		HeaderBegin:    0x68,
		HeaderEnd:      0x68,
		HeaderLength:   6,
		SaveDataLength: 2,
		CsBegin:        6,
		TailEnd: 0x16,
	}
	modelStc.startWork()
	return modelStc
}

func (r *ModelStc) startWork() {
	go func() {
		for {
			select {
			case data := <-r.readerChannel:
				err := grpool.Add(func() {
					r.digistData(data)
				})
				if err != nil {
					myLog.Error(err)
				}
			}
		}
	}()
}

//解析数据
func (r *ModelStc) digistData(buffer []byte) {
	time.Sleep(time.Second * 2)
	myLog.Println(hex.EncodeToString(buffer))
	//校验位
	csData := buffer[len(buffer)-2]
	csBytes := buffer[r.CsBegin : len(buffer)-2]
	sum := util.BytesSum(csBytes)
	if csData != sum {
		myLog.Error("校验失败:", hex.EncodeToString(buffer))
		return
	}
	if buffer[len(buffer)-1]!=r.TailEnd{
		myLog.Error("协议尾校验失败:",hex.EncodeToString(buffer))
	}
	fmt.Println("校验成功！")
}

func (r *ModelStc) GetName() string {
	return r.name
}

func (r *ModelStc) Unpack(buffer []byte) []byte {
	length := len(buffer)
	var i int
	for i = 0; i < length; i = i + 1 {
		if length < i+r.HeaderLength+r.SaveDataLength {
			break
		}
		if buffer[i] == r.HeaderBegin && buffer[i+r.HeaderLength-1] == r.HeaderEnd {
			countBytes := buffer[i+1 : i+1+r.SaveDataLength]
			messageLength := util.BytesToIntReverse(countBytes)>>2
			if length < i+r.HeaderLength+r.SaveDataLength+messageLength {
				break
			}
			data := buffer[i : i+r.HeaderLength+messageLength+2]
			r.readerChannel <- data
			i += r.HeaderLength + messageLength + 2 - 1
		}
	}
	if i == length {
		return make([]byte, 0)
	}
	return buffer[i:]
}
