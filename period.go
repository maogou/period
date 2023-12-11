package period

import (
	"time"
)

type Period struct {
	startDate    time.Time
	endDate      time.Time
	boundaryType string
}

func NewPeriod(startDate, endDate time.Time, boundaryType string) Period {
	if startDate.After(endDate) {
		startDate, endDate = endDate, startDate
	}

	if _, ok := boundaryTypes[boundaryType]; !ok {
		boundaryType = IncludeStartExcludeEnd
	}
	return Period{
		startDate:    startDate,
		endDate:      endDate,
		boundaryType: boundaryType,
	}
}

func NewDefaultPeriod(startDate, endDate time.Time) Period {
	if startDate.After(endDate) {
		startDate, endDate = endDate, startDate
	}
	return Period{
		startDate:    startDate,
		endDate:      endDate,
		boundaryType: IncludeStartExcludeEnd,
	}
}

func (p Period) FromPeriod(period Period, boundaryType string) Period {
	return Period{
		startDate:    period.startDate,
		endDate:      period.endDate,
		boundaryType: boundaryType,
	}
}

func (p Period) FromYear(year int, boundaryType string) Period {
	return Period{
		startDate:    time.Date(year, 1, 1, 0, 0, 0, 0, time.Local),
		endDate:      time.Date(year+1, 1, 1, 0, 0, 0, 0, time.Local),
		boundaryType: boundaryType,
	}
}

func (p Period) FromIsoYear(year int, boundaryType string) Period {
	return Period{
		startDate:    time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC),
		endDate:      time.Date(year+1, time.January, 1, 0, 0, 0, 0, time.UTC),
		boundaryType: boundaryType,
	}

}

func (p Period) FromSemester(year, semester int, boundaryType string) Period {
	startMonth := (semester-1)*6 + 1

	return Period{
		startDate:    time.Date(year, time.Month(startMonth), 1, 0, 0, 0, 0, time.Local),
		endDate:      time.Date(year, time.Month(startMonth+6), 1, 0, 0, 0, 0, time.Local),
		boundaryType: boundaryType,
	}
}

func (p Period) FromQuarter(year, quarter int, boundaryType string) Period {
	startMonth := (quarter-1)*3 + 1

	return Period{
		startDate:    time.Date(year, time.Month(startMonth), 1, 0, 0, 0, 0, time.Local),
		endDate:      time.Date(year, time.Month(startMonth+3), 1, 0, 0, 0, 0, time.Local),
		boundaryType: boundaryType,
	}

}

func (p Period) FromMonth(year, month int, boundaryType string) Period {
	return Period{
		startDate:    time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local),
		endDate:      time.Date(year, time.Month(month+1), 1, 0, 0, 0, 0, time.Local),
		boundaryType: boundaryType,
	}
}

func (p Period) FromDay(year, month, day int, boundaryType string) Period {
	return Period{
		startDate:    time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local),
		endDate:      time.Date(year, time.Month(month), day+1, 0, 0, 0, 0, time.Local),
		boundaryType: boundaryType,
	}

}

func (p Period) FromHour(year, month, day, hour int, boundaryType string) Period {
	return Period{
		startDate:    time.Date(year, time.Month(month), day, hour, 0, 0, 0, time.Local),
		endDate:      time.Date(year, time.Month(month), day, hour+1, 0, 0, 0, time.Local),
		boundaryType: boundaryType,
	}
}

func (p Period) FromMinute(year, month, day, hour, minute int, boundaryType string) Period {
	return Period{
		startDate:    time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.Local),
		endDate:      time.Date(year, time.Month(month), day, hour, minute+1, 0, 0, time.Local),
		boundaryType: boundaryType,
	}
}

func (p Period) FromSecond(year, month, day, hour, minute, second int, boundaryType string) Period {
	return Period{
		startDate:    time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Local),
		endDate:      time.Date(year, time.Month(month), day, hour, minute, second+1, 0, time.Local),
		boundaryType: boundaryType,
	}
}

func (p Period) WithBoundaryType(boundaryType string) Period {
	return p.BoundedBy(boundaryType)
}

func (p Period) WithDurationAfterStart(duration time.Duration) Period {
	return p.EndingOn(p.startDate.Add(duration))
}

func (p Period) WithDurationBeforeEnd(duration time.Duration) Period {
	return p.StartingOn(p.endDate.Add(-duration))
}

func (p Period) MoveStartDate(duration time.Duration) Period {
	return p.StartingOn(p.startDate.Add(duration))
}

func (p Period) MoveEndDate(duration time.Duration) Period {
	return p.EndingOn(p.endDate.Add(duration))
}

func (p Period) Move(duration time.Duration) Period {
	other := Period{
		startDate:    p.startDate.Add(duration),
		endDate:      p.endDate.Add(duration),
		boundaryType: p.boundaryType,
	}

	if p.Equals(other) {
		return p
	}

	return other
}

func (p Period) Expand(duration time.Duration) Period {
	other := Period{
		startDate:    p.startDate.Add(-duration),
		endDate:      p.endDate.Add(duration),
		boundaryType: p.boundaryType,
	}

	if p.Equals(other) {
		return p
	}

	return other
}

func (p Period) After(startDate time.Time, duration time.Duration, boundaryType string) Period {
	return Period{
		startDate:    startDate,
		endDate:      startDate.Add(duration),
		boundaryType: boundaryType,
	}
}

func (p Period) Before(endDate time.Time, duration time.Duration, boundaryType string) Period {
	return Period{
		startDate:    endDate.Add(-duration),
		endDate:      endDate,
		boundaryType: boundaryType,
	}
}

func (p Period) Around(sameDate time.Time, duration time.Duration, boundaryType string) Period {
	return Period{
		startDate:    sameDate.Add(-duration),
		endDate:      sameDate.Add(duration),
		boundaryType: boundaryType,
	}
}

func (p Period) StartingOn(startDate time.Time) Period {
	if p.startDate.Equal(startDate) {
		return p
	}

	return Period{
		startDate:    startDate,
		endDate:      p.endDate,
		boundaryType: p.boundaryType,
	}
}

func (p Period) EndingOn(endDate time.Time) Period {
	if p.endDate.Equal(endDate) {
		return p
	}

	return Period{
		startDate:    p.startDate,
		endDate:      endDate,
		boundaryType: p.boundaryType,
	}
}

func (p Period) BoundedBy(bound string) Period {
	if p.boundaryType == bound {
		return p
	}
	return Period{
		startDate:    p.startDate,
		endDate:      p.endDate,
		boundaryType: bound,
	}
}

func (p Period) BordersOnStart(other Period) bool {
	return p.endDate.Equal(other.startDate) && p.boundaryType[1:2]+other.boundaryType[0:1] != IncludeEnd+IncludeStart
}

func (p Period) BordersOnEnd(other Period) bool {
	return other.BordersOnStart(p)
}

func (p Period) MeetsOnEnd(other Period) bool {
	return p.startDate.Equal(other.endDate) &&
		boundaryIsStartIncluded(p.boundaryType) &&
		boundaryIsEndIncluded(other.boundaryType)
}

func (p Period) MeetsOnStart(other Period) bool {
	return p.endDate.Equal(other.startDate) &&
		boundaryIsEndIncluded(p.boundaryType) &&
		boundaryIsStartIncluded(other.boundaryType)
}

func (p Period) Abuts(other Period) bool {
	return p.BordersOnStart(other) || p.BordersOnEnd(other)
}

func (p Period) Meets(other Period) bool {
	return p.MeetsOnEnd(other) || p.MeetsOnStart(other)
}

func (p Period) Overlaps(other Period) bool {
	if p.Meets(other) {
		return true
	}

	return !p.Abuts(other) && p.startDate.Before(other.endDate) && p.endDate.After(other.startDate)
}

func (p Period) IsZero() bool {
	return p.startDate.IsZero() && p.endDate.IsZero()
}

func (p Period) Equals(other Period) bool {
	return p.startDate.Equal(other.startDate) && p.endDate.Equal(other.endDate) && p.boundaryType == other.boundaryType
}

func (p Period) IsEndedBy(other Period) bool {
	return p.endDate.Equal(other.endDate) && p.boundaryType[1] == other.boundaryType[1]
}

func (p Period) IsStartedBy(other Period) bool {
	return p.startDate.Equal(other.startDate) && p.boundaryType[0] == other.boundaryType[0]

}

func (p Period) IsEndIncluded() bool {
	return p.boundaryType[1:2] == IncludeEnd
}

func (p Period) IsStartIncluded() bool {
	return p.boundaryType[0:1] == IncludeStart
}

func (p Period) IsEndExcluded() bool {
	return p.boundaryType[1:2] == ExcludeEnd
}

func (p Period) IsStartExcluded() bool {
	return p.boundaryType[0:1] == ExcludeStart
}

func (p Period) Format(format string) string {
	boundaryType := IncludeStartExcludeEnd
	if len(p.boundaryType) > 0 {
		boundaryType = p.boundaryType
	}
	return boundaryType[0:1] + p.startDate.Format(format) + "," + p.endDate.Format(format) + boundaryType[1:2]
}

func (p Period) GetTimestampInterval() int64 {
	return p.timeDuration()
}

func (p Period) timeDuration() int64 {
	return p.endDate.Unix() - p.startDate.Unix()
}

func (p Period) GetTimeDuration() int64 {
	return p.timeDuration()
}

func (p Period) GetBoundaryType() string {
	if len(p.boundaryType) == 0 {
		return IncludeStartExcludeEnd
	}
	return p.boundaryType
}

func (p Period) GetStartDate() time.Time {
	return p.startDate
}

func (p Period) GetEndDate() time.Time {
	return p.endDate
}

func (p Period) IsBefore(other Period) bool {
	return p.endDate.Before(other.startDate) || p.endDate.Equal(other.startDate) && p.boundaryType[1] != other.boundaryType[0]
}

func (p Period) IsAfter(other Period) bool {
	return other.IsBefore(p)
}

func (p Period) dateInterval() time.Duration {
	return p.endDate.Sub(p.startDate)
}

func (p Period) GetDateInterval() time.Duration {
	return p.dateInterval()
}

func (p Period) DurationCompare(other Period) int {
	pNewDate := p.startDate.Add(p.dateInterval())
	otherNewDate := other.startDate.Add(other.dateInterval())

	if pNewDate.Before(otherNewDate) {
		return -1
	} else if pNewDate.After(otherNewDate) {
		return 1
	}

	return 0
}

func (p Period) DurationEquals(other Period) bool {
	return p.DurationCompare(other) == 0
}

func (p Period) DurationGreaterThan(other Period) bool {
	return p.DurationCompare(other) == 1
}

func (p Period) DurationLessThan(other Period) bool {
	return p.DurationCompare(other) == -1
}

func (p Period) containsDatePoint(datePoint time.Time, boundaryType string) bool {
	switch boundaryType {
	case ExcludeAll:
		return datePoint.After(p.startDate) && datePoint.Before(p.endDate)
	case IncludeAll:
		return (datePoint.Equal(p.startDate) || datePoint.After(p.startDate)) && (datePoint.Equal(p.endDate) || datePoint.Before(p.endDate))
	case ExcludeStartIncludeEnd:
		return datePoint.After(p.startDate) && (datePoint.Equal(p.endDate) || datePoint.Before(p.endDate))
	case IncludeStartExcludeEnd:
		fallthrough
	default:
		return (datePoint.Equal(p.startDate) || datePoint.After(p.startDate)) && datePoint.Before(p.endDate)
	}
}

func (p Period) containsInterval(other Period) bool {

	if p.startDate.Before(other.startDate) && p.endDate.After(other.endDate) {
		return true
	}

	if p.startDate.Equal(other.startDate) && p.endDate.Equal(other.endDate) {
		return p.boundaryType == other.boundaryType || p.boundaryType == IncludeAll
	}

	if p.startDate.Equal(other.startDate) {
		return (p.boundaryType[0] == other.boundaryType[0] || "[" == p.boundaryType[0:1]) && p.containsDatePoint(
			p.startDate.Add(other.GetDateInterval()), p.boundaryType,
		)
	}

	if p.endDate.Equal(other.endDate) {
		return (p.boundaryType[1] == other.boundaryType[1] || "]" == p.boundaryType[1:2]) && p.containsDatePoint(
			p.endDate.Add(-other.GetDateInterval()), p.boundaryType,
		)
	}

	return false
}

func (p Period) Contains(other Period) bool {
	return p.containsInterval(other)
}

func (p Period) Intersect(other Period) Period {
	if !p.Overlaps(other) {
		return Period{boundaryType: IncludeStartExcludeEnd}
	}

	startDate := p.startDate
	endDate := p.endDate
	pRune := []rune(p.boundaryType)
	otherRune := []rune(other.boundaryType)

	if other.startDate.After(p.startDate) {
		startDate = other.startDate
		pRune[0] = otherRune[0]
	}

	if other.endDate.Before(p.endDate) {
		endDate = other.endDate
		pRune[1] = otherRune[1]
	}

	intersect := Period{
		startDate:    startDate,
		endDate:      endDate,
		boundaryType: string(pRune),
	}

	if intersect.Equals(p) {
		return p
	}

	return intersect
}

func (p Period) Diff(other Period) []Period {
	if other.Equals(p) {
		return []Period{}
	}

	intersect := p.Intersect(other)
	if intersect.IsZero() {
		return []Period{}
	}

	merge := p.Merge(other)

	if merge.startDate.Equal(intersect.startDate) {
		boundary := boundaryIncludeStart(merge.boundaryType)
		if boundaryIsEndIncluded(intersect.boundaryType) {
			boundary = boundaryExcludeStart(merge.boundaryType)
		}

		return []Period{merge.StartingOn(intersect.endDate).BoundedBy(boundary)}
	}

	if merge.endDate.Equal(intersect.endDate) {

		boundary := boundaryIncludeEnd(merge.boundaryType)
		if boundaryIsStartIncluded(intersect.boundaryType) {
			boundary = boundaryExcludeEnd(merge.boundaryType)
		}

		return []Period{merge.EndingOn(intersect.startDate).BoundedBy(boundary)}
	}

	lastBoundary := boundaryIncludeEnd(merge.boundaryType)
	if boundaryIsStartIncluded(intersect.boundaryType) {
		lastBoundary = boundaryExcludeEnd(merge.boundaryType)
	}

	firstBoundary := boundaryIncludeStart(merge.boundaryType)
	if boundaryIsEndIncluded(intersect.boundaryType) {
		firstBoundary = boundaryExcludeStart(merge.boundaryType)
	}

	return []Period{
		merge.EndingOn(intersect.startDate).BoundedBy(lastBoundary),
		merge.StartingOn(intersect.endDate).BoundedBy(firstBoundary),
	}
}

func (p Period) Union(periods ...Period) Sequence {
	sequence := Sequence{intervals: append(periods, p)}

	return sequence.Unions()
}

func (p Period) Merge(others ...Period) Period {
	carry := p
	carryBoundary := []rune(carry.boundaryType)
	for _, other := range others {
		otherBoundary := []rune(other.boundaryType)
		if carry.startDate.After(other.startDate) {
			boundary := []rune{otherBoundary[0], carryBoundary[1]}
			carry = Period{
				startDate:    other.startDate,
				endDate:      carry.endDate,
				boundaryType: string(boundary),
			}
		}

		if carry.endDate.Before(other.endDate) {
			boundary := []rune{carryBoundary[0], otherBoundary[1]}
			carry = Period{
				startDate:    carry.startDate,
				endDate:      other.endDate,
				boundaryType: string(boundary),
			}
		}
	}

	return carry
}

func (p Period) Subtract(other Period) Sequence {
	if !p.Overlaps(other) {
		return Sequence{intervals: []Period{p}}
	}

	filter := func(item Period) bool {
		return !p.IsZero() && p.Overlaps(item)
	}

	diffs := p.Diff(other)
	var filtered []Period

	for _, diff := range diffs {
		if filter(diff) {
			filtered = append(filtered, diff)
		}
	}

	return Sequence{intervals: filtered}
}

func (p Period) Gap(other Period) Period {
	if p.Overlaps(other) {
		return Period{boundaryType: IncludeStartExcludeEnd}
	}

	bounds := "["
	if p.IsEndIncluded() {
		bounds = "("
	}

	if other.IsStartIncluded() {
		bounds += ")"
	} else {
		bounds += "]"
	}

	if other.startDate.After(p.startDate) {
		return Period{
			startDate:    p.endDate,
			endDate:      other.startDate,
			boundaryType: bounds,
		}
	}

	return Period{
		startDate:    other.endDate,
		endDate:      p.startDate,
		boundaryType: p.boundaryType,
	}
}
