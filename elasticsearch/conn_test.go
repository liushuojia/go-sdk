package elasticsearchConn

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestMysql(t *testing.T) {

	es := New().SetAddresses("http://127.0.0.1:9200")
	if err := es.Connect(); err != nil {
		log.Fatalln(err)
	}

	type document struct {
		Name string `json:"name"`
		Type int    `json:"type"`
	}

	var l []document

	fmt.Println(
		es.Query(
			"test001",
			strings.NewReader(`{"query": { "match": { "name" : "go" } }}`),
			&l,
		),
	)
	fmt.Println(l)
	//fmt.Println(
	//es.Update("test001", "0079", strings.NewReader(`{"doc": { "language": "Go" }}`)),
	//es.Update("test001", "0079", strings.NewReader(`{ ”language“: "Go" }`)),
	//)

	//queryTMP := `{ "query": { "match": { "title":"one" } } }`
	//fmt.Println(
	//	client.Search(
	//		client.Search.WithIndex("test001"),
	//		client.Search.WithBody(strings.NewReader(queryTMP)),
	//	),
	//)
	//type document struct {
	//	Name string `json:"name"`
	//	Type int    `json:"type"`
	//}
	//
	//var data document
	//fmt.Println(es.Get("test001", "0079", &data))
	//fmt.Println(data)
	//
	//data.Type = data.Type + 10000
	//
	//id, err := es.Save("test001", "0080", data)
	//fmt.Println("====================")
	//fmt.Println(id)
	//fmt.Println("====================")
	//fmt.Println(err)
	//fmt.Println("====================")
	//
	//fmt.Println("====================")
	//fmt.Println(es.Delete("test001", "0080"))
	//fmt.Println("====================")

	//fmt.Println(es.IndexDelete("my"))

	//info, err := es.Info()
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Println(info)

	//
	//var data document
	//fmt.Println(es.Get("test001", "007", &data))
	//fmt.Println(data)
	//
	//data.Type = data.Type + 10000
	//
	//fmt.Println(es.Save("test001", "0079", data))

	//fmt.Println(client.Indices.Create("my_index1"))
	//document := struct {
	//	Name string `json:"name"`
	//	Type int    `json:"type"`
	//}{
	//	Name: "go-elasticsearch221112222",
	//	Type: 44,
	//}
	//
	//fmt.Println(es.Save("test001", "003", document))
	//fmt.Println(es.Save("test001", "002", document))

	//fmt.Println(es.Save("test", "aa", document))
	//var buf bytes.Buffer
	//if err := json.NewEncoder(&buf).Encode(document); err != nil {
	//	log.Fatalf("Error encoding query: %s", err)
	//}

	//req := esapi.IndexRequest{
	//	Index: "test",
	//	//DocumentID: "10099",
	//	Body:    bytes.NewReader(buf.Bytes()),
	//	Refresh: "true",
	//}

	return

	//var (
	//	r  map[string]interface{}
	//	wg sync.WaitGroup
	//)

	//client, err := elasticsearch.NewClient(
	//	elasticsearch.Config{
	//		Addresses: []string{
	//			"http://127.0.0.1:9200",
	//		},
	//	},
	//)

	//if err != nil {
	//	// Handle error
	//	fmt.Printf("连接失败: %v\n", err)
	//} else {
	//	fmt.Println("连接成功")
	//}

	//fmt.Println(client.Indices.Create("my_index1"))
	//document := struct {
	//	Name string `json:"name"`
	//	Type int    `json:"type"`
	//}{
	//	Name: "go-elasticsearch",
	//	Type: 2,
	//}
	//var buf bytes.Buffer
	//if err := json.NewEncoder(&buf).Encode(document); err != nil {
	//	log.Fatalf("Error encoding query: %s", err)
	//}

	//req := esapi.IndexRequest{
	//	Index: "test",
	//	//DocumentID: "10099",
	//	Body:    bytes.NewReader(buf.Bytes()),
	//	Refresh: "true",
	//}
	//res, err := req.Do(context.Background(), client)
	//if err != nil {
	//	log.Fatalf("Error getting response: %s", err)
	//}
	//
	//fmt.Println(res)
	//fmt.Println(client.Indices.Delete([]string{"test"}))
	//fmt.Println(client.Delete("test", "1"))
	//fmt.Println(client.Delete("test", "1333"))

	//return
	//
	//fmt.Println(client.Get("test", "1"))
	//

	//response, err := client.Indices.Exists([]string{"my_index2"})
	//fmt.Println(err)
	//fmt.Println(response)
	//fmt.Println(response.StatusCode)
	//fmt.Println(response.Status())
	//fmt.Println(response.String())
	//fmt.Println(response.IsError())
	//fmt.Println(response.Warnings())
	//fmt.Println(response.HasWarnings())
	//
	//log.Println(strings.Repeat("~", 37))
	//response, err = client.Indices.Exists([]string{"my_index1"})
	//fmt.Println(err)
	//fmt.Println(response)
	//fmt.Println(response.StatusCode)
	//fmt.Println(response.Status())
	//fmt.Println(response.String())
	//fmt.Println(response.IsError())
	//fmt.Println(response.Warnings())
	//fmt.Println(response.HasWarnings())
	//
	//log.Println(strings.Repeat("~", 37))

	//fmt.Println(client.Indices.Create("my_index"))
	//fmt.Println(client.Indices.Create("my_index1"))
	//document = struct {
	//	Name string `json:"name"`
	//	Type int    `json:"type"`
	//}{
	//	Name: "go-elasticsearch",
	//	Type: 2,
	//}
	//data, _ = json.Marshal(document)
	//fmt.Println(client.Index("my_index1", bytes.NewReader(data)))
	//
	//fmt.Println(client.Index("my_index1", bytes.NewReader(data)))
	//
	//req = esapi.IndexRequest{
	//	Index:      "test",
	//	DocumentID: "10099",
	//	Body:       bytes.NewReader(data),
	//	Refresh:    "true",
	//}
	//res, err = req.Do(context.Background(), client)
	//if err != nil {
	//	log.Fatalf("Error getting response: %s", err)
	//}
	//defer res.Body.Close()
	//fmt.Println(res)

	//return

	// 2. Index documents concurrently
	//
	//for i, title := range []string{"Test One", "Test Two"} {
	//	wg.Add(1)
	//
	//	go func(i int, title string) {
	//		defer wg.Done()
	//
	//		// Build the request body.
	//		data, err := json.Marshal(struct {
	//			Title  string `json:"title"`
	//			Status int    `json:"status"`
	//		}{
	//			Title:  title,
	//			Status: i + 10,
	//		})
	//		if err != nil {
	//			log.Fatalf("Error marshaling document: %s", err)
	//		}
	//
	//		// Set up the request object.
	//		req := esapi.IndexRequest{
	//			Index:      "test",
	//			DocumentID: strconv.Itoa(i + 1),
	//			Body:       bytes.NewReader(data),
	//			Refresh:    "true",
	//		}
	//
	//		// Perform the request with the client.
	//		res, err := req.Do(context.Background(), client)
	//		if err != nil {
	//			log.Fatalf("Error getting response: %s", err)
	//		}
	//		defer res.Body.Close()
	//
	//		if res.IsError() {
	//			log.Printf("[%s] Error indexing document ID=%d", res.Status(), i+1)
	//		} else {
	//			// Deserialize the response into a map.
	//			var r map[string]interface{}
	//			if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
	//				log.Printf("Error parsing the response body: %s", err)
	//			} else {
	//				// Print the response status and indexed document version.
	//				log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
	//			}
	//		}
	//	}(i, title)
	//}
	//wg.Wait()
	//
	//log.Println(strings.Repeat("-", 37))
	// 3. Search for the indexed documents
	//
	// Build the request body.
	//var buf bytes.Buffer
	//query := map[string]interface{}{
	//	"query": map[string]interface{}{
	//		"match": map[string]interface{}{
	//			"title": "test",
	//		},
	//	},
	//}
	//if err := json.NewEncoder(&buf).Encode(query); err != nil {
	//	log.Fatalf("Error encoding query: %s", err)
	//}

	// Perform the search request.
	//res, err = client.Search(
	//	client.Search.WithContext(context.Background()),
	//	client.Search.WithIndex("test"),
	//	client.Search.WithBody(&buf),
	//	client.Search.WithTrackTotalHits(true),
	//	client.Search.WithPretty(),
	//)
	//if err != nil {
	//	log.Fatalf("Error getting response: %s", err)
	//}
	//defer res.Body.Close()
	//
	//if res.IsError() {
	//	var e map[string]interface{}
	//	if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
	//		log.Fatalf("Error parsing the response body: %s", err)
	//	} else {
	//		// Print the response status and error information.
	//		log.Fatalf("[%s] %s: %s",
	//			res.Status(),
	//			e["error"].(map[string]interface{})["type"],
	//			e["error"].(map[string]interface{})["reason"],
	//		)
	//	}
	//}

	//if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
	//	log.Fatalf("Error parsing the response body: %s", err)
	//}
	//// Print the response status, number of results, and request duration.
	//log.Printf(
	//	"[%s] %d hits; took: %dms",
	//	res.Status(),
	//	int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
	//	int(r["took"].(float64)),
	//)
	//// Print the ID and document source for each hit.
	//for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
	//	log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	//}
	//
	//log.Println(strings.Repeat("=", 37))
}
