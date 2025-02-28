package models

type JwtFetch struct {
	Jwt string `json:"jwt" binding:"required"`
}

const (
	Query_MasterDept = `SELECT id, name FROM public."Master_Dept"`
)

//("topic_code" text, "subject" text, "dept" text, "task_code" text, "task_name" text, "task_category" text, "generate_every" text, "priority" text, "estimasted_time_done" text, "assign_to" text, "created_date" text)
