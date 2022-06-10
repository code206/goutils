package convert

func stringsToInterfaces(strings []string) []interface{} {
	var interfaces = make([]interface{}, len(strings))

	for key, v := range strings {
		interfaces[key] = v
	}

	return interfaces
}
