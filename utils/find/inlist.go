package find

func InList[T comparable](list []T, key T) bool {
	for _, v := range list {
		if v == key {
			return true
		}
	}
	return false
}
