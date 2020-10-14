package mailer

type ProtocolState int

func (s ProtocolState) String() string {
	return [...]string{"SS", "FT"}[s]
}

const (
	SessionSetupState ProtocolState = 0
	FileTransferState ProtocolState = 1
)

type SessionSetupStageState int

func (s SessionSetupStageState) String() string {
	return [...]string{
		"SessionSetupConnInitState",
		"SessionSetupWaitConnState",
		"SessionSetupSendPasswdState",
		"SessionSetupWaitAddrState",
		"SessionSetupAuthRemoteState",
		"SessionSetupIfSecureState",
		"SessionSetupWaitOkState",
		"SessionSetupExitState",
	}[s]
}

const (
	SessionSetupConnInitState   SessionSetupStageState = 0
	SessionSetupWaitConnState   SessionSetupStageState = 1
	SessionSetupSendPasswdState SessionSetupStageState = 2
	SessionSetupWaitAddrState   SessionSetupStageState = 3
	SessionSetupAuthRemoteState SessionSetupStageState = 4
	SessionSetupIfSecureState   SessionSetupStageState = 5
	SessionSetupWaitOkState     SessionSetupStageState = 6
	SessionSetupExitState       SessionSetupStageState = 7
)