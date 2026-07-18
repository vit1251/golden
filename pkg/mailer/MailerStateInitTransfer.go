package mailer

import "time"

func mailerStateInitTransfer(mailer *Mailer) mailerStateFn {

    mailer.setTimer(120*time.Second)
    mailer.setRxState(RxWaitF)
    mailer.setTxState(TxGNF)

    return mailerStateSwitch
}
