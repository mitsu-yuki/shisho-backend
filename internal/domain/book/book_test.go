package book

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/mitsu-yuki/shisho-backend/pkg/ulid"
)

func TestNewBook(t *testing.T) {
	labelID := ulid.NewULID()
	publishID := ulid.NewULID()
	authorIDs := ulid.NewULID()
	type args struct {
		isbn         string
		labelID      string
		publishID    string
		title        string
		authorIDs string
		price        int
		explain      string
	}
	tests := []struct {
		name string
		args args
		want *Book
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				isbn: "9784758079211",
				labelID: labelID,
				publishID: publishID,
				title: "書籍タイトル",
				authorIDs: authorIDs,
				price: 800,
				explain: "書籍の説明",
			},
			want: &Book{
				isbn: "9784758079211",
				labelID: labelID,
				publishID: publishID,
				title: "書籍タイトル",
				authorIDs: authorIDs,
				price: 800,
				explain: "書籍の説明",
			},
			wantErr: false,
		},
		{
			name: "正常系: ISBNが無い場合",
			args: args{
				labelID: labelID,
				publishID: publishID,
				title: "書籍タイトル",
				authorIDs: authorIDs,
				price: 800,
				explain: "書籍の説明",
			},
			want: &Book{
				labelID: labelID,
				publishID: publishID,
				title: "書籍タイトル",
				authorIDs: authorIDs,
				price: 800,
				explain: "書籍の説明",
			},
			wantErr: false,
		},
		{
			name: "異常系: ISBNが不正",
			args: args{
				isbn: "9784758079210",
				labelID: labelID,
				publishID: publishID,
				title: "書籍タイトル",
				authorIDs: authorIDs,
				price: 800,
				explain: "書籍の説明",
			},
			want: nil,
			wantErr: true,
		},
		{
			name: "異常系: タイトルが不正",
			args: args{
				isbn: "9784758079210",
				labelID: "test",
				publishID: publishID,
				title: "",
				authorIDs: authorIDs,
				price: 800,
				explain: "書籍の説明",
			},
			want: nil,
			wantErr: true,
		},
		{
			name: "異常系: ラベルIDが不正",
			args: args{
				isbn: "9784758079210",
				labelID: "test",
				publishID: publishID,
				title: "書籍タイトル",
				authorIDs: authorIDs,
				price: 800,
				explain: "書籍の説明",
			},
			want: nil,
			wantErr: true,
		},
		{
			// 異常系:
			name: "異常系: 出版社IDが不正",
			args: args{
				isbn: "9784758079210",
				labelID: labelID,
				publishID: "test",
				title: "書籍タイトル",
				authorIDs: authorIDs,
				price: 800,
				explain: "書籍の説明",
			},
			want: nil,
			wantErr: true,
		},
		{
			// 異常系:
			name: "異常系: 著者リストIDが不正",
			args: args{
				isbn: "9784758079210",
				labelID: labelID,
				publishID: publishID,
				title: "書籍タイトル",
				authorIDs: "test",
				price: 800,
				explain: "書籍の説明",
			},
			want: nil,
			wantErr: true,
		},
	}
	ReleaseDay := time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)
	AddTime := time.Date(2020, 1, 2, 3, 4, 5, 6, time.Local)
	LastUpdateAt := time.Date(2020, 1, 2, 3, 4, 5, 6, time.Local)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBook(tt.args.isbn, tt.args.labelID, tt.args.publishID, tt.args.title, tt.args.authorIDs, ReleaseDay, tt.args.price, tt.args.explain, AddTime, LastUpdateAt)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			diff := cmp.Diff(
				got, tt.want,
				cmp.AllowUnexported(Book{}),
				cmpopts.IgnoreFields(Book{}, "id", "releaseDay", "createAt", "lastUpdateAt"),
			)
			if diff != "" {
				t.Errorf("NewBook() error = %v, wantErr %v. diff is %s", got, tt.wantErr, diff)
			}
		})
	}
}
