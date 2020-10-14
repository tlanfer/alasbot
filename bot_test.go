package alasbot

import (
	"fmt"
	"testing"
)

func TestBloodMoonMessage(t *testing.T) {
	type args struct {
		days               int
		hours              int
		minutes            int
		bloodMoonFrequency int
		offset             int
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args{
				1,
				10,
				10,
				7,
				0,
			},
			"The next bloodmoon will be on day 7.",
		},
		{
			args{
				600,
				10,
				10,
				7,
				0,
			},
			"The next bloodmoon will be on day 602.",
		},
		{
			args{
				603,
				10,
				10,
				7,
				0,
			},
			"The next bloodmoon will be on day 609.",
		},
		{
			args{
				600,
				10,
				10,
				14,
				595,
			},
			"The next bloodmoon will be on day 609.",
		},
		{
			args{
				603,
				10,
				10,
				14,
				595,
			},
			"The next bloodmoon will be on day 609.",
		},
		{
			args{
				609,
				10,
				10,
				14,
				595,
			},
			"The next bloodmoon will be today.",
		},
		{
			args{
				609,
				23,
				10,
				14,
				595,
			},
			"A bloodmoon is active!",
		},
		{
			args{
				610,
				02,
				10,
				14,
				595,
			},
			"A bloodmoon is active!",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.args), func(t *testing.T) {
			if got := BloodMoonMessage(tt.args.days, tt.args.hours, tt.args.minutes, tt.args.bloodMoonFrequency, tt.args.offset); got != tt.want {
				t.Errorf("BloodMoonMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
