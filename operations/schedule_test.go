package operations_test

import (
	"encoding/json"
	"testing"
	"time"

	omise "github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	. "github.com/omise/omise-go/operations"
	"github.com/omise/omise-go/schedule"
	r "github.com/stretchr/testify/require"
)

func TestCreateScheduleMarshal(t *testing.T) {
	testdata := []struct {
		req      *CreateSchedule
		expected string
	}{
		{
			req: &CreateSchedule{
				Every:     3,
				Period:    schedule.PeriodDay,
				StartDate: "2017-05-15",
				EndDate:   "2018-05-15",
				Customer:  "customer_id",
				Amount:    100000,
			},
			expected: `{"every":3,"period":"day","start_date":"2017-05-15","end_date":"2018-05-15","charge":{"customer":"customer_id","amount":100000}}`,
		},
		{
			req: &CreateSchedule{
				Every:  3,
				Period: schedule.PeriodWeek,
				Weekdays: []schedule.Weekday{
					schedule.Monday,
					schedule.Saturday,
				},
				StartDate: "2017-05-15",
				EndDate:   "2018-05-15",
				Customer:  "customer_id",
				Amount:    100000,
			},
			expected: `{"every":3,"period":"week","start_date":"2017-05-15","end_date":"2018-05-15","charge":{"customer":"customer_id","amount":100000},"on":{"weekdays":["monday","saturday"]}}`,
		},
		{
			req: &CreateSchedule{
				Every:       3,
				Period:      schedule.PeriodMonth,
				DaysOfMonth: []int{1, 15},
				StartDate:   "2017-05-15",
				EndDate:     "2018-05-15",
				Customer:    "customer_id",
				Amount:      100000,
			},
			expected: `{"every":3,"period":"month","start_date":"2017-05-15","end_date":"2018-05-15","charge":{"customer":"customer_id","amount":100000},"on":{"days_of_month":[1,15]}}`,
		},
		{
			req: &CreateSchedule{
				Every:          3,
				Period:         schedule.PeriodMonth,
				WeekdayOfMonth: "last_thursday",
				StartDate:      "2017-05-15",
				EndDate:        "2018-05-15",
				Customer:       "customer_id",
				Amount:         100000,
			},
			expected: `{"every":3,"period":"month","start_date":"2017-05-15","end_date":"2018-05-15","charge":{"customer":"customer_id","amount":100000},"on":{"weekday_of_month":"last_thursday"}}`,
		},
	}

	for _, td := range testdata {
		b, err := json.Marshal(td.req)
		r.Nil(t, err, "err should be nothing")
		r.Equal(t, td.expected, string(b))
	}
}

func TestSchedule(t *testing.T) {
	const (
		ScheduleID = "schd_57z9hj228pusa652nk1"
	)

	client := testutil.NewFixedClient(t)

	schd := &omise.Schedule{}
	client.MustDo(schd, &CreateSchedule{})
	r.Equal(t, ScheduleID, schd.ID)

	schds := &omise.ScheduleList{}
	client.MustDo(schds, &ListSchedules{})
	r.Len(t, schds.Data, 1)
	r.Equal(t, "schd_57zhl296uxc7yiun6xa", schds.Data[0].ID)
}

func TestCreateSchedule_Network(t *testing.T) {
	// CustomerID must have this customer in test server
	const CustomerID = `cust_57z9e1nce0wvbbkvef1`

	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)
	schd, create := &omise.Schedule{}, &CreateSchedule{
		Every:  3,
		Period: schedule.PeriodWeek,
		Weekdays: []schedule.Weekday{
			schedule.Monday,
			schedule.Saturday,
		},
		EndDate:  "2018-05-15",
		Customer: CustomerID,
		Amount:   100000,
	}
	client.MustDo(schd, create)
}

func TestListSchedules_Network(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)
	schds, list := &omise.ScheduleList{}, &ListSchedules{
		List{
			Limit: 100,
			From:  time.Date(2017, 5, 16, 0, 0, 0, 0, time.Local),
		},
	}
	client.MustDo(schds, list)

	t.Logf("Schedules Len: %d\n", len(schds.Data))
	t.Logf("%#v\n", schds)
}
