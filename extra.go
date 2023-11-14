package ferror

type Extra map[string]any

// isKeyInvalid a function to check key is not equal message, error, and status as they override the main key-value pairs.
func isKeyInvalid(key string) bool {
	return key == "message" || key == "status" || key == "error"
}
