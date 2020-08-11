package phone_number

import (
	"fmt"
	"reflect"
	"regexp"
)

type PhoneNumber string

func (pn PhoneNumber) Provider() string {
	return provider(fmt.Sprint(pn))
}

func Provider(pn interface{}) string {
	return provider(fmt.Sprint(pn))
}

func provider(pn string) string {
	var result string

	for i, iv := range providers {
		if reflect.TypeOf(iv).Kind() == reflect.Map {
			for j, jv := range iv.(map[string]interface{}) {
				result = check(j, jv, pn)
				if result != "" {
					return fmt.Sprintf("%v %v", i, result)
				}
			}
		} else {
			result = check(i, iv, pn)
			if result != "" {
				return result
			}
		}
	}
	return ""
}

func check(k, v, pn interface{}) string {
	matched, _ := regexp.MatchString(fmt.Sprint(v), fmt.Sprint(pn))
	if matched {
		return fmt.Sprint(k)
	}
	return ""
}
