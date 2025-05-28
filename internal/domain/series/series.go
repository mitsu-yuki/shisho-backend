package series

import (
	"fmt"
	"time"
	"unicode/utf8"

	errDomain "github.com/mitsu-yuki/shisho-backend/internal/domain/error"
	"github.com/mitsu-yuki/shisho-backend/pkg/ulid"
)

const (
	seriesNameLengthMin  = 1
	seriesBooksLengthMin = 1
)

type Series struct {
	id           string
	name         string
	books        SeriesBooks
	statusID     string
	createAt     time.Time
	lastUpdateAt time.Time
	deletedAt    *time.Time
}

func newSeries(
	id string,
	name string,
	books []SeriesBook,
	statusID string,
	createAt time.Time,
	lastUpdateAt time.Time,
	deletedAt *time.Time,
) (*Series, error) {
	// シリーズIDのバリデーション
	if !ulid.IsValid(id) {
		return nil, errDomain.NewError("シリーズIDが不正です")
	}

	// シリーズ名のバリデーション
	if utf8.RuneCountInString(name) < seriesNameLengthMin {
		return nil, errDomain.NewError(fmt.Sprintf("タイトル名は%d文字以上である必要があります", seriesNameLengthMin))
	}

	// シリーズが内包する作品数のバリデーション
	if len(books) < seriesBooksLengthMin {
		return nil, errDomain.NewError(fmt.Sprintf("シリーズ作品は%d作品以上である必要があります", seriesBooksLengthMin))
	}

	// ステータスのバリデーション
	if !ulid.IsValid(statusID) {
		return nil, errDomain.NewError("ステータスIDが不正です")
	}

	// 日付のバリデーション(lastUpdateAtのほうが後か)
	if lastUpdateAt.Before(createAt) {
		return nil, errDomain.NewError("更新日は作成日よりも後である必要があります")
	}

	// 削除フラグが立ってない もしくは 削除日は作成日よりも後であるか
	if deletedAt != nil && deletedAt.Before(createAt) {
		return nil, errDomain.NewError("削除日は作成日よりも後である必要があります")
	}
	return &Series{
		id:           id,
		name:         name,
		books:        books,
		statusID:     statusID,
		createAt:     createAt,
		lastUpdateAt: lastUpdateAt,
		deletedAt:    deletedAt,
	}, nil
}

func Reconstruct(
	id string,
	name string,
	books []SeriesBook,
	statusID string,
	createAt time.Time,
	lastUpdateAt time.Time,
	deletedAt *time.Time,
) (*Series, error) {
	return newSeries(id, name, books, statusID, createAt, lastUpdateAt, deletedAt)
}

func NewSeries(
	name string,
	books []SeriesBook,
	statusID string,
	createAt time.Time,
	lastUpdateAt time.Time,
	deletedAt *time.Time,
) (*Series, error) {
	return newSeries(ulid.NewULID(), name, books, statusID, createAt, lastUpdateAt, deletedAt)
}

func (s *Series) ID() string {
	return s.id
}

func (s *Series) Name() string {
	return s.name
}

func (s *Series) Books() []SeriesBook {
	return s.books
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

type SeriesBooks []SeriesBook

type SeriesBook struct {
	bookID string
}

func (s SeriesBooks) SeriesIDs() []string {
	var books []string
	for _, book := range s {
		books = append(books, book.bookID)
	}
	return books
}
