package domain

import (
	"time"
)

type Task struct {
	ID          string 			`json:"id" bson:"_id,omitempty"`
	Title       string 			`json:"title" bson:"title"`
	Description string 			`json:"description" bson:"description"`
	Stage	    string 			`json:"stage" bson:"stage"`
	CreatedAt   time.Time 		`json:"created_at" bson:"createdAt"`
	Deadline    *time.Time 		`json:"deadline" bson:"deadline"`
	CreatedBy   string 			`json:"created_by" bson:"createdBy"`
	AssignedTo  string 			`json:"assigned_to" bson:"assignedTo"`
	Tags		[]string 		`json:"tags" bson:"tags"`
}