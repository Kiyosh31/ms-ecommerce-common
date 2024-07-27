package utils

import (
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func WriteRpcError(err error, w http.ResponseWriter) {
	status := status.Convert(err)

	if status != nil {
		if status.Code() != codes.InvalidArgument {
			WriteError(w, http.StatusBadRequest, status.Message())
			return
		}
	}
}
