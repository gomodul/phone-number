package phonenumber

import (
	"fmt"
	"reflect"
	"regexp"
)

// PhoneNumber ...
type PhoneNumber string

// Provider ...
func (pn PhoneNumber) Provider() string {
	return provider(fmt.Sprint(pn))
}

// Provider ...
func Provider(pn interface{}) string {
	return provider(fmt.Sprint(pn))
}

func provider(pn string) string {
	var result string
	for _, iv := range providers {
		if reflect.TypeOf(iv.Regex).Kind() == reflect.Map {
			for j, jv := range iv.Regex.(map[string]interface{}) {
				result = check(j, jv, pn)
				if result != "" {
					return fmt.Sprintf("%v %v", iv.Name, result)
				}
			}
		} else {
			result = check(iv.Name, iv.Regex, pn)
			if result != "" {
				return result
			}
		}
	}
	return ""
}

func check(k, v, pn interface{}) string {
	matched, err := regexp.MatchString(fmt.Sprint(v), fmt.Sprint(pn))
	if err != nil || !matched {
		return ""
	}
	return fmt.Sprint(k)
}
