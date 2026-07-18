package mailer

import (
    "log"
    "reflect"
)

type CaseRoutine struct {
    Case      reflect.SelectCase
    Callback  func(ok bool) mailerStateFn
}

func mailerStateSwitch(mailer *Mailer) mailerStateFn {

    /* RxState is RxDone and TxState is TxDone */
    if mailer.rxState == RxDone && mailer.txState == TxDone {
	log.Printf("session complete")
	return mailerStateEnd
    }

    cases := make([]CaseRoutine, 0)

    /* Timer Expired */
    cases = append(cases, CaseRoutine{
        Case: reflect.SelectCase{
            Dir:  reflect.SelectRecv,
	    Chan: reflect.ValueOf(mailer.timerCh),
        },
        Callback: func(ok bool) mailerStateFn {
            log.Printf("Timeout")
            mailer.report.SetStatus("Session timeout")
            return mailerStateEnd
        },
    })

    /* Data available in Input Buffer */
    if mailer.rxState != RxDone {
        cases = append(cases, CaseRoutine{
            Case: reflect.SelectCase{
  	        Dir:  reflect.SelectRecv,
	        Chan: reflect.ValueOf(mailer.stream.InFrameReady),
            },
            Callback: func(ok bool) mailerStateFn {
                if ok {
                    mailer.rxRoutineResult = ReceiveRoutine(mailer)
                    return mailerStateReceive
                } else {
	            return mailerStateEnd
	        }
            },
        })
    }

    /* Free space exists in output buffer */
    if mailer.txState != TxDone {
        cases = append(cases, CaseRoutine{
            Case: reflect.SelectCase{
                Dir:  reflect.SelectSend,
                Chan: reflect.ValueOf(mailer.stream.OutFrameReady),
                Send: reflect.ValueOf(0),
            },
            Callback: func(ok bool) mailerStateFn{
                mailer.txRoutineResult = TransmitRoutine(mailer)
                return mailerStateTransmit
            },
        })
    }

    // chosen - индекс сработавшего кейса в срезе cases
    // value  - полученное значение (тип reflect.Value)
    // ok     - открыт ли канал (false, если закрыт)
    var selectCases []reflect.SelectCase
    for _, c := range cases {
	selectCases = append(selectCases, c.Case)
    }
    chosen, _, ok := reflect.Select(selectCases)

    userCase := cases[chosen]
    return userCase.Callback(ok)

}
