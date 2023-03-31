package poker

import (
	"reflect"
	"testing"
)

func TestCheckInputLength(t *testing.T) {
	type args struct {
		t   string
		err string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
		wantError  string
	}{
		{
			name:       "The test is successful(NORMAL)",
			args:       args{"H1 Q3 H2 H4 D5", ""},
			wantOutput: "H1 Q3 H2 H4 D5",
			wantError:  "",
		},
		{
			name:       "The test is successful(ERROR)",
			args:       args{"H1 Q3 H2 H4 D5 Q3", ""},
			wantOutput: "H1 Q3 H2 H4 D5 Q3",
			wantError:  "400(Bad Request): 入力値の長さが異なります",
		},

		// Failしたテスト（例）
		// {
		// 	name:       "The test is failed(NORMAL)",
		// 	args:       args{"H1 Q3 H2 H4 D5", ""},
		// 	wantOutput: "",
		// 	wantError:  "400(Bad Request): 入力値の長さが異なります",
		// },
		// {
		// 	name:       "The test is failed(ERROR)",
		// 	args:       args{"H1 Q3 H2 H4 D5 Q3", ""},
		// 	wantOutput: "H1 Q3 H2 H4 D5",
		// 	wantError:  "",
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CheckInputLength(tt.args.t, tt.args.err)
			if got != tt.wantOutput {
				t.Errorf("CheckInputLength() got = %v, want %v", got, tt.wantOutput)
			}
			if got1 != tt.wantError {
				t.Errorf("CheckInputLength() got1 = %v, want %v", got1, tt.wantError)
			}
		})
	}
}

func TestExDuplicatedCards(t *testing.T) {
	type args struct {
		cards string
		err   string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		// Passしたテスト。正常系・異常系ともに確認
		{
			name:  "The test is successful(NORMAL)",
			args:  args{"H1 Q3 H2 H4 D5", ""},
			want:  "H1 Q3 H2 H4 D5",
			want1: "",
		},
		{
			name:  "The test is successful(ERROR)",
			args:  args{"H1 Q3 H2 H2 D5", ""},
			want:  "H1 Q3 H2 H2 D5",
			want1: "400(Bad Request): 同じカードを複数枚入力しています",
		},

		// Failしたテスト（例）
		// {
		// 	name:  "The test is failed(NORMAL)",
		// 	args:  args{"H1 Q3 H2 H4 D5", ""},
		// 	want:  "H1 Q3 H2 H4 D5",
		// 	want1: "400(Bad Request): 同じカードを複数枚入力しています",
		// },
		// {
		// 	name:  "The test is failed(ERROR)",
		// 	args:  args{"H1 Q3 H2 H2 D5", ""},
		// 	want:  "H1 Q3 H2 H2 D5",
		// 	want1: "",
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ExDuplicatedCards(tt.args.cards, tt.args.err)
			if got != tt.want {
				t.Errorf("ExDuplicatedCards() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ExDuplicatedCards() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMakeIntSlice(t *testing.T) {
	type args struct {
		cards string
		err   string
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 string
	}{
		{
			name:  "The test is successful(NORMAL)",
			args:  args{"H1 Q3 H2 H4 D5", ""},
			want:  []int{1, 3, 2, 4, 5},
			want1: "",
		},
		{
			name:  "The test is successful(ERROR)",
			args:  args{"H1 Q3 H2 H4 D5 Q34", ""},
			want:  []int{1, 3, 2, 4, 5, 34},
			want1: "400(Bad Request): 入力されたカードの値が異常です",
		},

		// Failしたテスト（例）
		// {
		// 	name:  "The test is failed(NORMAL)",
		// 	args:  args{"H1 Q3 H2 H4 D5", ""},
		// 	want:  []int{0},
		// 	want1: "400(Bad Request): 同じカードを複数枚入力しています",
		// },
		// {
		// 	name:  "The test is failed(ERROR)",
		// 	args:  args{"H1 Q3 H2 H4 D5 Q3", ""},
		// 	want:  []int{1, 3, 2, 4, 5},
		// 	want1: "",
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := MakeIntSlice(tt.args.cards, tt.args.err)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeIntSlice() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MakeIntSlice() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
