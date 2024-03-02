package model

type AppRequest struct {
	Parent string `json:"parent_id" binding:"required"`
	Child string `json:"child_id" binding:"required"`
	Semester string `json:"desired_semester_id" binding:"required"`
}