package msg

import "testing"

func TestSubjectCompactorSingle(t *testing.T) {

	subject := "Re: RU.GOLANG"

	sc := NewSubjectCompactor()
	newSubject := sc.Compact(subject)

	t.Logf("subject = %s", subject)
	t.Logf("newSubject = %s", newSubject)

}

func TestSubjectCompactorWithMultiply(t *testing.T) {

	subject := "RE: RE: RE: RE: RE: RE: RU.GOLANG"

	sc := NewSubjectCompactor()
	newSubject := sc.Compact(subject)

	t.Logf("subject = %s", subject)
	t.Logf("newSubject = %s", newSubject)

}

func TestSubjectCompactorWithCounter(t *testing.T) {

	subject := "RE[1]: RU.GOLANG"

	sc := NewSubjectCompactor()
	newSubject := sc.Compact(subject)

	t.Logf("subject = %s", subject)
	t.Logf("newSubject = %s", newSubject)

}
