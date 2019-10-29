package entity

type Meeting struct {
	Sponsor      string
	Participator []string
	Start, End   Date
	Title        string
}

func GetSponsor(a Meeting) string {
	return a.Sponsor
}
func GetParticipator(a Meeting) []string {
	return a.Participator
}

func GetStart(a Meeting) Date {
	return a.Start
}
func GetEnd(a Meeting) Date {
	return a.End
}
func GetTitle(a Meeting) string {
	return a.Title
}
