package ui2

import (
	"bufio"
	"fmt"
	"go.uber.org/dig"
	"io"
	"log"

	"github.com/gliderlabs/ssh"
)

type Service struct {
}

func NewService(c *dig.Container) *Service {
	return new(Service)
}

func (session *Service) handleSession(s ssh.Session) {

	//s.User()

	io.WriteString(s, fmt.Sprintf("\x1B[H"))
	io.WriteString(s, fmt.Sprintf("\x1B[2J"))

	io.WriteString(s, fmt.Sprintf("▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒\n"))
	io.WriteString(s, fmt.Sprintf("                                                                           \n"))
	io.WriteString(s, fmt.Sprintf(" ::::::::   ::::::::  :::        :::::::::  :::::::::: ::::    :::         \n"))
	io.WriteString(s, fmt.Sprintf(":+:    :+: :+:    :+: :+:        :+:    :+: :+:        :+:+:   :+:         \n"))
	io.WriteString(s, fmt.Sprintf("+:+        +:+    +:+ +:+        +:+    +:+ +:+        :+:+:+  +:+         \n"))
	io.WriteString(s, fmt.Sprintf(":#:        +#+    +:+ +#+        +#+    +:+ +#++:++#   +#+ +:+ +#+         \n"))
	io.WriteString(s, fmt.Sprintf("+#+   +#+# +#+    +#+ +#+        +#+    +#+ +#+        +#+  +#+#+#         \n"))
	io.WriteString(s, fmt.Sprintf("#+#    #+# #+#    #+# #+#        #+#    #+# #+#        #+#   #+#+#         \n"))
	io.WriteString(s, fmt.Sprintf(" ########   ########  ########## #########  ########## ###    ####         \n"))
	io.WriteString(s, fmt.Sprintf("                                                                           \n"))
	io.WriteString(s, fmt.Sprintf("                                                                           \n"))
	io.WriteString(s, fmt.Sprintf("       :::::::::   :::::::: ::::::::::: ::::    ::: :::::::::::            \n"))
	io.WriteString(s, fmt.Sprintf("       :+:    :+: :+:    :+:    :+:     :+:+:   :+:     :+:                \n"))
	io.WriteString(s, fmt.Sprintf("       +:+    +:+ +:+    +:+    +:+     :+:+:+  +:+     +:+                \n"))
	io.WriteString(s, fmt.Sprintf("       +#++:++#+  +#+    +:+    +#+     +#+ +:+ +#+     +#+                \n"))
	io.WriteString(s, fmt.Sprintf("       +#+        +#+    +#+    +#+     +#+  +#+#+#     +#+                \n"))
	io.WriteString(s, fmt.Sprintf("       #+#        #+#    #+#    #+#     #+#   #+#+#     #+#                \n"))
	io.WriteString(s, fmt.Sprintf("       ###         ######## ########### ###    ####     ###                \n"))
	io.WriteString(s, fmt.Sprintf("                                                                           \n"))
	io.WriteString(s, fmt.Sprintf("▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒\n"))
	io.WriteString(s, fmt.Sprintf("                             Version 1.2.12                                \n"))
	io.WriteString(s, fmt.Sprintf("▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒\n"))

	text, err := bufio.NewReader(s).ReadString('\n')
	pty, _, check := s.Pty()

	log.Printf("Terminal size: %d", pty.Window.Height)
	log.Printf("some value = %v", check)

	if err != nil {
		panic("GetLines: " + err.Error())
	}

	io.WriteString(s, fmt.Sprintf("ton texte %s\n", text))
}

func (s *Service) Start() {
	ssh.ListenAndServe("127.0.0.1:2222", s.handleSession)
}
