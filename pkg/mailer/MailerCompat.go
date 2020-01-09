package mailer

import(
	"os/exec"
	"os"
	"bufio"
	"log"
	"fmt"
	"io"
)

type MailerCompat struct {
	OutboundPath   string          /* Outbound path        */
	Files        []string          /* File strings         */
}

func NewMailerCompat() (*MailerCompat) {
	mc := new(MailerCompat)
	mc.OutboundPath = "/var/spool/ftn/outb/139f0018.dlo"
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

func (self *MailerCompat) RecoverTransmitQueue() (error) {

	//
	log.Printf("Recove in %s", self.OutboundPath)

	stream, err1 := os.Open(self.OutboundPath)
	log.Printf("err = %+v", err1)
	if err1 != nil {
		return err1
	}
	defer stream.Close()

	//
	scanner := bufio.NewScanner(stream)
	for scanner.Scan() {
		newLine := scanner.Text()
		log.Printf("Queue: rec = %s", newLine)   //RegisterTransmitFile
		if newLine[0] == '^' {
			newFile := string(newLine[1:])
			self.Files = append(self.Files, newFile) // RegisterTransmitFile(newLine)
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func (self *MailerCompat) CopyFile(src string, dst string) (error) {

	log.Printf("Copy %s -> %s", src, dst)

	/* Open */
	in, err1 := os.Open(src)
	if err1 != nil {
		return err1
	}
	defer in.Close()

	/* Create */
	out, err2 := os.Create(dst)
	if err2 != nil {
		return err2
	}
	defer out.Close()

	/* Copy */
	size, err3 := io.Copy(out, in)
	log.Printf("write: size = %d", size)
	if err3 != nil {
		return err3
	}

	return nil
}

func (self *MailerCompat) SaveTransmitQueue() (error) {

	stream, err1 := os.Create(self.OutboundPath)
	if err1 != nil {
		return err1
	}
	defer stream.Close()

	for _, newFile := range self.Files {
		newRec := fmt.Sprintf("^%s\n", newFile)
		stream.WriteString(newRec)
	}

	return nil
}

func (self *MailerCompat) registerTransmitFile(filename string) (error) {
	self.Files = append(self.Files, filename)
	return nil
}

func (self *MailerCompat) TransmitFile(pktName string) (error) {


	dirName := "/var/spool/ftn/outb/139f0018.sep"
	log.Printf("Make dir %s", dirName)
	err1 := os.MkdirAll(dirName, os.ModeDir)
	if err1 != nil {
		return err1
	}

	newFile := fmt.Sprintf("%s/%s", dirName, pktName)
	log.Printf("Prepare new path: dst = %s", newFile)

	/* Make copy */
	log.Printf("Make copy: %s -> %s", pktName, newFile)
	err2 := self.CopyFile(pktName, newFile)
	if err2 != nil {
		return err2
	}

	/* Recover queue */
	self.RecoverTransmitQueue()
	log.Printf("Recover: %+v", self.Files)

	/* Register new item in queue */
	self.registerTransmitFile(newFile)
	log.Printf("Register: %s", newFile)

	/* Save queue */
	self.SaveTransmitQueue()
	log.Printf("Store: %+v", self.Files)

	return nil
}
