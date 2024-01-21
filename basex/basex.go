package basex

//类似于base64的，可以自定义的一个“加密”
//参考 url https://www.cnblogs.com/Mishell/p/12241872.html
import (
	"bytes"
	"math/big"
)

type BaseXByte []byte

var baseBytes = BaseXByte("123456789abcdefghijklmnopqrstuvwxyz")

type BaseX struct {
	bytes []byte
	blen  int64
}

func NewBaseX(bytes ...[]byte) *BaseX {
	if len(bytes) > 0 {
		baseBytes = bytes[0]
	}
	return &BaseX{
		bytes: baseBytes,
		blen:  int64(len(baseBytes)),
	}
}
func (s BaseXByte) String() string {
	return string(s)
}

// Encode 编码
func (e *BaseX) Encode(input string) BaseXByte {
	x := big.NewInt(0).SetBytes(BaseXByte(input))
	base := big.NewInt(e.blen)
	//fmt.Println(base)
	zero := big.NewInt(0)
	mod := &big.Int{}
	var result BaseXByte
	// 被除数/除数=商……余数
	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod)
		result = append(result, e.bytes[mod.Int64()])
	}
	e.reverse(result)
	return result
}

// Decode 解码
func (e *BaseX) Decode(input []byte) BaseXByte {
	result := big.NewInt(0)
	for _, b := range input {
		charIndex := bytes.IndexByte(e.bytes, b)
		result.Mul(result, big.NewInt(e.blen))
		result.Add(result, big.NewInt(int64(charIndex)))
	}
	decoded := result.Bytes()
	if input[0] == e.bytes[0] {
		decoded = append([]byte{0x00}, decoded...)
	}
	return decoded
}

// ReverseBytes 翻转字节
func (e *BaseX) reverse(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}
