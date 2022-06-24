package influxdb

import (
	"fmt"
	"log"
	"time"

	"github.com/deyring/argos/models"
	"github.com/deyring/argos/resultsink"
	_ "github.com/influxdata/influxdb1-client"
	client "github.com/influxdata/influxdb1-client/v2"
	influxdb "github.com/influxdata/influxdb1-client/v2"
)

type influxSink struct {
	client   influxdb.Client
	database string
}

func New(host, user, password, database string) resultsink.Sink {
	client, err := influxdb.NewHTTPClient(influxdb.HTTPConfig{
		Addr:     host,
		Username: user,
		Password: password,
	})
	if err != nil {
		log.Fatalf("initiatializing influxdb client failed: %v", err)
	}

	return &influxSink{
		client:   client,
		database: database,
	}
}

func (s *influxSink) Handle(result *models.Result) error {
	bp, err := influxdb.NewBatchPoints(client.BatchPointsConfig{
		Database:  s.database,
		Precision: "ns",
	})

	now := time.Now()

	for _, transaction := range result.TransactionResults {

		tags := map[string]string{"config": fmt.Sprintf("%s", result.Name), "transaction": transaction.Name}
		fields := map[string]interface{}{
			"success": boolToInt(transaction.Success),
		}

		pt, err := client.NewPoint("argos.transaction", tags, fields, now)
		if err != nil {
			log.Fatal(err)
		}
		bp.AddPoint(pt)

		for _, endpoint := range transaction.EndpointCheckResults {
			tags := map[string]string{"config": fmt.Sprintf("%s", result.Name), "transaction": transaction.Name, "endpoint": endpoint.Name, "url": endpoint.URL}
			fields := map[string]interface{}{
				"success":                boolToInt(endpoint.Success),
				"connect_duration":       int64(endpoint.ConnectDuration),
				"tls_handshake_duration": int64(endpoint.TLSHandshakeDuration),
				"first_byte_duration":    int64(endpoint.FirstByteDuration),
				"total_duration":         int64(endpoint.TotalDuration),
			}

			pt, err := client.NewPoint("argos.transaction.endpoint", tags, fields, now)
			if err != nil {
				log.Fatal(err)
			}
			bp.AddPoint(pt)
		}
	}

	// Write data
	err = s.client.Write(bp)
	if err != nil {
		log.Fatalf("failed to write to influxdb: %v", err)
		return err
	}

	return nil
}

func boolToInt(in bool) int {
	if in {
		return 1
	}
	return 0
}
