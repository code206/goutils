package convert

func strings2Interfaces(strings []string) []interface{} {
	var interfaces = make([]interface{}, len(strings))

	for key, v := range strings {
		interfaces[key] = v
	}

	return interfaces
}
