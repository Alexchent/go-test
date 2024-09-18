package unionfind

import "testing"

func TestEquationsPossible(t *testing.T) {
	type args struct {
		equations []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "t1", args: args{equations: []string{"a==b", "a!=b"}}, want: false},
		{name: "t2", args: args{equations: []string{"b==a", "a==b"}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EquationsPossible(tt.args.equations); got != tt.want {
				t.Errorf("EquationsPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}
