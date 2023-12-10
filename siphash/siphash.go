package siphash

import "github.com/dchest/siphash"

func SipHash(token string) uint64 {
	return siphash.Hash(1, 1, []byte(token))

}
