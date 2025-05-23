package author

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewAuthor(t *testing.T) {
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
		name    string
		args    args
		want    *Author
		wantErr bool
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
			want: &Author{
				name:         "test",
				namePhonic:   "テスト",
				createAt:     now,
				lastUpdateAt: now,
				deletedAt:    nil,
			},
			wantErr: false,
		},
		{
			name: "異常系: nameが不正",
			args: args{
				name:         "",
				namePhonic:   "テスト",
				createAt:     now,
				lastUpdateAt: later,
				deletedAt:    &later,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "異常系: namePhonicが不正",
			args: args{
				name:         "test",
				namePhonic:   "てすと",
				createAt:     now,
				lastUpdateAt: later,
				deletedAt:    nil,
			},
			want:    nil,
			wantErr: true,
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
			want:    nil,
			wantErr: true,
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
			want:    nil,
			wantErr: true,
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
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAuthor(tt.args.name, tt.args.namePhonic, tt.args.createAt, tt.args.lastUpdateAt, tt.args.deletedAt)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAuthor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			diff := cmp.Diff(
				got, tt.want,
				cmp.AllowUnexported(Author{}),
				cmpopts.IgnoreFields(Author{}, "id"),
			)

			if diff != "" {
				t.Errorf("NewAuthor() = %v, want = %v, error is %s", got, tt.want, diff)
			}
		})
	}
}
