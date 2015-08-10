package service

import (
	. "github.com/talbright/go-desk/types"
	. "github.com/talbright/go-desk/resource"
	"net/http"
)

type InsightsV3Service struct {
	client *Client
}

func (s *InsightsV3Service) Report(query *InsightsV3Query) (*InsightsV3Report, *http.Response, error) {
	restful := Restful{}
	report := &InsightsV3Report{}
	resp, err := restful.
		Post("insights3/reports").
		Body(query).
		Json(report).
		Client(s.client).
		Do()

	return report, resp, err
}
