package mailer

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/vit1251/golden/pkg/common"
	"log"
	"net"
	"runtime"
	"time"
)

func (self *Mailer) openSession(host string) (error) {



	return nil
}

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

func (self *Mailer) processSessionSetupState() {
	if (self.sessionSetupState == SessionSetupConnInitState) {
		/* Switch state */
		self.SetSessionSetupState(SessionSetupWaitConnState)
		/* Start connection */
		if conn, err := net.DialTimeout("tcp", self.ServerAddr, time.Millisecond*1200); err != nil {
			log.Printf("Report no connection")
			self.SetSessionSetupState(SessionSetupExitState)
		} else {
			/* Save connection */
			self.conn = conn
			/* Create reader and writer */
			self.reader = bufio.NewReader(conn)
			self.writer = bufio.NewWriter(conn)
			/* Start frame processing */
			go self.processRX()
			go self.processTX()
			/* Connection establish routine: send info and out address */
			//self.writeInfo("SYS", "Vitold Station")
			//self.writeInfo("ZYZ", "Vitold Sedyshev")
			//self.writeInfo("LOC", "Saint-Petersburg, Russia")
			self.writeInfo("NDL", "115200,TCP,BINKP")
			self.writeInfo("TIME", self.GetTime())
			self.writeInfo("OS", self.GetPlatform())
			appName := "GoldenMailer"
			appVersion :=common.GetVersion()
			self.writeInfo("VER", fmt.Sprintf("%s/%s", appName, appVersion))
			self.writeAddress(self.addr)
			// TODO - setup timeout ...
			/* Change state */
			self.SetSessionSetupState(SessionSetupWaitAddrState)
		}
	} else if (self.sessionSetupState == SessionSetupWaitAddrState) {
		nextFrame := <-self.inDataFrames
		if nextFrame.Command {
			if nextFrame.CommandFrame.CommandID == M_NUL {
				key, value, err1 := self.parseInfoFrame(nextFrame.CommandFrame.Body)
				if err1 != nil {
					panic(err1)
				}
				log.Printf("Remote side option: name = %s value = %s", key, value)
				if bytes.Equal(key, []byte("OPT")) {
					self.parseInfoOptFrame(value)
				}
			} else if nextFrame.CommandFrame.CommandID == M_ADR {
				self.SetSessionSetupState(SessionSetupAuthRemoteState)
				self.writePassword(self.respAuthorization)
				self.SetSessionSetupState(SessionSetupIfSecureState)
			} else {
				log.Panicf("Unexpected Frame in state: %v", self.sessionSetupState)
			}
		}
	} else if (self.sessionSetupState == SessionSetupIfSecureState) {
		nextFrame := <-self.inDataFrames
		log.Printf("Auth complete: frame = %+v", nextFrame)
		if nextFrame.Command {
			if nextFrame.CommandFrame.CommandID == M_OK {
				self.SetProtocolState(FileTransferState)
				self.SetFileTransferState(FileTransferInitTransferState)
			}
		} else {
			log.Panicf("Unexpected frame")
		}
	} else {
		log.Panicf("Wrong session setup state: %v", self.sessionSetupState)
	}
}
