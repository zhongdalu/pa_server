//@Auth:zdl
package deer

import (
	"bytes"
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
	Header         []byte
	HeaderLength   int
	SaveDataLength int
	CsBegin        int
}

func NewSub(na string) *ModelStc {
	modelStc := &ModelStc{
		name:           na,
		readerChannel:  make(chan []byte, 100),
		Header:         []byte{0xeb, 0x90, 0xeb, 0x90, 0xeb, 0x90},
		HeaderLength:   6,
		SaveDataLength: 1,
		CsBegin:        7,
	}
	//开启解析数据协程
	modelStc.startWork()
	return modelStc
}

func (r *ModelStc) GetName() string {
	return r.name
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
	csData := buffer[len(buffer)-1]
	csBytes := buffer[r.CsBegin : len(buffer)-1]
	sum := util.BytesSum(csBytes)
	if csData != sum {
		myLog.Error("校验失败:", hex.EncodeToString(buffer))
		return
	}
	fmt.Println("校验成功！")
}

func (r *ModelStc) Unpack(buffer []byte) []byte {
	length := len(buffer)
	var i int
	for i = 0; i < length; i = i + 1 {
		if length < i+r.HeaderLength+r.SaveDataLength {
			break
		}
		if bytes.Equal(buffer[i:i+r.HeaderLength], r.Header) {
			countBytes := buffer[i+r.HeaderLength : i+r.HeaderLength+r.SaveDataLength]
			messageLength := util.BytesToInt(countBytes)
			if length < i+r.HeaderLength+r.SaveDataLength+messageLength {
				break
			}
			data := buffer[i : i+r.HeaderLength+r.SaveDataLength+messageLength] //+r.HeaderLength+r.SaveDataLength
			r.readerChannel <- data
			i += r.HeaderLength + r.SaveDataLength + messageLength - 1
		}
	}
	if i == length {
		return make([]byte, 0)
	}
	return buffer[i:]
}
