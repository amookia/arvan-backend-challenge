package response

type (
	PutObjectError struct {
		Err error `json:"error"`
	}
)
