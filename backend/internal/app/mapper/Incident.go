package mapper

import (
	"hackathon-tg-bot/internal/app/model/entity"
	"hackathon-tg-bot/internal/app/model/response"
)

func IncidentToIncidentResponse(incident *entity.Incident) *response.Incident {
	r := &response.Incident{
		Id:   incident.Id,
		Db:   DatabaseToDatabaseResponse(incident.Db),
		Date: incident.Date,
	}

	return r
}

func IncidentToIncidentResponses(incidents *[]entity.Incident) *[]response.Incident {
	rs := make([]response.Incident, 0)

	for _, incident := range *incidents {
		rs = append(rs, *IncidentToIncidentResponse(&incident))
	}

	return &rs
}
