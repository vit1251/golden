package msg

import "testing"

func TestSubjectCompactorSingle(t *testing.T) {

	subject := "Re: RU.GOLANG"

	sc := NewSubjectCompactor()
	newSubject := sc.Compact(subject)

	t.Logf("subject = %s", subject)
	t.Logf("newSubject = %s", newSubject)

}

func TestSubjectCompactorMultiply(t *testing.T) {

	subject := "RE: RE: RE: RE: RE: RE: RU.GOLANG"

	sc := NewSubjectCompactor()
	newSubject := sc.Compact(subject)

	t.Logf("subject = %s", subject)
	t.Logf("newSubject = %s", newSubject)

}

func TestSubjectCompactorWithCounterSingle(t *testing.T) {

	subject := "RE[1]: RU.GOLANG"

	sc := NewSubjectCompactor()
	newSubject := sc.Compact(subject)

	t.Logf("subject = %s", subject)
	t.Logf("newSubject = %s", newSubject)

}

func TestSubjectCompactorWithCounterMultiply(t *testing.T) {

	subject := "Re[3]: RE[2]: RU.GOLANG"

	sc := NewSubjectCompactor()
	newSubject := sc.Compact(subject)

	t.Logf("subject = %s", subject)
	t.Logf("newSubject = %s", newSubject)

}
