package haiku

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

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
			want: "",
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

func TestWithSeperator(t *testing.T) {
	got := haiku{
		seperator: "ORIGINAL",
	}
	want := haiku{
		seperator: "WANT",
	}

	WithSeperator("WANT")(&got)

	if diff := cmp.Diff(want, got, cmp.AllowUnexported(haiku{})); diff != "" {
		t.Errorf("WithSeperator(): mismatch (-want, +got):\n%v\n", diff)
	}
}

func TestWithComponents(t *testing.T) {
	c1 := func() string {
		return "one"
	}
	c2 := func() string {
		return "two"
	}

	got := haiku{}
	want := haiku{
		components: []string{"one", "two"},
	}

	WithComponents(c1, c2)(&got)

	if diff := cmp.Diff(want, got, cmp.AllowUnexported(haiku{})); diff != "" {
		t.Errorf("WithComponents(): mismatch (-want, +got):\n%v\n", diff)
	}
}

func TestCustom(t *testing.T) {
	order := []string{}
	o1 := func(h *haiku) {
		order = append(order, "o1")
	}
	o2 := func(h *haiku) {
		order = append(order, "o2")
	}
	o3 := func(h *haiku) {
		order = append(order, "o3")
	}
	_ = Custom(o1, o2, o3)
	if diff := cmp.Diff([]string{"o1", "o2", "o3"}, order, cmp.AllowUnexported(haiku{})); diff != "" {
		t.Errorf("Custom(): option invocation order mismatch (-want, +got):\n%v\n", diff)
	}
}
