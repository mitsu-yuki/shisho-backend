package book

import (
	"time"
	"unicode/utf8"

	errDomain "github.com/mitsu-yuki/shisho-backend/internal/domain/error"
	"github.com/mitsu-yuki/shisho-backend/pkg/checkdigit"
	"github.com/mitsu-yuki/shisho-backend/pkg/ulid"
)

const (
	// タイトルの最小値
	titleLengthMin = 1
)

type Book struct {
	id             string
	isbn           string
	labelID        string
	publishID      string
	title          string
	authorIDs   string
	releaseDay     time.Time
	price          int
	explain        string
	createAt        time.Time
	lastUpdateAt time.Time
}

func newBook(
	id string,
	isbn string,
	labelID string,
	publishID string,
	title string,
	authorIDs string,
	releaseDay time.Time,
	price int,
	explain string,
	createAt time.Time,
	lastUpdateAt time.Time,
) (*Book, error) {
	// レーベルIDのバリデーション
	if !ulid.IsValid(labelID) {
		return nil, errDomain.NewError("labelID is invalid.")
	}

	// 出版社IDのバリデーション
	if !ulid.IsValid(publishID) {
		return nil, errDomain.NewError("publishID is invalid.")
	}

	// 著者リストIDのバリデーション
	if !ulid.IsValid(authorIDs) {
		return nil, errDomain.NewError("authorIDs is invalid")
	}

	// タイトルのバリデーション
	if utf8.RuneCountInString(title) < titleLengthMin {
		return nil, errDomain.NewError("title is invalid")
	}

	// ISBNがある場合には有効なISBNか調べる
	if isbn != "" && !checkdigit.ISBN13IsValid(isbn) {
		return nil, errDomain.NewError("ISBN is invalid")
	}

	return &Book{
		id:             id,
		isbn:           isbn,
		labelID:        labelID,
		publishID:      publishID,
		title:          title,
		authorIDs:   authorIDs,
		releaseDay:     releaseDay,
		price:          price,
		explain:        explain,
		createAt:        createAt,
		lastUpdateAt: lastUpdateAt,
	}, nil
}

func Reconstruct(
	id string,
	isbn string,
	labelID string,
	publishID string,
	title string,
	authorIDs string,
	releaseDay time.Time,
	price int,
	explain string,
	createAt time.Time,
	lastUpdateAt time.Time,
) (*Book, error) {
	return newBook(
		id,
		isbn,
		labelID,
		publishID,
		title,
		authorIDs,
		releaseDay,
		price,
		explain,
		createAt,
		lastUpdateAt,
	)
}

func NewBook(
	isbn string,
	labelID string,
	publishID string,
	title string,
	authorIDs string,
	releaseDay time.Time,
	price int,
	explain string,
	createAt time.Time,
	lastUpdateAt time.Time,
) (*Book, error) {
	return newBook(
		ulid.NewULID(),
		isbn,
		labelID,
		publishID,
		title,
		authorIDs,
		releaseDay,
		price,
		explain,
		createAt,
		lastUpdateAt,
	)
}
