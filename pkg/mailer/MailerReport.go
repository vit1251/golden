package mailer

import (
	"log"
	"time"
)

type MailerReport struct {
	sessionID    int64     /* Session number                      */
	startSession time.Time /* Mailer start session Date and Time  */
	stopSession  time.Time /* Mailer stop session Date and Time   */
	status       string    /* Mailer status                       */
	remoteIdent  string    /* Mailer remote ident                 */
}

func NewMailerReport() *MailerReport {
	report := new(MailerReport)
	report.sessionID = 0
	report.startSession = time.Now()
	report.stopSession = time.Now()
	report.status = "N/A"
	report.remoteIdent = ""
	return report
}

func (self *MailerReport) GetSessionID() int64 {
	return self.sessionID
}

func (self *MailerReport) SetSessionID(sessionId int64) {
	self.sessionID = sessionId
}

func (self *MailerReport) SetSessionStart(now time.Time) {
	self.startSession = now
}

func (self *MailerReport) GetSessionStart() time.Time {
	return self.startSession
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

func (self *MailerReport) SetSessionStop(now time.Time) {
	self.stopSession = now
}

func (self *MailerReport) GetSessionStop() time.Time {
	return self.stopSession
}

func (self *MailerReport) SetStatus(s string) {
	self.status = s
}

func (self *MailerReport) GetStatus() string {
	return self.status
}

func (self *MailerReport) SetRemoteIdent(remoteIdent string) {
	self.remoteIdent = remoteIdent
}
