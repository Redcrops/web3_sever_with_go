package core

import(
	"bytes"
	"testing"
	"time"

	"github.com/RedCrops/web3_sever_with_go/types"
)

func TestHead_Encode_Decode(t *testing.T){
	h := &Header{
		Version: 1,
		PrevBlock: types.RandomHash(),
		Timestamp: time.Now().UnixNano(),
		Height: 10,
		Nonce: 989394,
	}
}