package auth

import (
	"log"
	"bytes"
	"testing"
)

// Password here is tanstaaftanstaaf
// S: "CRAM-MD5-f0315b074d728d483d6887d0182fc328"
// C: "CRAM-MD5-56be002162a4a15ba7a9064f0c93fd00"

func TestDigest(t *testing.T) {

	a := NewAuthorizer()
	a.SetChallengeData([]byte("f0315b074d728d483d6887d0182fc328"))
	a.SetSecret([]byte("tanstaaftanstaaf"))

	expected := []byte("56be002162a4a15ba7a9064f0c93fd00")

	actual, err := a.CalculateDigest()
	if err != nil {
		panic(err)
	}

	log.Printf("actual = %s expected = %s", actual, expected)

	res := bytes.Compare(expected, actual)
	if res != 0 {
		t.Errorf("Wrong digest calculation: actual = %q expected = %q", actual, expected)
	}

}
