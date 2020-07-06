package helper

import "github.com/qasir-id/qasirworker"

var (
	Dispatcher = qasirworker.NewDispatcher(10, 9999999)
	TimeZone   = "Asia/Jakarta"
)
