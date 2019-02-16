//@Auth:zdl
package deer

import (
	"bytes"
	"encoding/hex"
	"errors"
	"gitee.com/sky_big/workmod/util"
)

//头
type HeadMode struct {
	BeginSign []byte
	EndSign   []byte
	Pos       int
	Len       int
}

//长度位
type LenMode struct {
	Pos         int
	Len         int
	CalBeginPos int
	CalEndPos   int
}

//校验位
type CsMode struct {
	Len int
	CalBeginPos int
}

//Model
type Eb90 struct {
	data     []byte
	HeadMode HeadMode
	LenMode  LenMode
	CsMode   CsMode
}

type ProtocolFace interface {
	//校验头
	CalHead() error
	//获取长度
	GetLen() error
	//计算校验和
	CalCs() error
}

func NewEb90(data []byte) *Eb90 {
	return &Eb90{
		data: data,
		HeadMode: HeadMode{
			BeginSign: []byte{0xeb, 0x90, 0xeb},
			EndSign:   []byte{0x90, 0xeb, 0x90},
			Pos:       0,
			Len:       6,
		},
		LenMode: LenMode{
			Pos:         6,
			Len:         1,
			CalBeginPos: 7,
		},
		CsMode: CsMode{
			Len: 1,
			CalBeginPos:7,
		},
	}
}

func (r *Eb90) CalHead() error {
	head := r.data[r.HeadMode.Pos:r.HeadMode.Pos+r.HeadMode.Len]
	if head==nil{
		return errors.New("获取head数据失败")
	}
	pHead:=append(r.HeadMode.BeginSign,r.HeadMode.EndSign...)
	if !bytes.Equal(head,pHead){
		return  errors.New("head 不一致"+hex.EncodeToString(head)+"$"+hex.EncodeToString(pHead))
	}
	return nil
}

func (r*Eb90)GetLen()(int,error){
	lenData := r.data[r.LenMode.Pos:r.LenMode.Pos+r.LenMode.Len]
	if lenData==nil {
		return 0,errors.New("获取长度数据失败")
	}
	len:=util.BytesToUInt16(lenData)
	return len,nil
}


func (r *Eb90) CalCs() {
	//csData:=r.data[r.CsMode.CalBeginPos:len(r.data)-1]


}

