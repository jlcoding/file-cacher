package common

const SHA1_PATH = "Sha1Path"

func BuildKey(parts []string) string {
	var key string
	for _, part := range parts {
		key += part + ":"
	}
	return key[0 : len(key)-1]
}
