package operations_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/omise/omise-go/v2"
	"github.com/omise/omise-go/v2/internal/testutil"
	r "github.com/stretchr/testify/require"
)

var futureDate = omise.DateString(time.Now().AddDate(1, 0, 0).Format("2006-01-02"))

func TestCreateChargeScheduleMarshal(t *testing.T) {
	testdata := []struct {
		req      *omise.CreateScheduleParams
		expected string
	}{
		{
			req: &omise.CreateScheduleParams{
				Every:     3,
				Period:    omise.ScheduleDay,
				StartDate: "2017-05-15",
				EndDate:   "2018-05-15",
				Charge: &omise.ChargeScheduleParams{
					Customer: "customer_id",
					Amount:   100000,
				},
			},
			expected: `{"charge":{"amount":100000,"customer":"customer_id"},"end_date":"2018-05-15","every":3,"period":"day","start_date":"2017-05-15"}`,
		},
		{
			req: &omise.CreateScheduleParams{
				Every:  3,
				Period: omise.ScheduleWeek,
				On: &omise.ScheduleOnParams{
					Weekdays: []omise.Weekdays{
						omise.ScheduleOnMonday,
						omise.ScheduleOnSaturday,
					},
				},
				StartDate: "2017-05-15",
				EndDate:   "2018-05-15",
				Charge: &omise.ChargeScheduleParams{
					Customer: "customer_id",
					Amount:   100000,
				},
			},
			expected: `{"charge":{"amount":100000,"customer":"customer_id"},"end_date":"2018-05-15","every":3,"on":{"weekdays":["monday","saturday"]},"period":"week","start_date":"2017-05-15"}`,
		},
		{
			req: &omise.CreateScheduleParams{
				Every:  3,
				Period: omise.ScheduleMonth,
				On: &omise.ScheduleOnParams{
					DaysOfMonth: []int{1, 15},
				},
				StartDate: "2017-05-15",
				EndDate:   "2018-05-15",
				Charge: &omise.ChargeScheduleParams{
					Customer: "customer_id",
					Amount:   100000,
				},
			},
			expected: `{"charge":{"amount":100000,"customer":"customer_id"},"end_date":"2018-05-15","every":3,"on":{"days_of_month":[1,15]},"period":"month","start_date":"2017-05-15"}`,
		},
		{
			req: &omise.CreateScheduleParams{
				Every:  3,
				Period: omise.ScheduleMonth,
				On: &omise.ScheduleOnParams{
					WeekdayOfMonth: "last_thursday",
				},
				StartDate: "2017-05-15",
				EndDate:   "2018-05-15",
				Charge: &omise.ChargeScheduleParams{
					Customer: "customer_id",
					Amount:   100000,
				},
			},
			expected: `{"charge":{"amount":100000,"customer":"customer_id"},"end_date":"2018-05-15","every":3,"on":{"weekday_of_month":"last_thursday"},"period":"month","start_date":"2017-05-15"}`,
		},
	}

	for _, td := range testdata {
		b, err := json.Marshal(td.req)
		r.Nil(t, err, "err should be nothing")
		r.Equal(t, td.expected, string(b))
	}
}

func TestCreateChargeSchedule_Network(t *testing.T) {
	// CustomerID must have this customer in test server
	const CustomerID = `cust_57z9e1nce0wvbbkvef1`

	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)
	create := &omise.CreateScheduleParams{
		Every:  3,
		Period: omise.ScheduleWeek,
		On: &omise.ScheduleOnParams{
			Weekdays: []omise.Weekdays{
				omise.ScheduleOnMonday,
				omise.ScheduleOnSaturday,
			},
		},
		EndDate: futureDate,
		Charge: &omise.ChargeScheduleParams{
			Customer: CustomerID,
			Amount:   100000,
		},
	}

	client.Schedule().Create(context.Background(), create)
}

func TestCreateTransferScheduleMarshal(t *testing.T) {
	testdata := []struct {
		req      *omise.CreateScheduleParams
		expected string
	}{
		{
			req: &omise.CreateScheduleParams{
				Every:     3,
				Period:    omise.ScheduleDay,
				StartDate: "2017-05-15",
				EndDate:   "2018-05-15",
				Transfer: &omise.TransferScheduleParams{
					Recipient: "recipient_id",
					Amount:    100000,
				},
			},
			expected: `{"end_date":"2018-05-15","every":3,"period":"day","start_date":"2017-05-15","transfer":{"amount":100000,"recipient":"recipient_id"}}`,
		},
		{
			req: &omise.CreateScheduleParams{
				Every:  3,
				Period: omise.ScheduleWeek,
				On: &omise.ScheduleOnParams{
					Weekdays: []omise.Weekdays{
						omise.ScheduleOnMonday,
						omise.ScheduleOnSaturday,
					},
				},
				StartDate: "2017-05-15",
				EndDate:   "2018-05-15",
				Transfer: &omise.TransferScheduleParams{
					Recipient:           "recipient_id",
					PercentageOfBalance: 20.35,
				},
			},
			expected: `{"end_date":"2018-05-15","every":3,"on":{"weekdays":["monday","saturday"]},"period":"week","start_date":"2017-05-15","transfer":{"percentage_of_balance":20.35,"recipient":"recipient_id"}}`,
		},
		{
			req: &omise.CreateScheduleParams{
				Every:  3,
				Period: omise.ScheduleMonth,
				On: &omise.ScheduleOnParams{
					DaysOfMonth: []int{1, 15},
				},
				StartDate: "2017-05-15",
				EndDate:   "2018-05-15",
				Transfer: &omise.TransferScheduleParams{
					Recipient: "recipient_id",
					Amount:    100000,
				},
			},
			expected: `{"end_date":"2018-05-15","every":3,"on":{"days_of_month":[1,15]},"period":"month","start_date":"2017-05-15","transfer":{"amount":100000,"recipient":"recipient_id"}}`,
		},
		{
			req: &omise.CreateScheduleParams{
				Every:  3,
				Period: omise.ScheduleMonth,
				On: &omise.ScheduleOnParams{
					WeekdayOfMonth: "last_thursday",
				},
				StartDate: "2017-05-15",
				EndDate:   "2018-05-15",
				Transfer: &omise.TransferScheduleParams{
					Recipient:           "recipient_id",
					PercentageOfBalance: 50.55,
				},
			},
			expected: `{"end_date":"2018-05-15","every":3,"on":{"weekday_of_month":"last_thursday"},"period":"month","start_date":"2017-05-15","transfer":{"percentage_of_balance":50.55,"recipient":"recipient_id"}}`,
		},
	}

	for _, td := range testdata {
		b, err := json.Marshal(td.req)
		r.Nil(t, err, "err should be nothing")
		r.Equal(t, td.expected, string(b))
	}
}

func TestCreateTransferSchedule_Network(t *testing.T) {
	// RecipientID must have this recipient in test server
	const RecipientID = `recp_57z9e1nce0wvbbkvef1`

	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)
	create := &omise.CreateScheduleParams{
		Every:  3,
		Period: omise.ScheduleWeek,
		On: &omise.ScheduleOnParams{
			Weekdays: []omise.Weekdays{
				omise.ScheduleOnMonday,
				omise.ScheduleOnSaturday,
			},
		},
		EndDate: futureDate,

		Transfer: &omise.TransferScheduleParams{
			Recipient: RecipientID,
			Amount:    100000,
		},
	}
	client.Schedule().Create(context.Background(), create)
}

func TestCreateSchedule(t *testing.T) {
	const (
		ScheduleID = "schd_57z9hj228pusa652nk1"
	)

	client := testutil.NewFixedClient(t)

	schd, _ := client.Schedule().Create(context.Background(), &omise.CreateScheduleParams{})
	r.Equal(t, ScheduleID, schd.ID)
}

func TestListSchedule(t *testing.T) {
	client := testutil.NewFixedClient(t)
	schds, _ := client.Schedule().List(context.Background(), &omise.ListSchedulesParams{})

	r.Len(t, schds.Data, 2)

	r.Equal(t, "schd_57zhl296uxc7yiun6xa", schds.Data[0].ID)
	r.NotNil(t, schds.Data[0].Charge)
	r.Nil(t, schds.Data[0].Transfer)

	r.Equal(t, "schd_57zhl296uxc7yiun6xx", schds.Data[1].ID)
	r.NotNil(t, schds.Data[1].Transfer)
	r.Nil(t, schds.Data[1].Charge)
}

func TestListSchedules_Network(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)
	list := omise.ListSchedulesParams{
		ListParams: omise.ListParams{
			Limit: 100,
			From:  time.Date(2017, 5, 16, 0, 0, 0, 0, time.Local),
		},
	}

	schds, _ := client.Schedule().List(context.Background(), &list)

	t.Logf("Schedules Len: %d\n", len(schds.Data))
	t.Logf("%#v\n", schds)
}

func TestListChargeSchedules(t *testing.T) {
	client := testutil.NewFixedClient(t)
	schds, _ := client.Charge().ListSchedules(context.Background(), &omise.ListChargeSchedulesParams{})

	r.Len(t, schds.Data, 1)

	r.Equal(t, "schd_57zhl296uxc7yiun6xa", schds.Data[0].ID)
}

func TestListChargeSchedules_Network(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)
	list := omise.ListChargeSchedulesParams{
		ListParams: omise.ListParams{
			Limit: 100,
			From:  time.Date(2017, 5, 16, 0, 0, 0, 0, time.Local),
		},
	}

	schds, _ := client.Charge().ListSchedules(context.Background(), &list)

	t.Logf("Schedules Len: %d\n", len(schds.Data))
	t.Logf("%#v\n", schds)
}

func TestListTransferSchedules(t *testing.T) {
	client := testutil.NewFixedClient(t)
	schds, _ := client.Transfer().ListSchedules(context.Background(), &omise.ListTransferSchedulesParams{})

	r.Len(t, schds.Data, 1)

	r.Equal(t, "schd_57zhl296uxc7yiun6xx", schds.Data[0].ID)
}

func TestListTransferSchedules_Network(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)
	list := omise.ListTransferSchedulesParams{
		ListParams: omise.ListParams{
			Limit: 100,
			From:  time.Date(2017, 5, 16, 0, 0, 0, 0, time.Local),
		},
	}

	schds, _ := client.Transfer().ListSchedules(context.Background(), &list)

	t.Logf("Schedules Len: %d\n", len(schds.Data))
	t.Logf("%#v\n", schds)
}

func TestListScheduleOccurrences(t *testing.T) {
	client := testutil.NewFixedClient(t)
	occurrences, _ := client.Schedule().ListOccurrences(context.Background(), &omise.ListScheduleOccurrencesParams{ScheduleID: "schd_57z9hj228pusa652nk1"})

	r.Len(t, occurrences.Data, 2)

	r.Equal(t, "occu_588fwvt863w1w24et5u", occurrences.Data[0].ID)
	r.Equal(t, "occu_588fwvt863w1w24et7u", occurrences.Data[1].ID)
}

func TestListScheduleOccurrences_Network(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)
	list := omise.ListScheduleOccurrencesParams{
		ScheduleID: "schd_57z9hj228pusa652nk1",
		ListParams: omise.ListParams{
			Limit: 100,
			From:  time.Date(2017, 5, 16, 0, 0, 0, 0, time.Local),
		},
	}

	occurrences, _ := client.Schedule().ListOccurrences(context.Background(), &list)

	t.Logf("Occurrences Len: %d\n", len(occurrences.Data))
	t.Logf("%#v\n", occurrences)
}

func TestRetrieveSchedule(t *testing.T) {
	ScheduleID := "schd_57z9hj228pusa652nk1"

	client := testutil.NewFixedClient(t)
	schd, _ := client.Schedule().Retrieve(context.Background(), &omise.RetrieveScheduleParams{ScheduleID: ScheduleID})
	r.Equal(t, ScheduleID, schd.ID)
	r.Nil(t, schd.Transfer)
	r.Equal(t, int64(100000), schd.Charge.Amount)
	r.Equal(t, omise.ScheduleActive, schd.Status)
	r.Len(t, schd.NextOccurrencesOn, 30)

	ScheduleID = "schd_57z9hj228pusa652nk2"

	schd, _ = client.Schedule().Retrieve(context.Background(), &omise.RetrieveScheduleParams{ScheduleID: ScheduleID})
	r.Equal(t, ScheduleID, schd.ID)
	r.Nil(t, schd.Charge)
	r.Equal(t, int64(100000), *schd.Transfer.Amount)
	r.Equal(t, omise.ScheduleActive, schd.Status)
	r.Len(t, schd.NextOccurrencesOn, 30)
}

func TestRetrieveSchedule_Network(t *testing.T) {
	// ScheduleID must have this schedule in test server
	ScheduleID := "schd_57z9hj228pusa652nk1"

	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)
	schd, _ := client.Schedule().Retrieve(context.Background(), &omise.RetrieveScheduleParams{ScheduleID: ScheduleID})

	t.Logf("%#v\n", schd)
}

func TestDestroySchedule(t *testing.T) {
	ScheduleID := "schd_57z9hj228pusa652nk1"

	client := testutil.NewFixedClient(t)
	schd, _ := client.Schedule().Destroy(context.Background(), &omise.DestroyScheduleParams{ScheduleID: ScheduleID})
	r.Equal(t, ScheduleID, schd.ID)
	r.Nil(t, schd.Transfer)
	r.Equal(t, int64(100000), schd.Charge.Amount)
	r.Equal(t, omise.ScheduleDeleted, schd.Status)
	r.Len(t, schd.NextOccurrencesOn, 30)

	ScheduleID = "schd_57z9hj228pusa652nk2"

	schd, _ = client.Schedule().Destroy(context.Background(), &omise.DestroyScheduleParams{ScheduleID: ScheduleID})
	r.Equal(t, ScheduleID, schd.ID)
	r.Nil(t, schd.Charge)
	r.Equal(t, int64(100000), *schd.Transfer.Amount)
	r.Equal(t, omise.ScheduleDeleted, schd.Status)
	r.Len(t, schd.NextOccurrencesOn, 30)
}

func TestDestroySchedule_Network(t *testing.T) {
	// ScheduleID must have this schedule in test server
	ScheduleID := "schd_57z9hj228pusa652nk1"

	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)
	schd, _ := client.Schedule().Destroy(context.Background(), &omise.DestroyScheduleParams{ScheduleID: ScheduleID})

	t.Logf("%#v\n", schd)
}
