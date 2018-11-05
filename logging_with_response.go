package handlers

import (
	"io"
	"net/http"
)

func writeResponseLog(writer io.Writer, params LogFormatterParams) {
	buf := buildCommonLogLine(params.Request, params.URL, params.TimeStamp, params.StatusCode, params.Size)
	buf = append(buf, ` "`...)
	buf = appendQuoted(buf, params.Request.Referer())
	buf = append(buf, `" "`...)
	buf = appendQuoted(buf, params.Request.UserAgent())
	buf = append(buf, `" `...)
	buf = appendQuoted(buf, params.Duration.String())
	buf = append(buf, ` "`...)
	buf = appendQuoted(buf, string(params.Body))
	buf = append(buf, '"', '\n')
	writer.Write(buf)
}

// CombinedLoggingHandler return a http.Handler that wraps h and logs requests to out in
// Apache Combined Log Format.
//
// See http://httpd.apache.org/docs/2.2/logs.html#combined for a description of this format.
//
// LoggingHandler always sets the ident field of the log to -
func ResponseLoggingHandler(out io.Writer, h http.Handler) http.Handler {
	return loggingHandler{out, h, writeResponseLog}
}
