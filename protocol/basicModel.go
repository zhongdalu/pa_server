//@Auth:zdl
package protocol

//拆包
type ProtoSub interface {
	GetName() string
	Unpack(buffer []byte) []byte
}

type StandardStc struct {
	DestinationIP      string `json:"DestinationIP"`
	DestinationMoudule string `json:"DestinationMoudule"`
	DestinationPort    int    `json:"DestinationPort"`
	Details            struct {
		Data       string `json:"Data"`
		ReturnCode string `json:"ReturnCode"`
		Type       string `json:"Type"`
		FacPort    int    `json:"fac_port"`
		Socket     int    `json:"socket"`
	} `json:"Details"`
	ProtocolCode string `json:"ProtocolCode"`
	SendDt       string `json:"SendDt"`
	SourceIP     string `json:"SourceIP"`
	SourceModule string `json:"SourceModule"`
}
