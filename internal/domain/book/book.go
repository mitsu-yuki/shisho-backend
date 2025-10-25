package book

import (
	"fmt"
	"time"
	"unicode/utf8"

	errDomain "github.com/mitsu-yuki/shisho-backend/internal/domain/error"
	"github.com/mitsu-yuki/shisho-backend/pkg/checkdigit"
	"github.com/mitsu-yuki/shisho-backend/pkg/ulid"
)

const (
	// タイトルの最小値
	titleLengthMin = 1
	// 著者の最小数
	bookAuthorsLengthMin = 1
	priceMin             = 0
)

type Book struct {
	id           string
	isbn         *string
	labelID      string
	publishID    string
	title        string
	authorIDs    BookAuthors
	releaseDay   time.Time
	price        int
	explain      string
	createAt     time.Time
	lastUpdateAt time.Time
	deletedAt    *time.Time
}

func newBook(
	id string,
	isbn *string,
	labelID string,
	publishID string,
	title string,
	authorIDs []BookAuthor,
	releaseDay time.Time,
	price int,
	explain string,
	createAt time.Time,
	lastUpdateAt time.Time,
	deletedAt *time.Time,
) (*Book, error) {
	// ISBNがある場合には有効なISBNか調べる
	if isbn != nil && !checkdigit.ISBN13IsValid(*isbn) {
		return nil, errDomain.NewError("ISBNが不正です")
	}

	// レーベルIDのバリデーション
	if !ulid.IsValid(labelID) {
		return nil, errDomain.NewError("レーベルIDが不正です")
	}

	// 出版社IDのバリデーション
	if !ulid.IsValid(publishID) {
		return nil, errDomain.NewError("出版社IDが不正です")
	}

	// タイトルのバリデーション
	if utf8.RuneCountInString(title) < titleLengthMin {
		return nil, errDomain.NewError(fmt.Sprintf("タイトル名は%d文字以上である必要があります", titleLengthMin))
	}

	// 著者リストIDのバリデーション
	if len(authorIDs) < bookAuthorsLengthMin {
		return nil, errDomain.NewError(fmt.Sprintf("著者は%d人以上である必要があります", bookAuthorsLengthMin))
	}

	// 発売日のバリデーション
	if releaseDay.IsZero() {
		return nil, errDomain.NewError("発売日はゼロ値以外である必要があります")
	}

	// 金額のバリデーション
	if price < priceMin {
		return nil, errDomain.NewError(fmt.Sprintf("金額は%d円以上である必要があります", priceMin))
	}
	// 日付のバリデーション(lastUpdateAtのほうが後か)
	if lastUpdateAt.Before(createAt) {
		return nil, errDomain.NewError("更新日は作成日よりも後である必要があります")
	}

	// 削除フラグが立ってない もしくは 削除日は作成日よりも後であるか
	if deletedAt != nil && deletedAt.Before(createAt) {
		return nil, errDomain.NewError("削除日は作成日よりも後である必要があります")
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
		deletedAt:    deletedAt,
	}, nil
}

func Reconstruct(
	id string,
	isbn *string,
	labelID string,
	publishID string,
	title string,
	authorIDs []BookAuthor,
	releaseDay time.Time,
	price int,
	explain string,
	createAt time.Time,
	lastUpdateAt time.Time,
	deletedAt *time.Time,
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
		deletedAt,
	)
}

func NewBook(
	isbn *string,
	labelID string,
	publishID string,
	title string,
	authorIDs []BookAuthor,
	releaseDay time.Time,
	price int,
	explain string,
	createAt time.Time,
	lastUpdateAt time.Time,
	deletedAt *time.Time,
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
		deletedAt,
	)
}

func (b *Book) ID() string {
	return b.id
}

func (b *Book) ISBN() string {
	if b.isbn == nil {
		return ""
	}
	return *b.isbn
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

func (b *Book) DeletedAt() *time.Time {
	return b.deletedAt
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
