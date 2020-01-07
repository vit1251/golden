package mailer

import(
	"os/exec"
	"log"
)

type MailerCompat struct {
}

func NewMailerCompat() (*MailerCompat) {
	mc := new(MailerCompat)
	return mc
}

func (self *MailerCompat) Check() (error) {
	mailer := exec.Command("/opt/binkd-1.0.4/binkd", "-P", "2:5023/24@fidonet", "-p", "/etc/binkd/binkd.cfg")
	err := mailer.Start()
	if err != nil {
		return err
	}
	log.Printf("Waiting for command to finish...")
	err = mailer.Wait()
	log.Printf("mailer = %s", mailer.String() )
	return err
}
