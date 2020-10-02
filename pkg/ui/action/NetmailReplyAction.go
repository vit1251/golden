package action

import "net/http"

type NetmailReplyAction struct {
	Action
}

func NewNetmailReplyAction() *NetmailReplyAction {
	return new(NetmailReplyAction)
}

func (self *NetmailReplyAction)  ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
