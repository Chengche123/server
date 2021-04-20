package env

import (
	"os"
	"testing"
)

func TestGetEnvOrDefault(t *testing.T) {
	cases := []struct {
		name string
		env  string
		def  string
		want string
	}{
		{
			name: "preset_test",
			env:  "TEST_ENV",
			def:  "",
			want: "TEST_ENV",
		},
		{
			name: "unset_test",
			env:  "NO_ENV",
			def:  "def",
			want: "def",
		},
	}

	os.Setenv("TEST_ENV", "TEST_ENV")

	for _, cs := range cases {
		t.Run(cs.name, func(t *testing.T) {
			val := GetEnvOrDefault(cs.env, cs.def)

			if val != cs.want {
				t.Errorf("GetEnvOrDefault(%q, %q) = %q, want %q", cs.env, cs.def, val, cs.want)
			}
		})
	}

}
