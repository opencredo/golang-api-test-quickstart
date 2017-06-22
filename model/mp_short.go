package model

type MpShort struct {
	MemberId     string `json:"member_id"`
	PersonId     string `json:"person_id"`
	Name         string `json:"name"`
	Party        string `json:"party"`
	Constituency string `json:"constituency"`
	Office       []Office `json:"office"`
}