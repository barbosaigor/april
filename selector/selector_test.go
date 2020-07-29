package selector

import (
	"testing"
)

const exp = `some text for testing. /{Hello world}\`

func TestPrefixMatch(t *testing.T) {
	if !Match(exp, "some", Prefix) {
		t.Error("Error: Prefix match not matched")
	} else {
		t.Log("Matched with sucess")
	}

	if Match(exp, "txt", Prefix) {
		t.Error("Error: Prefix match not matched")
	} else {
		t.Log("Matched with sucess")
	}
}

func TestInfixMatch(t *testing.T) {
	if !Match(exp, "text", Infix) {
		t.Error("Error: Infix match not matched")
	} else {
		t.Log("Matched with sucess")
	}

	if !Match(exp, "world}", Infix) {
		t.Error("Error: Infix match not matched")
	} else {
		t.Log("Matched with sucess")
	}

	if Match(exp, "txt", Infix) {
		t.Error("Error: Infix match not matched")
	} else {
		t.Log("Matched with sucess")
	}
}

func TestPostfixMatch(t *testing.T) {
	txt := `testing. /{Hello world}\`
	if !Match(exp, txt, Postfix) {
		t.Errorf("Error: Postfix match not matched: %s", txt)
	} else {
		t.Log("Matched with sucess")
	}
	txt = `ld}\`
	if !Match(exp, txt, Postfix) {
		t.Errorf("Error: Postfix match not matched: %s", txt)
	} else {
		t.Log("Matched with sucess")
	}
	txt = "testing. /{Hello world}"
	if Match(exp, txt, Postfix) {
		t.Errorf("Error: Postfix match not matched: %s", txt)
	} else {
		t.Log("Matched with sucess")
	}
}

func TestExactMatch(t *testing.T) {
	if !Match(exp, exp, Exact) {
		t.Errorf("Error: All match not matched: %s", exp)
	} else {
		t.Log("Matched with sucess")
	}
	txt := `ld}\`
	if Match(exp, txt, Exact) {
		t.Errorf("Error: All match not matched: %s", txt)
	} else {
		t.Log("Matched with sucess")
	}
}
