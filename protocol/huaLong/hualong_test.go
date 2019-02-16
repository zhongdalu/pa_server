//@Auth:zdl
package huaLong

import (
	"encoding/hex"
	"fmt"
	"github.com/gogf/gf/g/os/glog"
	"gitee.com/sky_big/workmod"
	"testing"
)

func TestFunc(t *testing.T) {
	glog.Println("11111111")
	workmod.Init()
	str := "682B002B00684B43810000570B6204FFD616"
    data,err:=hex.DecodeString(str)
    if err!=nil {
    	fmt.Println(err)
	}
	proto := NewSub("6060")
	proto.Unpack(data)
	select {
	}
}
