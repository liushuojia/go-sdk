package elasticsearchConn

import "time"

type Info struct {
	Name        string      `json:"name"`
	ClusterName string      `json:"cluster_name"`
	ClusterUUID string      `json:"cluster_uuid"`
	Version     InfoVersion `json:"version"`
	Tagline     string      `json:"tagline"`
}
type InfoVersion struct {
	Number                           string    `json:"number"`
	BuildFlavor                      string    `json:"build_flavor"`
	BuildType                        string    `json:"build_type"`
	BuildHash                        string    `json:"build_hash"`
	BuildDate                        time.Time `json:"build_date"`
	BuildSnapshot                    bool      `json:"build_snapshot"`
	LuceneVersion                    string    `json:"lucene_version"`
	MinimumWireCompatibilityVersion  string    `json:"minimum_wire_compatibility_version"`
	MinimumIndexCompatibilityVersion string    `json:"minimum_index_compatibility_version"`
}

type Response struct {
	Index         string         `json:"_index"`
	Type          string         `json:"_type"`
	Id            string         `json:"_id"`
	Version       int64          `json:"_version"`
	Result        string         `json:"result,omitempty"` // created | updated
	ForcedRefresh bool           `json:"forced_refresh,omitempty"`
	Found         bool           `json:"found,omitempty"`
	SeqNo         int64          `json:"_seq_no"`
	PrimaryTerm   int64          `json:"_primary_term"`
	Shards        ResponseShards `json:"_shards,omitempty"`
	Source        interface{}    `json:"_source,omitempty"`
}

type ResponseShards struct {
	Total      int64 `json:"total"`
	Successful int64 `json:"successful"`
	Skipped    int64 `json:"skipped,omitempty"`
	Failed     int64 `json:"failed"`
}

type QueryResponse struct {
	Took     int64          `json:"took"`
	TimedOut bool           `json:"timed_out"`
	Shards   ResponseShards `json:"_shards"`
	Hits     struct {
		Total struct {
			Value    int64  `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index  string      `json:"_index"`
			Type   string      `json:"_type"`
			Id     string      `json:"_id"`
			Score  float64     `json:"_score"`
			Source interface{} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}
