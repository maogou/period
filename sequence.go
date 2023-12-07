package period

import (
	"sort"
)

type Sequence struct {
	intervals []Period
}

func NewSequence(periods ...Period) Sequence {
	return Sequence{
		intervals: periods,
	}
}

func (s Sequence) Sort(callback func(Period, Period) bool) Sequence {
	sort.Slice(
		s.intervals, func(i, j int) bool {
			return callback(s.intervals[i], s.intervals[j])
		},
	)

	return s
}

func (s Sequence) Sorted(callback func(Period, Period) int64) Sequence {
	periods := s.intervals

	sort.Slice(
		periods, func(i, j int) bool {
			return callback(periods[i], periods[j]) < 0
		},
	)

	return Sequence{intervals: periods}
}

func (s Sequence) Reduce(callback func(Sequence, Period) Sequence, initial Sequence) Sequence {
	for _, p := range s.intervals {
		initial = callback(initial, p)
	}

	return initial
}

func (s Sequence) subtractOne(sequence Sequence, interval Period) Sequence {
	if sequence.IsEmpty() {
		return sequence
	}

	reducer := func(sequence Sequence, period Period) Sequence {
		subtract := period.Subtract(interval)
		if !subtract.IsEmpty() {
			return sequence.Push(subtract.intervals...)
		}

		return sequence
	}

	return sequence.Reduce(reducer, Sequence{})
}

func (s Sequence) Subtract(sequence Sequence) Sequence {
	if s.IsEmpty() {
		return s
	}

	newSequence := sequence.Reduce(s.subtractOne, s)

	if newSequence.Equals(s) {
		return s
	}

	return newSequence
}

func (s Sequence) calculateUnion(sequence Sequence, period Period) Sequence {
	if sequence.IsEmpty() {
		if period.IsZero() {
			return sequence
		}
		return sequence.Push(period)
	}

	index := sequence.Count() - 1

	interval := sequence.Get(index)

	if interval.Overlaps(period) {
		return sequence.Set(index, interval.Merge(period))
	}
	return sequence.Push(period)
}

func (s Sequence) Set(offset int, period Period) Sequence {
	index := s.filterOffset(offset)

	if -1 == index && len(s.intervals) == 0 {
		return s.Push(period)
	}

	if -1 == index && len(s.intervals) > 0 {
		return s.Push(period)
	}

	s.intervals[offset] = period

	return s
}

func (s Sequence) Unshift(periods ...Period) Sequence {
	s.intervals = append(periods, s.intervals...)

	return s
}

func (s Sequence) OffsetSet(offset int, period Period) Sequence {

	if offset > len(s.intervals) {
		return s.Push(period)
	}

	return s.Set(offset, period)
}

func (s Sequence) OffsetUnset(offset int) Sequence {
	return s.Remove(offset)
}

func (s Sequence) Push(periods ...Period) Sequence {
	s.intervals = append(s.intervals, periods...)

	return s
}

func (s Sequence) Insert(offset int, period Period, periods ...Period) Sequence {
	if offset == 0 {
		return s.Unshift(append([]Period{period}, periods...)...)
	}

	if len(s.intervals) == offset {
		return s.Push(append([]Period{period}, periods...)...)
	}

	index := s.filterOffset(offset)

	if -1 == index {
		return s
	}

	periods = append([]Period{period}, periods...)

	s.intervals = append(s.intervals[:index], append(periods, s.intervals[index:]...)...)

	return s
}

func (s Sequence) Remove(offset int) Sequence {
	offset = s.filterOffset(offset)

	if -1 == offset {
		return s
	}

	if len(s.intervals) == 1 && 0 == offset {
		return Sequence{}
	}

	s.intervals = append(s.intervals[:offset], s.intervals[offset+1:]...)

	return s
}

func (s Sequence) Filter(callback func(Period) bool) Sequence {
	var intervals []Period

	for _, p := range s.intervals {
		if callback(p) {
			intervals = append(intervals, p)
		}
	}

	return Sequence{intervals: intervals}
}

func (s Sequence) Map(transform func(Period) Period) Sequence {
	var intervals []Period

	for _, p := range s.intervals {
		intervals = append(intervals, transform(p))
	}

	return Sequence{intervals: intervals}
}

func (s Sequence) Unions() Sequence {
	otherSequence := s.Sorted(s.sortByStartDate).Reduce(s.calculateUnion, Sequence{})

	if otherSequence.Equals(s) {
		return s
	}

	return otherSequence
}

func (s Sequence) Gaps() Sequence {
	sequence := Sequence{}
	interval := Period{}

	sortedSequence := s.Sorted(s.sortByStartDate)

	for _, period := range sortedSequence.intervals {

		if interval.IsZero() {
			interval = period
			continue
		}

		if !interval.Overlaps(period) && !interval.Abuts(period) {
			sequence = sequence.Push(interval.Gap(period))
		}

		if !interval.Contains(period) {
			interval = period
		}
	}

	return sequence
}

func (s Sequence) Intersections() Sequence {
	var (
		sequence              Sequence
		current               Period
		isPreviouslyContained bool
	)

	sortedSequence := s.Sorted(s.sortByStartDate)

	for _, period := range sortedSequence.intervals {
		if current.IsZero() {
			current = period
			continue
		}

		isContained := current.Contains(period)

		if isContained && isPreviouslyContained && !sequence.Contains(period) {
			sequence = sequence.Push(current.Intersect(current))
			continue
		}

		if current.Overlaps(period) && !sequence.Contains(current) {
			sequence = sequence.Push(current.Intersect(period))
		}

		isPreviouslyContained = isContained

		if !isContained {
			current = period
		}
	}

	return sequence
}

func (s Sequence) Clear() Sequence {
	s.intervals = []Period{}

	return s
}

func (s Sequence) IndexOf(other Period) int {
	for i, p := range s.intervals {
		if p.Equals(other) {
			return i
		}
	}

	return -1
}

func (s Sequence) filterOffset(offset int) int {
	total := len(s.intervals)

	if 0 == total {
		return -1
	}

	if 0 > total+offset {
		return -1
	}

	if 0 > total-offset-1 {
		return -1
	}

	if 0 > offset {
		return total + offset
	}

	return offset
}

func (s Sequence) Count() int {
	return len(s.intervals)
}

func (s Sequence) totalTimeDuration() int64 {
	var total int64

	for _, p := range s.intervals {
		total += p.GetTimestampInterval()
	}

	return total
}

func (s Sequence) GetTotalTimestampInterval() int64 {
	return s.totalTimeDuration()
}

func (s Sequence) sortByStartDate(period, other Period) int64 {
	if period.startDate.Equal(other.startDate) {
		return 0
	} else if period.startDate.Before(other.startDate) {
		return -1
	}

	return 1
}

func (s Sequence) Contains(periods ...Period) bool {
	for _, p := range periods {
		if -1 == s.IndexOf(p) {
			return false
		}
	}

	return 0 != len(periods)
}

func (s Sequence) Equals(other Sequence) bool {
	if s.Count() != other.Count() {
		return false
	}

	for offset, p := range s.intervals {
		if !p.Equals(other.Get(offset)) {
			return false
		}
	}

	return true
}

func (s Sequence) IsEmpty() bool {
	return len(s.intervals) == 0
}

func (s Sequence) Every(callback func(Period, int) bool) bool {
	for offset, p := range s.intervals {
		if !callback(p, offset) {
			return false
		}
	}

	return true
}

func (s Sequence) Some(callback func(Period, int) bool) bool {
	for offset, p := range s.intervals {
		if callback(p, offset) {
			return true
		}
	}

	return false

}

func (s Sequence) OffsetExists(offset int) bool {
	return -1 != s.filterOffset(offset)
}

func (s Sequence) OffsetGet(offset int) Period {
	return s.Get(offset)
}

func (s Sequence) Get(offset int) Period {
	offset = s.filterOffset(offset)

	if -1 == offset {
		return Period{}
	}

	return s.intervals[offset]
}

func (s Sequence) GetInterval() []Period {
	if s.IsEmpty() {
		return []Period{}
	}

	return s.intervals
}
