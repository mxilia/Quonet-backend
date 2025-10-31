package dto

import "github.com/mxilia/Conflux-backend/internal/entities"

func ToThreadResponse(thread *entities.Thread) *ThreadResponse {
	return &ThreadResponse{
		ID:        thread.ID,
		Title:     thread.Title,
		Posts:     thread.Posts,
		CreatedAt: thread.CreatedAt,
	}
}

func ToThreadResponseList(threads []*entities.Thread) []*ThreadResponse {
	res := make([]*ThreadResponse, 0, len(threads))
	for _, t := range threads {
		res = append(res, ToThreadResponse(t))
	}
	return res
}
