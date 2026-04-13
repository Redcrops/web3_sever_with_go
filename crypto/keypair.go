package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"math/big"

	"github.com/Redcrops/web3_sever_with_go/types"
)

type PrivateKey struct {
	key *ecdsa.PrivateKey
}

func (k PrivateKey) Sign(data []byte) (*Signature, error) {
	r, s, err := ecdsa.Sign(rand.Reader, k.key, data)
	if err != nil {
		return nil, err
	}
	return &Signature{
		r: r,
		s: s,
	}, nil
}
func GeneratePrivateKey() PrivateKey {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	return PrivateKey{key: key}
}

func (k PrivateKey) PublicKey() PublicKey {
	return PublicKey{
		key: &k.key.PublicKey,
	}
}

type PublicKey struct {
	key *ecdsa.PublicKey
}

func (k PublicKey) ToSlice() []byte {
	// 这里的 .X 和 .Y 是 Go 语言中的结构体字段访问语法。
	// k.key 是一个 *ecdsa.PublicKey 类型的指针，PublicKey 结构体有 X, Y 坐标 (大椭圆曲线点) 字段。
	// 通过 k.key.X 和 k.key.Y 访问这些字段。
	// 如果你看到划线，可能是由于 IDE 检测到 MarshallCompressed 不是有效的函数（应为 MarshalCompressed），或者包没有正确导入。
	return elliptic.MarshalCompressed(k.key.Curve, k.key.X, k.key.Y)
}
func (k PublicKey) Address() types.Address {
	h := sha256.Sum256(k.ToSlice())
	return types.NewAddressFromBytes(h[len(h)-20:])
}

type Signature struct {
	r, s *big.Int
}

func (sig Signature) Verify(publicKey PublicKey, data []byte) bool {
	// Verify方法用于验证数字签名是否有效。
	// 其中 r 和 s 是签名的两个组成部分（都是大整数），
	// 它们是在对消息进行签名时由椭圆曲线数字签名算法（ECDSA）生成的。
	// 这里将签名人公钥publicKey、原始数据data和签名的两个部分r、s传递给ecdsa.Verify方法进行验证。
	// ECDSA签名校验的内部原理如下：
	// 1. 首先根据传入的原始数据（data）和椭圆曲线公钥参数，计算消息的哈希值（实现在ecdsa.Verify内部可自定义）。
	// 2. 签名由两个大整数 (r, s) 组成，校验流程如下：
	//    a) 校验 r, s 是否在合法范围（0, N)，N 是椭圆曲线的阶（order）。
	//    b) 计算消息哈希转为整数 e。
	//    c) 计算 s 的模逆 w = s^-1 mod N。
	//    d) 计算 u1 = e * w mod N，u2 = r * w mod N。
	//    e) 计算曲线上的两个点的加和：点1为基点G乘以u1，点2为公钥乘以u2，得到点 (x, y)。
	//    f) 校验 x 坐标对 N 取模，是否等于 r。
	// 如果相等，则验证通过，说明该签名是由该公钥对应的私钥对该数据签名所得，否则验证失败。
	return ecdsa.Verify(publicKey.key, data, sig.r, sig.s)
}
