package mailer

//type escapeExample struct {
//	Name string
//	Result string
//}
//
//func TestEscape(t *testing.T) {
//
//	var ok bool = true
//	var examples []escapeExample
//	examples = append(examples, escapeExample{
//		Name: "hello 123.txt",
//		Result: "hello\\x20123.txt",
//	})
//	examples = append(examples, escapeExample{
//		Name: "hello-123.txt",
//		Result: "hello-123.txt",
//	})
//	examples = append(examples, escapeExample{
//		Name: "hello|123",
//		Result: "hello\\x7c123",
//	})
//
//	for idx, example := range examples {
//		newName := escape(example.Name, MODERN)
//		if newName != example.Result {
//			t.Logf("Error on %d example: %s -> %s but expect %s", idx, example.Name, newName, example.Result)
//			ok = false
//		}
//	}
//
//	if !ok {
//		t.Fatal("Too many errors")
//	}
//
//}
//
//type unescapeExample struct {
//	Name string
//	Result string
//}
//
//func TestUnescape(t *testing.T) {
//
//	var ok bool = true
//	var examples []unescapeExample
//	examples = append(examples, unescapeExample{
//		Name: "hello\\x20123.txt",
//		Result: "hello 123.txt",
//	})
//
//	for idx, example := range examples {
//		newName, _ := unescape(example.Name)
//		if newName != example.Result {
//			t.Logf("Error on %d example: %s -> %s but expect %s", idx, example.Name, newName, example.Result)
//			ok = false
//		}
//	}
//
//	if !ok {
//		t.Fatal("Too many errors")
//	}
//
//}