package dto

import "github.com/mxilia/Conflux-backend/internal/entities"

func ToThreadResponse(thread *entities.Thread) *ThreadResponse {
	return &ThreadResponse{
		Title: thread.Title,
	}
}

func ToThreadResponseList(threads []*entities.Thread) []*ThreadResponse {
	res := make([]*ThreadResponse, 0, len(threads))
	for _, t := range threads {
		res = append(res, ToThreadResponse(t))
	}
	return res
}
