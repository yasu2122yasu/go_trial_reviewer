package calc

import "testing"

func TestCalc(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"テストケース：成功", args{3, 5}, 13},
		{"テストケース：失敗", args{4, 1}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Calc(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}
