package exceptions

import "errors"

var NotFoundException = errors.New("error: Not Found")

var BadRequestException = errors.New("error: Bad Request")

var InternalServerError = errors.New("error: Internal Server Error")