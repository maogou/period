package period

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoundaryIsStartIncluded(t *testing.T) {
	tests := []struct {
		name     string
		boundary string
		want     bool
	}{
		{
			name:     "BoundaryIsStartIncluded_WithIncludeStartExcludeEnd",
			boundary: IncludeStartExcludeEnd,
			want:     true,
		},
		{
			name:     "BoundaryIsStartIncluded_WithIncludeAll",
			boundary: IncludeAll,
			want:     true,
		},
		{
			name:     "BoundaryIsStartIncluded_WithExcludeStartIncludeEnd",
			boundary: ExcludeStartIncludeEnd,
			want:     false,
		},
		{
			name:     "BoundaryIsStartIncluded_WithExcludeAll",
			boundary: ExcludeAll,
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := boundaryIsStartIncluded(tt.boundary)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestBoundaryIsEndIncluded(t *testing.T) {
	tests := []struct {
		name     string
		boundary string
		want     bool
	}{
		{
			name:     "BoundaryIsEndIncluded_WithExcludeStartIncludeEnd",
			boundary: ExcludeStartIncludeEnd,
			want:     true,
		},
		{
			name:     "BoundaryIsEndIncluded_WithIncludeAll",
			boundary: IncludeAll,
			want:     true,
		},
		{
			name:     "BoundaryIsEndIncluded_WithIncludeStartExcludeEnd",
			boundary: IncludeStartExcludeEnd,
			want:     false,
		},
		{
			name:     "BoundaryIsEndIncluded_WithExcludeAll",
			boundary: ExcludeAll,
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := boundaryIsEndIncluded(tt.boundary)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestBoundaryEqualsStart(t *testing.T) {
	tests := []struct {
		name  string
		self  string
		other string
		want  bool
	}{
		{
			name:  "BoundaryEqualsStart_WithIncludeAllAndIncludeAll",
			self:  IncludeAll,
			other: IncludeAll,
			want:  true,
		},
		{
			name:  "BoundaryEqualsStart_WithIncludeAllAndExcludeAll",
			self:  IncludeAll,
			other: ExcludeAll,
			want:  false,
		},
		{
			name:  "BoundaryEqualsStart_WithIncludeStartExcludeEndAndIncludeAll",
			self:  IncludeStartExcludeEnd,
			other: IncludeAll,
			want:  true,
		},
		{
			name:  "BoundaryEqualsStart_WithIncludeStartExcludeEndAndExcludeAll",
			self:  IncludeStartExcludeEnd,
			other: ExcludeAll,
			want:  false,
		},
		{
			name:  "BoundaryEqualsStart_WithExcludeStartIncludeEndAndIncludeAll",
			self:  ExcludeStartIncludeEnd,
			other: IncludeAll,
			want:  false,
		},
		{
			name:  "BoundaryEqualsStart_WithExcludeStartIncludeEndAndExcludeAll",
			self:  ExcludeStartIncludeEnd,
			other: ExcludeAll,
			want:  true,
		},
		{
			name:  "BoundaryEqualsStart_WithExcludeAllAndIncludeAll",
			self:  ExcludeAll,
			other: IncludeAll,
			want:  false,
		},
		{
			name:  "BoundaryEqualsStart_WithExcludeAllAndExcludeAll",
			self:  ExcludeAll,
			other: ExcludeAll,
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := boundaryEqualsStart(tt.self, tt.other)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestBoundaryEqualsEnd(t *testing.T) {
	tests := []struct {
		name  string
		self  string
		other string
		want  bool
	}{
		{
			name:  "BoundaryEqualsEnd_WithIncludeAllAndIncludeAll",
			self:  IncludeAll,
			other: IncludeAll,
			want:  true,
		},
		{
			name:  "BoundaryEqualsEnd_WithIncludeAllAndExcludeAll",
			self:  IncludeAll,
			other: ExcludeAll,
			want:  false,
		},
		{
			name:  "BoundaryEqualsEnd_WithIncludeAllAndIncludeStartExcludeEnd",
			self:  IncludeAll,
			other: IncludeStartExcludeEnd,
			want:  false,
		},
		{
			name:  "BoundaryEqualsEnd_WithIncludeAllAndExcludeStartIncludeEnd",
			self:  IncludeAll,
			other: ExcludeStartIncludeEnd,
			want:  true,
		},
		{
			name:  "BoundaryEqualsEnd_WithExcludeStartIncludeEndAndIncludeAll",
			self:  ExcludeStartIncludeEnd,
			other: IncludeAll,
			want:  true,
		},
		{
			name:  "BoundaryEqualsEnd_WithExcludeStartIncludeEndAndExcludeAll",
			self:  ExcludeStartIncludeEnd,
			other: ExcludeAll,
			want:  false,
		},
		{
			name:  "BoundaryEqualsEnd_WithExcludeStartIncludeEndAndIncludeStartExcludeEnd",
			self:  ExcludeStartIncludeEnd,
			other: IncludeStartExcludeEnd,
			want:  false,
		},
		{
			name:  "BoundaryEqualsEnd_WithExcludeStartIncludeEndAndExcludeStartIncludeEnd",
			self:  ExcludeStartIncludeEnd,
			other: ExcludeStartIncludeEnd,
			want:  true,
		},
		{
			name:  "BoundaryEqualsEnd_WithIncludeStartExcludeEndAndIncludeAll",
			self:  IncludeStartExcludeEnd,
			other: IncludeAll,
			want:  false,
		},
		{
			name:  "BoundaryEqualsEnd_WithIncludeStartExcludeEndAndExcludeAll",
			self:  IncludeStartExcludeEnd,
			other: ExcludeAll,
			want:  true,
		},
		{
			name:  "BoundaryEqualsEnd_WithIncludeStartExcludeEndAndIncludeStartExcludeEnd",
			self:  IncludeStartExcludeEnd,
			other: IncludeStartExcludeEnd,
			want:  true,
		},
		{
			name:  "BoundaryEqualsEnd_WithIncludeStartExcludeEndAndExcludeStartIncludeEnd",
			self:  IncludeStartExcludeEnd,
			other: ExcludeStartIncludeEnd,
			want:  false,
		},
		{
			name:  "BoundaryEqualsEnd_WithExcludeAllAndIncludeAll",
			self:  ExcludeAll,
			other: IncludeAll,
			want:  false,
		},
		{
			name:  "BoundaryEqualsEnd_WithExcludeAllAndExcludeAll",
			self:  ExcludeAll,
			other: ExcludeAll,
			want:  true,
		},
		{
			name:  "BoundaryEqualsEnd_WithExcludeAllAndIncludeStartExcludeEnd",
			self:  ExcludeAll,
			other: IncludeStartExcludeEnd,
			want:  true,
		},
		{
			name:  "BoundaryEqualsEnd_WithExcludeAllAndExcludeStartIncludeEnd",
			self:  ExcludeAll,
			other: ExcludeStartIncludeEnd,
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := boundaryEqualsEnd(tt.self, tt.other)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestBoundaryIncludeStart(t *testing.T) {
	tests := []struct {
		name     string
		boundary string
		want     string
	}{
		{
			name:     "BoundaryIncludeStart_WithExcludeAll",
			boundary: ExcludeAll,
			want:     IncludeStartExcludeEnd,
		},
		{
			name:     "BoundaryIncludeStart_WithExcludeStartIncludeEnd",
			boundary: ExcludeStartIncludeEnd,
			want:     IncludeAll,
		},
		{
			name:     "BoundaryIncludeStart_WithIncludeStartExcludeEnd",
			boundary: IncludeStartExcludeEnd,
			want:     IncludeStartExcludeEnd,
		},
		{
			name:     "BoundaryIncludeStart_WithIncludeAll",
			boundary: IncludeAll,
			want:     IncludeAll,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := boundaryIncludeStart(tt.boundary)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestBoundaryIncludeEnd(t *testing.T) {
	tests := []struct {
		name     string
		boundary string
		want     string
	}{
		{
			name:     "BoundaryIncludeEnd_WithExcludeAll",
			boundary: ExcludeAll,
			want:     ExcludeStartIncludeEnd,
		},
		{
			name:     "BoundaryIncludeEnd_WithIncludeStartExcludeEnd",
			boundary: IncludeStartExcludeEnd,
			want:     IncludeAll,
		},
		{
			name:     "BoundaryIncludeEnd_WithIncludeAll",
			boundary: IncludeAll,
			want:     IncludeAll,
		},
		{
			name:     "BoundaryIncludeEnd_WithExcludeStartIncludeEnd",
			boundary: ExcludeStartIncludeEnd,
			want:     ExcludeStartIncludeEnd,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := boundaryIncludeEnd(tt.boundary)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestBoundaryExcludeStart(t *testing.T) {
	tests := []struct {
		name     string
		boundary string
		want     string
	}{
		{
			name:     "BoundaryExcludeStart_WithIncludeAll",
			boundary: IncludeAll,
			want:     ExcludeStartIncludeEnd,
		},
		{
			name:     "BoundaryExcludeStart_WithIncludeStartExcludeEnd",
			boundary: IncludeStartExcludeEnd,
			want:     ExcludeAll,
		},
		{
			name:     "BoundaryExcludeStart_WithExcludeStartIncludeEnd",
			boundary: ExcludeStartIncludeEnd,
			want:     ExcludeStartIncludeEnd,
		},
		{
			name:     "BoundaryExcludeStart_WithExcludeAll",
			boundary: ExcludeAll,
			want:     ExcludeAll,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := boundaryExcludeStart(tt.boundary)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestBoundaryExcludeEnd(t *testing.T) {
	tests := []struct {
		name     string
		boundary string
		want     string
	}{
		{
			name:     "BoundaryExcludeEnd_WithIncludeAll",
			boundary: IncludeAll,
			want:     IncludeStartExcludeEnd,
		},
		{
			name:     "BoundaryExcludeEnd_WithExcludeStartIncludeEnd",
			boundary: ExcludeStartIncludeEnd,
			want:     ExcludeAll,
		},
		{
			name:     "BoundaryExcludeEnd_WithIncludeStartExcludeEnd",
			boundary: IncludeStartExcludeEnd,
			want:     IncludeStartExcludeEnd,
		},
		{
			name:     "BoundaryExcludeEnd_WithExcludeAll",
			boundary: ExcludeAll,
			want:     ExcludeAll,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := boundaryExcludeEnd(tt.boundary)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestBoundaryReplaceStart(t *testing.T) {
	tests := []struct {
		name  string
		self  string
		other string
		want  string
	}{
		{
			name:  "BoundaryReplaceStart_WithIncludeAll",
			self:  ExcludeAll,
			other: IncludeAll,
			want:  IncludeStartExcludeEnd,
		},
		{
			name:  "BoundaryReplaceStart_WithExcludeStartIncludeEnd",
			self:  ExcludeAll,
			other: ExcludeStartIncludeEnd,
			want:  ExcludeAll,
		},
		{
			name:  "BoundaryReplaceStart_WithIncludeStartExcludeEnd",
			self:  ExcludeAll,
			other: IncludeStartExcludeEnd,
			want:  IncludeStartExcludeEnd,
		},
		{
			name:  "BoundaryReplaceStart_WithExcludeAll",
			self:  ExcludeAll,
			other: ExcludeAll,
			want:  ExcludeAll,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := boundaryReplaceStart(tt.self, tt.other)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestBoundaryReplaceEnd(t *testing.T) {
	tests := []struct {
		name  string
		self  string
		other string
		want  string
	}{
		{
			name:  "BoundaryReplaceEnd_WithIncludeAll",
			self:  ExcludeAll,
			other: IncludeAll,
			want:  ExcludeStartIncludeEnd,
		},
		{
			name:  "BoundaryReplaceEnd_WithExcludeStartIncludeEnd",
			self:  ExcludeAll,
			other: ExcludeStartIncludeEnd,
			want:  ExcludeStartIncludeEnd,
		},
		{
			name:  "BoundaryReplaceEnd_WithIncludeStartExcludeEnd",
			self:  ExcludeAll,
			other: IncludeStartExcludeEnd,
			want:  ExcludeAll,
		},
		{
			name:  "BoundaryReplaceEnd_WithExcludeAll",
			self:  ExcludeAll,
			other: ExcludeAll,
			want:  ExcludeAll,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := boundaryReplaceEnd(tt.self, tt.other)
				assert.Equal(t, tt.want, got)
			},
		)
	}
}
