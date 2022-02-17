package models

type Query struct {
	Results []Location `json:"results"`
}

func (q *Query) GetMostConfidenceLocation() Location {
	confidence, i := 0, 0
	for idx, location := range q.Results {
		if location.Confidence > confidence {
			confidence = location.Confidence
			i = idx
		}
	}
	return q.Results[i]
}
