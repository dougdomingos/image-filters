package api

const (
	// HTTP_Ok signals that the requested operation was successful. 
	HTTP_Ok = 200

	// HTTP_Bad_Request signals that the client request is malformed (e.g., a
	// missing or invalid parameter).
	HTTP_Bad_Request = 400

	// HTTP_Not_Found signals that a requested filter pipeline does not exist.
	HTTP_Not_Found = 404
)