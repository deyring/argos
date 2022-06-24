package runner

import (
	"errors"
	"io"
	"time"

	"github.com/deyring/argos/models"
	"github.com/deyring/argos/resultsink"
	"github.com/deyring/argos/resultsink/sinkfactory"
	"github.com/deyring/argos/runner/executor"
	"github.com/deyring/argos/utils"
)

type Runner struct {
	logger      utils.Logger
	config      *models.Config
	resultSinks []resultsink.Sink
}

func New(logger utils.Logger, configFileReader io.Reader) (*Runner, error) {
	config := &models.Config{}
	if err := config.Load(configFileReader); err != nil {
		return nil, err
	}

	resultSinks, err := initResultSinks(config)
	if err != nil {
		return nil, err
	}

	return &Runner{
		logger:      logger,
		config:      config,
		resultSinks: resultSinks,
	}, nil
}

func initResultSinks(config *models.Config) ([]resultsink.Sink, error) {
	resultSinks := []resultsink.Sink{}

	for _, sinkConfig := range config.Outputs {
		sink, err := initResultSink(sinkConfig)
		if err != nil {
			return nil, err
		}
		resultSinks = append(resultSinks, sink)
	}

	return resultSinks, nil
}

func initResultSink(output models.Output) (resultsink.Sink, error) {
	switch output.Type {
	case models.OutputTypeStdOut:
		return sinkfactory.GetNewStdoutSink(), nil
	case models.OutputTypeInfluxDB:
		return sinkfactory.GetNewInfluxDBSink(output.Host, output.User, output.Password, output.Database, output.Insecure), nil
	default:
		return nil, errors.New("unknown output type")
	}
}

func (r *Runner) Run() error {
	r.logger.Infof("running checks on %s. Execution strategy: %s", r.config.Name, r.config.Execute)

	if r.config.Execute == models.ExecuteTypeLoop {
		for {
			result, err := r.runCheck()
			if err != nil {
				return err
			}

			for _, sink := range r.resultSinks {
				if err := sink.Handle(result); err != nil {
					return err
				}
			}

			time.Sleep(time.Duration(r.config.Sleep) * time.Second)
		}
	} else {
		result, err := r.runCheck()
		if err != nil {
			return err
		}

		for _, sink := range r.resultSinks {
			if err := sink.Handle(result); err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *Runner) runCheck() (*models.Result, error) {
	result := &models.Result{
		Name:               r.config.Name,
		TransactionResults: []models.TransactionResult{},
	}
	for _, transaction := range r.config.Transactions {
		transactionResult, err := r.runTransaction(&transaction)
		if err != nil {
			return nil, err
		}

		result.TransactionResults = append(result.TransactionResults, *transactionResult)
	}
	return result, nil
}

func (r *Runner) runTransaction(transaction *models.Transaction) (*models.TransactionResult, error) {
	r.logger.Debugf("running checks on transaction %s", transaction.Name)
	transactionResult := &models.TransactionResult{
		Name:                 transaction.Name,
		Success:              true,
		EndpointCheckResults: []models.EndpointCheckResult{},
	}

	for _, check := range transaction.Checks {
		result, err := r.executeCheck(&check)
		if err != nil {
			return nil, err
		}

		if !result.Success {
			transactionResult.Success = false
		}

		transactionResult.EndpointCheckResults = append(transactionResult.EndpointCheckResults, *result)
	}
	return transactionResult, nil
}

func (r *Runner) executeCheck(check *models.EndpointCheck) (*models.EndpointCheckResult, error) {
	r.logger.Debugf("executing check %s", check.Name)
	executorInstance, err := executor.New(check)
	if err != nil {
		return nil, err
	}

	result, err := executorInstance.Run()
	if err != nil {
		r.logger.Warningf("check %s failed: %s", check.Name, err.Error())
		return &models.EndpointCheckResult{
			Name:    check.Name,
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	success, err := check.AssertResult(result)
	if err != nil {
		return nil, err
	}

	result.Success = success

	return result, nil
}
