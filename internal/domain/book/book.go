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
	id           string
	isbn         string
	labelID      string
	publishID    string
	title        string
	authorIDs    BookAuthors
	releaseDay   time.Time
	price        int
	explain      string
	createAt     time.Time
	lastUpdateAt time.Time
}

func newBook(
	id string,
	isbn string,
	labelID string,
	publishID string,
	title string,
	authorIDs []BookAuthor,
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
	if len(authorIDs) < 1 {
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
		id:           id,
		isbn:         isbn,
		labelID:      labelID,
		publishID:    publishID,
		title:        title,
		authorIDs:    authorIDs,
		releaseDay:   releaseDay,
		price:        price,
		explain:      explain,
		createAt:     createAt,
		lastUpdateAt: lastUpdateAt,
	}, nil
}

func Reconstruct(
	id string,
	isbn string,
	labelID string,
	publishID string,
	title string,
	authorIDs []BookAuthor,
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
	authorIDs []BookAuthor,
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

func (b *Book) ID() string {
	return b.id
}

func (b *Book) ISBN() string {
	return b.isbn
}

func (b *Book) LabelID() string {
	return b.labelID
}

func (b *Book) PublishID() string {
	return b.publishID
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) AuthorIDs() []string {
	var authorIDs []string
	for _, author := range b.authorIDs {
		authorIDs = append(authorIDs, author.authorID)
	}
	return authorIDs
}

func (b *Book) ReleaseDay() time.Time {
	return b.releaseDay
}

func (b *Book) Price() int {
	return b.price
}

func (b *Book) Explain() string {
	return b.explain
}

func (b *Book) CreateAt() time.Time {
	return b.createAt
}

func (b *Book) LastUpdateAt() time.Time {
	return b.lastUpdateAt
}

type BookAuthors []BookAuthor

type BookAuthor struct {
	authorID string
}

func (b BookAuthors) AuthorIDs() []string {
	var authorIDs []string
	for _, author := range b {
		authorIDs = append(authorIDs, author.authorID)
	}
	return authorIDs
}
