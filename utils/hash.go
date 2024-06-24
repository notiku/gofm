package utils

import (
	"crypto/md5"
	"encoding/hex"
	"sort"
	"strings"
)

// Hashes a request object using the MD5 algorithm
func HashRequest(obj map[string]string, secretKey string) string {
	// Extract and sort the keys
	keys := make([]string, 0, len(obj))
	for k := range obj {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Concatenate the keys and values
	var sb strings.Builder
	for _, k := range keys {
		sb.WriteString(k)
		sb.WriteString(obj[k])
	}
	sb.WriteString(secretKey)

	// Hash the concatenated string
	hasher := md5.New()
	hasher.Write([]byte(sb.String()))
	hashed := hex.EncodeToString(hasher.Sum(nil))

	return hashed
}
