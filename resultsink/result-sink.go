package resultsink

import "github.com/deyring/argos/models"

type Sink interface {
	Handle(result *models.Result) error
}
