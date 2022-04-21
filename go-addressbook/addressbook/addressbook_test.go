package addressbook

import (
	"testing"
)

func TestCreatePerson(t *testing.T) {
	person := CreatePerson(
		1,
		"John Doe",
		"john@doe.com",
		&PhoneNumber{
			Number: "012399123123",
			Type:   Person_HOME,
		},
		&PhoneNumber{
			Number: "123123001230",
			Type:   Person_WORK,
		},
	)
	got := person.GetEmail()
	want := "john@doe.com"
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
	got = person.GetName()
	want = "John Doe"
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}

	got = person.GetPhones()[0].Number
	want = "012399123123"
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
	got = person.GetPhones()[0].Type.String()
	want = Person_HOME.String()
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}

}
func BenchmarkCreatePerson(t *testing.B) {
	person := CreatePerson(
		1,
		"John Doe",
		"john@doe.com",
		&PhoneNumber{
			Number: "012399123123",
			Type:   Person_HOME,
		},
		&PhoneNumber{
			Number: "123123001230",
			Type:   Person_WORK,
		},
	)
	got := person.GetEmail()
	want := "john@doe.com"
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
	got = person.GetName()
	want = "John Doe"
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}

	got = person.GetPhones()[0].Number
	want = "012399123123"
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
	got = person.GetPhones()[0].Type.String()
	want = Person_HOME.String()
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}

}

func BenchmarkCreateAddressBook(t *testing.B) {
	var got, want interface{}

	person1 := CreatePerson(
		1,
		"John Doe",
		"john@doe.com",
		&PhoneNumber{
			Number: "012399123123",
			Type:   Person_HOME,
		},
		&PhoneNumber{
			Number: "123123001230",
			Type:   Person_WORK,
		},
	)
	person2 := CreatePerson(
		2,
		"Jean Doe",
		"jean@doe.com",
		&PhoneNumber{
			Number: "012381230123",
			Type:   Person_HOME,
		},
		&PhoneNumber{
			Number: "1239900991230",
			Type:   Person_WORK,
		},
	)

	addressbook := CreateAddressBook(person1, person2)

	got = len(addressbook.GetPerson())
	want = 2
	if got != want {
		t.Errorf("Length of address book: Got %q, expecting %q", got, want)
	}
	got = addressbook.GetPerson()[0].GetName()
	want = "John Doe"
	if got != want {
		t.Errorf("Expecting %q, got %q", want, got)
	}
}
func TestCreateAddressBook(t *testing.T) {
	var got, want interface{}

	person1 := CreatePerson(
		1,
		"John Doe",
		"john@doe.com",
		&PhoneNumber{
			Number: "012399123123",
			Type:   Person_HOME,
		},
		&PhoneNumber{
			Number: "123123001230",
			Type:   Person_WORK,
		},
	)
	person2 := CreatePerson(
		2,
		"Jean Doe",
		"jean@doe.com",
		&PhoneNumber{
			Number: "012381230123",
			Type:   Person_HOME,
		},
		&PhoneNumber{
			Number: "1239900991230",
			Type:   Person_WORK,
		},
	)

	addressbook := CreateAddressBook(person1, person2)

	got = len(addressbook.GetPerson())
	want = 2
	if got != want {
		t.Errorf("Length of address book: Got %q, expecting %q", got, want)
	}
	got = addressbook.GetPerson()[0].GetName()
	want = "John Doe"
	if got != want {
		t.Errorf("Expecting %q, got %q", want, got)
	}
}
