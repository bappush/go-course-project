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
