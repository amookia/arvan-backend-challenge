package response

type (
	Error struct {
		Message string   `json:"message"`
		Errors  []string `json:"errors"`
	}
)
