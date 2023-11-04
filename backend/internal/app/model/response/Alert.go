package response

type Alert struct {
	Id       int
	Incident *Incident
	IsSent   *bool
	//User
}
