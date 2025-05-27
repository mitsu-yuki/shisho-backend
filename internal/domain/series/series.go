package series

import (
	"time"
	"unicode/utf8"

	errDomain "github.com/mitsu-yuki/shisho-backend/internal/domain/error"
	"github.com/mitsu-yuki/shisho-backend/pkg/ulid"
)

const (
	seriesNameLengthMin = 1
)

type Series struct {
	id           string
	name         string
	seriesIDs    SeriesBooks
	statusID       string
	createAt     time.Time
	lastUpdateAt time.Time
}

type SeriesBooks []SeriesBook

type SeriesBook struct {
	bookID string
}

func newSeries(
	id string,
	name string,
	seriesIDs SeriesBooks,
	statusID string,
	createAt time.Time,
	lastUpdateAt time.Time,
) (*Series, error) {
	// シリーズIDのバリデーション
	if !ulid.IsValid(id) {
		return nil, errDomain.NewError("seriesID is invalid.")
	}

	// シリーズ名のバリデーション
	if utf8.RuneCountInString(name) < seriesNameLengthMin {
		return nil, errDomain.NewError("seriesName is invalid.")
	}

	// シリーズが内包する作品数のバリデーション
	if len(seriesIDs) < 1 {
		return nil, errDomain.NewError("seriesIDs is invalid")
	}

	// ステータスのバリデーション
	if !ulid.IsValid(statusID) {
		return nil, errDomain.NewError("statusID is invalid.")
	}

	return &Series{
		id:           id,
		name:         name,
		seriesIDs:    seriesIDs,
		statusID:       statusID,
		createAt:     createAt,
		lastUpdateAt: lastUpdateAt,
	}, nil
}

func Reconstruct(
	id string,
	name string,
	seriesIDs SeriesBooks,
	statusID string,
	createAt time.Time,
	lastUpdateAt time.Time,
) (*Series, error) {
	return newSeries(id, name, seriesIDs, statusID, createAt, lastUpdateAt)
}

func NewSeries(
	name string,
	seriesIDs SeriesBooks,
	statusID string,
	createAt time.Time,
	lastUpdateAt time.Time,
) (*Series, error) {
	return newSeries(ulid.NewULID(), name, seriesIDs, statusID, createAt, lastUpdateAt)
}

func (s *Series) ID() string {
	return s.id
}

func (s *Series) Name() string {
	return s.name
}

func (s *Series) StatusID() string {
	return s.statusID
}

func (s *Series) CreateAt() time.Time {
	return s.createAt
}

func (s *Series) LastUpdateAt() time.Time {
	return s.lastUpdateAt
}

func (s SeriesBooks) SeriesIDs() []string {
	var seriesIDs []string
	for _, seriesID := range s {
		seriesIDs = append(seriesIDs, seriesID.bookID)
	}
	return seriesIDs
}
