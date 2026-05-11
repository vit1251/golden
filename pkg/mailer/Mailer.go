package mailer

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/vit1251/golden/pkg/mailer/auth"
	stream2 "github.com/vit1251/golden/pkg/mailer/stream"
	"github.com/vit1251/golden/pkg/mailer/util"
	"github.com/vit1251/golden/pkg/queue"
	"github.com/vit1251/golden/pkg/registry"
	"log"
	"os"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"time"
	"runtime/trace"
	"context"
	"errors"
)

var ErrConnectionTimeout = errors.New("connection timeout")
var ErrReadTimeout = errors.New("read timeout")
var ErrWriteTimeout = errors.New("write timeout")
var ErrDisconnectTimeout = errors.New("disconnect timeout")

type Mailer struct {
	registry *registry.Container /* ???                         */

	activeState mailerStateFn /* Mailer state                */

	rxState RxState /* RX FSM                      */
	txState TxState /* TX FSM                      */

	stream *stream2.MailerStream /* ???                         */

	reader *bufio.Reader /* RX network stream           */
	writer *bufio.Writer /* TX network stream           */

	connectionCount int            /* Active session count        */
	wait            sync.WaitGroup /* Sync                        */

	addr       string /* Network address             */
	secret     string /* Secret password             */
	ServerAddr string /* Server IPv4 or FQDN address */

	inboundDirectory  string /* ???                         */
	outboundDirectory string /* ???                         */

	respAuthorization string /* ???                         */

	recvStream *os.File /* Stream using in Rx routines */
	sendStream *os.File /* Stream using in Tx routines */

	readSize  int64 /* Size incoming download      */
	writeSize int64 /* Size outgoing upload        */

	connComplete chan int /* ???                         */
	recvUnix     int      /* ???                         */

	sendName *queue.FileEntry /* Upload entry                */
	recvName *queue.FileEntry /* Download entry              */

	workInbound  string /* ???                         */
	workOutbound string /* ???                         */

	work string /* ???                         */

	InFileCount  int
	OutFileCount int

	workPath   string
	systemName string
	userName   string
	location   string

	rxRoutineResult ReceiveRoutineResult
	txRoutineResult TransmitRoutineResult

	outboundQueue []queue.FileEntry
	inboundQueue  []queue.FileEntry

	queue *util.TheQueue /* TheQueue      */

	pendingFiles util.Directory
	chunk        []byte
	report       *MailerReport /* Mailer report             */
}

func NewMailer(r *registry.Container) *Mailer {
	return &Mailer{
		connComplete: make(chan int),
		registry: r,
		queue: util.NewTheQueue(),
		connectionCount: 0,
	}
}

func (m *Mailer) GetReport() *MailerReport {
	return m.report
}

func (m *Mailer) writeTrafic(mail int, data int) {
	raw := fmt.Sprintf("TRF %d %d", mail, data)
	m.stream.WriteComment(raw)
}

func (m *Mailer) SetTempOutbound(workOutbound string) {
	m.workOutbound = workOutbound
}

func (m *Mailer) SetAddr(addr string) {
	m.addr = addr
}

func (m *Mailer) SetSecret(secret string) {
	m.secret = secret
}

func (m *Mailer) Start() (error, *MailerReport) {

	if m.connectionCount > 0 {
		return fmt.Errorf("fido session alredy in progress"), nil
	}

	// Initialize new report
	m.report = NewMailerReport()

	// Add wait
	m.wait.Add(1)

	// Play!
	go m.run()

	return nil, m.report
}

func (m *Mailer) IsTransmitting() bool {
	return m.sendName != nil
}

func (m *Mailer) IsReceiving() bool {
	return m.recvName != nil
}

func makeFuncName(i interface{}) string {
	wideName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	parts := strings.Split(wideName, ".")
	ourName := parts[len(parts)-1]
	return ourName
}

func (m *Mailer) run() {

        // Trace
        f, err := os.Create("trace.out")
        if err != nil {
		log.Printf("mailer: No trace")
        }
        defer f.Close()
        trace.Start(f)
        defer trace.Stop()

	ctx := context.Background()
	ctx, task := trace.NewTask(ctx, "run")
	defer task.End()

	m.report.SetSessionStart(time.Now())

	/* Reset active state */
	m.activeState = mailerStateStart

	/* Start processing */
	log.Printf("mailer: Start mailer routine")
	for {
		trace.Logf(ctx, "mailer", "state = %s", makeFuncName(m.activeState))
		newState := m.activeState(m)
		trace.Logf(ctx, "mailer", "change state: %s -> %s", makeFuncName(m.activeState), makeFuncName(newState))
		m.activeState = newState
		/* Stop processing when done */
		if newState == nil {
			log.Printf("mailer: Reach Exit state")
			break
		}
	}
	log.Printf("mailer: Stop mailer routine")

	/* Update report */
	m.report.SetSessionStop(time.Now())

	/* Close connection */
	m.wait.Done()

	/* Remove counter */
	m.connectionCount = m.connectionCount - 1

}

func (m *Mailer) Wait() *MailerReport {

	/* Wait session complete */
	m.wait.Wait()

	/* Dump report */
	m.report.Dump()

	return m.report
}

func (m *Mailer) SetServerAddr(addr string) {
	if !strings.Contains(addr, ":") {
		defaultPort := 24554
		addr = fmt.Sprintf("%s:%d", addr, defaultPort)
	}
	m.ServerAddr = addr
}

func (m *Mailer) SetInboundDirectory(inb string) {
	m.inboundDirectory = inb
}

func (m *Mailer) SetOutboundDirectory(outb string) {
	m.outboundDirectory = outb
}

func (m *Mailer) SetTempInbound(workInbound string) {
	m.workInbound = workInbound
}

func (m *Mailer) SetTemp(work string) {
	m.work = work
}

func (m *Mailer) GetWorkOutbound() string {
	return m.workOutbound
}

func (m *Mailer) GetAddr() string {
	return m.addr
}

func (m *Mailer) GetSystemName() string {
	return m.systemName
}

func (m *Mailer) GetUserName() string {
	return m.userName
}

func (m *Mailer) GetLocation() string {
	return m.location
}

func (m *Mailer) SetLocation(location string) {
	m.location = location
}

func (self *Mailer) SetUserName(name string) {
	self.userName = name
}

func (m *Mailer) SetStationName(name string) {
	m.systemName = name
}

func (m *Mailer) AddOutbound(path queue.FileEntry) {
	m.outboundQueue = append(m.outboundQueue, path)
}

func (m *Mailer) createAuthorization(chData []byte) string {
	a := auth.NewAuthorizer()
	a.SetChallengeData(chData)
	a.SetSecret([]byte(m.secret))
	responseDigest, err := a.CalculateDigest()
	if err != nil {
		panic(err)
	}
	password := fmt.Sprintf("%s-%s-%s", "CRAM", "MD5", responseDigest)
	return password
}

func (m *Mailer) processNulOptFrame(rawOptions []byte) {

	log.Printf("Mailer: Remote server option: %s", rawOptions)

	/* Split options */
	options := bytes.Fields(rawOptions)
	for _, option := range options {
		if bytes.HasPrefix(option, []byte("CRAM-")) {
			parts := bytes.SplitN(option, []byte("-"), 3)
			authScheme := parts[1]
			if bytes.Equal(authScheme, []byte("MD5")) {
				authDigest := parts[2]
				log.Printf("Use %s as digest", authDigest)
				m.respAuthorization = m.createAuthorization(authDigest)
			} else {
				log.Panicf("Wrong mechanism: authScheme = %s", authScheme)
			}
		}
	}

}

func (m *Mailer) processNulFrame(nextFrame stream2.Frame) {

	packet := nextFrame.CommandFrame.Body
	values := bytes.SplitN(packet, []byte(" "), 2)

	if len(values) == 2 {

		key := values[0]
		value := values[1]

		log.Printf("Mailer: Remote side M_NUL packet: name = %s value = %s", key, value)

		if bytes.Equal(key, []byte("OPT")) {
			m.processNulOptFrame(value)
		}

	} else {
		log.Printf("Mailer: Remote side M_NUL parse error: packet = %s", packet)
	}

}

func (m *Mailer) GetWork() string {
	return m.work
}

func (m *Mailer) IsReceiveName(name string) bool {
	if m.recvName != nil {
		if name == m.recvName.Name {
			return true
		}
	}
	return false
}

func (m *Mailer) IsTransmitName(name string) bool {
	if m.sendName != nil {
		if name == m.sendName.Name {
			return true
		}
	}
	return false
}

func (m *Mailer) readFrame() (stream2.Frame, error) {
	var timeout time.Duration = 15 * time.Second
	select {
	case frame := <-m.stream.InFrame:
		return frame, nil
	case <-time.After(timeout):
		return stream2.Frame{}, ErrReadTimeout
	}
}

func (m *Mailer) writeFrame(f stream2.Frame) error {
//	var timeout time.Duration = 15 * time.Second
	return nil
}

func (m *Mailer) connect() error {
	var timeout time.Duration = 15 * time.Second
	select {
	case <-m.stream.InFrameReady:
		return nil

	case <-time.After(timeout):
		return ErrConnectionTimeout
	}
}

func (m *Mailer) disconnect(timeout int16) error {
	return nil
}
