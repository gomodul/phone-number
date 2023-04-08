package phonenumber

import (
	"fmt"
	"strings"
	"testing"
)

func TestPhoneNumber_Provider_Telkomsel(t *testing.T) {
	pns := []PhoneNumber{
		"085100000000",
		"085200000000",
		"0853000000000",
		"082310000000",

		"08110000000",

		"081200000000",
		"081300000000",
		"082100000000",
		"082200000000",
	}
	for _, pn := range pns {
		if !strings.Contains(pn.Provider(), "TELKOMSEL") {
			t.Fatal("invalid value")
		}
		fmt.Println(pn.Provider())
	}
}

func TestPhoneNumber_Provider_Telkomsel_AS(t *testing.T) {
	pns := []PhoneNumber{
		"085100000000",
		"085200000000",
		"0853000000000",

		"082310000000",
	}
	for _, pn := range pns {
		if !strings.Contains(pn.Provider(), "AS") {
			t.Fatal("invalid value")
		}
		fmt.Println(pn.Provider())
	}
}

func TestPhoneNumber_Provider_Telkomsel_Simpati(t *testing.T) {
	pns := []PhoneNumber{
		"081200000000",
		"081300000000",
		"082100000000",
		"082200000000",
	}
	for _, pn := range pns {
		if !strings.Contains(pn.Provider(), "SIMPATI") {
			t.Fatal("invalid value")
		}
		fmt.Println(pn.Provider())
	}
}

func TestPhoneNumber_Provider_Telkomsel_Halo(t *testing.T) {
	pns := []PhoneNumber{
		"08110000000",
	}
	for _, pn := range pns {
		if !strings.Contains(pn.Provider(), "HALO") {
			t.Fatal("invalid value")
		}
		fmt.Println(pn.Provider())
	}
}

func TestPhoneNumber_Provider_Indosat_OOREDOO(t *testing.T) {
	pns := []PhoneNumber{
		"085700000000",
		"085800000000",
	}
	for _, pn := range pns {
		if pn.Provider() != "INDOSAT OOREDOO" {
			t.Fatal("invalid value")
		}
		fmt.Println(pn.Provider())
	}
}

func TestPhoneNumber_Provider_XL(t *testing.T) {
	pns := []PhoneNumber{
		"087868884504",
	}
	for _, pn := range pns {
		if pn.Provider() != "XL" {
			t.Fatal("invalid value", pn.Provider())
		}
		fmt.Println(pn.Provider())
	}
}

func TestPhoneNumber_Provider_Axis(t *testing.T) {
	pns := []PhoneNumber{
		"083862081192",
	}
	for _, pn := range pns {
		if pn.Provider() != "AXIS" {
			t.Fatal("invalid value", pn.Provider())
		}
		fmt.Println(pn.Provider())
	}
}

func TestPhoneNumber_Provider_Three(t *testing.T) {
	pns := []PhoneNumber{
		"089520032657",
	}
	for _, pn := range pns {
		if pn.Provider() != "THREE" {
			t.Fatal("invalid value", pn.Provider())
		}
		fmt.Println(pn.Provider())
	}
}

func TestPhoneNumber_Provider_SMARTFREN(t *testing.T) {
	pns := []PhoneNumber{
		"088211122152",
	}
	for _, pn := range pns {
		if pn.Provider() != "SMARTFREN" {
			t.Fatal("invalid value", pn.Provider())
		}
		fmt.Println(pn.Provider())
	}
}

func TestPhoneNumber_Provider_TELKOM(t *testing.T) {
	pns := []PhoneNumber{
		"0215386950",
	}
	for _, pn := range pns {
		if pn.Provider() != "TELKOM" {
			t.Fatal("invalid value", pn.Provider())
		}
		fmt.Println(pn.Provider())
	}
}

func TestPhoneNumber_Provider_Notfound(t *testing.T) {
	pns := []PhoneNumber{
		"079520032657",
	}
	for _, pn := range pns {
		result := pn.Provider()
		if result != "" || Provider(pn) != result {
			t.Fatal("invalid value")
		}
		fmt.Println(pn.Provider())
	}
}
