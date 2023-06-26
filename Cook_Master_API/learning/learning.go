package learning

type Learning struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	Start_Time    string  `json:"start_time"`
	End_Time      string  `json:"end_time"`
	Instructor_ID int     `json:"instructor_id"`
}
