package myFile

import "testing"

func TestPrintFile(t *testing.T) {
	type args struct {
		fullFilename string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "t1", args: args{fullFilename: "test.txt"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintFile(tt.args.fullFilename)
		})
	}
}
