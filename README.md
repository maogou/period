Period
============

[![Author](http://img.shields.io/badge/author-@maogou-blue.svg?style=flat-square)](https://github.com/maogou)
[![Latest Version](https://img.shields.io/github/release/maogou/period.svg?style=flat-square)](https://github.com/maogou/period/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/maogou/period.svg)](https://pkg.go.dev/github.com/maogou/period)
[![License](https://img.shields.io/github/license/maogou/period.svg)](https://github.com/maogou/period/blob/master/LICENSE)

Project Introduction
-------

Period is a project written in Go language and implements [thephpleague/period](https://github.com/thephpleague/period). It is mainly used for handling and operating time periods. It provides a series of methods such as union, intersection, and difference to facilitate users to perform various operations on time periods.

> ⚠️ **Warning**: Currently, Go (1.21.x) language does not support enumeration types. If Go language supports enumeration types in future versions, we will refactor the methods in the `bounds.go` file.

[简体中文介绍](https://github.com/maogou/period/blob/main/README_zh.md)

Main Features
-------

- **Create Time Period**: Users can create a time period, including start time and end time.
- **Time Period Union**: Users can calculate the union of two or more time periods.
- **Time Period Intersection**: Users can calculate the intersection of two or more time periods.
- **Time Period Difference**: Users can calculate the difference of two time periods.

How to Use
-------

First, you need to import the Period project into your Go project. You can add the following dependencies to your `go.mod` file:

```go
require (
github.com/maogou/period v1.0.0
)
```

Then, you can use the following code to create a time period:

```go
package main

import (
	"github.com/maogou/period"
	"time"
)

func main() {
	p1 := period.Period{
		startDate:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
		endDate:      time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
		boundaryType: period.IncludeStartExcludeEnd,
	}
	
	p2 := period.Period{
		startDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local),
		endDate:      time.Date(2023, 1, 4, 0, 0, 0, 0, time.Local),
		boundaryType: period.IncludeStartExcludeEnd,
	}
	
	result := p1.Union(p2)
	fmt.Println(result)
}
```

Methods
-------

The following are the main methods of the `Period` struct:

- `GetTimestampInterval()`: Returns the timestamp interval of the period.
- `Subtract(Period)`: Subtracts another period from the current period and returns a new period.
- `Merge(Period)`: Merges the current period with another period and returns a new period.
- `Intersect(Period)`: Returns the intersection of the current period and another period.
- `Overlaps(Period)`: Determines whether the current period overlaps with another period.
- `Abuts(Period)`: Determines whether the current period is adjacent to another period.
- `Contains(Period)`: Determines whether the current period contains another period.
- `Equals(Period)`: Determines whether the current period is equal to another period.
- `Gap(Period)`: Returns the gap between the current period and another period.
- `Diff(Period)`: Returns the difference between the current period and another period.
- `Union(Period...)`: Gets the union of the current period collection.
- `IsZero()`: Determines whether the current period is zero.

The following are the main methods of the `Sequence` struct:

- `NewSequence(Period...)`: Creates a new sequence of periods.
- `totalTimeDuration()`: Returns the total time duration of the sequence.
- `GetTotalTimestampInterval()`: Returns the total timestamp interval of the sequence.
- `Sort(func(Period, Period) bool)`: Sorts the sequence based on a given comparison function.
- `Contains(Period...)`: Determines whether the sequence contains a given period.
- `Subtract(Sequence)`: Subtracts another sequence from the current sequence and returns a new sequence.
- `Equals(Sequence)`: Determines whether the current sequence is equal to another sequence.
- `Unions()`: Returns the union of the sequence.
- `Gaps()`: Returns the gaps in the sequence.
- `IsEmpty()`: Determines whether the sequence is empty.
- `IndexOf(Period)`: Returns the index of a given period in the sequence.
- `Count()`: Returns the number of periods in the sequence.
- `Get(int)`: Returns the period at a given index in the sequence.
- `Set(int, Period)`: Sets a period at a given index in the sequence.
- `Push(Period...)`: Adds a period to the end of the sequence.
- `Intersections()`: Returns the intersections of the sequence.
- `Remove(int)`: Removes the period at a given index in the sequence.
- `Filter(func(Period) bool)`: Filters the sequence based on a given filter function.
- `Map(func(Period) Period)`: Maps the sequence based on a given mapping function.
- `Every(func(Period, int) bool)`: Determines whether every period in the sequence satisfies a given condition.
- `Some(func(Period, int) bool)`: Determines whether some periods in the sequence satisfy a given condition.
- `Clear()`: Clears the sequence.

Testing
-------

The Period project includes a series of unit tests to ensure the quality and correctness of the code. You can run these tests by executing the following command:

```bash
go test ./... -v -count=1
```

Code Coverage (100%)
-------

We use Go's built-in testing tools to track code coverage. You can generate and view the coverage report by following these steps:

```bash
 go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out
``` 

Contribution
-------

We welcome any form of contribution, including but not limited to submitting issues, providing feedback, proposing new feature suggestions, and improving code. If you have any questions or suggestions, feel free to submit an issue or pull request on GitHub.

License
-------

The Period project is licensed under the MIT License. For details, please see the [LICENSE](LICENSE) file.