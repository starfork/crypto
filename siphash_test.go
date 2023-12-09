package gocrypto

import (
	"fmt"
	"testing"
)

func TestSipHash(t *testing.T) {
	fmt.Println(SipHash("abcd"))
}
