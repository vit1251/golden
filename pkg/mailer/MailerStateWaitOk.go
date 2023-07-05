package mailer

import (
	"fmt"
	"github.com/vit1251/golden/pkg/mailer/stream"
	"log"
)

func mailerStateWaitOkProcessCommandFrame(mailer *Mailer, nextFrame stream.Frame) mailerStateFn {
	command := nextFrame.CommandFrame.CommandID

	if command == stream.M_NUL {
		log.Printf("Warning: ... NUL packet during authorization state ...")
		return mailerStateWaitOk
	} else if command == stream.M_OK {
		log.Printf("Auth - OK")
		return mailerStateOpts
	} else if command == stream.M_ERR {
		log.Printf("AUTH - ERROR: err = %+v", nextFrame.CommandFrame.Body)
		mailer.report.SetStatus(fmt.Sprintf("Authorization error: reason = %+v", nextFrame.CommandFrame.Body))
	}

	return nil
}

func mailerStateWaitOkProcessFrame(mailer *Mailer, nextFrame stream.Frame) mailerStateFn {

	if nextFrame.IsCommandFrame() {
		return mailerStateWaitOkProcessCommandFrame(mailer, nextFrame)
	} else {
		log.Printf("Unexpected frame: frame = %+v", nextFrame)
	}

	return mailerStateWaitOk
}

func mailerStateWaitOk(mailer *Mailer) mailerStateFn {

	select {

	case nextFrame := <-mailer.stream.InFrame:
		return mailerStateWaitOkProcessFrame(mailer, nextFrame)

		//	case <-mailer.WaitAddrTimeout:
		//		log.Printf("Timeout!")

	}

	return mailerStateWaitOk

}
