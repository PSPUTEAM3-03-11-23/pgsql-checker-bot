package entity

type Alert struct {
	Id       int
	Incident *Incident
	IsSent   *bool
	//User
}
