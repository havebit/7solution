package greeting

import "testing"

func TestGreetBob(t *testing.T) {
	given := "Bob"
	expect := "Hello, Bob."

	actual := greet(given)

	if actual != expect {
		t.Errorf("geven name %q, expect %q, but actual %q\n", given, expect, actual)
	}
}

func TestGreetJack(t *testing.T) {
	given := "Jack"
	expect := "Hello, Jack."

	actual := greet(given)

	if actual != expect {
		t.Errorf("geven name %q, expect %q, but actual %q\n", given, expect, actual)
	}
}

func TestGreetAnnonymous(t *testing.T) {
	expect := "Hello, my friend."

	actual := greet()

	if actual != expect {
		t.Errorf("no geven name, expect %q, but actual %q\n", expect, actual)
	}
}

func TestGreetShout(t *testing.T) {
	given := "JERRY"
	expect := "HELLO, JERRY."

	actual := greet(given)

	if actual != expect {
		t.Errorf("geven %q, expect %q, but actual %q\n", given, expect, actual)
	}
}

func TestGreetMultipleNames(t *testing.T) {
	given := []string{"Amy", "Brian", "Charlotte"}
	expect := "Hello, Amy, Brian and Charlotte."

	actual := greet(given...)

	if actual != expect {
		t.Errorf("geven %q, expect %q, but actual %q\n", given, expect, actual)
	}
}
