# Simple Automate Absent with Go

* https://heroku.com
* google calender API

## Support 

- www.talenta.co


## ENV
```
EMAIL=xxx
PASSWORD=xx
DEVICE_ID=xxx
LATITUDE=-6.3443904
LONGITUDE=106.8705516
SECRET_KEY=xxx          //for middleware service 
API_KEY_GOOGLE_CALENDER=xxxx
GOOGLE_CALENDER_ID=xxxx //by default use calenderId indonesia, for more calenderid visit https://gist.github.com/mattn/1438183
CRON_CHECKIN="1 8 * * 1,2,3,4,5" // https://crontab.guru/
CRON_CHECKOUT="1 8 * * 1,2,3,4,5" // https://crontab.guru/
```

## Setup
* SetUp heroku 
* SetUp Env file
* Change file selfie.JPG to your own photo

## API
```
/ping --> ping service
/checkin --> checkin absent
/checkout --> checkout absent
```

## Running with Your Cron
```go
//by default serve rest api
// CRON_CHECKIN = "1 8 * * 1,2,3,4,5" --> “At 08:01 on Monday, Tuesday,Wednesday, Thursday, and Friday.”
// CRON_CHECKOUT = "1 20 * * 1,2,3,4,5" --> “At 20:01 on Monday, Tuesday,Wednesday, Thursday, and Friday.”
func main() {
	log.Println("Start")
	helper.Dispatcher.Run()
	cron.RunJob()
	//running with rest api
	//gate.Route()

	log.Println("Finish")
}
```

