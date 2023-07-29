package response

type (
	PutObjectError struct {
		Err string `json:"error"`
	}
	PutObject struct {
		Message  string `json:"message"`
		ObjectId string `json:"objectId"`
	}
)
