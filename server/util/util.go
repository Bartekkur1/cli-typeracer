package util

import "bytes"

func LooksLikeJSON(data []byte) bool {
	trimmed := bytes.TrimSpace(data)
	return len(trimmed) > 0 && (trimmed[0] == '{' && trimmed[len(trimmed)-1] == '}' || trimmed[0] == '[' && trimmed[len(trimmed)-1] == ']')
}
