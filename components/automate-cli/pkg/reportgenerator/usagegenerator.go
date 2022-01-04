package usagegenerator

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gocarina/gocsv"
	elastic "gopkg.in/olivere/elastic.v6"
)

type ConvergeInfo struct {
	NodeID             string `json:"entity_uuid"`
	Environment        string `json:"environment"`
	FQDN               string `json:"fqdn"`
	HostName           string `json:"hostname"`
	IPAddress          string `json:"ipaddress"`
	MacAddress         string `json:"macaddress"`
	Platform           string `json:"platform"`
	PlatformFamily     string `json:"platform_family"`
	PlatformVersion    string `json:"platform_version"`
	UptimeSeconds      int64  `json:"uptime_seconds"`
	Status             string `json:"client_run_status"`
	StartTime          string `json:"start_time"`
	EndTime            string `json:"end_time"`
	SerialNumber       string `json:"dmi_system_serial_number"`
	TotalResourceCount int    `json:"total_resource_count"`
}

type ComplianceInfo []interface{}

func GenerateNodeCount(esHostName string, esPort string, startTime time.Time, endTime time.Time) {
	elasticSearchURL := fmt.Sprintf("http://%s:%s", esHostName, esPort)
	client, err := elastic.NewClient(
		elastic.SetURL(elasticSearchURL),
		elastic.SetSniff(false),
	)
	if err != nil {
		fmt.Println("Elastic error : ", err)
		os.Exit(1)
	}

	queryElasticSearchNodeCount(client, startTime, endTime)
}

func GenerateNodeRunReport(esHostName string, esPort string, startTime time.Time, endTime time.Time) {
	elasticSearchURL := fmt.Sprintf("http://%s:%s", esHostName, esPort)
	client, err := elastic.NewClient(
		elastic.SetURL(elasticSearchURL),
		elastic.SetSniff(false),
	)
	if err != nil {
		fmt.Println("Elastic error : ", err)
		os.Exit(1)
	}

	queryElasticSearchNodeReport(client, startTime, endTime)
}

func queryElasticSearchNodeCount(client *elastic.Client, startTime time.Time, endTime time.Time) {
	filename := fmt.Sprintf("nodecount_%s_%s.csv", startTime.Format("2006-01-02"), endTime.Format("2006-01-02"))
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Println("Failed to open Node Count report")
		os.Exit(1)
	}
	defer f.Close()
	writer := gocsv.DefaultCSVWriter(f)
	//writer := struct2csv.NewWriter(f)
	err = writer.Write([]string{"Start Time", "End Time", "Unique Node Count"})
	if err != nil {
		fmt.Println("Error while writing csv: ", err)
		os.Exit(1)
	}

	st := endTime.Add(24 * time.Hour)
	for {
		et := st
		st = et.Add(24 * -time.Hour)

		err = writer.Write([]string{"Day Count"})
		if err != nil {
			fmt.Println("Error while writing csv: ", err)
			os.Exit(1)
		}
		dayCount, ok := getUniqueCounts(client, st, et)
		if ok {
			err = writer.Write([]string{st.Format(time.RFC3339), et.Format(time.RFC3339), fmt.Sprintf("%f", *dayCount.Value)})
			if err != nil {
				fmt.Println("Error while writing csv: ", err)
				os.Exit(1)
			}
		}
		t := et.Add(-time.Hour)
		if t.Before(st) {
			t = st
		}
		err = writer.Write([]string{"Hourly Count"})
		if err != nil {
			fmt.Println("Error while writing csv: ", err)
			os.Exit(1)
		}
		writer.Flush()
		for {
			metric, ok := getUniqueCounts(client, t, et)
			if ok && *metric.Value > 0 {
				err = writer.Write([]string{t.Format(time.RFC3339), et.Format(time.RFC3339), fmt.Sprintf("%f", *metric.Value)})
				if err != nil {
					fmt.Println("Error while writing csv: ", err)
					os.Exit(1)
				}
				writer.Flush()
			}

			if t == st {
				break
			}

			et = t
			t = et.Add(-time.Hour)
			if t.Before(st) {
				t = st
			}
		}
		err = writer.Write([]string{"", "", ""})
		if err != nil {
			fmt.Println("Error while writing csv: ", err)
			os.Exit(1)
		}
		writer.Flush()
		if st == startTime {
			break
		}
	}
	writer.Flush()
	fmt.Println("The node count report is available in: ", filename)
}

func getUniqueCounts(client *elastic.Client, startTime time.Time, endTime time.Time) (*elastic.AggregationValueMetric, bool) {
	rangeQuery := elastic.NewRangeQuery("end_time").
		Format("yyyy-MM-dd||yyyy-MM-dd-HH:mm:ss||yyyy-MM-dd'T'HH:mm:ssZ")
	rangeQuery.Gte(startTime)
	rangeQuery.Lte(endTime)

	aggr := elastic.NewCardinalityAggregation().Field("entity_uuid")
	searchService := client.Search().
		Query(rangeQuery).
		Index("converge-history-*").
		Size(100).
		Aggregation("nodes_count", aggr)

	searchResult, err := searchService.Do(context.Background())
	if err != nil {
		fmt.Println("Error in query: ", err)
		os.Exit(1)
	}

	metric, ok := searchResult.Aggregations.ValueCount("nodes_count")
	return metric, ok
}

func (s ComplianceInfo) mapAt(index int) map[string]interface{} {
	if len(s) <= index {
		return nil
	}

	return s[index].(map[string]interface{})
}

func queryElasticSearchNodeReport(client *elastic.Client, startTime time.Time, endTime time.Time) {
	filename := fmt.Sprintf("nodeinfo_%s_%s.csv", startTime.Format("2006-01-02"), endTime.Format("2006-01-02"))
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Println("Failed to open Node Detail report")
		os.Exit(1)
	}
	defer f.Close()

	//	writer := struct2csv.NewWriter(f)
	writer := gocsv.DefaultCSVWriter(f)

	var sourceFields = []string{
		"entity_uuid",
		"environment",
		"fqdn",
		"hostname",
		"ipaddress",
		"macaddress",
		"platform",
		"platform_family",
		"platform_version",
		"uptime_seconds",
		"client_run_status",
		"start_time",
		"end_time",
		"dmi_system_serial_number",
		"total_resource_count",
	}

	t := endTime.Add(24 * time.Hour)
	//t := endTime.Add(-time.Hour * 24)
	if t.Before(startTime) {
		t = startTime
	}

	for {
		rangeQuery := elastic.NewRangeQuery("end_time").
			Format("yyyy-MM-dd||yyyy-MM-dd-HH:mm:ss||yyyy-MM-dd'T'HH:mm:ssZ")
		rangeQuery.Gte(t)
		rangeQuery.Lte(endTime)

		fetchSource := elastic.NewFetchSourceContext(true).Include(sourceFields...)
		searchService := client.Search().
			Query(rangeQuery).
			Index("converge-history-*").
			Size(10000).
			FetchSourceContext(fetchSource)

		searchResult, err := searchService.Do(context.Background())
		if err != nil {
			fmt.Println("Error in query: ", err)
			os.Exit(1)
		}

		if searchResult.Hits.TotalHits > 0 {

			for ind, hit := range searchResult.Hits.Hits {
				var s ConvergeInfo
				err = json.Unmarshal(*hit.Source, &s)
				if err != nil {
					fmt.Println("Json marshal err :", err)
					os.Exit(1)
				}

				if ind == 0 {
					err = gocsv.Marshal([]ConvergeInfo{s}, f)
					if err != nil {
						fmt.Println("Error while writing csv: ", err)
						os.Exit(1)
					}
				} else {
					err = gocsv.MarshalWithoutHeaders([]ConvergeInfo{s}, f)
					if err != nil {
						fmt.Println("Error while writing csv: ", err)
						os.Exit(1)
					}
					writer.Flush()
				}
			}
		}
		if t == startTime {
			break
		}

		endTime = t
		t = endTime.Add(-24 * time.Hour)
		if t.Before(startTime) {
			t = startTime
		}
	}
	fmt.Println("The details of the runs can be found in : ", filename)
}

func GenerateComplianceNodeCount(esHostName string, esPort string, startTime time.Time, endTime time.Time) {
	elasticSearchURL := fmt.Sprintf("http://%s:%s", esHostName, esPort)
	client, err := elastic.NewClient(
		elastic.SetURL(elasticSearchURL),
		elastic.SetSniff(false),
	)
	if err != nil {
		fmt.Println("Elastic error : ", err)
		os.Exit(1)
	}

	queryElasticSearchComplianceNodeCount(client, startTime, endTime)
}

func GenerateComplianceNodeRunReport(esHostName string, esPort string, startTime time.Time, endTime time.Time) {
	elasticSearchURL := fmt.Sprintf("http://%s:%s", esHostName, esPort)
	client, err := elastic.NewClient(
		elastic.SetURL(elasticSearchURL),
		elastic.SetSniff(false),
	)
	if err != nil {
		fmt.Println("Elastic error : ", err)
		os.Exit(1)
	}

	queryElasticSearchComplianceNodeReport(client, startTime, endTime)
}

func queryElasticSearchComplianceNodeCount(client *elastic.Client, startTime time.Time, endTime time.Time) {
	filename := fmt.Sprintf("compliancenodecount_%s_%s.csv", startTime.Format("2006-01-02"), endTime.Format("2006-01-02"))
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Println("Failed to open Node Count report")
		os.Exit(1)
	}
	defer f.Close()
	writer := gocsv.DefaultCSVWriter(f)
	//writer := struct2csv.NewWriter(f)
	err = writer.Write([]string{"Start Time", "End Time", "Unique Node Count"})
	if err != nil {
		fmt.Println("Error while writing csv: ", err)
		os.Exit(1)
	}

	st := endTime.Add(24 * time.Hour)
	for {
		et := st
		st = et.Add(24 * -time.Hour)

		err = writer.Write([]string{"Day Count"})
		if err != nil {
			fmt.Println("Error while writing csv: ", err)
			os.Exit(1)
		}
		dayCount, ok := getUniqueComplianceCounts(client, st, et)
		if ok {
			err = writer.Write([]string{st.Format(time.RFC3339), et.Format(time.RFC3339), fmt.Sprintf("%f", *dayCount.Value)})
			if err != nil {
				fmt.Println("Error while writing csv: ", err)
				os.Exit(1)
			}
		}
		t := et.Add(-time.Hour)
		if t.Before(st) {
			t = st
		}
		err = writer.Write([]string{"Hourly Count"})
		if err != nil {
			fmt.Println("Error while writing csv: ", err)
			os.Exit(1)
		}
		writer.Flush()
		for {
			metric, ok := getUniqueComplianceCounts(client, t, et)
			if ok && *metric.Value > 0 {
				err = writer.Write([]string{t.Format(time.RFC3339), et.Format(time.RFC3339), fmt.Sprintf("%f", *metric.Value)})
				if err != nil {
					fmt.Println("Error while writing csv: ", err)
					os.Exit(1)
				}
				writer.Flush()
			}

			if t == st {
				break
			}

			et = t
			t = et.Add(-time.Hour)
			if t.Before(st) {
				t = st
			}
		}
		err = writer.Write([]string{"", "", ""})
		if err != nil {
			fmt.Println("Error while writing csv: ", err)
			os.Exit(1)
		}
		writer.Flush()
		if st == startTime {
			break
		}
	}
	writer.Flush()
	fmt.Println("The node count report is available in: ", filename)
}

func getUniqueComplianceCounts(client *elastic.Client, startTime time.Time, endTime time.Time) (*elastic.AggregationValueMetric, bool) {
	rangeQuery := elastic.NewRangeQuery("end_time").
		Format("yyyy-MM-dd||yyyy-MM-dd-HH:mm:ss||yyyy-MM-dd'T'HH:mm:ssZ")
	rangeQuery.Gte(startTime)
	rangeQuery.Lte(endTime)

	aggr := elastic.NewCardinalityAggregation().Field("node_uuid")
	searchService := client.Search().
		Query(rangeQuery).
		Index("comp-7-r-*").
		Size(100).
		Aggregation("nodes_count", aggr)

	searchResult, err := searchService.Do(context.Background())
	if err != nil {
		fmt.Println("Error in query: ", err)
		os.Exit(1)
	}

	metric, ok := searchResult.Aggregations.ValueCount("nodes_count")
	return metric, ok
}

func queryElasticSearchComplianceNodeReport(client *elastic.Client, startTime time.Time, endTime time.Time) {
	filename := fmt.Sprintf("complianceinfo_%s_%s.csv", startTime.Format("2006-01-02"), endTime.Format("2006-01-02"))
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Println("Failed to open Compliance scan detail report")
		os.Exit(1)
	}
	defer f.Close()

	w := gocsv.DefaultCSVWriter(f)

	var sourceField = []string{
		"node_uuid",
		"version",
		"node_name",
		"environment",
		"platform.name",
		"controls_sums.total",
		"controls_sums.passed.total",
		"controls_sums.skipped.total",
		"controls_sums.failed.total",
		"profiles.name",
		"profiles.title",
		"profiles.version",
		"end_time",
	}

	t := endTime.Add(24 * time.Hour)

	if t.Before(startTime) {
		t = startTime
	}
	header := true
	for {
		rangeQuery := elastic.NewRangeQuery("end_time").
			Format("yyyy-MM-dd||yyyy-MM-dd-HH:mm:ss||yyyy-MM-dd'T'HH:mm:ssZ")
		rangeQuery.Gte(t)
		rangeQuery.Lte(endTime)

		fetchSource := elastic.NewFetchSourceContext(true).Include(sourceField...)

		searchService := client.Search().
			Query(rangeQuery).
			Index("comp-7-r-*").
			Size(10000).
			FetchSourceContext(fetchSource)

		searchResult, err := searchService.Do(context.Background())
		if err != nil {
			fmt.Println("Error in query: ", err)
			os.Exit(1)
		}

		if searchResult.Hits.TotalHits > 0 {
			if header == true {
				headers := []string{"Node_ID", "Version", "Node_Name", "Environment", "End_Time", "Platform__Name", "Controls_Sums__Total", "Controls_Sums__Passed", "Controls_Sums__Skipped", "Controls_Sums__Failed", "Profiles__Count"}
				if err := w.Write(headers); err != nil {
					fmt.Println("Error while writing csv: ", err)
				}
				w.Flush()
				if err := w.Error(); err != nil {
					fmt.Println("Error while writing csv: ", err)
				}
			}

			for _, hit := range searchResult.Hits.Hits {
				
				var payload []byte

				payload = []byte("[")
				comma := []byte(",")
				payload = append(payload, *hit.Source...)
				payload = append(payload, comma...)
				payload = payload[:len(payload)-1]
				payload = append(payload, []byte("]")...)
				var s ComplianceInfo

				err = json.Unmarshal(payload, &s)
				if err != nil {
					fmt.Println("Json marshal err :", err)
					os.Exit(1)
				}

				for i := 0; i < len(s); i++ {
					record := []string{
						s.mapAt(i)["node_uuid"].(string),
						s.mapAt(i)["version"].(string),
						s.mapAt(i)["node_name"].(string),
						s.mapAt(i)["environment"].(string),
						s.mapAt(i)["end_time"].(string),
						s.mapAt(i)["platform"].(map[string]interface{})["name"].(string),
						fmt.Sprint(s.mapAt(i)["controls_sums"].(map[string]interface{})["total"]),
						fmt.Sprint(s.mapAt(i)["controls_sums"].(map[string]interface{})["passed"].(map[string]interface{})["total"]),
						fmt.Sprint(s.mapAt(i)["controls_sums"].(map[string]interface{})["skipped"].(map[string]interface{})["total"]),
						fmt.Sprint(s.mapAt(i)["controls_sums"].(map[string]interface{})["failed"].(map[string]interface{})["total"]),
						strconv.Itoa(len(s.mapAt(i)["profiles"].([]interface{}))),
					}
					if err := w.Write(record); err != nil {
						fmt.Println("Error while writing csv: ", err)

					}
					w.Flush()
					header = false
					if err := w.Error(); err != nil {
						fmt.Println("Error while writing csv: ", err)

					}

				}
			}
		}
		if t == startTime {
			break
		}

		endTime = t
		t = endTime.Add(-24 * time.Hour)
		if t.Before(startTime) {
			t = startTime
		}
	}
	fmt.Println("The details of the runs can be found in : ", filename)
}
