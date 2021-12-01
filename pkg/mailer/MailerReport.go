package mailer

import (
	"log"
	"time"
)

type MailerReport struct {
	startSession time.Time /* Mailer start session Date and Time  */
	stopSession  time.Time /* Mailer stop session Date and Time   */
	status       string    /* Mailer status                       */
	remoteIdent  string    /* Mailer remote ident                 */
}

func NewMailerReport() *MailerReport {
	return new(MailerReport)
}

func (self *MailerReport) SetStartSession(now time.Time) {
	self.startSession = now
}

func (self MailerReport) GetDuration() time.Duration {
	return self.stopSession.Sub(self.startSession)
}

func (self MailerReport) Dump() {

	log.Printf("--- Mailer session report (QSL) ---\n"+
		"    Start session: %+v\n"+
		"     Stop session: %+v\n"+
		" Session duration: %+v\n"+
		"           Status: %+v\n"+
		"     Remote ident: %+v",
		self.startSession,
		self.stopSession,
		self.GetDuration(),
		self.status,
		self.remoteIdent,
	)

}

func (self *MailerReport) SetStopSession(now time.Time) {
	self.stopSession = now
}

func (self *MailerReport) SetStatus(s string) {
	self.status = s
}

func (self *MailerReport) SetRemoteIdent(remoteIdent string) {
	self.remoteIdent = remoteIdent
}
