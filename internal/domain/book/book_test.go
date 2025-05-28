package book

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/mitsu-yuki/shisho-backend/pkg/ulid"
)

func TestNewBook(t *testing.T) {
	validISBN := "9784758079211"
	invalidISBN := "9784758079212"
	labelID := ulid.NewULID()
	publishID := ulid.NewULID()
	authorID1 := ulid.NewULID()
	authorID2 := ulid.NewULID()
	now := time.Now()
	earlier := now.Add(-1 * time.Hour)
	later := now.Add(1 * time.Hour)
	type args struct {
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
	tests := []struct {
		name       string
		args       args
		want       *Book
		wantErr    bool
		wantErrStr string
	}{
		{
			name: "正常系",
			args: args{
				isbn:      &validISBN,
				labelID:   labelID,
				publishID: publishID,
				title:     "書籍タイトル",
				authorIDs: []BookAuthor{
					{
						authorID: authorID1,
					},
					{
						authorID: authorID2,
					},
				},
				releaseDay:   now,
				price:        800,
				explain:      "書籍の説明",
				createAt:     now,
				lastUpdateAt: now,
				deletedAt:    nil,
			},
			want: &Book{
				isbn:      &validISBN,
				labelID:   labelID,
				publishID: publishID,
				title:     "書籍タイトル",
				authorIDs: []BookAuthor{
					{
						authorID: authorID1,
					},
					{
						authorID: authorID2,
					},
				},
				releaseDay:   now,
				price:        800,
				explain:      "書籍の説明",
				createAt:     now,
				lastUpdateAt: now,
				deletedAt:    nil,
			},
			wantErr:    false,
			wantErrStr: "",
		},
		{
			name: "正常系: ISBNがnil",
			args: args{
				isbn:      nil,
				labelID:   labelID,
				publishID: publishID,
				title:     "書籍タイトル",
				authorIDs: []BookAuthor{
					{
						authorID: authorID1,
					},
					{
						authorID: authorID2,
					},
				},
				releaseDay:   now,
				price:        800,
				explain:      "書籍の説明",
				createAt:     now,
				lastUpdateAt: now,
				deletedAt:    nil,
			},
			want: &Book{
				isbn:      nil,
				labelID:   labelID,
				publishID: publishID,
				title:     "書籍タイトル",
				authorIDs: []BookAuthor{
					{
						authorID: authorID1,
					},
					{
						authorID: authorID2,
					},
				},
				releaseDay:   now,
				price:        800,
				explain:      "書籍の説明",
				createAt:     now,
				lastUpdateAt: now,
				deletedAt:    nil,
			},
			wantErr:    false,
			wantErrStr: "",
		},
		{
			name: "異常系: ISBNが不正",
			args: args{
				isbn:      &invalidISBN,
				labelID:   labelID,
				publishID: publishID,
				title:     "書籍タイトル",
				authorIDs: []BookAuthor{
					{
						authorID: authorID1,
					},
				},
				releaseDay:   now,
				price:        800,
				explain:      "書籍の説明",
				createAt:     now,
				lastUpdateAt: later,
				deletedAt:    nil,
			},
			want:       nil,
			wantErr:    true,
			wantErrStr: "ISBNが不正です",
		},
		{
			name: "異常系: レーベルIDが不正",
			args: args{
				isbn:      &validISBN,
				labelID:   "labelID",
				publishID: publishID,
				title:     "書籍タイトル",
				authorIDs: []BookAuthor{
					{
						authorID: authorID1,
					},
				},
				releaseDay:   now,
				price:        800,
				explain:      "書籍の説明",
				createAt:     now,
				lastUpdateAt: later,
				deletedAt:    nil,
			},
			want:       nil,
			wantErr:    true,
			wantErrStr: "レーベルIDが不正です",
		},
		{
			name: "異常系: 出版社IDが不正",
			args: args{
				isbn:      &validISBN,
				labelID:   labelID,
				publishID: "publishID",
				title:     "書籍タイトル",
				authorIDs: []BookAuthor{
					{
						authorID: authorID1,
					},
				},
				releaseDay:   now,
				price:        800,
				explain:      "書籍の説明",
				createAt:     now,
				lastUpdateAt: later,
				deletedAt:    nil,
			},
			want:       nil,
			wantErr:    true,
			wantErrStr: "出版社IDが不正です",
		},
		{
			name: "異常系: タイトルが不正",
			args: args{
				isbn:      &validISBN,
				labelID:   labelID,
				publishID: publishID,
				title:     "",
				authorIDs: []BookAuthor{
					{
						authorID: authorID1,
					},
				},
				releaseDay:   now,
				price:        800,
				explain:      "書籍の説明",
				createAt:     now,
				lastUpdateAt: later,
				deletedAt:    nil,
			},
			want:       nil,
			wantErr:    true,
			wantErrStr: fmt.Sprintf("タイトル名は%d文字以上である必要があります", titleLengthMin),
		},
		{
			name: "異常系: 著者リストIDが不正",
			args: args{
				isbn:         &validISBN,
				labelID:      labelID,
				publishID:    publishID,
				title:        "書籍タイトル",
				authorIDs:    []BookAuthor{},
				releaseDay:   now,
				price:        800,
				explain:      "書籍の説明",
				createAt:     now,
				lastUpdateAt: later,
				deletedAt:    nil,
			},
			want:       nil,
			wantErr:    true,
			wantErrStr: fmt.Sprintf("著者は%d人以上である必要があります", bookAuthorsLengthMin),
		},
		{
			name: "異常系: 発売日が不正",
			args: args{
				isbn:      &validISBN,
				labelID:   labelID,
				publishID: publishID,
				title:     "書籍タイトル",
				authorIDs: []BookAuthor{
					{
						authorID: authorID1,
					},
				},
				releaseDay:   time.Time{},
				price:        800,
				explain:      "書籍の説明",
				createAt:     now,
				lastUpdateAt: later,
				deletedAt:    nil,
			},
			want:       nil,
			wantErr:    true,
			wantErrStr: "発売日はゼロ値以外である必要があります",
		},
		{
			name: "異常系: 金額が不正",
			args: args{
				isbn:      &validISBN,
				labelID:   labelID,
				publishID: publishID,
				title:     "書籍タイトル",
				authorIDs: []BookAuthor{
					{
						authorID: authorID1,
					},
				},
				releaseDay:   now,
				price:        -1,
				explain:      "書籍の説明",
				createAt:     now,
				lastUpdateAt: later,
				deletedAt:    nil,
			},
			want:       nil,
			wantErr:    true,
			wantErrStr: fmt.Sprintf("金額は%d円以上である必要があります", priceMin),
		},
		{
			name: "異常系: 更新日が不正",
			args: args{
				isbn:      &validISBN,
				labelID:   labelID,
				publishID: publishID,
				title:     "書籍タイトル",
				authorIDs: []BookAuthor{
					{
						authorID: authorID1,
					},
				},
				releaseDay:   now,
				price:        800,
				explain:      "書籍の説明",
				createAt:     now,
				lastUpdateAt: earlier,
				deletedAt:    nil,
			},
			want:       nil,
			wantErr:    true,
			wantErrStr: "更新日は作成日よりも後である必要があります",
		},
		{
			name: "異常系: 削除日が不正",
			args: args{
				isbn:      &validISBN,
				labelID:   labelID,
				publishID: publishID,
				title:     "書籍タイトル",
				authorIDs: []BookAuthor{
					{
						authorID: authorID1,
					},
				},
				releaseDay:   now,
				price:        800,
				explain:      "書籍の説明",
				createAt:     now,
				lastUpdateAt: later,
				deletedAt:    &earlier,
			},
			want:       nil,
			wantErr:    true,
			wantErrStr: "削除日は作成日よりも後である必要があります",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBook(
				tt.args.isbn,
				tt.args.labelID,
				tt.args.publishID,
				tt.args.title,
				tt.args.authorIDs,
				tt.args.releaseDay,
				tt.args.price,
				tt.args.explain,
				tt.args.createAt,
				tt.args.lastUpdateAt,
				tt.args.deletedAt,
			)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBook() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil && err.Error() != tt.wantErrStr {
				if diff := cmp.Diff(err.Error(), tt.wantErrStr); diff != "" {
					t.Errorf("got: %v, want: %s.\n error is %s", err.Error(), tt.wantErrStr, diff)
				}
			}
			diff := cmp.Diff(
				got, tt.want,
				cmp.AllowUnexported(Book{}, BookAuthor{}),
				cmpopts.IgnoreFields(Book{}, "id"),
			)

			if diff != "" {
				t.Errorf("NewBook() = %v, want = %v.\n error is %s", got, tt.want, diff)
			}
		})
	}
}
