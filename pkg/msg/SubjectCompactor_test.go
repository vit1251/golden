package msg

import "testing"

func TestSubjectCompactor(t *testing.T) {
	subject := "RE: RE: RE: RE: RE: RE: RU.GOLANG"

	sc := NewSubjectCompactor()
	newSubject := sc.Compact(subject)

	t.Logf("subject = %s", subject)
	t.Logf("newSubject = %s", newSubject)

}

