//@Auth:zdl
//通讯协议处理，主要处理封包和解包的过程
package protocol

import (
	"com.deer.pa_server/protocol/deer"
	"com.deer.pa_server/protocol/huaLong"
	"errors"
	"gitee.com/sky_big/workmod/myLog"
)

var protoMap map[string]ProtoSub

func init() {
	protoMap = make(map[string]ProtoSub)
}

//解包
func Unpack(buffer []byte, facPort string) []byte {
	st, err := GetSubSington(facPort)
	if err != nil {
		myLog.Error(err)
		return []byte{}
	}
	return st.Unpack(buffer)
}

func GetSubSington(facPort string) (ProtoSub, error) {
	proto, ok := protoMap[facPort]
	if ok {
		return proto, nil
	}
	switch facPort {
	case "6060":
		proto = deer.NewSub(facPort)
	case "1":
		proto = huaLong.NewSub(facPort)
	}
	if proto == nil {
		return nil, errors.New("缺少" + facPort + "解析协议")
	}
	protoMap[facPort] = proto
	return proto, nil
}
