package enrollment

import (
	"context"

	"github.com/SanGameDev/go_lib_response/response"
)

type (
	Controller func(ctx context.Context, request interface{}) (interface{}, error)

	Endpoints struct {
		Create Controller
	}

	CreateReq struct {
		UserID   string `json:"user_id"`
		CourseID string `json:"course_id"`
	}
)

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(s),
	}
}

func makeCreateEndpoint(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(CreateReq)

		if req.UserID == "" {
			return nil, response.BadRequest(ErrCourseIDRequired.Error())
		}

		if req.CourseID == "" {
			return nil, response.BadRequest(ErrUserIDRequired.Error())
		}

		enroll, err := s.Create(ctx, req.UserID, req.CourseID)
		if err != nil {
			return nil, response.InternalServerError(err.Error())
		}

		return response.Created("success", enroll, nil), nil
	}
}
