package lifts

type Lift struct {
	Id     int    `json:"id"`
	UserId int    `json:"userId"`
	Name   string `json:"name"`
	Date   string `json:"date"`
	Weight int    `json:"weight"`
	Reps   int    `json:"reps"`
}
