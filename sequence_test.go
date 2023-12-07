package period

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewSequence(t *testing.T) {
	tests := []struct {
		name     string
		periods  []Period
		expected Sequence
	}{
		{
			name:     "NewSequence_WithNoPeriods",
			periods:  []Period{},
			expected: Sequence{intervals: []Period{}},
		},
		{
			name: "NewSequence_WithSinglePeriod",
			periods: []Period{
				{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
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
				},
			},
		},
		{
			name: "NewSequence_WithMultiplePeriods",
			periods: []Period{
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
				got := NewSequence(tt.periods...)
				assert.Equal(t, tt.expected, got)
			},
		)
	}
}

func TestSequenceGetInterval(t *testing.T) {
	tests := []struct {
		name     string
		sequence Sequence
		want     []Period
	}{
		{
			name:     "GetInterval_WithEmptySequence",
			sequence: NewSequence(),
			want:     []Period{},
		},
		{
			name: "GetInterval_WithNonEmptySequence",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			want: []Period{
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
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.sequence.GetInterval()
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestSequenceClear(t *testing.T) {
	tests := []struct {
		name     string
		sequence Sequence
		want     int
	}{
		{
			name:     "Clear_WithEmptySequence",
			sequence: NewSequence(),
			want:     0,
		},
		{
			name: "Clear_WithNonEmptySequence",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.sequence.Clear()
				assert.Equal(t, tt.want, len(got.intervals))
			},
		)
	}
}

func TestSequenceTotalTimeDuration(t *testing.T) {
	tests := []struct {
		name     string
		sequence Sequence
		want     int64
	}{
		{
			name:     "TotalTimeDuration_WithEmptySequence",
			sequence: NewSequence(),
			want:     0,
		},
		{
			name: "TotalTimeDuration_WithSinglePeriod",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			want: 86400, // 24 hours in seconds
		},
		{
			name: "TotalTimeDuration_WithMultiplePeriods",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			want: 172800, // 48 hours in seconds
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.sequence.totalTimeDuration()
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestSequenceGetTotalTimestampInterval(t *testing.T) {
	tests := []struct {
		name     string
		sequence Sequence
		want     int64
	}{
		{
			name:     "GetTotalTimestampInterval_WithEmptySequence",
			sequence: NewSequence(),
			want:     0,
		},
		{
			name: "GetTotalTimestampInterval_WithSinglePeriod",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			want: 86400, // 24 hours in seconds
		},
		{
			name: "GetTotalTimestampInterval_WithMultiplePeriods",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			want: 172800, // 48 hours in seconds
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.sequence.GetTotalTimestampInterval()
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestSequenceSortByStartDate(t *testing.T) {
	tests := []struct {
		name     string
		p1       Period
		p2       Period
		sequence Sequence
		want     int64
	}{
		{
			name: "SortByStartDate_WithEqualDates",
			p1: Period{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
			},
			p2: Period{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
			},
			sequence: NewSequence(),
			want:     0,
		},
		{
			name: "SortByStartDate_WithP1BeforeP2",
			p1: Period{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
			},
			p2: Period{
				startDate: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
			},
			sequence: NewSequence(),
			want:     -1,
		},
		{
			name: "SortByStartDate_WithP2BeforeP1",
			p1: Period{
				startDate: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
			},
			p2: Period{
				startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
			},
			sequence: NewSequence(),
			want:     1,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.sequence.sortByStartDate(tt.p1, tt.p2)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestSequenceSort(t *testing.T) {
	tests := []struct {
		name     string
		sequence Sequence
		callback func(Period, Period) bool
		want     Sequence
	}{
		{
			name:     "Sort_WithEmptySequence",
			sequence: NewSequence(),
			callback: func(p1, p2 Period) bool {
				return p1.startDate.Before(p2.startDate)
			},
			want: NewSequence(),
		},
		{
			name: "Sort_WithSinglePeriodInSequence",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			callback: func(p1, p2 Period) bool {
				return p1.startDate.Before(p2.startDate)
			},
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
		{
			name: "Sort_WithMultiplePeriodsInSequence",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			callback: func(p1, p2 Period) bool {
				return p1.startDate.Before(p2.startDate)
			},
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.sequence.Sort(tt.callback)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestSequenceSorted(t *testing.T) {
	tests := []struct {
		name     string
		sequence Sequence
		callback func(Period, Period) int64
		want     Sequence
	}{
		{
			name:     "Sorted_WithEmptySequence",
			sequence: NewSequence(),
			callback: func(p1, p2 Period) int64 {
				return p1.startDate.Sub(p2.startDate).Milliseconds()
			},
			want: NewSequence(),
		},
		{
			name: "Sorted_WithSinglePeriod",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			callback: func(p1, p2 Period) int64 {
				return p1.startDate.Sub(p2.startDate).Milliseconds()
			},
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
		{
			name: "Sorted_WithMultiplePeriods",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			callback: func(p1, p2 Period) int64 {
				return p1.startDate.Sub(p2.startDate).Milliseconds()
			},
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.sequence.Sorted(tt.callback)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestSequenceContains(t *testing.T) {
	tests := []struct {
		name     string
		sequence Sequence
		periods  []Period
		want     bool
	}{
		{
			name:     "Contains_WithEmptySequence",
			sequence: NewSequence(),
			periods:  []Period{},
			want:     false,
		},
		{
			name: "Contains_WithSinglePeriodInSequence",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			periods: []Period{
				{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			},
			want: true,
		},
		{
			name: "Contains_WithMultiplePeriodsInSequence",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			periods: []Period{
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
			want: true,
		},
		{
			name: "Contains_WithPeriodNotInSequence",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			periods: []Period{
				{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.sequence.Contains(tt.periods...)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestSequenceReduce(t *testing.T) {
	tests := []struct {
		name     string
		sequence Sequence
		callback func(Sequence, Period) Sequence
		initial  Sequence
		want     Sequence
	}{
		{
			name:     "Reduce_WithEmptySequence",
			sequence: NewSequence(),
			callback: func(s Sequence, p Period) Sequence {
				return s.Push(p)
			},
			initial: NewSequence(),
			want:    NewSequence(),
		},
		{
			name: "Reduce_WithSinglePeriod",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			callback: func(s Sequence, p Period) Sequence {
				return s.Push(p)
			},
			initial: NewSequence(),
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
		{
			name: "Reduce_WithMultiplePeriods",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			callback: func(s Sequence, p Period) Sequence {
				return s.Push(p)
			},
			initial: NewSequence(),
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.sequence.Reduce(tt.callback, tt.initial)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestSequenceSubtractOne(t *testing.T) {
	tests := []struct {
		name     string
		sequence Sequence
		period   Period
		want     Sequence
	}{
		{
			name:     "SubtractOne_WithEmptySequence",
			sequence: NewSequence(),
			period:   Period{},
			want:     NewSequence(),
		},
		{
			name: "SubtractOne_WithSinglePeriodInSequence",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			period: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			want: NewSequence(),
		},
		{
			name: "SubtractOne_WithMultiplePeriodsInSequence",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			period: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.sequence.subtractOne(tt.sequence, tt.period)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestSequenceSubtract(t *testing.T) {
	tests := []struct {
		name     string
		sequence Sequence
		subtract Sequence
		want     Sequence
	}{
		{
			name:     "Subtract_WithEmptySequence",
			sequence: NewSequence(),
			subtract: NewSequence(),
			want:     NewSequence(),
		},
		{
			name: "Subtract_WithSinglePeriodInSequence",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			subtract: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			want: NewSequence(),
		},
		{
			name: "SequenceSubtract_WithMultiplePeriodsInSequence",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			subtract: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
		{
			name: "SequenceSubtract_WithMultiplePeriodsEmptySequence",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			subtract: NewSequence(),
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.sequence.Subtract(tt.subtract)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestSequenceCalculateUnion(t *testing.T) {
	tests := []struct {
		name     string
		sequence Sequence
		period   Period
		want     Sequence
	}{
		{
			name:     "CalculateUnion_WithEmptySequence",
			sequence: NewSequence(),
			period:   Period{},
			want:     NewSequence(),
		},
		{
			name: "CalculateUnion_WithSinglePeriodInSequence",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			period: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
		{
			name: "CalculateUnion_WithMultiplePeriodsInSequence",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			period: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.sequence.calculateUnion(tt.sequence, tt.period)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestSequenceUnions(t *testing.T) {
	tests := []struct {
		name     string
		sequence Sequence
		want     Sequence
	}{
		{
			name:     "Unions_WithEmptySequence",
			sequence: NewSequence(),
			want:     NewSequence(),
		},
		{
			name: "Unions_WithSinglePeriodInSequence",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
		{
			name: "Unions_WithMultiplePeriodsInSequence",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.sequence.Unions()
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestSequenceEquals(t *testing.T) {
	tests := []struct {
		name     string
		sequence Sequence
		other    Sequence
		want     bool
	}{
		{
			name:     "Equals_WithEmptySequences",
			sequence: NewSequence(),
			other:    NewSequence(),
			want:     true,
		},
		{
			name: "Equals_WithIdenticalSinglePeriodSequences",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			other: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			want: true,
		},
		{
			name: "Equals_WithDifferentSinglePeriodSequences",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			other: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			want: false,
		},
		{
			name: "Equals_WithIdenticalMultiplePeriodSequences",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			other: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			want: true,
		},
		{
			name: "Equals_WithDifferentMultiplePeriodSequences",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			other: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 4, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.sequence.Equals(tt.other)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestSequenceGaps(t *testing.T) {
	tests := []struct {
		name     string
		sequence Sequence
		want     Sequence
	}{
		{
			name:     "Gaps_WithEmptySequence",
			sequence: NewSequence(),
			want:     NewSequence(),
		},
		{
			name: "Gaps_WithSinglePeriodInSequence",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			want: NewSequence(),
		},
		{
			name: "Gaps_WithMultiplePeriodsInSequence_NoGaps",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			want: NewSequence(),
		},
		{
			name: "Gaps_WithMultiplePeriodsInSequence_WithGaps",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 4, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.sequence.Gaps()
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestSequenceIsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		sequence Sequence
		want     bool
	}{
		{
			name:     "IsEmpty_WithEmptySequence",
			sequence: NewSequence(),
			want:     true,
		},
		{
			name: "IsEmpty_WithSinglePeriodInSequence",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			want: false,
		},
		{
			name: "IsEmpty_WithMultiplePeriodsInSequence",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.sequence.IsEmpty()
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestSequenceIndexOf(t *testing.T) {
	tests := []struct {
		name     string
		sequence Sequence
		period   Period
		want     int
	}{
		{
			name:     "IndexOf_WithEmptySequence",
			sequence: NewSequence(),
			period:   Period{},
			want:     -1,
		},
		{
			name: "IndexOf_WithSinglePeriodInSequence_PeriodExists",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			period: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			want: 0,
		},
		{
			name: "IndexOf_WithSinglePeriodInSequence_PeriodDoesNotExist",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			period: Period{
				startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			want: -1,
		},
		{
			name: "IndexOf_WithMultiplePeriodsInSequence_PeriodExists",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			period: Period{
				startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			want: 1,
		},
		{
			name: "IndexOf_WithMultiplePeriodsInSequence_PeriodDoesNotExist",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			period: Period{
				startDate:    time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 4, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.sequence.IndexOf(tt.period)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestSequenceFilterOffset(t *testing.T) {
	tests := []struct {
		name     string
		sequence Sequence
		offset   int
		want     int
	}{
		{
			name:     "FilterOffset_WithEmptySequence",
			sequence: NewSequence(),
			offset:   0,
			want:     -1,
		},
		{
			name: "FilterOffset_WithSinglePeriodInSequence_PositiveOffset",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			offset: 1,
			want:   -1,
		},
		{
			name: "FilterOffset_WithSinglePeriodInSequence_NegativeOffset",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			offset: -1,
			want:   0,
		},
		{
			name: "FilterOffset_WithMultiplePeriodsInSequence_PositiveOffset",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			offset: 2,
			want:   -1,
		},
		{
			name: "FilterOffset_WithMultiplePeriodsInSequence_PositiveOffsetLessThanSequenceLength",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			offset: -7,
			want:   -1,
		},
		{
			name: "FilterOffset_WithMultiplePeriodsInSequence_NegativeOffset",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			offset: -1,
			want:   1,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.sequence.filterOffset(tt.offset)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestSequenceOffsetGet(t *testing.T) {
	tests := []struct {
		name     string
		sequence Sequence
		offset   int
		want     Period
	}{
		{
			name:     "OffsetGet_WithEmptySequence",
			sequence: NewSequence(),
			offset:   0,
			want:     Period{},
		},
		{
			name: "OffsetGet_WithSinglePeriodInSequence_ValidOffset",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			offset: 0,
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name: "OffsetGet_WithSinglePeriodInSequence_InvalidOffset",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			offset: 1,
			want:   Period{},
		},
		{
			name: "OffsetGet_WithMultiplePeriodsInSequence_ValidOffset",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			offset: 1,
			want: Period{
				startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name: "OffsetGet_WithMultiplePeriodsInSequence_InvalidOffset",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			offset: 2,
			want:   Period{},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.sequence.OffsetGet(tt.offset)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestSequenceOffsetSet(t *testing.T) {
	tests := []struct {
		name     string
		sequence Sequence
		offset   int
		period   Period
		want     Sequence
	}{
		{
			name:     "OffsetSet_WithEmptySequence",
			sequence: NewSequence(),
			offset:   0,
			period: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
		{
			name: "OffsetSet_WithSinglePeriodInSequence_OffsetEqualLength",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			offset: 1,
			period: Period{
				startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
		{
			name: "OffsetSet_WithSinglePeriodInSequence_OffsetGreaterThanLength",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			offset: 2,
			period: Period{
				startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
		{
			name: "OffsetSet_WithSinglePeriodInSequence_OffsetLessThanLength",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			offset: 0,
			period: Period{
				startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.sequence.OffsetSet(tt.offset, tt.period)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestSequenceOffsetUnset(t *testing.T) {
	tests := []struct {
		name     string
		sequence Sequence
		offset   int
		want     Sequence
	}{
		{
			name:     "OffsetUnset_WithEmptySequence",
			sequence: NewSequence(),
			offset:   0,
			want:     NewSequence(),
		},
		{
			name: "OffsetUnset_WithSinglePeriodInSequence_ValidOffset",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			offset: 0,
			want:   NewSequence(),
		},
		{
			name: "OffsetUnset_WithSinglePeriodInSequence_InvalidOffset",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			offset: 1,
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
		{
			name: "OffsetUnset_WithMultiplePeriodsInSequence_ValidOffset",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			offset: 1,
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
		{
			name: "OffsetUnset_WithMultiplePeriodsInSequence_InvalidOffset",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			offset: 2,
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.sequence.OffsetUnset(tt.offset)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestSequenceOffsetExists(t *testing.T) {
	tests := []struct {
		name     string
		sequence Sequence
		offset   int
		want     bool
	}{
		{
			name:     "OffsetExists_WithEmptySequence",
			sequence: NewSequence(),
			offset:   0,
			want:     false,
		},
		{
			name: "OffsetExists_WithSinglePeriodInSequence_ValidOffset",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			offset: 0,
			want:   true,
		},
		{
			name: "OffsetExists_WithSinglePeriodInSequence_InvalidOffset",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			offset: 1,
			want:   false,
		},
		{
			name: "OffsetExists_WithMultiplePeriodsInSequence_ValidOffset",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			offset: 1,
			want:   true,
		},
		{
			name: "OffsetExists_WithMultiplePeriodsInSequence_InvalidOffset",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			offset: 2,
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.sequence.OffsetExists(tt.offset)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestSequenceCount(t *testing.T) {
	tests := []struct {
		name     string
		sequence Sequence
		want     int
	}{
		{
			name:     "Count_WithEmptySequence",
			sequence: NewSequence(),
			want:     0,
		},
		{
			name: "Count_WithSinglePeriodInSequence",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			want: 1,
		},
		{
			name: "Count_WithMultiplePeriodsInSequence",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.sequence.Count()
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestSequenceGet(t *testing.T) {
	tests := []struct {
		name     string
		sequence Sequence
		offset   int
		want     Period
	}{
		{
			name:     "Get_WithEmptySequence",
			sequence: NewSequence(),
			offset:   0,
			want:     Period{},
		},
		{
			name: "Get_WithSinglePeriodInSequence_ValidOffset",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			offset: 0,
			want: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name: "Get_WithSinglePeriodInSequence_InvalidOffset",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			offset: 1,
			want:   Period{},
		},
		{
			name: "Get_WithMultiplePeriodsInSequence_ValidOffset",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			offset: 1,
			want: Period{
				startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
		},
		{
			name: "Get_WithMultiplePeriodsInSequence_InvalidOffset",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			offset: 2,
			want:   Period{},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.sequence.Get(tt.offset)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestSequenceSet(t *testing.T) {
	tests := []struct {
		name     string
		sequence Sequence
		offset   int
		period   Period
		want     Sequence
	}{
		{
			name:     "Set_WithEmptySequenceAndZeroOffset",
			sequence: NewSequence(),
			offset:   0,
			period: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			want: NewSequence(
				Period{startDate: time.Date(
					2023, 1, 1, 0, 0, 0, 0, time.Local,
				), endDate: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local), boundaryType: IncludeStartExcludeEnd},
			),
		},
		{
			name:     "Set_WithEmptySequenceAndNonZeroOffset",
			sequence: NewSequence(),
			offset:   1,
			period: Period{startDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 2, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			want: NewSequence(
				Period{startDate: time.Date(
					2023, 1, 1, 0, 0, 0, 0, time.Local,
				), endDate: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local), boundaryType: IncludeStartExcludeEnd},
			),
		},
		{
			name: "Set_WithNonEmptySequenceAndValidOffset",
			sequence: NewSequence(
				Period{startDate: time.Date(
					2023, 1, 1, 0, 0, 0, 0, time.Local,
				), endDate: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local), boundaryType: IncludeStartExcludeEnd},
			),
			offset: 0,
			period: Period{startDate: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 3, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			want: NewSequence(
				Period{startDate: time.Date(
					2023, 1, 2, 0, 0, 0, 0, time.Local,
				), endDate: time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local), boundaryType: IncludeStartExcludeEnd},
			),
		},
		{
			name: "Set_WithNonEmptySequenceAndInvalidOffset",
			sequence: NewSequence(
				Period{startDate: time.Date(
					2023, 1, 1, 0, 0, 0, 0, time.Local,
				), endDate: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local), boundaryType: IncludeStartExcludeEnd},
			),
			offset: 1,
			period: Period{startDate: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local), endDate: time.Date(
				2023, 1, 3, 0, 0, 0, 0, time.Local,
			), boundaryType: IncludeStartExcludeEnd},
			want: NewSequence(
				Period{startDate: time.Date(
					2023, 1, 1, 0, 0, 0, 0, time.Local,
				), endDate: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local), boundaryType: IncludeStartExcludeEnd},
				Period{startDate: time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local), endDate: time.Date(
					2023, 1, 3, 0, 0, 0, 0, time.Local,
				), boundaryType: IncludeStartExcludeEnd},
			),
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.sequence.Set(tt.offset, tt.period)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestSequenceUnshift(t *testing.T) {
	tests := []struct {
		name     string
		sequence Sequence
		periods  []Period
		want     Sequence
	}{
		{
			name:     "Unshift_WithEmptySequence",
			sequence: NewSequence(),
			periods: []Period{
				{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			},
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
		{
			name: "Unshift_WithNonEmptySequence",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			periods: []Period{
				{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			},
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.sequence.Unshift(tt.periods...)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestSequencePush(t *testing.T) {
	tests := []struct {
		name     string
		sequence Sequence
		periods  []Period
		want     Sequence
	}{
		{
			name:     "Push_WithEmptySequence",
			sequence: NewSequence(),
			periods: []Period{
				{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			},
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
		{
			name: "Push_WithNonEmptySequence",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			periods: []Period{
				{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			},
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.sequence.Push(tt.periods...)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestSequenceInsert(t *testing.T) {
	tests := []struct {
		name     string
		sequence Sequence
		offset   int
		period   Period
		periods  []Period
		want     Sequence
	}{
		{
			name:     "Insert_WithEmptySequenceAndZeroOffset",
			sequence: NewSequence(),
			offset:   0,
			period: Period{
				startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
		{
			name: "Insert_WithNonEmptySequenceAndValidOffset",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			offset: 0,
			period: Period{
				startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
		{
			name: "Insert_WithNonEmptySequenceAndInvalidOffset",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			offset: 1,
			period: Period{
				startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
		{
			name: "Insert_WithNonEmptySequenceAndNotInvalidOffset",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			offset: 2,
			period: Period{
				startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
		{
			name: "Insert_WithNonEmptySequenceAndNotInvalidOffsetAndMultiplePeriods",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 4, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			offset: 1,
			period: Period{
				startDate:    time.Date(2023, 1, 5, 0, 0, 0, 0, time.Local),
				endDate:      time.Date(2023, 1, 6, 0, 0, 0, 0, time.Local),
				boundaryType: IncludeStartExcludeEnd,
			},
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 5, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 6, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 4, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.sequence.Insert(tt.offset, tt.period, tt.periods...)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestSequenceRemove(t *testing.T) {
	tests := []struct {
		name     string
		sequence Sequence
		offset   int
		want     Sequence
	}{
		{
			name:     "Remove_WithEmptySequence",
			sequence: NewSequence(),
			offset:   0,
			want:     NewSequence(),
		},
		{
			name: "Remove_WithSinglePeriodInSequence_ValidOffset",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			offset: 0,
			want:   NewSequence(),
		},
		{
			name: "Remove_WithSinglePeriodInSequence_InvalidOffset",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			offset: 1,
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
		{
			name: "Remove_WithMultiplePeriodsInSequence_ValidOffset",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			offset: 1,
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
		{
			name: "Remove_WithMultiplePeriodsInSequence_InvalidOffset",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			offset: 2,
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.sequence.Remove(tt.offset)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestSequenceFilter(t *testing.T) {
	tests := []struct {
		name     string
		sequence Sequence
		callback func(Period) bool
		want     Sequence
	}{
		{
			name:     "Filter_WithEmptySequence",
			sequence: NewSequence(),
			callback: func(p Period) bool {
				return p.startDate.Year() == 2023
			},
			want: NewSequence(),
		},
		{
			name: "Filter_WithSinglePeriodInSequence_MatchingCallback",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			callback: func(p Period) bool {
				return p.startDate.Year() == 2023
			},
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
		{
			name: "Filter_WithSinglePeriodInSequence_NonMatchingCallback",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			callback: func(p Period) bool {
				return p.startDate.Year() == 2022
			},
			want: NewSequence(),
		},
		{
			name: "Filter_WithMultiplePeriodsInSequence_SomeMatchingCallback",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2022, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			callback: func(p Period) bool {
				return p.startDate.Year() == 2023
			},
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.sequence.Filter(tt.callback)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestSequenceMap(t *testing.T) {
	tests := []struct {
		name      string
		sequence  Sequence
		transform func(Period) Period
		want      Sequence
	}{
		{
			name:     "Map_WithEmptySequence",
			sequence: NewSequence(),
			transform: func(p Period) Period {
				return Period{
					startDate:    p.startDate.Add(time.Hour * 24),
					endDate:      p.endDate.Add(time.Hour * 24),
					boundaryType: p.boundaryType,
				}
			},
			want: NewSequence(),
		},
		{
			name: "Map_WithSinglePeriodInSequence",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			transform: func(p Period) Period {
				return Period{
					startDate:    p.startDate.Add(time.Hour * 24),
					endDate:      p.endDate.Add(time.Hour * 24),
					boundaryType: p.boundaryType,
				}
			},
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
		{
			name: "Map_WithMultiplePeriodsInSequence",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			transform: func(p Period) Period {
				return Period{
					startDate:    p.startDate.Add(time.Hour * 24),
					endDate:      p.endDate.Add(time.Hour * 24),
					boundaryType: p.boundaryType,
				}
			},
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 4, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.sequence.Map(tt.transform)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestSequenceEvery(t *testing.T) {
	tests := []struct {
		name     string
		sequence Sequence
		callback func(Period, int) bool
		want     bool
	}{
		{
			name:     "Every_WithEmptySequence",
			sequence: NewSequence(),
			callback: func(p Period, i int) bool {
				return p.startDate.Year() == 2023
			},
			want: true,
		},
		{
			name: "Every_WithSinglePeriodInSequence_MatchingCallback",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			callback: func(p Period, i int) bool {
				return p.startDate.Year() == 2023
			},
			want: true,
		},
		{
			name: "Every_WithSinglePeriodInSequence_NonMatchingCallback",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			callback: func(p Period, i int) bool {
				return p.startDate.Year() == 2022
			},
			want: false,
		},
		{
			name: "Every_WithMultiplePeriodsInSequence_AllMatchingCallback",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			callback: func(p Period, i int) bool {
				return p.startDate.Year() == 2023
			},
			want: true,
		},
		{
			name: "Every_WithMultiplePeriodsInSequence_SomeMatchingCallback",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2022, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			callback: func(p Period, i int) bool {
				return p.startDate.Year() == 2023
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.sequence.Every(tt.callback)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestSequenceSome(t *testing.T) {
	tests := []struct {
		name     string
		sequence Sequence
		callback func(Period, int) bool
		want     bool
	}{
		{
			name:     "Some_WithEmptySequence",
			sequence: NewSequence(),
			callback: func(p Period, i int) bool {
				return p.startDate.Year() == 2023
			},
			want: false,
		},
		{
			name: "Some_WithSinglePeriodInSequence_MatchingCallback",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			callback: func(p Period, i int) bool {
				return p.startDate.Year() == 2023
			},
			want: true,
		},
		{
			name: "Some_WithSinglePeriodInSequence_NonMatchingCallback",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			callback: func(p Period, i int) bool {
				return p.startDate.Year() == 2022
			},
			want: false,
		},
		{
			name: "Some_WithMultiplePeriodsInSequence_SomeMatchingCallback",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2022, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			callback: func(p Period, i int) bool {
				return p.startDate.Year() == 2023
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.sequence.Some(tt.callback)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestSequenceIntersections(t *testing.T) {
	tests := []struct {
		name     string
		sequence Sequence
		want     Sequence
	}{
		{
			name:     "Intersections_WithEmptySequence",
			sequence: NewSequence(),
			want:     NewSequence(),
		},
		{
			name: "Intersections_WithSinglePeriodInSequence_NoIntersection",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			want: NewSequence(),
		},
		{
			name: "Intersections_WithMultiplePeriodsInSequence_SomeIntersections",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 4, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
		{
			name: "Intersections_WithMultiplePeriodsInSequence_NoIntersections",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 4, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			want: NewSequence(),
		},
		{
			name: "Intersections_WithMultiplePeriodsInSequence_OD0",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 31, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 10, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 15, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 10, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 31, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 10, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 15, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},

				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 31, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
		{
			name: "Intersections_WithMultiplePeriodsInSequence_PeriodCross",
			sequence: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 31, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 2, 10, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 2, 20, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 3, 01, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 3, 31, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 1, 20, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 3, 10, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
			want: NewSequence(
				Period{
					startDate:    time.Date(2023, 1, 20, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 1, 31, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 2, 10, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 2, 20, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
				Period{
					startDate:    time.Date(2023, 3, 1, 0, 0, 0, 0, time.Local),
					endDate:      time.Date(2023, 3, 10, 0, 0, 0, 0, time.Local),
					boundaryType: IncludeStartExcludeEnd,
				},
			),
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := tt.sequence.Intersections()
				assert.Equal(t, tt.want, got)
			},
		)
	}
}
