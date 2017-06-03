package tu


const symPrefix = "Go-"

func varName(pkgName string, varName string) string {
	totalLen := len(symPrefix) + len(pkgName) + len(".") + len(varName)
	buf := make([]byte, totalLen)

	pos := 0
	copy(buf[pos:], symPrefix)
	pos += len(symPrefix)
	copy(buf[pos:], pkgName)
	pos += len(pkgName)
	copy(buf[pos:], ".")
	pos += len(".")
	copy(buf[pos:], varName)

	return string(buf)
}