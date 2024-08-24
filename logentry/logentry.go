package logentry

import "time"

type LogEntry struct {
	Timestamp      time.Time `json:"timestamp"`
	ServiceName    string    `json:"service_name"`
	LogLevel       string    `json:"log_level"`
	CorrelationID  string    `json:"correlation_id"`
	ErrorMessage   string    `json:"error_message"`
	Error          string    `json:"error"`
	UserContext    string    `json:"user_context"`
	HTTPMethod     string    `json:"http_method"`
	URL            string    `json:"url"`
	ResponseStatus int       `json:"response_status"`
	Payload        string    `json:"payload"`
	ExecutionTime  int64     `json:"execution_time"`
}
