package period

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewPeriod(t *testing.T) {
	type args struct {
		startDate    time.Time
		endDate      time.Time
		boundaryType string
	}
	tests := []struct {
		name string
		args args
		want Period
	}{
		{
			name: "NewPeriod_WithStartDateBeforeEndDate",
			args: args{
				startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeAll,
			},
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeAll,
			},
		},
		{
			name: "NewPeriod_WithInvalidBoundaryType",
			args: args{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: "Invalid",
			},
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := NewPeriod(tt.args.startDate, tt.args.endDate, tt.args.boundaryType)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestNewDefaultPeriod(t *testing.T) {
	type args struct {
		startDate time.Time
		endDate   time.Time
	}
	tests := []struct {
		name string
		args args
		want Period
	}{
		{
			name: "NewDefaultPeriod_WithStartDateBeforeEndDate",
			args: args{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
			},
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name: "NewDefaultPeriod_WithStartDateAfterEndDate",
			args: args{
				startDate: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
			},
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := NewDefaultPeriod(tt.args.startDate, tt.args.endDate)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestNewIncludeAllPeriod(t *testing.T) {
	type args struct {
		startDate time.Time
		endDate   time.Time
	}
	tests := []struct {
		name string
		args args
		want Period
	}{
		{
			name: "NewIncludeAllPeriod_ValidDates",
			args: args{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
			},
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeAll,
			},
		},
		{
			name: "NewIncludeAllPeriod_StartDateAfterEndDate",
			args: args{
				startDate: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
			},
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeAll,
			},
		},
		{
			name: "NewIncludeAllPeriod_SameStartAndEndDate",
			args: args{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
			},
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeAll,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := NewIncludeAllPeriod(tt.args.startDate, tt.args.endDate)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestBordersOnStart(t *testing.T) {
	type args struct {
		other Period
	}
	tests := []struct {
		name string
		p    Period
		args args
		want bool
	}{
		{
			name: "BordersOnStart_WithSameEndDateAndStartDate",
			p:    Period{endDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), boundaryType: IncludeStartExcludeEnd},
			args: args{other: Period{startDate: time.Date(
				2023, 1, 1, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd}},
			want: true,
		},
		{
			name: "BordersOnStart_WithDifferentEndDateAndStartDate",
			p:    Period{endDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), boundaryType: IncludeStartExcludeEnd},
			args: args{other: Period{startDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd}},
			want: false,
		},
		{
			name: "BordersOnStart_WithSameEndDateAndStartDateButDifferentBoundaryType",
			p:    Period{endDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), boundaryType: IncludeAll},
			args: args{other: Period{startDate: time.Date(
				2023, 1, 1, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeAll}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				assert.Equal(t, tt.want, tt.p.BordersOnStart(tt.args.other))
			},
		)
	}
}

func TestBordersOnEnd(t *testing.T) {
	type args struct {
		other Period
	}
	tests := []struct {
		name string
		p    Period
		args args
		want bool
	}{
		{
			name: "BordersOnEnd_WithSameStartDateAndEndDate",
			p: Period{startDate: time.Date(
				2023, 1, 1, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			args: args{other: Period{endDate: time.Date(
				2023, 1, 1, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd}},
			want: true,
		},
		{
			name: "BordersOnEnd_WithDifferentStartDateAndEndDate",
			p: Period{startDate: time.Date(
				2023, 1, 1, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			args: args{other: Period{endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd}},
			want: false,
		},
		{
			name: "BordersOnEnd_WithSameStartDateAndEndDateButDifferentBoundaryType",
			p:    Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), boundaryType: IncludeAll},
			args: args{other: Period{endDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), boundaryType: IncludeAll}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				assert.Equal(t, tt.want, tt.p.BordersOnEnd(tt.args.other))
			},
		)
	}
}

func TestAbuts(t *testing.T) {
	type args struct {
		other Period
	}
	tests := []struct {
		name string
		p    Period
		args args
		want bool
	}{
		{
			name: "Abuts_WithAdjacentPeriods",
			p: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			args: args{other: Period{startDate: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 3, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd}},
			want: true,
		},
		{
			name: "Abuts_WithNonAdjacentPeriods",
			p: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			args: args{other: Period{startDate: time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 4, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd}},
			want: false,
		},
		{
			name: "Abuts_WithOverlappingPeriods",
			p: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 3, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			args: args{other: Period{startDate: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 4, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				assert.Equal(t, tt.want, tt.p.Abuts(tt.args.other))
			},
		)
	}
}

func TestOverlaps(t *testing.T) {
	type args struct {
		other Period
	}
	tests := []struct {
		name string
		p    Period
		args args
		want bool
	}{
		{
			name: "Overlaps_WithOverlappingPeriods",
			p: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 3, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			args: args{other: Period{startDate: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 4, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd}},
			want: true,
		},
		{
			name: "Overlaps_WithNonOverlappingPeriods",
			p: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			args: args{other: Period{startDate: time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 4, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd}},
			want: false,
		},
		{
			name: "Overlaps_WithAdjacentPeriods",
			p: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			args: args{other: Period{startDate: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 3, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				assert.Equal(t, tt.want, tt.p.Overlaps(tt.args.other))
			},
		)
	}
}

func TestIsZero(t *testing.T) {
	tests := []struct {
		name string
		p    Period
		want bool
	}{
		{
			name: "IsZero_WithZeroPeriod",
			p:    Period{startDate: time.Time{}, endDate: time.Time{}, boundaryType: IncludeStartExcludeEnd},
			want: true,
		},
		{
			name: "IsZero_WithNonZeroPeriod",
			p: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				assert.Equal(t, tt.want, tt.p.IsZero())
			},
		)
	}
}

func TestEquals(t *testing.T) {
	tests := []struct {
		name  string
		p     Period
		other Period
		want  bool
	}{
		{
			name: "Equals_WithEqualPeriods",
			p: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			other: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			want: true,
		},
		{
			name: "Equals_WithDifferentStartDates",
			p: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			other: Period{startDate: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			want: false,
		},
		{
			name: "Equals_WithDifferentEndDates",
			p: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			other: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 3, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			want: false,
		},
		{
			name: "Equals_WithDifferentBoundaryTypes",
			p: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			other: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeAll},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				assert.Equal(t, tt.want, tt.p.Equals(tt.other))
			},
		)
	}
}

func TestStartingOn(t *testing.T) {
	type args struct {
		startDate time.Time
	}
	tests := []struct {
		name string
		p    Period
		args args
		want Period
	}{
		{
			name: "StartingOn_WithSameStartDate",
			p: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			args: args{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)},
			want: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
		},
		{
			name: "StartingOn_WithDifferentStartDate",
			p: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			args: args{startDate: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local)},
			want: Period{startDate: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				assert.Equal(t, tt.want, tt.p.StartingOn(tt.args.startDate))
			},
		)
	}
}

func TestIsEndedBy(t *testing.T) {
	tests := []struct {
		name  string
		p     Period
		other Period
		want  bool
	}{
		{
			name:  "IsEndedBy_WithSameEndDateAndBoundaryType",
			p:     Period{endDate: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local), boundaryType: IncludeStartExcludeEnd},
			other: Period{endDate: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local), boundaryType: IncludeStartExcludeEnd},
			want:  true,
		},
		{
			name:  "IsEndedBy_WithDifferentEndDate",
			p:     Period{endDate: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local), boundaryType: IncludeStartExcludeEnd},
			other: Period{endDate: time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local), boundaryType: IncludeStartExcludeEnd},
			want:  false,
		},
		{
			name:  "IsEndedBy_WithDifferentBoundaryType",
			p:     Period{endDate: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local), boundaryType: IncludeStartExcludeEnd},
			other: Period{endDate: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local), boundaryType: IncludeAll},
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				assert.Equal(t, tt.want, tt.p.IsEndedBy(tt.other))
			},
		)
	}
}

func TestIsStartedBy(t *testing.T) {
	tests := []struct {
		name  string
		p     Period
		other Period
		want  bool
	}{
		{
			name: "IsStartedBy_WithSameStartDateAndBoundaryType",
			p: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			other: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			want: true,
		},
		{
			name: "IsStartedBy_WithDifferentStartDate",
			p: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			other: Period{startDate: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 3, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			want: false,
		},
		{
			name: "IsStartedBy_WithDifferentBoundaryType",
			p: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			other: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeAll},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				assert.Equal(t, tt.want, tt.p.IsStartedBy(tt.other))
			},
		)
	}
}

func TestEndingOn(t *testing.T) {
	tests := []struct {
		name    string
		p       Period
		endDate time.Time
		want    Period
	}{
		{
			name: "EndingOn_WithSameEndDate",
			p: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			endDate: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
			want: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
		},
		{
			name: "EndingOn_WithDifferentEndDate",
			p: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			endDate: time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
			want: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 3, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				assert.Equal(t, tt.want, tt.p.EndingOn(tt.endDate))
			},
		)
	}
}

func TestBoundedBy(t *testing.T) {
	tests := []struct {
		name         string
		p            Period
		boundaryType string
		want         Period
	}{
		{
			name: "BoundedBy_WithSameBoundaryType",
			p: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			boundaryType: IncludeStartExcludeEnd,
			want: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
		},
		{
			name: "BoundedBy_WithDifferentBoundaryType",
			p: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			boundaryType: IncludeAll,
			want: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeAll},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				assert.Equal(t, tt.want, tt.p.BoundedBy(tt.boundaryType))
			},
		)
	}
}

func TestIsEndIncluded(t *testing.T) {
	tests := []struct {
		name string
		p    Period
		want bool
	}{
		{
			name: "IsEndIncluded_WithIncludeEndBoundaryType",
			p:    Period{boundaryType: IncludeStartExcludeEnd},
			want: false,
		},
		{
			name: "IsEndIncluded_WithExcludeEndBoundaryType",
			p:    Period{boundaryType: IncludeAll},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				assert.Equal(t, tt.want, tt.p.IsEndIncluded())
			},
		)
	}
}

func TestIsStartIncluded(t *testing.T) {
	tests := []struct {
		name string
		p    Period
		want bool
	}{
		{
			name: "IsStartIncluded_WithIncludeStartBoundaryType",
			p:    Period{boundaryType: IncludeStartExcludeEnd},
			want: true,
		},
		{
			name: "IsStartIncluded_WithExcludeStartBoundaryType",
			p:    Period{boundaryType: ExcludeStartIncludeEnd},
			want: false,
		},
		{
			name: "IsStartIncluded_WithExcludeAllBoundaryType",
			p:    Period{boundaryType: ExcludeAll},
			want: false,
		},
		{
			name: "IsStartIncluded_WithIncludeAllBoundaryType",
			p:    Period{boundaryType: IncludeAll},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				assert.Equal(t, tt.want, tt.p.IsStartIncluded())
			},
		)
	}
}

func TestIsEndExcluded(t *testing.T) {
	tests := []struct {
		name string
		p    Period
		want bool
	}{
		{
			name: "IsEndExcluded_WithExcludeEndBoundaryType",
			p:    Period{boundaryType: IncludeStartExcludeEnd},
			want: true,
		},
		{
			name: "IsEndExcluded_WithIncludeEndBoundaryType",
			p:    Period{boundaryType: IncludeAll},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				assert.Equal(t, tt.want, tt.p.IsEndExcluded())
			},
		)
	}
}

func TestIsStartExcluded(t *testing.T) {
	tests := []struct {
		name string
		p    Period
		want bool
	}{
		{
			name: "IsStartExcluded_WithExcludeStartBoundaryType",
			p:    Period{boundaryType: ExcludeStartIncludeEnd},
			want: true,
		},
		{
			name: "IsStartExcluded_WithIncludeStartBoundaryType",
			p:    Period{boundaryType: IncludeAll},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				assert.Equal(t, tt.want, tt.p.IsStartExcluded())
			},
		)
	}
}

func TestFormat(t *testing.T) {
	tests := []struct {
		name   string
		p      Period
		format string
		want   string
	}{
		{
			name: "Format_WithYYYYMMDD",
			p: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			format: "2006-01-02",
			want:   "[2023-01-01,2023-01-02)",
		},
		{
			name: "Format_WithDDMMYYYY",
			p: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeAll},
			format: "02-01-2006",
			want:   "[01-01-2023,02-01-2023]",
		},
		{
			name: "Format_WithMMDDYYYY",
			p: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: ExcludeAll},
			format: "01-02-2006",
			want:   "(01-01-2023,01-02-2023)",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				assert.Equal(t, tt.want, tt.p.Format(tt.format))
			},
		)
	}
}

func TestGetTimestampInterval(t *testing.T) {
	tests := []struct {
		name string
		p    Period
		want int64
	}{
		{
			name: "GetTimestampInterval_WithOneDayPeriod",
			p: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			want: int64(24 * 60 * 60),
		},
		{
			name: "GetTimestampInterval_WithTwoDaysPeriod",
			p: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 3, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			want: int64(48 * 60 * 60),
		},
		{
			name: "GetTimestampInterval_WithZeroPeriod",
			p:    Period{startDate: time.Time{}, endDate: time.Time{}, boundaryType: IncludeStartExcludeEnd},
			want: int64(0),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				assert.Equal(t, tt.want, tt.p.GetTimestampInterval())
			},
		)
	}
}

func TestTimeDuration(t *testing.T) {
	tests := []struct {
		name string
		p    Period
		want int64
	}{
		{
			name: "TimeDuration_WithOneDayPeriod",
			p: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			want: int64(24 * 60 * 60),
		},
		{
			name: "TimeDuration_WithTwoDaysPeriod",
			p: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 3, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			want: int64(48 * 60 * 60),
		},
		{
			name: "TimeDuration_WithZeroPeriod",
			p:    Period{startDate: time.Time{}, endDate: time.Time{}, boundaryType: IncludeStartExcludeEnd},
			want: int64(0),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				assert.Equal(t, tt.want, tt.p.timeDuration())
			},
		)
	}
}

func TestGetBoundaryType(t *testing.T) {
	tests := []struct {
		name string
		p    Period
		want string
	}{
		{
			name: "GetBoundaryType_WithIncludeStartExcludeEnd",
			p:    Period{boundaryType: IncludeStartExcludeEnd},
			want: IncludeStartExcludeEnd,
		},
		{
			name: "GetBoundaryType_WithExcludeStartIncludeEnd",
			p:    Period{boundaryType: ExcludeStartIncludeEnd},
			want: ExcludeStartIncludeEnd,
		},
		{
			name: "GetBoundaryType_WithExcludeAll",
			p:    Period{boundaryType: ExcludeAll},
			want: ExcludeAll,
		},
		{
			name: "GetBoundaryType_WithIncludeAll",
			p:    Period{boundaryType: IncludeAll},
			want: IncludeAll,
		},
		{
			name: "GetBoundaryType_WithDefaultBoundaryType",
			p:    Period{},
			want: IncludeStartExcludeEnd,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				assert.Equal(t, tt.want, tt.p.GetBoundaryType())
			},
		)
	}
}

func TestGetStartDate(t *testing.T) {
	tests := []struct {
		name string
		p    Period
		want time.Time
	}{
		{
			name: "GetStartDate_WithValidStartDate",
			p: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			want: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
		},
		{
			name: "GetStartDate_WithZeroStartDate",
			p: Period{startDate: time.Time{}, endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			want: time.Time{},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				assert.Equal(t, tt.want, tt.p.GetStartDate())
			},
		)
	}
}

func TestGetEndDate(t *testing.T) {
	tests := []struct {
		name string
		p    Period
		want time.Time
	}{
		{
			name: "GetEndDate_WithValidEndDate",
			p: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			want: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
		},
		{
			name: "GetEndDate_WithZeroEndDate",
			p: Period{startDate: time.Date(
				2023, 1, 1, 0, 0, 0, 0, time.Local,
			), endDate: time.Time{}, boundaryType: IncludeStartExcludeEnd},
			want: time.Time{},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				assert.Equal(t, tt.want, tt.p.GetEndDate())
			},
		)
	}
}

func TestGap(t *testing.T) {
	tests := []struct {
		name  string
		p     Period
		other Period
		want  Period
	}{
		{
			name: "Gap_WithOverlappingPeriods",
			p: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 3, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			other: Period{startDate: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 4, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			want: Period{boundaryType: IncludeStartExcludeEnd},
		},
		{
			name: "Gap_WithNonOverlappingPeriods",
			p: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			other: Period{startDate: time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 4, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			want: Period{startDate: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 3, 0, 0, 0, 0, time.Local,
			), boundaryType: "[)"},
		},
		{
			name: "Gap_WithAdjacentPeriods",
			p: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			other: Period{startDate: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 3, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			want: Period{startDate: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: "[)"},
		},
		{
			name: "Gap_WithAdjacentIncludeAllPeriods",
			p: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeAll},
			other: Period{startDate: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 3, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			want: Period{boundaryType: "[)"},
		},
		{
			name: "Gap_WithAdjacentIncludeAllPeriodsAndDiffBoundaryType",
			p: Period{startDate: time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 5, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeAll},
			other: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: ExcludeStartIncludeEnd},
			want: Period{startDate: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 3, 0, 0, 0, 0, time.Local,
			), boundaryType: "[]"},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.p.Gap(tt.other)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestWithBoundaryType(t *testing.T) {
	tests := []struct {
		name         string
		p            Period
		boundaryType string
		want         Period
	}{
		{
			name: "WithBoundaryType_SameBoundaryType",
			p: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 12, 31, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			boundaryType: IncludeStartExcludeEnd,
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 12, 31, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name: "WithBoundaryType_DifferentBoundaryType",
			p: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 12, 31, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			boundaryType: IncludeAll,
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 12, 31, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeAll,
			},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.p.WithBoundaryType(tt.boundaryType)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestWithDurationAfterStart(t *testing.T) {
	tests := []struct {
		name     string
		p        Period
		duration time.Duration
		want     Period
	}{
		{
			name: "WithDurationAfterStart_OneHour",
			p: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 1, 1, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			duration: time.Hour,
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 1, 1, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name: "WithDurationAfterStart_ZeroDuration",
			p: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 1, 1, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			duration: 0,
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.p.WithDurationAfterStart(tt.duration)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestWithDurationBeforeEnd(t *testing.T) {
	tests := []struct {
		name     string
		p        Period
		duration time.Duration
		want     Period
	}{
		{
			name: "WithDurationBeforeEnd_OneHour",
			p: Period{
				startDate:    time.Date(2023, 1, 1, 10, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			duration: time.Hour,
			want: Period{
				startDate:    time.Date(2023, 1, 1, 23, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name: "WithDurationBeforeEnd_ZeroDuration",
			p: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			duration: 0,
			want: Period{
				startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.p.WithDurationBeforeEnd(tt.duration)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestMoveStartDate(t *testing.T) {
	tests := []struct {
		name     string
		p        Period
		duration time.Duration
		want     Period
	}{
		{
			name: "MoveStartDate_PositiveDuration",
			p: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			duration: time.Hour * 24,
			want: Period{
				startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name: "MoveStartDate_NegativeDuration",
			p: Period{
				startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			duration: -time.Hour * 24,
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name: "MoveStartDate_ZeroDuration",
			p: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			duration: 0,
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.p.MoveStartDate(tt.duration)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestMoveEndDate(t *testing.T) {
	tests := []struct {
		name     string
		p        Period
		duration time.Duration
		want     Period
	}{
		{
			name: "MoveEndDate_PositiveDuration",
			p: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			duration: time.Hour * 24,
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name: "MoveEndDate_NegativeDuration",
			p: Period{
				startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			duration: -time.Hour * 24,
			want: Period{
				startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name: "MoveEndDate_ZeroDuration",
			p: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			duration: 0,
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.p.MoveEndDate(tt.duration)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestMove(t *testing.T) {
	tests := []struct {
		name     string
		p        Period
		duration time.Duration
		want     Period
	}{
		{
			name: "Move_PositiveDuration",
			p: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			duration: time.Hour * 24,
			want: Period{
				startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name: "Move_NegativeDuration",
			p: Period{
				startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			duration: -time.Hour * 24,
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name: "Move_ZeroDuration",
			p: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			duration: 0,
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.p.Move(tt.duration)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestExpand(t *testing.T) {
	tests := []struct {
		name     string
		p        Period
		duration time.Duration
		want     Period
	}{
		{
			name: "Expand_PositiveDuration",
			p: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			duration: time.Hour * 24,
			want: Period{
				startDate:    time.Date(2022, 12, 31, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name: "Expand_NegativeDuration",
			p: Period{
				startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			duration: -time.Hour * 24,
			want: Period{
				startDate:    time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name: "Expand_ZeroDuration",
			p: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			duration: 0,
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.p.Expand(tt.duration)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestIsBefore(t *testing.T) {
	tests := []struct {
		name  string
		p     Period
		other Period
		want  bool
	}{
		{
			name: "IsBefore_WithEndDateBeforeOtherStartDate",
			p: Period{
				endDate:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			other: Period{
				startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			want: true,
		},
		{
			name: "IsBefore_WithEndDateEqualToOtherStartDateAndDifferentBoundaryType",
			p: Period{
				endDate:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			other: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeAll,
			},
			want: true,
		},
		{
			name: "IsBefore_WithEndDateEqualToOtherStartDateAndSameBoundaryType",
			p: Period{
				endDate:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			other: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			want: true,
		},
		{
			name: "IsBefore_WithEndDateAfterOtherStartDate",
			p: Period{
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			other: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.p.IsBefore(tt.other)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestIsAfter(t *testing.T) {
	tests := []struct {
		name  string
		p     Period
		other Period
		want  bool
	}{
		{
			name: "IsAfter_WithEndDateAfterOtherStartDate",
			p: Period{
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			other: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			want: true,
		},
		{
			name: "IsAfter_WithEndDateEqualToOtherStartDateAndDifferentBoundaryType",
			p: Period{
				endDate:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			other: Period{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, 1, 7, 0, 0, 0, 0, time.Local),

				boundaryType: IncludeAll,
			},
			want: false,
		},
		{
			name: "IsAfter_WithEndDateEqualToOtherStartDateAndSameBoundaryType",
			p: Period{
				endDate:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			other: Period{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, 1, 7, 0, 0, 0, 0, time.Local),

				boundaryType: IncludeStartExcludeEnd,
			},
			want: false,
		},
		{
			name: "IsAfter_WithEndDateBeforeOtherStartDate",
			p: Period{
				endDate:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			other: Period{
				startDate: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, 1, 7, 0, 0, 0, 0, time.Local),

				boundaryType: IncludeStartExcludeEnd,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				assert.Equal(t, tt.want, tt.p.IsAfter(tt.other))
			},
		)
	}
}

func TestDateInterval(t *testing.T) {
	tests := []struct {
		name     string
		period   Period
		expected time.Duration
	}{
		{
			name: "DateInterval_WithOneDayDifference",
			period: Period{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
			},
			expected: 24 * time.Hour,
		},
		{
			name: "DateInterval_WithSameStartAndEndDate",
			period: Period{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
			},
			expected: 0,
		},
		{
			name: "DateInterval_WithEndDateBeforeStartDate",
			period: Period{
				startDate: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
			},
			expected: -24 * time.Hour,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.period.dateInterval()
				assert.Equal(t, tt.expected, got)
			},
		)
	}
}

func TestDurationCompare(t *testing.T) {
	tests := []struct {
		name     string
		period   Period
		other    Period
		expected int
	}{
		{
			name: "DurationCompare_WithEqualDurations",
			period: Period{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
			},
			other: Period{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
			},
			expected: 0,
		},
		{
			name: "DurationCompare_WithPeriodDurationLessThanOther",
			period: Period{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
			},
			other: Period{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
			},
			expected: -1,
		},
		{
			name: "DurationCompare_WithPeriodDurationGreaterThanOther",
			period: Period{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
			},
			other: Period{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
			},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.period.DurationCompare(tt.other)
				assert.Equal(t, tt.expected, got)
			},
		)
	}
}

func TestDurationEquals(t *testing.T) {
	tests := []struct {
		name     string
		period   Period
		other    Period
		expected bool
	}{
		{
			name: "DurationEquals_WithEqualDurations",
			period: Period{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
			},
			other: Period{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
			},
			expected: true,
		},
		{
			name: "DurationEquals_WithDifferentDurations",
			period: Period{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
			},
			other: Period{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.period.DurationEquals(tt.other)
				assert.Equal(t, tt.expected, got)
			},
		)
	}
}

func TestDurationGreaterThan(t *testing.T) {
	tests := []struct {
		name     string
		period   Period
		other    Period
		expected bool
	}{
		{
			name: "DurationGreaterThan_WithGreaterDuration",
			period: Period{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
			},
			other: Period{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
			},
			expected: true,
		},
		{
			name: "DurationGreaterThan_WithEqualDuration",
			period: Period{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
			},
			other: Period{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
			},
			expected: false,
		},
		{
			name: "DurationGreaterThan_WithLesserDuration",
			period: Period{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
			},
			other: Period{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.period.DurationGreaterThan(tt.other)
				assert.Equal(t, tt.expected, got)
			},
		)
	}
}

func TestDurationLessThan(t *testing.T) {
	tests := []struct {
		name     string
		period   Period
		other    Period
		expected bool
	}{
		{
			name: "DurationLessThan_WithLesserDuration",
			period: Period{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
			},
			other: Period{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
			},
			expected: true,
		},
		{
			name: "DurationLessThan_WithEqualDuration",
			period: Period{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
			},
			other: Period{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
			},
			expected: false,
		},
		{
			name: "DurationLessThan_WithGreaterDuration",
			period: Period{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
			},
			other: Period{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.period.DurationLessThan(tt.other)
				assert.Equal(t, tt.expected, got)
			},
		)
	}
}

func TestAfter(t *testing.T) {
	tests := []struct {
		name         string
		startDate    time.Time
		duration     time.Duration
		boundaryType string
		want         Period
	}{
		{
			name:         "After_WithPositiveDuration",
			startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
			duration:     time.Hour * 24,
			boundaryType: IncludeStartExcludeEnd,
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name:         "After_WithZeroDuration",
			startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
			duration:     0,
			boundaryType: IncludeStartExcludeEnd,
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name:         "After_WithNegativeDuration",
			startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
			duration:     -time.Hour * 24,
			boundaryType: IncludeStartExcludeEnd,
			want: Period{
				startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				p := Period{}
				got := p.After(tt.startDate, tt.duration, tt.boundaryType)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestBefore(t *testing.T) {
	tests := []struct {
		name         string
		endDate      time.Time
		duration     time.Duration
		boundaryType string
		want         Period
	}{
		{
			name:         "Before_WithPositiveDuration",
			endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
			duration:     time.Hour * 24,
			boundaryType: IncludeStartExcludeEnd,
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name:         "Before_WithZeroDuration",
			endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
			duration:     0,
			boundaryType: IncludeStartExcludeEnd,
			want: Period{
				startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name:         "Before_WithNegativeDuration",
			endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
			duration:     -time.Hour * 24,
			boundaryType: IncludeStartExcludeEnd,
			want: Period{
				startDate:    time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				p := Period{}
				got := p.Before(tt.endDate, tt.duration, tt.boundaryType)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestAround(t *testing.T) {
	tests := []struct {
		name         string
		sameDate     time.Time
		duration     time.Duration
		boundaryType string
		want         Period
	}{
		{
			name:         "Around_PositiveDuration",
			sameDate:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
			duration:     time.Hour * 24,
			boundaryType: IncludeStartExcludeEnd,
			want: Period{
				startDate:    time.Date(2022, 12, 31, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name:         "Around_ZeroDuration",
			sameDate:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
			duration:     0,
			boundaryType: IncludeStartExcludeEnd,
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name:         "Around_NegativeDuration",
			sameDate:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
			duration:     -time.Hour * 24,
			boundaryType: IncludeStartExcludeEnd,
			want: Period{
				endDate:      time.Date(2022, 12, 31, 0, 0, 0, 0, time.Local),
				startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				p := Period{}
				got := p.Around(tt.sameDate, tt.duration, tt.boundaryType)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestFromPeriod(t *testing.T) {
	tests := []struct {
		name         string
		period       Period
		boundaryType string
		want         Period
	}{
		{
			name: "FromPeriod_WithValidPeriodAndBoundaryType",
			period: Period{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
			},
			boundaryType: IncludeStartExcludeEnd,
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name: "FromPeriod_WithValidPeriodAndDifferentBoundaryType",
			period: Period{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
			},
			boundaryType: IncludeAll,
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeAll,
			},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.period.FromPeriod(tt.period, tt.boundaryType)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestIntersect(t *testing.T) {
	tests := []struct {
		name     string
		period   Period
		other    Period
		expected Period
	}{
		{
			name: "Intersect_WithOverlappingPeriods",
			period: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			other: Period{
				startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 4, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			expected: Period{
				startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name: "Intersect_WithAdjacentPeriods",
			period: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			other: Period{
				startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			expected: Period{boundaryType: IncludeStartExcludeEnd},
		},
		{
			name: "Intersect_WithAdjacentSamePeriods",
			period: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			other: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			expected: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name: "Intersect_WithOverlappingPeriods_SameBoundaryType",
			period: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			other: Period{
				startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 4, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			expected: Period{
				startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name: "Intersect_WithOverlappingPeriods_DifferentBoundaryType",
			period: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			other: Period{
				startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 4, 0, 0, 0, 0, time.Local),
				boundaryType: ExcludeStartIncludeEnd,
			},
			expected: Period{
				startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: "()",
			},
		},
		{
			name: "Intersect_WithNonOverlappingPeriods",
			period: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			other: Period{
				startDate:    time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 4, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			expected: Period{boundaryType: IncludeStartExcludeEnd},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.period.Intersect(tt.other)
				assert.Equal(t, tt.expected, got)
			},
		)
	}
}

func TestUnion(t *testing.T) {
	tests := []struct {
		name     string
		period   Period
		others   []Period
		expected Sequence
	}{
		{
			name: "WithMultipleOverlappingPeriods",
			period: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			others: []Period{
				{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 4, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			},
			expected: Sequence{
				intervals: []Period{
					{
						startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
						endDate:      time.Date(2023, 1, 4, 0, 0, 0, 0, time.Local),
						boundaryType: IncludeStartExcludeEnd,
					},
				},
			},
		},
		{
			name: "WithNonOverlappingPeriods",
			period: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			others: []Period{
				{
					startDate:    time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 4, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			},
			expected: Sequence{
				intervals: []Period{
					{
						startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
						endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
						boundaryType: IncludeStartExcludeEnd,
					},
					{
						startDate:    time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
						endDate:      time.Date(2023, 1, 4, 0, 0, 0, 0, time.Local),
						boundaryType: IncludeStartExcludeEnd,
					},
				},
			},
		},
		{
			name: "WithAdjacentPeriods",
			period: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			others: []Period{
				{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			},
			expected: Sequence{
				intervals: []Period{
					{
						startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
						endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
						boundaryType: IncludeStartExcludeEnd,
					},
					{
						startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
						endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
						boundaryType: IncludeStartExcludeEnd,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.period.Union(tt.others...)
				assert.Equal(t, tt.expected, got)
			},
		)
	}
}

func TestDiff(t *testing.T) {
	tests := []struct {
		name     string
		period   Period
		other    Period
		expected []Period
	}{
		{
			name: "Diff_WithOverlappingPeriods",
			period: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 4, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			other: Period{
				startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 5, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			expected: []Period{
				{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				{
					startDate:    time.Date(2023, 1, 4, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 5, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			},
		},
		{
			name: "Diff_WithOverlappingPeriodsAndDiffBoundaryType",
			period: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 4, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeAll,
			},
			other: Period{
				startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 5, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			expected: []Period{
				{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				{
					startDate:    time.Date(2023, 1, 4, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 5, 0, 0, 0, 0, time.Local),
					boundaryType: ExcludeAll,
				},
			},
		},
		{
			name: "Diff_WithOverlappingSameStartDatePeriods",
			period: Period{
				startDate:    time.Date(2023, 3, 1, 10, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 6, 1, 13, 0, 0, 0, time.Local),
				boundaryType: IncludeAll,
			},
			other: Period{
				startDate:    time.Date(2023, 3, 1, 10, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 8, 1, 14, 0, 0, 0, time.Local),
				boundaryType: IncludeAll,
			},
			expected: []Period{
				{
					startDate:    time.Date(2023, 6, 1, 13, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 8, 1, 14, 0, 0, 0, time.Local),
					boundaryType: ExcludeStartIncludeEnd,
				},
			},
		},
		{
			name: "Diff_WithOverlappingSameEndDatePeriods",
			period: Period{
				startDate:    time.Date(2023, 3, 5, 10, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 8, 1, 14, 0, 0, 0, time.Local),
				boundaryType: IncludeAll,
			},
			other: Period{
				startDate:    time.Date(2023, 3, 1, 10, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 8, 1, 14, 0, 0, 0, time.Local),
				boundaryType: IncludeAll,
			},
			expected: []Period{
				{
					startDate:    time.Date(2023, 3, 1, 10, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 3, 5, 10, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			},
		},
		{
			name: "Diff_WithNonOverlappingPeriods",
			period: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			other: Period{
				startDate:    time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 4, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			expected: []Period{},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.period.Diff(tt.other)
				assert.Equal(t, tt.expected, got)
			},
		)
	}
}

func TestMerge(t *testing.T) {
	tests := []struct {
		name     string
		period   Period
		others   []Period
		expected Period
	}{
		{
			name: "Merge_WithOverlappingPeriods",
			period: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			others: []Period{
				{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 4, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			},
			expected: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 4, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name: "Merge_WithNonOverlappingPeriods",
			period: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			others: []Period{
				{
					startDate:    time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 4, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			},
			expected: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 4, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name: "Merge_WithAdjacentPeriods",
			period: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			others: []Period{
				{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			},
			expected: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.period.Merge(tt.others...)
				assert.Equal(t, tt.expected, got)
			},
		)
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name     string
		period   Period
		other    Period
		expected Sequence
	}{
		{
			name: "Subtract_WithNonOverlappingPeriods",
			period: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			other: Period{
				startDate:    time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 4, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			expected: Sequence{intervals: []Period{
				{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			}},
		},
		{
			name: "Subtract_WithOverlappingPeriods",
			period: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 4, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			other: Period{
				startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			expected: Sequence{intervals: []Period{
				{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				{
					startDate:    time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 4, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			}},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.period.Subtract(tt.other)
				assert.Equal(t, tt.expected, got)
			},
		)
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		name   string
		period Period
		other  Period
		want   bool
	}{
		{
			name: "Contains_WithContainedPeriod",
			period: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			other: Period{
				startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			want: true,
		},
		{
			name: "Contains_WithNonContainedPeriod",
			period: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			other: Period{
				startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.period.Contains(tt.other)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestContainsInterval(t *testing.T) {
	tests := []struct {
		name   string
		period Period
		other  Period
		want   bool
	}{
		{
			name: "ContainsInterval_WithContainedPeriod",
			period: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			other: Period{
				startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			want: true,
		},
		{
			name: "ContainsInterval_WithFullContainedPeriod",
			period: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 5, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			other: Period{
				startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			want: true,
		},
		{
			name: "ContainsInterval_WithNonContainedPeriod",
			period: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			other: Period{
				startDate:    time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 4, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			want: false,
		},
		{
			name: "ContainsInterval_WithSamePeriods",
			period: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			other: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			want: true,
		},
		{
			name: "ContainsInterval_WithSameStartDateAndSameBoundaryFullContainedPeriod",
			period: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 9, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			other: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.period.containsInterval(tt.other)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestContainsDatePoint(t *testing.T) {
	tests := []struct {
		name         string
		period       Period
		datePoint    time.Time
		boundaryType string
		want         bool
	}{
		{
			name: "ContainsDatePoint_WithExcludeAllBoundaryType",
			period: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: ExcludeAll,
			},
			datePoint:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
			boundaryType: ExcludeAll,
			want:         true,
		},
		{
			name: "ContainsDatePoint_WithIncludeAllBoundaryType",
			period: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeAll,
			},
			datePoint:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
			boundaryType: IncludeAll,
			want:         true,
		},
		{
			name: "ContainsDatePoint_WithExcludeStartIncludeEndBoundaryType",
			period: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: ExcludeStartIncludeEnd,
			},
			datePoint:    time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
			boundaryType: ExcludeStartIncludeEnd,
			want:         true,
		},
		{
			name: "ContainsDatePoint_WithIncludeStartExcludeEndBoundaryType",
			period: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			datePoint:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
			boundaryType: IncludeStartExcludeEnd,
			want:         true,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.period.containsDatePoint(tt.datePoint, tt.boundaryType)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestFromQuarter(t *testing.T) {
	tests := []struct {
		name         string
		year         int
		quarter      int
		boundaryType string
		want         Period
	}{
		{
			name:         "FromQuarter_FirstQuarter",
			year:         2023,
			quarter:      1,
			boundaryType: IncludeStartExcludeEnd,
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 4, 1, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name:         "FromQuarter_SecondQuarter",
			year:         2023,
			quarter:      2,
			boundaryType: IncludeStartExcludeEnd,
			want: Period{
				startDate:    time.Date(2023, 4, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 7, 1, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name:         "FromQuarter_ThirdQuarter",
			year:         2023,
			quarter:      3,
			boundaryType: IncludeStartExcludeEnd,
			want: Period{
				startDate:    time.Date(2023, 7, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 10, 1, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name:         "FromQuarter_FourthQuarter",
			year:         2023,
			quarter:      4,
			boundaryType: IncludeStartExcludeEnd,
			want: Period{
				startDate:    time.Date(2023, 10, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name:         "FromQuarter_InvalidQuarter",
			year:         2023,
			quarter:      5,
			boundaryType: IncludeStartExcludeEnd,
			want: Period{
				startDate:    time.Date(2023, 13, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 16, 1, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				p := Period{}
				got := p.FromQuarter(tt.year, tt.quarter, tt.boundaryType)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestFromSemester(t *testing.T) {
	tests := []struct {
		name         string
		year         int
		semester     int
		boundaryType string
		want         Period
	}{
		{
			name:         "FromSemester_FirstSemester",
			year:         2023,
			semester:     1,
			boundaryType: IncludeStartExcludeEnd,
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 7, 1, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name:         "FromSemester_SecondSemester",
			year:         2023,
			semester:     2,
			boundaryType: IncludeStartExcludeEnd,
			want: Period{
				startDate:    time.Date(2023, 7, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name:         "FromSemester_InvalidSemester",
			year:         -2023,
			semester:     3,
			boundaryType: IncludeStartExcludeEnd,
			want: Period{
				startDate:    time.Date(-2022, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(-2022, 7, 1, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				p := Period{}
				got := p.FromSemester(tt.year, tt.semester, tt.boundaryType)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestFromIsoYear(t *testing.T) {
	tests := []struct {
		name         string
		year         int
		boundaryType string
		want         Period
	}{
		{
			name:         "FromIsoYear_ValidYear",
			year:         2023,
			boundaryType: IncludeStartExcludeEnd,
			want: Period{
				startDate:    time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
				endDate:      time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name:         "FromIsoYear_NegativeYear",
			year:         -2023,
			boundaryType: IncludeStartExcludeEnd,
			want: Period{
				startDate:    time.Date(-2023, time.January, 1, 0, 0, 0, 0, time.UTC),
				endDate:      time.Date(-2022, time.January, 1, 0, 0, 0, 0, time.UTC),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name:         "FromIsoYear_ZeroYear",
			year:         0,
			boundaryType: IncludeStartExcludeEnd,
			want: Period{
				startDate:    time.Date(0, time.January, 1, 0, 0, 0, 0, time.UTC),
				endDate:      time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				p := Period{}
				got := p.FromIsoYear(tt.year, tt.boundaryType)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestFromYear(t *testing.T) {
	tests := []struct {
		name         string
		year         int
		boundaryType string
		want         Period
	}{
		{
			name:         "FromYear_ValidYear",
			year:         2023,
			boundaryType: IncludeStartExcludeEnd,
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name:         "FromYear_NegativeYear",
			year:         -2023,
			boundaryType: IncludeStartExcludeEnd,
			want: Period{
				startDate:    time.Date(-2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(-2022, 1, 1, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name:         "FromYear_ZeroYear",
			year:         0,
			boundaryType: IncludeStartExcludeEnd,
			want: Period{
				startDate:    time.Date(0, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(1, 1, 1, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				p := Period{}
				got := p.FromYear(tt.year, tt.boundaryType)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestFromMonth(t *testing.T) {
	tests := []struct {
		name         string
		year         int
		month        int
		boundaryType string
		want         Period
	}{
		{
			name:         "FromMonth_ValidMonth",
			year:         2023,
			month:        1,
			boundaryType: IncludeStartExcludeEnd,
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 2, 1, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name:         "FromMonth_InvalidMonth",
			year:         2023,
			month:        13,
			boundaryType: IncludeStartExcludeEnd,
			want: Period{
				startDate:    time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2024, 2, 1, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name:         "FromMonth_DifferentBoundaryType",
			year:         2023,
			month:        1,
			boundaryType: IncludeAll,
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 2, 1, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeAll,
			},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				p := Period{}
				got := p.FromMonth(tt.year, tt.month, tt.boundaryType)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestFromDay(t *testing.T) {
	tests := []struct {
		name         string
		year         int
		month        int
		day          int
		boundaryType string
		want         Period
	}{
		{
			name:         "FromDay_ValidInput",
			year:         2023,
			month:        1,
			day:          1,
			boundaryType: IncludeStartExcludeEnd,
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name:         "FromDay_DifferentBoundaryType",
			year:         2023,
			month:        1,
			day:          1,
			boundaryType: IncludeAll,
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeAll,
			},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				p := Period{}
				got := p.FromDay(tt.year, tt.month, tt.day, tt.boundaryType)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestFromMinute(t *testing.T) {
	tests := []struct {
		name         string
		year         int
		month        int
		day          int
		hour         int
		minute       int
		boundaryType string
		want         Period
	}{
		{
			name:         "FromMinute_ValidInput",
			year:         2023,
			month:        1,
			day:          1,
			hour:         0,
			minute:       0,
			boundaryType: IncludeStartExcludeEnd,
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 1, 0, 1, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name:         "FromMinute_DifferentBoundaryType",
			year:         2023,
			month:        1,
			day:          1,
			hour:         0,
			minute:       0,
			boundaryType: IncludeAll,
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 1, 0, 1, 0, 0, time.Local),
				boundaryType: IncludeAll,
			},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				p := Period{}
				got := p.FromMinute(tt.year, tt.month, tt.day, tt.hour, tt.minute, tt.boundaryType)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestFromSecond(t *testing.T) {
	tests := []struct {
		name         string
		year         int
		month        int
		day          int
		hour         int
		minute       int
		second       int
		boundaryType string
		want         Period
	}{
		{
			name:         "FromSecond_ValidInput",
			year:         2023,
			month:        1,
			day:          1,
			hour:         0,
			minute:       0,
			second:       0,
			boundaryType: IncludeStartExcludeEnd,
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 1, 0, 0, 1, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name:         "FromSecond_DifferentBoundaryType",
			year:         2023,
			month:        1,
			day:          1,
			hour:         0,
			minute:       0,
			second:       0,
			boundaryType: IncludeAll,
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 1, 0, 0, 1, 0, time.Local),
				boundaryType: IncludeAll,
			},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				p := Period{}
				got := p.FromSecond(tt.year, tt.month, tt.day, tt.hour, tt.minute, tt.second, tt.boundaryType)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestFromHour(t *testing.T) {
	tests := []struct {
		name         string
		year         int
		month        int
		day          int
		hour         int
		boundaryType string
		want         Period
	}{
		{
			name:         "FromHour_ValidInput",
			year:         2023,
			month:        1,
			day:          1,
			hour:         0,
			boundaryType: IncludeStartExcludeEnd,
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 1, 1, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name:         "FromHour_DifferentBoundaryType",
			year:         2023,
			month:        1,
			day:          1,
			hour:         0,
			boundaryType: IncludeAll,
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 1, 1, 0, 0, 0, time.Local),
				boundaryType: IncludeAll,
			},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				p := Period{}
				got := p.FromHour(tt.year, tt.month, tt.day, tt.hour, tt.boundaryType)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestGetTimeDuration(t *testing.T) {
	tests := []struct {
		name     string
		period   Period
		expected int64
	}{
		{
			name: "GetTimeDuration_WithOneDayPeriod",
			period: Period{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
			},
			expected: 24 * 60 * 60,
		},
		{
			name: "GetTimeDuration_WithTwoDaysPeriod",
			period: Period{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:   time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
			},
			expected: 48 * 60 * 60,
		},
		{
			name: "GetTimeDuration_WithZeroPeriod",
			period: Period{
				startDate: time.Time{},
				endDate:   time.Time{},
			},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.period.GetTimeDuration()
				assert.Equal(t, tt.expected, got)
			},
		)
	}
}
