package venue

type NewVenueInput struct {
	Name     string `json:"name"`
	Capacity int32  `json:"capacity"`
	City     string `json:"city"`
	State    string `json:"state"`
}
