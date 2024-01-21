package basex_test

import (
	"fmt"
	"testing"

	"github.com/starfork/go-crypto/basex"
)

func TestX(t *testing.T) {
	text := "abcdefg"
	x1 := basex.NewBaseX([]byte("123456789abcdefghijklmnopqrstuvwxyz"))
	enc1 := x1.Encode(text)
	fmt.Println(enc1)
	fmt.Println(x1.Decode(enc1))

	x2 := basex.NewBaseX([]byte("xN2mXyIJzu93gCVlAnfoS1LRTDasKBpr6vUi4QcqOEdtjY8hPZMbH57GF0wWek"))
	enc2 := x2.Encode(text)
	fmt.Println(enc2)
	fmt.Println(x2.Decode(enc2))

}
