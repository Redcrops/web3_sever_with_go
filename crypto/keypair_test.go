package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeyPair_Sign_Verify(t *testing.T) {
	privKey := GeneratePrivateKey()
	publicKey := privKey.PublicKey()

	// 签名用于证明消息的真实性和完整性，保证消息未被篡改且确实来自签名者。
	// 例如在实际应用场景中，A 向 B 发送交易数据，A 用自己的私钥对数据签名，B 可以用 A 的公钥验证签名，从而确认消息确实由 A 发出且未被修改。
	msg := []byte("Alice pays Bob 10 coins")
	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)

	b := sig.Verify(publicKey, msg)
	assert.True(t, b, "signature should be verified")

	otherPrivKey := GeneratePrivateKey()
	otherPublicKey := otherPrivKey.PublicKey()
	b = sig.Verify(otherPublicKey, msg)
	assert.False(t, b, "signature should not be verified")
}
