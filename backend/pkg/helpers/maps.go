package helpers

func MapKeyContains(m map[string]interface{}, values []string) bool {
	for k := range m {
		for _, v := range values {
			if k == v {
				return true
			}
		}
	}
	return false
}
