package counters

const (
	DefaultKey = ""
	LikesKey   = "likes"
	ViewsKey   = "views"
)

var allowedKeys = map[string]bool{
	DefaultKey: true,
	LikesKey:   true,
	ViewsKey:   true,
}

func IsKeyAllowed(key string) bool {
	_, ok := allowedKeys[key]
	return ok
}

func GetAllowedKeys() []string {
	keys := make([]string, 0, len(allowedKeys))
	for k := range allowedKeys {
		keys = append(keys, k)
	}

	return keys
}
