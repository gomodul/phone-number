package phone_number

import (
	"fmt"
	"testing"
)

func TestPhoneNumber_Provider_Telkomsel_AS(t *testing.T) {
	pn := PhoneNumber("082318000000")
	result := pn.Provider()
	fmt.Println(result)
	if result != "TELKOMSEL AS" {
		t.Fatal("invalid value")
	}
}

func TestPhoneNumber_Provider_Telkomsel_Simpati(t *testing.T) {
	pn := PhoneNumber("082299552272")
	result := pn.Provider()
	fmt.Println(result)
	if result != "TELKOMSEL SIMPATI" {
		t.Fatal("invalid value")
	}
}
func TestPhoneNumber_Provider_Telkomsel_Halo(t *testing.T) {
	pn := PhoneNumber("08115590999")
	result := pn.Provider()
	fmt.Println(result)
	if result != "TELKOMSEL HALO" {
		t.Fatal("invalid value")
	}
}

func TestPhoneNumber_Provider_Indosat_OOREDOO(t *testing.T) {
	pn := PhoneNumber("085890005171")
	result := pn.Provider()
	fmt.Println(result)
	if result != "INDOSAT OOREDOO" {
		t.Fatal("invalid value")
	}
}

func TestPhoneNumber_Provider_XL(t *testing.T) {
	pn := PhoneNumber("087868884504")
	result := pn.Provider()
	fmt.Println(result)
	if result != "XL" {
		t.Fatal("invalid value")
	}
}

func TestPhoneNumber_Provider_Axis(t *testing.T) {
	pn := PhoneNumber("083862081192")
	result := pn.Provider()
	fmt.Println(result)
	if result != "AXIS" {
		t.Fatal("invalid value")
	}
}

func TestPhoneNumber_Provider_Three(t *testing.T) {
	pn := PhoneNumber("089520032657")
	result := pn.Provider()
	fmt.Println(result)
	if result != "THREE" {
		t.Fatal("invalid value")
	}
}

func TestPhoneNumber_Provider_Notfound(t *testing.T) {
	pn := PhoneNumber("079520032657")
	result := pn.Provider()
	fmt.Println(result)
	if result != "" || Provider(pn) != result {
		t.Fatal("invalid value")
	}
}
