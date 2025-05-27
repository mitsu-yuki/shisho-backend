package label

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewLabel(t *testing.T) {
	now := time.Now()
	earlier := now.Add(-1 * time.Hour)
	later := now.Add(1 * time.Hour)
	type args struct {
		name         string
		namePhonic   string
		createAt     time.Time
		lastUpdateAt time.Time
		deletedAt    *time.Time
	}

	tests := []struct {
		name       string
		args       args
		want       *Label
		wantErr    bool
		wantErrStr string
	}{
		{
			name: "正常系",
			args: args{
				name:         "test",
				namePhonic:   "テスト",
				createAt:     now,
				lastUpdateAt: now,
				deletedAt:    nil,
			},
			want: &Label{
				name:         "test",
				namePhonic:   "テスト",
				createAt:     now,
				lastUpdateAt: now,
				deletedAt:    nil,
			},
			wantErr:    false,
			wantErrStr: "",
		},
		{
			name: "異常系: nameが不正",
			args: args{
				name:         "",
				namePhonic:   "テスト",
				createAt:     now,
				lastUpdateAt: now,
				deletedAt:    &later,
			},
			want:       nil,
			wantErr:    true,
			wantErrStr: fmt.Sprintf("レーベル名は%d文字以上である必要があります", nameLengthMin),
		},
		{
			name: "異常系: namePhonicが不正",
			args: args{
				name:         "test",
				namePhonic:   "",
				createAt:     now,
				lastUpdateAt: later,
				deletedAt:    nil,
			},
			want:       nil,
			wantErr:    true,
			wantErrStr: fmt.Sprintf("レーベル名読みは%d文字以上である必要があります", namePhonicLengthMin),
		},
		{
			name: "異常系: namePhonicが空",
			args: args{
				name:         "test",
				namePhonic:   "",
				createAt:     now,
				lastUpdateAt: now,
				deletedAt:    nil,
			},
			want:       nil,
			wantErr:    true,
			wantErrStr: fmt.Sprintf("レーベル名読みは%d文字以上である必要があります", namePhonicLengthMin),
		},
		{
			name: "異常系: namePhonicがカタカナでない",
			args: args{
				name:         "test",
				namePhonic:   "てすと",
				createAt:     now,
				lastUpdateAt: later,
				deletedAt:    &later,
			},
			want:       nil,
			wantErr:    true,
			wantErrStr: "レーベル名読みはカタカナである必要があります",
		},
		{
			name: "異常系: 作成日が不正",
			args: args{
				name:         "test",
				namePhonic:   "テスト",
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
				name:         "test",
				namePhonic:   "テスト",
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
			got, err := NewLabel(tt.args.name, tt.args.namePhonic, tt.args.createAt, tt.args.lastUpdateAt, tt.args.deletedAt)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewLabel() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil && err.Error() != tt.wantErrStr {
				if diff := cmp.Diff(err.Error(), tt.wantErrStr); diff != "" {
					t.Errorf("got: %v, want: %s.\n error is %s", err.Error(), tt.wantErrStr, diff)
				}
			}
			diff := cmp.Diff(
				got, tt.want,
				cmp.AllowUnexported(Label{}),
				cmpopts.IgnoreFields(Label{}, "id"),
			)

			if diff != "" {
				t.Errorf("NewLabel() = %v, want = %v.\n error is %s", got, tt.want, diff)
			}
		})
	}
}
