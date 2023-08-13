package lifts

type Lift struct {
	Id       int    `json:"id"`
	UserId   int    `json:"userId"`
	Name     string `json:"name"`
	LiftDate string `json:"liftDate"`
	Weight   int    `json:"weight"`
	Reps     int    `json:"reps"`
}
