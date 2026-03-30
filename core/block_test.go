package core

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/Redcrops/web3_sever_with_go/types"
	"github.com/stretchr/testify/assert"
)

func TestHead_Encode_Decode(t *testing.T) {
	h := &Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Timestamp: time.Now().UnixNano(),
		Height:    10,
		Nonce:     989394,
	}
	buf := &bytes.Buffer{}
	h.EncodeBinary(buf)
	assert.Nil(t, h.EncodeBinary(buf))

	hDecode := &Header{}
	assert.Nil(t, hDecode.DecodeBinary(buf))
	assert.Equal(t, h, hDecode)
}

func TestBlock_Encode_Decode(t *testing.T) {
	b := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: time.Now().UnixNano(),
			Height:    10,
			Nonce:     989394,
		},
		Trasactions: nil,
	}
	buf := &bytes.Buffer{}
	assert.Nil(t, b.EncodeBinary(buf))

	block := &Block{}
	assert.Nil(t, block.DecodeBinary(buf))
	assert.Equal(t, b, block)
}

func Test_Block_Hash(t *testing.T) {
	b := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: time.Now().UnixNano(),
			Height:    10,
			Nonce:     989394,
		},
		Trasactions: []Trasaction{},
	}
	h := b.Hash()

	fmt.Println("block hash is ", h)
	assert.False(t, h.IsZero())
}
