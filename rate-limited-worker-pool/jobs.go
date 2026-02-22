package main

import "time"

type Job struct {
	Id       int
	Payload  string
	Duration time.Duration
}

var Jobs = []Job{
	{Id: 1, Payload: "Process image", Duration: time.Millisecond * 300},
	{Id: 2, Payload: "Send email", Duration: time.Millisecond * 400},
	{Id: 3, Payload: "Generate report", Duration: time.Millisecond * 500},
	{Id: 4, Payload: "Cleanup temp files", Duration: time.Millisecond * 600},
	{Id: 5, Payload: "Backup database", Duration: time.Millisecond * 700},
	{Id: 6, Payload: "Validate input", Duration: time.Millisecond * 800},
	{Id: 7, Payload: "Resize photo", Duration: time.Millisecond * 900},
	{Id: 8, Payload: "Authenticate user", Duration: time.Millisecond * 1000},
	{Id: 9, Payload: "Send notification", Duration: time.Millisecond * 350},
	{Id: 10, Payload: "Compress log", Duration: time.Millisecond * 450},
	{Id: 11, Payload: "Process image", Duration: time.Millisecond * 550},
	{Id: 12, Payload: "Send email", Duration: time.Millisecond * 650},
	{Id: 13, Payload: "Generate report", Duration: time.Millisecond * 750},
	{Id: 14, Payload: "Cleanup temp files", Duration: time.Millisecond * 850},
	{Id: 15, Payload: "Backup database", Duration: time.Millisecond * 950},
	{Id: 16, Payload: "Validate input", Duration: time.Millisecond * 300},
	{Id: 17, Payload: "Resize photo", Duration: time.Millisecond * 400},
	{Id: 18, Payload: "Authenticate user", Duration: time.Millisecond * 500},
	{Id: 19, Payload: "Send notification", Duration: time.Millisecond * 600},
	{Id: 20, Payload: "Compress log", Duration: time.Millisecond * 700},
	{Id: 21, Payload: "Process image", Duration: time.Millisecond * 800},
	{Id: 22, Payload: "Send email", Duration: time.Millisecond * 900},
	{Id: 23, Payload: "Generate report", Duration: time.Millisecond * 1000},
	{Id: 24, Payload: "Cleanup temp files", Duration: time.Millisecond * 350},
	{Id: 25, Payload: "Backup database", Duration: time.Millisecond * 450},
	{Id: 26, Payload: "Validate input", Duration: time.Millisecond * 550},
	{Id: 27, Payload: "Resize photo", Duration: time.Millisecond * 650},
	{Id: 28, Payload: "Authenticate user", Duration: time.Millisecond * 750},
	{Id: 29, Payload: "Send notification", Duration: time.Millisecond * 850},
	{Id: 30, Payload: "Compress log", Duration: time.Millisecond * 950},
	{Id: 31, Payload: "Process image", Duration: time.Millisecond * 300},
	{Id: 32, Payload: "Send email", Duration: time.Millisecond * 400},
	{Id: 33, Payload: "Generate report", Duration: time.Millisecond * 500},
	{Id: 34, Payload: "Cleanup temp files", Duration: time.Millisecond * 600},
	{Id: 35, Payload: "Backup database", Duration: time.Millisecond * 700},
	{Id: 36, Payload: "Validate input", Duration: time.Millisecond * 800},
	{Id: 37, Payload: "Resize photo", Duration: time.Millisecond * 900},
	{Id: 38, Payload: "Authenticate user", Duration: time.Millisecond * 1000},
	{Id: 39, Payload: "Send notification", Duration: time.Millisecond * 350},
	{Id: 40, Payload: "Compress log", Duration: time.Millisecond * 450},
	{Id: 41, Payload: "Process image", Duration: time.Millisecond * 550},
	{Id: 42, Payload: "Send email", Duration: time.Millisecond * 650},
	{Id: 43, Payload: "Generate report", Duration: time.Millisecond * 750},
	{Id: 44, Payload: "Cleanup temp files", Duration: time.Millisecond * 850},
	{Id: 45, Payload: "Backup database", Duration: time.Millisecond * 950},
	{Id: 46, Payload: "Validate input", Duration: time.Millisecond * 300},
	{Id: 47, Payload: "Resize photo", Duration: time.Millisecond * 400},
	{Id: 48, Payload: "Authenticate user", Duration: time.Millisecond * 500},
	{Id: 49, Payload: "Send notification", Duration: time.Millisecond * 600},
	{Id: 50, Payload: "Compress log", Duration: time.Millisecond * 700},
}
