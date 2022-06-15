package stdout

import (
	"os"

	"github.com/deyring/argos/models"
	resultsink "github.com/deyring/argos/resultsink"
	"github.com/jedib0t/go-pretty/v6/table"
)

type stdoutSink struct {
}

func (s *stdoutSink) Handle(result *models.Result) error {

	printTable(result)

	/* printEndpointCheckResult := func(result models.EndpointCheckResult) string {
		return fmt.Sprintf("\t%s: %s\n\tStats: StatusCode: %d, Connect: %s, TLS Handshake: %s, First Byte: %s, Total Duration: %s\n", result.Name,
			boolToEmoji(result.Success),
			result.StatusCode,
			result.ConnectDuration,
			result.TLSHandshakeDuration,
			result.FirstByteDuration,
			result.TotalDuration,
		)
	}

	printTransactionResult := func(result *models.TransactionResult) string {
		var endpointPrint string
		for _, endpoint := range result.EndpointCheckResults {
			endpointPrint += printEndpointCheckResult(endpoint)
		}
		return fmt.Sprintf("%s: %s\n%s", result.Name, boolToEmoji(result.Success), endpointPrint)
	}

	fmt.Println("\n************************************************")
	fmt.Printf("Results for %s\n\n", result.Name)

	for _, transactionResult := range result.TransactionResults {
		fmt.Println(printTransactionResult(&transactionResult))
	}

	fmt.Println("\n************************************************") */

	return nil
}

func printTable(result *models.Result) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Transaction", "Endpoint Name", "Success", "StatusCode", "Connect", "TLS Handshake", "First Byte", "Total"})
	for _, transactionResult := range result.TransactionResults {
		for _, endpointResult := range transactionResult.EndpointCheckResults {
			t.AppendRow(table.Row{transactionResult.Name, endpointResult.Name, boolToEmoji(endpointResult.Success), endpointResult.StatusCode, endpointResult.ConnectDuration, endpointResult.TLSHandshakeDuration, endpointResult.FirstByteDuration, endpointResult.TotalDuration})
		}
	}
	t.Render()
}

func boolToEmoji(success bool) string {
	if success {
		return "✅"
	}
	return "❌"
}

func New() resultsink.Sink {
	return &stdoutSink{}
}
