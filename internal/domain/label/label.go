package label

import (
	"fmt"
	"time"
	"unicode/utf8"

	errDomain "github.com/mitsu-yuki/shisho-backend/internal/domain/error"
	"github.com/mitsu-yuki/shisho-backend/pkg/text"
	"github.com/mitsu-yuki/shisho-backend/pkg/ulid"
)

const (
	nameLengthMin       = 1
	namePhonicLengthMin = 1
)

type Label struct {
	id           string
	name         string
	namePhonic   string
	createAt     time.Time
	lastUpdateAt time.Time
	deletedAt    *time.Time
}

func newLabel(
	id string,
	name string,
	namePhonic string,
	createAt time.Time,
	lastUpdateAt time.Time,
	deletedAt *time.Time,
) (*Label, error) {
	// IDのバリデーション
	if !ulid.IsValid(id) {
		return nil, errDomain.NewError("レーベルIDが不正です")
	}

	// レーベル名のバリデーション
	if utf8.RuneCountInString(name) < nameLengthMin {
		return nil, errDomain.NewError(fmt.Sprintf("レーベル名は%d文字以上である必要があります", nameLengthMin))
	}

	// レーベル名(読み)のバリデーション
	if utf8.RuneCountInString(namePhonic) < namePhonicLengthMin {
		return nil, errDomain.NewError(fmt.Sprintf("レーベル名読みは%d文字以上である必要があります", namePhonicLengthMin))
	}

	if !text.IsKatakana(namePhonic) {
		return nil, errDomain.NewError("レーベル名読みはカタカナである必要があります")
	}

	// 日付のバリデーション(lastUpdateAtのほうが後か)
	if lastUpdateAt.Before(createAt) {
		return nil, errDomain.NewError("更新日は作成日よりも後である必要があります")
	}

	// 削除フラグが立ってない もしくは 削除日は作成日よりも後であるか
	if deletedAt != nil && deletedAt.Before(createAt) {
		return nil, errDomain.NewError("削除日は作成日よりも後である必要があります")
	}

	return &Label{
		id:           id,
		name:         name,
		namePhonic:   namePhonic,
		createAt:     createAt,
		lastUpdateAt: lastUpdateAt,
		deletedAt:    deletedAt,
	}, nil
}

func Reconstruct(
	id string,
	name string,
	namePhonic string,
	createAt time.Time,
	lastUpdateAt time.Time,
	deletedAt *time.Time,
) (*Label, error) {
	return newLabel(id, name, namePhonic, createAt, lastUpdateAt, deletedAt)
}

func NewLabel(
	name string,
	namePhonic string,
	createAt time.Time,
	lastUpdateAt time.Time,
	deletedAt *time.Time,
) (*Label, error) {
	return newLabel(ulid.NewULID(), name, namePhonic, createAt, lastUpdateAt, deletedAt)
}
