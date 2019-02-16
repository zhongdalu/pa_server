//@Auth:zdl
package deer

import (
	"encoding/hex"
	"fmt"
	"gitee.com/sky_big/workmod"
	"github.com/gogf/gf/g/os/glog"
	"runtime"
	"testing"
)

func TestFunc(t *testing.T) {
	workmod.Init()
	glog.Println("11111111")
	str := "eb90eb90eb90059fda02007b"//eb90eb90eb90059fda02007b
    data,err:=hex.DecodeString(str)
    if err!=nil {
    	fmt.Println(err)
	}
	proto := NewSub("6060")
	proto.Unpack(data)

	a:=runtime.GOOS
	select {
	}
}
