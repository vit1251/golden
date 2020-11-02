package msg

import (
	"fmt"
	"strings"
)

type SubjectCompactor struct {
}

func NewSubjectCompactor() *SubjectCompactor {
	return new(SubjectCompactor)
}

func (self SubjectCompactor) hasPrefix(subject string) bool {
	return strings.HasPrefix(subject, "RE:")
}

func (self SubjectCompactor) Compact(subject string) string {

	var level int = 0
	for self.hasPrefix(subject) {
		subject = self.trimPrefix(subject)
		subject = strings.Trim(subject, " ")
		level += 1
	}

	newSubject := fmt.Sprintf("RE[%d]: %s", level, subject)

	return newSubject

}

func (self SubjectCompactor) trimPrefix(subject string) string {
	subject = strings.TrimPrefix(subject, "RE:")
	return subject
}
