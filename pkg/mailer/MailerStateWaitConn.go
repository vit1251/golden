package mailer

import (
	"log"
	"fmt"
	"github.com/vit1251/golden/internal/common"
	"time"
)

func makeSystemTime() string {
	now := time.Now().Format(time.RFC822)
	return now
}

func makeOperationSystemName() string {
	return commonfunc.GetPlatform()
}

func makeVersionString() string {
	appName := "GoldenMailer"
	appVersion := commonfunc.GetVersion()
	protocolVersion := "binkp/1.0"
	return fmt.Sprintf("%s/%s %s", appName, appVersion, protocolVersion)
}

func mailerStateWaitConnProcessWelcome(mailer *Mailer) {

	/* Send M_NUL frames with system info (optional) */
	if username := mailer.GetUserName(); username != "" {
		mailer.stream.WriteInfo("ZYZ", username)
	}
	if stationName := mailer.GetSystemName(); stationName != "" {
		mailer.stream.WriteInfo("SYS", stationName)
	}
	if location := mailer.GetLocation(); location != "" {
		mailer.stream.WriteInfo("LOC", location)
	}
	mailer.stream.WriteInfo("TIME", makeSystemTime())
	mailer.stream.WriteInfo("OS", makeOperationSystemName())
	mailer.stream.WriteInfo("VER", makeVersionString())

	/* Send M_ADR frame with system address */
	mailer.stream.WriteAddress(mailer.GetAddr())

}

func mailerStateWaitConn(mailer *Mailer) mailerStateFn {
	err := mailer.connect()
	if err != nil {
		log.Printf("mailer: connection timeout")
		return mailerStateEnd
	}
	mailerStateWaitConnProcessWelcome(mailer)
	return mailerStateAdditionalStep
}
