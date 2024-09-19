package creator

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

type Creator struct {
	id           string
	name         string
	namePhonic   string
	createAt     time.Time
	lastUpdateAt time.Time
}

func newCreator(
	id string,
	name string,
	namePhonic string,
	createAt time.Time,
	lastUpdateAt time.Time,
) (*Creator, error) {
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

	return &Creator{
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
) (*Creator, error) {
	return newCreator(id, name, namePhonic, createAt, lastUpdateAt)
}

func NewCreator(
	name string,
	namePhonic string,
	createAt time.Time,
	lastUpdateAt time.Time,
) (*Creator, error) {
	return newCreator(ulid.NewULID(), name, namePhonic, createAt, lastUpdateAt)
}
