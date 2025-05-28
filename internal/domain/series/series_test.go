package series

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/mitsu-yuki/shisho-backend/pkg/ulid"
)

func TestNewSeries(t *testing.T) {
	bookID1 := ulid.NewULID()
	bookID2 := ulid.NewULID()
	statusID := ulid.NewULID()
	now := time.Now()
	earlier := now.Add(-1 * time.Hour)
	later := now.Add(1 * time.Hour)
	type args struct {
		name         string
		books        SeriesBooks
		statusID     string
		createAt     time.Time
		lastUpdateAt time.Time
		deletedAt    *time.Time
	}
	tests := []struct {
		name       string
		args       args
		want       *Series
		wantErr    bool
		wantErrStr string
	}{
		{
			name: "正常系",
			args: args{
				name: "テスト",
				books: []SeriesBook{
					{
						bookID: bookID1,
					},
					{
						bookID: bookID2,
					},
				},
				statusID:     statusID,
				createAt:     earlier,
				lastUpdateAt: now,
				deletedAt:    &later,
			},
			want: &Series{
				name: "テスト",
				books: []SeriesBook{
					{
						bookID: bookID1,
					},
					{
						bookID: bookID2,
					},
				},
				statusID:     statusID,
				createAt:     earlier,
				lastUpdateAt: now,
				deletedAt:    &later,
			},
			wantErr:    false,
			wantErrStr: "",
		},
		{
			name: "異常系: nameが不正",
			args: args{
				name: "",
				books: []SeriesBook{
					{
						bookID: bookID1,
					},
				},
				statusID:     statusID,
				createAt:     now,
				lastUpdateAt: now,
				deletedAt:    nil,
			},
			want:       nil,
			wantErr:    true,
			wantErrStr: fmt.Sprintf("タイトル名は%d文字以上である必要があります", seriesNameLengthMin),
		},
		{
			name: "異常系: 作品リストIDが不正",
			args: args{
				name:         "テスト",
				books:        []SeriesBook{},
				statusID:     statusID,
				createAt:     now,
				lastUpdateAt: now,
				deletedAt:    nil,
			},
			want:       nil,
			wantErr:    true,
			wantErrStr: fmt.Sprintf("シリーズ作品は%d作品以上である必要があります", seriesBooksLengthMin),
		},
		{
			name: "異常系: ステータスIDが不正",
			args: args{
				name: "テスト",
				books: []SeriesBook{
					{
						bookID: bookID1,
					},
				},
				statusID:     "",
				createAt:     now,
				lastUpdateAt: now,
				deletedAt:    nil,
			},
			want:       nil,
			wantErr:    true,
			wantErrStr: "ステータスIDが不正です",
		},
		{
			name: "異常系: 更新日が不正",
			args: args{
				name: "テスト",
				books: []SeriesBook{
					{
						bookID: bookID1,
					},
				},
				statusID:     statusID,
				createAt:     later,
				lastUpdateAt: now,
				deletedAt:    nil,
			},
			want:       nil,
			wantErr:    true,
			wantErrStr: "更新日は作成日よりも後である必要があります",
		},
		{
			name: "異常系: 削除日が不正",
			args: args{
				name: "テスト",
				books: []SeriesBook{
					{
						bookID: bookID1,
					},
				},
				statusID:     statusID,
				createAt:     now,
				lastUpdateAt: now,
				deletedAt:    &earlier,
			},
			want:       nil,
			wantErr:    true,
			wantErrStr: "削除日は作成日よりも後である必要があります",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSeries(tt.args.name, tt.args.books, tt.args.statusID, tt.args.createAt, tt.args.lastUpdateAt, tt.args.deletedAt)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSeries() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil && err.Error() != tt.wantErrStr {
				if diff := cmp.Diff(err.Error(), tt.wantErrStr); diff != "" {
					t.Errorf("got: %v, want: %s.\n error is %s", err.Error(), tt.wantErrStr, diff)
				}
			}
			diff := cmp.Diff(
				got, tt.want,
				cmp.AllowUnexported(Series{}, SeriesBook{}),
				cmpopts.IgnoreFields(Series{}, "id"),
			)

			if diff != "" {
				t.Errorf("NewSeries() = %v, want = %v.\n error is %s", got, tt.want, diff)
			}
		})
	}
}
