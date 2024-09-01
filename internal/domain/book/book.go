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
	authorListID   string
	releaseDay     time.Time
	price          int
	explain        string
	addTime        time.Time
	lastUpdateTime time.Time
}

func newBook(
	id string,
	isbn string,
	labelID string,
	publishID string,
	title string,
	authorListID string,
	releaseDay time.Time,
	price int,
	explain string,
	addTime time.Time,
	lastUpdateTime time.Time,
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
	if !ulid.IsValid(authorListID) {
		return nil, errDomain.NewError("authorListID is invalid")
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
		authorListID:   authorListID,
		releaseDay:     releaseDay,
		price:          price,
		explain:        explain,
		addTime:        addTime,
		lastUpdateTime: lastUpdateTime,
	}, nil
}

func Reconstruct(
	id string,
	isbn string,
	labelID string,
	publishID string,
	title string,
	authorListID string,
	releaseDay time.Time,
	price int,
	explain string,
	addTime time.Time,
	lastUpdateTime time.Time,
) (*Book, error) {
	return newBook(
		id,
		isbn,
		labelID,
		publishID,
		title,
		authorListID,
		releaseDay,
		price,
		explain,
		addTime,
		lastUpdateTime,
	)
}

func NewBook(
	isbn string,
	labelID string,
	publishID string,
	title string,
	authorListID string,
	releaseDay time.Time,
	price int,
	explain string,
	addTime time.Time,
	lastUpdateTime time.Time,
) (*Book, error) {
	return newBook(
		ulid.NewULID(),
		isbn,
		labelID,
		publishID,
		title,
		authorListID,
		releaseDay,
		price,
		explain,
		addTime,
		lastUpdateTime,
	)
}
