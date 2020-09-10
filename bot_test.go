package alasbot

import "testing"

func Test_bloodMoonMessage(t *testing.T) {
	type args struct {
		days    int
		hours   int
		minutes int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"normal day",
			args{1, 6, 3},
			"The next bloodmoon will be on day 7.",
		},
		{
			"bloodmoon day, early",
			args{7, 14, 12},
			"The next bloodmoon will be today.",
		},
		{
			"bloodmoon day, early",
			args{7, 23, 12},
			"A bloodmoon is active!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bloodMoonMessage(tt.args.days, tt.args.hours, tt.args.minutes); got != tt.want {
				t.Errorf("bloodMoonMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
