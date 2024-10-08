// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package common

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsUnknown(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_Unknown.String() && e.Code == 500
}

func ErrorUnknown(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_Unknown.String(), fmt.Sprintf(format, args...))
}

func IsBadRequest(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_BAD_REQUEST.String() && e.Code == 400
}

func ErrorBadRequest(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_BAD_REQUEST.String(), fmt.Sprintf(format, args...))
}

func IsNotUnauthorized(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_NOT_Unauthorized.String() && e.Code == 401
}

func ErrorNotUnauthorized(format string, args ...interface{}) *errors.Error {
	return errors.New(401, ErrorReason_NOT_Unauthorized.String(), fmt.Sprintf(format, args...))
}

func IsNotForbidden(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_NOT_Forbidden.String() && e.Code == 403
}

func ErrorNotForbidden(format string, args ...interface{}) *errors.Error {
	return errors.New(403, ErrorReason_NOT_Forbidden.String(), fmt.Sprintf(format, args...))
}

func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_NOT_FOUND.String() && e.Code == 404
}

func ErrorNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, ErrorReason_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsConflict(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_Conflict.String() && e.Code == 409
}

func ErrorConflict(format string, args ...interface{}) *errors.Error {
	return errors.New(409, ErrorReason_Conflict.String(), fmt.Sprintf(format, args...))
}
