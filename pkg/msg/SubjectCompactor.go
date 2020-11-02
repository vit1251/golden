package msg

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type SubjectCompactor struct {
}

func NewSubjectCompactor() *SubjectCompactor {
	return new(SubjectCompactor)
}

func (self SubjectCompactor) Compact(subject string) string {

	var newSubject string = subject
	var level int = 0

	/* Parse without number */
	for {
		re1 := regexp.MustCompile(`^[Rr][Ee]\:`)
		match1 := re1.FindStringSubmatch(newSubject)
		if match1 != nil {
			newSubject = newSubject[3:]
			newSubject = strings.TrimLeft(newSubject, " ")
			level += 1
		} else {
			break
		}
	}

	/* Parse with number */
	re2 := regexp.MustCompile(`^[Rr][Ee]\[(\d+)\]:`)
	match2 := re2.FindStringSubmatch(subject)
	if match2 != nil {
		size := len(match2[0])
		newSubject = newSubject[size:]
		newSubject = strings.TrimLeft(newSubject, " ")
		fmt.Printf("re = %q", match2)
		num, _ := strconv.Atoi(match2[1])
		level += num
	}

	/* Remove leading spaces */
	if level == 0 {
		newSubject = fmt.Sprintf("RE: %s", newSubject)
	} else {
		newSubject = fmt.Sprintf("RE[%d]: %s", level + 1, newSubject)
	}

	return newSubject

}

