package runner

import (
	"fmt"

	"github.com/deyring/argos/models"
	"github.com/deyring/argos/runner/executor"
	"github.com/deyring/argos/utils"
)

type Runner struct {
	logger utils.Logger
	config *models.Config
}

func New(logger utils.Logger, configFilename string) (*Runner, error) {
	config := &models.Config{}
	if err := config.Load(configFilename); err != nil {
		return nil, err
	}

	return &Runner{
		logger: logger,
		config: config,
	}, nil
}

func (r *Runner) Run() error {
	r.logger.Debugf("running checks on %s", r.config.Name)
	for _, transaction := range r.config.Transactions {
		results, err := r.runTransaction(&transaction)
		if err != nil {
			return err
		}

		for _, result := range results {
			fmt.Printf("Result %s:\nSuccess: %v\nStatusCode:%d\nTLS Handshake: %s\nTTFB: %s\nTotal Duration:%s", transaction.Name, result.Success, result.StatusCode, result.TLSHandshakeDuration, result.FirstByteDuration, result.TotalDuration)
		}
	}
	return nil
}

func (r *Runner) runTransaction(transaction *models.Transaction) ([]*models.EndpointCheckResult, error) {
	r.logger.Debugf("running checks on transaction %s", transaction.Name)
	checkResults := []*models.EndpointCheckResult{}
	for _, check := range transaction.Checks {
		result, err := r.executeCheck(&check)
		if err != nil {
			return nil, err
		}
		checkResults = append(checkResults, result)
	}
	return checkResults, nil
}

func (r *Runner) executeCheck(check *models.EndpointCheck) (*models.EndpointCheckResult, error) {
	r.logger.Debugf("executing check %s", check.Name)
	executorInstance, err := executor.New(check)
	if err != nil {
		return nil, err
	}

	result, err := executorInstance.Run()
	if err != nil {
		return nil, err
	}

	success, err := check.AssertResult(result)
	if err != nil {
		return nil, err
	}

	result.Success = success

	return result, nil
}
