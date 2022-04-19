package api

func validataInput(iban string) bool {
	if iban == "" {
		return false
	}
	return true
}
