package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
    name := "Gladys"
    want := regexp.MustCompile(`\b`+name+`\b`)
    msg, err := Hello("Gladys")
    if !want.MatchString(msg) || err != nil {
        t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
    msg, err := Hello("")
    if msg != "" || err == nil {
        t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}

func TestHelloMultiple(t *testing.T){
	names := []string{"Ace", "Mark", "Eduardo"}
	msgs, err := HelloMultiple(names)

	if(err != nil){
		t.Fatalf(`HelloMultiple([]string{"Ace", "Mark", "Eduardo"}) = %q, %v, error should be nil`, msgs, err)
	}

	for _,name := range names {
		want := regexp.MustCompile(`\b`+name+`\b`)
		if !want.MatchString(msgs[name]) {
			t.Fatalf(`HelloMultiple([]string{"Ace", "Mark", "Eduardo"}) = %q, %v, want match for %#q, nil`, msgs[name], err, want)
		}
	}

}

func TestHelloMultipleEmpty(t *testing.T){
	names := []string{"Ace", "", "Eduardo"}
	msgs, err := HelloMultiple(names)

	if(err == nil){
		t.Fatalf(`HelloMultiple([]string{"Ace", "", "Eduardo"}) = %q, %v, error should NOT be nil`, msgs, err)
	}
}