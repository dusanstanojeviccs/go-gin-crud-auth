package lifts

import "sync"

type liftRepository struct {
	mu    sync.Mutex
	lifts []*Lift
}

func (this *liftRepository) findAll(userId int) []*Lift {
	this.mu.Lock()
	defer this.mu.Unlock()

	filteredLifts := []*Lift{}

	for _, lift := range this.lifts {
		if lift.UserId == userId {
			filteredLifts = append(filteredLifts, lift)
		}
	}

	return filteredLifts
}

func (this *liftRepository) findById(id int, userId int) *Lift {
	this.mu.Lock()
	defer this.mu.Unlock()

	for _, lift := range this.lifts {
		if lift.Id == id && lift.UserId == userId {
			return lift
		}
	}
	return nil
}

func (this *liftRepository) delete(id int, userId int) {
	this.mu.Lock()
	defer this.mu.Unlock()

	newLifts := []*Lift{}

	for _, lift := range this.lifts {
		if lift.Id != id || lift.UserId != userId {
			newLifts = append(newLifts, lift)
		}
	}
	this.lifts = newLifts
}

func (this *liftRepository) create(lift *Lift, userId int) {
	this.mu.Lock()
	defer this.mu.Unlock()

	lift.Id = len(this.lifts) + 1
	lift.UserId = userId
	this.lifts = append(this.lifts, lift)
}

func (this *liftRepository) update(lift *Lift, userId int) {
	this.mu.Lock()
	defer this.mu.Unlock()

	lift.UserId = userId

	for i, existingLift := range this.lifts {
		if existingLift.Id == lift.Id && existingLift.UserId == userId {
			this.lifts[i] = lift
		}
	}
}

var LiftRepository = liftRepository{lifts: []*Lift{}}
