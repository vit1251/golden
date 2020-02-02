package mailer

import (
	"testing"
	"log"
)

// Password here is tanstaaftanstaaf
// S: "CRAM-MD5-f0315b074d728d483d6887d0182fc328"
// C: "CRAM-MD5-56be002162a4a15ba7a9064f0c93fd00"

func TestDigest(t *testing.T) {

	a := NewAuthorizer()

	a.SetChallengeData("f0315b074d728d483d6887d0182fc328")
	a.SetSecret("tanstaaftanstaaf")

	expected := "56be002162a4a15ba7a9064f0c93fd00"

	actual, err := a.CalculateDigest()
	if err != nil {
		panic(err)
	}

	log.Printf("actual = %s expected = %s", actual, expected)

	if actual != expected {
		t.Errorf("Wrong digest calculation: actual = %q expected = %q", actual, expected)
	}

}
