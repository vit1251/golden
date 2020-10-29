package auth

import (
	"crypto/md5"
	"encoding/hex"
	"log"
)

var ipad []byte = []byte{
	0x036, 0x036, 0x036, 0x036, 0x036, 0x036, 0x036, 0x036,
	0x036, 0x036, 0x036, 0x036, 0x036, 0x036, 0x036, 0x036,
	0x036, 0x036, 0x036, 0x036, 0x036, 0x036, 0x036, 0x036,
	0x036, 0x036, 0x036, 0x036, 0x036, 0x036, 0x036, 0x036,
	0x036, 0x036, 0x036, 0x036, 0x036, 0x036, 0x036, 0x036,
	0x036, 0x036, 0x036, 0x036, 0x036, 0x036, 0x036, 0x036,
	0x036, 0x036, 0x036, 0x036, 0x036, 0x036, 0x036, 0x036,
	0x036, 0x036, 0x036, 0x036, 0x036, 0x036, 0x036, 0x036,
}

var opad []byte = []byte{
	0x05C, 0x05C, 0x05C, 0x05C, 0x05C, 0x05C, 0x05C, 0x05C,
	0x05C, 0x05C, 0x05C, 0x05C, 0x05C, 0x05C, 0x05C, 0x05C,
	0x05C, 0x05C, 0x05C, 0x05C, 0x05C, 0x05C, 0x05C, 0x05C,
	0x05C, 0x05C, 0x05C, 0x05C, 0x05C, 0x05C, 0x05C, 0x05C,
	0x05C, 0x05C, 0x05C, 0x05C, 0x05C, 0x05C, 0x05C, 0x05C,
	0x05C, 0x05C, 0x05C, 0x05C, 0x05C, 0x05C, 0x05C, 0x05C,
	0x05C, 0x05C, 0x05C, 0x05C, 0x05C, 0x05C, 0x05C, 0x05C,
	0x05C, 0x05C, 0x05C, 0x05C, 0x05C, 0x05C, 0x05C, 0x05C,
}

type Authorizer struct {
	challengeData []byte  /*  */
	secret        []byte
}

func NewAuthorizer() (*Authorizer) {
	a := new(Authorizer)
	return a
}

func (self *Authorizer) convertByteToString(source []byte) []byte {
	return source
}

func (self *Authorizer) convertStringToByte(source []byte) []byte {
	return source
}

func (self *Authorizer) SetSecret(secret string) {
	self.secret = []byte(secret)
}

func (self *Authorizer) SetChallengeData(challengeData string) {
	if newChallengeData, err := hex.DecodeString(challengeData); err != nil {
		panic(err)
	} else {
		self.challengeData = newChallengeData
	}
}

//               +-------------A4--------------+
//               |                             |
//               |                      +------A3-----+
//               |                      |             |
//              A1                     A2             |
//               |                      |             |
//   HASH((secret XOR opad), HASH((secret XOR ipad), challengedata))
//
//   where HASH is chosen hash function, ipad and opad are 36 hex and 5C
//   hex (as defined in [Keyed]) and secret is a password null-padded to
//   a length of 64 bytes. If the password is longer than 64 bytes, the
//   hash-function digest of the password is used as an input (16-byte
//   for [MD5] and 20-byte for [SHA-1]) to the keyed hashed calculation.

func (self *Authorizer) CalculateDigest() (string, error) {

	/* A1: secret XOR opad (k_opad) */
	var A1 []byte
	for i := 0; i < 64; i++ {
		A1 = append(A1, '\x00')
	}
	for idx, secretByte := range self.secret {
		A1[idx] = secretByte
	}
	for i := 0; i < 64; i++ {
		A1[i] ^= '\x5C';
	}
	log.Printf("A1 = %x", A1)

	/* A2: secret XOR ipad (k_ipad) */
	var A2 []byte
	for i := 0; i < 64; i++ {
		A2 = append(A2, '\x00')
	}
	for idx, secretByte := range self.secret {
		A2[idx] = secretByte
	}
	for i := 0; i < 64; i++ {
		A2[i] ^= '\x36';
	}
	log.Printf("A2 = %x", A2)

	/* A3: HASH(A2, challengedata) */
	h3 := md5.New()
	h3.Write(A2)
	h3.Write(self.challengeData)
	A3 := h3.Sum(nil)
	log.Printf("A3 = %x", A3)

	/* A4: HASH(A1, A3) */
	h4 := md5.New()
	h4.Write(A1)
	h4.Write(A3)
	A4 := h4.Sum(nil)
	log.Printf("A4 = %x", A4)

	/* Make string representation */
	result := hex.EncodeToString(A4)

	return result, nil
}
