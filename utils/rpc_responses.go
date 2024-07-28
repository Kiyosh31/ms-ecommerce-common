package utils

import (
	"fmt"
	"net/http"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
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

func FieldValidation(field string, err error) *errdetails.BadRequest_FieldViolation {
	return &errdetails.BadRequest_FieldViolation{
		Field:       field,
		Description: err.Error(),
	}
}

func InvalidArgumentError(violations []*errdetails.BadRequest_FieldViolation) error {
	badRequest := &errdetails.BadRequest{FieldViolations: violations}
	statusInvalid := status.New(codes.InvalidArgument, "Invalid parameters")

	statusDetails, err := statusInvalid.WithDetails(badRequest)
	if err != nil {
		return statusInvalid.Err()
	}

	return statusDetails.Err()
}

func ParseInterfaceToString(word interface{}) (string, error) {
	s, ok := word.(string)
	if !ok {
		return "nil", fmt.Errorf("error while parsing interface to string")
	}

	return s, nil
}
