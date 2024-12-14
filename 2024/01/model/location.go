package model

type Locations struct {
	GroupOneLocationIDs []int
	GroupTwoLocationIDs []int
}

func NewLocations() (locations *Locations) {
	return &Locations{
		GroupOneLocationIDs: []int{},
		GroupTwoLocationIDs: []int{},
	}
}

func (l *Locations) AddGroupOneLocationID(locationID int) {
	l.GroupOneLocationIDs = append(l.GroupOneLocationIDs, locationID)
}

func (l *Locations) AddGroupTwoLocationID(locationID int) {
	l.GroupTwoLocationIDs = append(l.GroupTwoLocationIDs, locationID)
}
