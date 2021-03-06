package linodego

import (
	log "github.com/Sirupsen/logrus"
	"time"
)

type CustomTime struct {
	time.Time
}

const ctLayout = "2006-01-02 15:04:05.0"

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	if len(b) == 0 {
		return
	}
	loc, _ := time.LoadLocation("America/New_York")
	ct.Time, err = time.ParseInLocation(ctLayout, string(b), loc)
	return
}

func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(ct.Time.Format(ctLayout)), nil
}

var nilTime = (time.Time{}).UnixNano()

func (ct *CustomTime) IsSet() bool {
	return ct.UnixNano() != nilTime
}

type CustomShortTime struct {
	time.Time
}

const ctLayout2 = "2006-01-02 15:04:05"

func (ct *CustomShortTime) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	if len(b) == 0 {
		return
	}
	loc, _ := time.LoadLocation("America/New_York")
	ct.Time, err = time.ParseInLocation(ctLayout2, string(b), loc)
	return
}

func (ct *CustomShortTime) MarshalJSON() ([]byte, error) {
	return []byte(ct.Time.Format(ctLayout2)), nil
}

func (ct *CustomShortTime) IsSet() bool {
	return ct.UnixNano() != nilTime
}

// expose a means to wait for a linode's pending jobs to complete
func WaitForPendingJobs(client *Client, linodeid int) {
	for {
		jobListResp, err := client.Job.List(linodeid, 0, true)
		if err != nil {
			log.Fatal(err)
		}
		if len(jobListResp.Jobs) == 0 {
			break
		}
		time.Sleep(time.Second * 5)
	}
}
