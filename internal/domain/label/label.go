package label

import (
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
}

func newLabel(
	id string,
	name string,
	namePhonic string,
	createAt time.Time,
	lastUpdateAt time.Time,
) (*Label, error) {
	// 名前のバリデーション
	if utf8.RuneCountInString(name) < nameLengthMin {
		return nil, errDomain.NewError("name is invalid")
	}

	// 名前(読み)のバリデーション
	if utf8.RuneCountInString(namePhonic) < namePhonicLengthMin || !text.IsKatakana(namePhonic) {
		return nil, errDomain.NewError("name is invalid")
	}

	// 日付のバリデーション(lastUpdateAtのほうが後か)
	if createAt.After(lastUpdateAt) {
		return nil, errDomain.NewError("createAt and lastUpdateAt are invalid")
	}

	return &Label{
		id:           id,
		name:         name,
		namePhonic:   namePhonic,
		createAt:     createAt,
		lastUpdateAt: lastUpdateAt,
	}, nil
}

func Reconstruct(
	id string,
	name string,
	namePhonic string,
	createAt time.Time,
	lastUpdateAt time.Time,
) (*Label, error) {
	return newLabel(id, name, namePhonic, createAt, lastUpdateAt)
}

func NewLabel(
	name string,
	namePhonic string,
	createAt time.Time,
	lastUpdateAt time.Time,
) (*Label, error) {
	return newLabel(ulid.NewULID(), name, namePhonic, createAt, lastUpdateAt)
}
