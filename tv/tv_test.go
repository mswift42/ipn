package tv

import (
	"reflect"
	"testing"
)

func TestNewProgramme(t *testing.T) {
	type args struct {
		title     string
		subtitle  string
		synopsis  string
		pid       string
		thumbnail string
		url       string
	}
	tests := []struct {
		name string
		args args
		want *Programme
	}{
		{
			name: "prog1",
			args: args{"prog1", "series 1: episode 1", "an equisite programme", "p00", "http://thumbnail.url", "http://programme.url"},
			want: &Programme{"prog1", "series 1: episode 1", "an equisite programme", "p00", "http://thumbnail.url", "http://programme.url", 0},
		},
	}
	for _, tt := range tests {
		if got := NewProgramme(tt.args.title, tt.args.subtitle, tt.args.synopsis, tt.args.pid, tt.args.thumbnail, tt.args.url); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. NewProgramme() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
