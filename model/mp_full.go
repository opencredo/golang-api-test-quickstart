package model

type Mp struct {
	MemberId      string `json:"member_id"`
	House         string `json:"house"`
	Constituency  string `json:"constituency"`
	Party         string `json:"party"`
	EnteredHouse  string `json:"entered_house"`
	LeftHouse     string `json:"left_house"`
	EnteredReason string `json:"entered_reason"`
	LeftReason    string `json:"left_reason"`
	PersonId      string `json:"person_id"`
	LastUpdate    string `json:"lastupdate"`
	Title         string `json:"title"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	FullName      string `json:"full_name"`
	Url           string `json:"url"`
	Image         string `json:"image"`
	ImageHeight   int `json:"image_height"`
	ImageWidth    int `json:"image_width"`
}