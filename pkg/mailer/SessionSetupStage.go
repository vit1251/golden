package mailer

import (
	"bytes"
	"fmt"
	"log"
	"runtime"
	"time"
)

func (self *Mailer) GetPlatform() string {
	if runtime.GOOS == "windows" {
		return "Windows"
	} else if runtime.GOOS == "linux" {
	 	return "Linux"
	}
	return "Unknown"
}

func (self *Mailer) GetTime() string {
	// Sun, 26 Jan 2020 18:02:17 +0300
	now := time.Now().Format(time.RFC1123Z)
	log.Printf("Time is %s", now)
	return now
}

func (self *Mailer) createAuthorization(chData []byte) (string) {
	a := NewAuthorizer()
	a.SetChallengeData(string(chData))
	a.SetSecret(self.secret)
	responseDigest, err := a.CalculateDigest()
	if err != nil {
		panic(err)
	}
	password := fmt.Sprintf("%s-%s-%s", "CRAM", "MD5", responseDigest)
	return password
}

func (self *Mailer) parseInfoFrame(body []byte) ([]byte, []byte, error) {
	values := bytes.SplitN(body, []byte(" "), 2)
	return values[0], values[1], nil
}

func (self *Mailer) parseInfoOptFrame(rawOptions []byte)  {
	log.Printf("Remote server option: %s", rawOptions)

	/* Split options */
	options := bytes.Fields(rawOptions)
	for _, option := range options {
		if bytes.HasPrefix(option, []byte("CRAM-")) {
			parts := bytes.SplitN(option, []byte("-"), 3)
			authScheme := parts[1]
			if bytes.Equal(authScheme, []byte("MD5")) {
				authDigest := parts[2]
				self.respAuthorization = self.createAuthorization(authDigest)
			} else {
				log.Panicf("Wrong mechanism: authScheme = %s", authScheme)
			}
		}
	}
}
