package types

type AttendanceCreate struct {
	Side  string `json:"side"`
	Name  string `json:"name"`
	Meal  string `json:"meal"`
	Count int    `json:"count"`
}

type Attendance struct {
	Id        int    `json:"id"`
	Side      string `json:"side"`
	Name      string `json:"name"`
	Meal      string `json:"meal"`
	Count     int    `json:"count"`
	Timestamp uint64 `json:"timestamp"`
}
