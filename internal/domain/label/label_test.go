package label

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewLabel(t *testing.T) {
	type args struct {
		name         string
		namePhonic   string
		createAt     time.Time
		lastUpdateAt time.Time
	}

	tests := []struct {
		name    string
		args    args
		want    *Label
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				name:       "test",
				namePhonic: "テスト",
				createAt:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
				lastUpdateAt:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
			},
			want: &Label{
				name:       "test",
				namePhonic: "テスト",
				createAt:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
				lastUpdateAt:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
			},
			wantErr: false,
		},
		{
			name: "異常系: nameが不正",
			args: args{
				name: "",
				namePhonic: "テスト",
				createAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
				lastUpdateAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
			},
			want: nil,
			wantErr: true,
		},
		{
			name: "異常系: namePhonicが不正",
			args: args{
				name: "test",
				namePhonic: "てすと",
				createAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
				lastUpdateAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
			},
			want: nil,
			wantErr: true,
		},
		{
			name: "異常系: namePhonicが空",
			args: args{
				name: "test",
				namePhonic: "",
				createAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
				lastUpdateAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
			},
			want: nil,
			wantErr: true,
		},
		{
			name: "異常系: 日付が不正",
			args: args{
				name: "test",
				namePhonic: "テスト",
				createAt: time.Date(2020, 1, 1, 0, 0, 0, 1, time.Local),
				lastUpdateAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
			},
			want: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err :=  NewLabel(tt.args.name, tt.args.namePhonic, tt.args.createAt, tt.args.lastUpdateAt)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewLabel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			diff := cmp.Diff(
				got, tt.want,
				cmp.AllowUnexported(Label{}),
				cmpopts.IgnoreFields(Label{}, "id"),
			)

			if diff != "" {
				t.Errorf("NewLabel() = %v, want = %v, error is %s", got, tt.want, diff)
			}
		})
	}
}
