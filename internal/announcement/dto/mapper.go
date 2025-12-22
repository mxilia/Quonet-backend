package dto

import "github.com/mxilia/Quonet-backend/internal/entities"

func ToAnnouncementResponse(announcement *entities.Announcement) *AnnouncementResponse {
	return &AnnouncementResponse{
		ID:        announcement.ID,
		AuthorID:  announcement.AuthorID,
		Content:   announcement.Content,
		Author:    announcement.Author,
		CreatedAt: announcement.CreatedAt,
	}
}

func ToAnnouncementResponseList(announcements []*entities.Announcement) []*AnnouncementResponse {
	res := make([]*AnnouncementResponse, 0, len(announcements))
	for _, announcement := range announcements {
		res = append(res, ToAnnouncementResponse(announcement))
	}
	return res
}
