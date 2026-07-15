package mailer

import (
    "log"
    "time"
)

type MailerReport struct {
    sessionID    int64      /* Session number                      */
    startSession time.Time  /* Mailer start session Date and Time  */
    stopSession  time.Time  /* Mailer stop session Date and Time   */
    inFileCount  int        /* Received file count                 */
    outFileCount int        /* Sended file count                   */
    inSize       int64      /* Received bytes                      */
    outSize      int64      /* Sended bytes                        */
    status       string     /* Mailer status                       */
    remoteIdent  string     /* Mailer remote ident                 */
}

func NewMailerReport() *MailerReport {
    report := new(MailerReport)
    report.sessionID = 0
    report.startSession = time.Now()
    report.stopSession = time.Now()
    report.inFileCount = 0
    report.outFileCount = 0
    report.inSize = 0
    report.outSize = 0
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

func (r *MailerReport) SetSessionStop(now time.Time) { r.stopSession = now }
func (r *MailerReport) GetSessionStop() time.Time { return r.stopSession }
func (r *MailerReport) SetStatus(s string) { r.status = s }
func (r *MailerReport) GetStatus() string { return r.status }
func (r *MailerReport) SetRemoteIdent(remoteIdent string) { r.remoteIdent = remoteIdent }

func (r *MailerReport) GetInFileCount() int { return r.inFileCount }
func (r *MailerReport) SetInFileCount(v int) { r.inFileCount = v }
func (r *MailerReport) GetOutFileCount() int { return r.outFileCount }
func (r *MailerReport) SetOutFileCount(v int) { r.outFileCount = v }

func (r *MailerReport) GetInSize() int64 { return r.inSize }
func (r *MailerReport) SetInSize(v int64) { r.inSize = v }
func (r *MailerReport) GetOutSize() int64 { return r.outSize }
func (r *MailerReport) SetOutSize(v int64) { r.outSize = v }
