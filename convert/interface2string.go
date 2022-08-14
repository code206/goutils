package convert

func Interface2String(i interface{}) string {
	if i != nil {
		return i.(string)
	} else {
		return ""
	}
}
