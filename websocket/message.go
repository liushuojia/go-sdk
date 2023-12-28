package websocketConn

type Message struct {
	ID      int64  `json:"id,omitempty"`      // ID
	UUID    string `json:"uuid,omitempty"`    // UUID
	Action  string `json:"action,omitempty"`  // 类型 heartbeat publish subscribe
	Title   string `json:"title,omitempty"`   // 标题
	Content string `json:"content,omitempty"` // 内容
	Link    string `json:"link,omitempty"`    // 链接
}
