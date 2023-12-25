package utils

// IDString
// 字符串前面补0  长度为length
func IDString(id string, length int) string {
	if len(id) >= length {
		return id
	}
	id = "000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000" + id
	return id[len(id)-length:]
}
