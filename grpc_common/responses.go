package grpc

import (
	"net/http"

	"github.com/Kiyosh31/ms-ecommerce-common/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ManageRpcErrors(err error, w http.ResponseWriter) {
	status := status.Convert(err)

	if status != nil {
		if status.Code() != codes.InvalidArgument {
			json.WriteError(w, http.StatusBadRequest, status.Message())
			return
		}
	}
}
