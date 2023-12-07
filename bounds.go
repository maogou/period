package period

// boundaryTypes  边界类型
const (
	IncludeStartExcludeEnd = "[)"
	ExcludeStartIncludeEnd = "(]"
	ExcludeAll             = "()"
	IncludeAll             = "[]"

	IncludeStart = "["
	IncludeEnd   = "]"
	ExcludeStart = "("
	ExcludeEnd   = ")"
)

var boundaryTypes = map[string]int{
	IncludeStartExcludeEnd: 1,
	ExcludeStartIncludeEnd: 1,
	ExcludeAll:             1,
	IncludeAll:             1,
}

// TODO if go support enum, this function can be removed

func boundaryIsStartIncluded(boundary string) bool {
	switch boundary {
	case IncludeStartExcludeEnd, IncludeAll:
		return true
	default:
		return false
	}
}

func boundaryIsEndIncluded(boundary string) bool {
	switch boundary {
	case ExcludeStartIncludeEnd, IncludeAll:
		return true
	default:
		return false
	}
}

func boundaryEqualsStart(self, other string) bool {
	switch self {
	case IncludeAll, IncludeStartExcludeEnd:
		return boundaryIsStartIncluded(other)
	default:
		return !boundaryIsStartIncluded(other)
	}
}

func boundaryEqualsEnd(self, other string) bool {
	switch self {
	case IncludeAll, ExcludeStartIncludeEnd:
		return boundaryIsEndIncluded(other)
	default:
		return !boundaryIsEndIncluded(other)
	}
}

func boundaryIncludeStart(boundary string) string {
	switch boundary {
	case ExcludeAll:
		return IncludeStartExcludeEnd
	case ExcludeStartIncludeEnd:
		return IncludeAll
	default:
		return boundary
	}
}

func boundaryIncludeEnd(boundary string) string {
	switch boundary {
	case ExcludeAll:
		return ExcludeStartIncludeEnd
	case IncludeStartExcludeEnd:
		return IncludeAll
	default:
		return boundary
	}
}

func boundaryExcludeStart(boundary string) string {
	switch boundary {
	case IncludeAll:
		return ExcludeStartIncludeEnd
	case IncludeStartExcludeEnd:
		return ExcludeAll
	default:
		return boundary
	}
}

func boundaryExcludeEnd(boundary string) string {
	switch boundary {
	case IncludeAll:
		return IncludeStartExcludeEnd
	case ExcludeStartIncludeEnd:
		return ExcludeAll
	default:
		return boundary
	}
}

func boundaryReplaceStart(self, other string) string {
	if boundaryIsStartIncluded(other) {
		return boundaryIncludeStart(self)
	}

	return boundaryExcludeStart(self)
}

func boundaryReplaceEnd(self, other string) string {
	if boundaryIsEndIncluded(other) {
		return boundaryIncludeEnd(self)
	}

	return boundaryExcludeEnd(self)
}
