package haiku

import "testing"

func TestHaikuString(t *testing.T) {
	for tn, tc := range map[string]struct {
		haiku haiku
		want  string
	}{
		"zero": {},
		"no seperator": {
			haiku: haiku{
				components: []string{"one", "two"},
			},
			want: "onetwo",
		},
		"no components": {
			haiku: haiku{
				seperator: "-",
			},
			want: "onetwo",
		},
		"usual": {
			haiku: haiku{
				seperator:  "-",
				components: []string{"one", "two"},
			},
			want: "one-two",
		},
	} {
		t.Run(tn, func(t *testing.T) {
			if got := tc.haiku.String(); got != tc.want {
				t.Errorf("haiku.String(): got: %q want: %q", got, tc.want)
			}
		})
	}
}
