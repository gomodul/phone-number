package phonenumber

var providers = []struct {
	Name  string
	Regex interface{}
}{
	{
		Name: "TELKOMSEL",
		Regex: map[string]interface{}{
			"AS":      "^(\\+62|\\+0|0|62)8(5[123]|23)[0-9]{7,9}$",
			"HALO":    "^(\\+62|\\+0|0|62)8(11)[0-9]{6,7}$",
			"SIMPATI": "^(\\+62|\\+0|0|62)8(1[23]|2[12])[0-9]{8,9}$",
		},
	},
	{
		Name: "INDOSAT",
		Regex: map[string]interface{}{
			"OOREDOO": "^(\\+62|\\+0|0|62)8(1[46]|5[5678])[0-9]{5,9}$",
		},
	},
	{
		Name:  "XL",
		Regex: "^(\\+62|\\+0|0|62)8(1[789]|59|7[46789])[0-9]{5,9}$",
	},
	{
		Name:  "AXIS",
		Regex: "^(\\+62|\\+0|0|62)8(3[1238])[0-9]{5,9}$",
	},
	{
		Name:  "THREE",
		Regex: "^(\\+62|\\+0|0|62)8(9[56789])[0-9]{7,9}$",
	},
	{
		Name:  "SMARTFREN",
		Regex: "^(\\+62|\\+0|0|62)88[1-9](\\d{6,9})",
	},
	{
		Name:  "TELKOM",
		Regex: "^(\\d{3,4})(\\d{3})(\\d{3,4})$",
	},
	{
		Name:  "BANK",
		Regex: "^(\\d{4})(\\d{4})(\\d{4})(\\d{4})$",
	},
}
