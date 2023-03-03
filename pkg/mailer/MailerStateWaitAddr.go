package mailer

import (
	"github.com/vit1251/golden/pkg/mailer/stream"
	"log"
)

func mailerStateWaitAddrProcessFrame(mailer *Mailer, nextFrame stream.Frame) mailerStateFn {

	/* M_ADR frame received */
	if nextFrame.IsCommandFrame() {
		if nextFrame.CommandFrame.CommandID == stream.M_ADR {
			return mailerStateAuthRemote
		}
	}

	/* M_BSY frame received */
	if nextFrame.IsCommandFrame() {
		if nextFrame.CommandFrame.CommandID == stream.M_BSY {
			mailer.report.SetStatus("Remote system is BUSY")
			log.Printf("Remote system is BUSY")
			return mailerStateEnd
		}
	}

	/* M_ERR frame received */
	if nextFrame.IsCommandFrame() {
		if nextFrame.CommandFrame.CommandID == stream.M_BSY {
			log.Printf("Remote system is ERROR")
			return mailerStateEnd
		}
	}

	/* M_NUL frame received */
	if nextFrame.IsCommandFrame() {
		if nextFrame.CommandFrame.CommandID == stream.M_NUL {
			mailer.processNulFrame(nextFrame)
		}
	}

	return mailerStateWaitAddr
}

func mailerStateWaitAddr(mailer *Mailer) mailerStateFn {

	nextFrame := <-mailer.stream.InFrame
	return mailerStateWaitAddrProcessFrame(mailer, nextFrame)

}
