package reflection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := map[string]struct {
		Input         interface{}
		ExpectedCalls []string
	}{
		"struct with one string field": {
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		"struct with two string fields": {
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		"struct with not string field": {
			struct {
				Name string
				Age  int
			}{"Chris", 24},
			[]string{"Chris"},
		},
		"struct with nested struct :o": {
			Person{Name: "Chris", Profile: Profile{Age: 24, City: "London"}},
			[]string{"Chris", "London"},
		},
		"pointer to things": {
			&Person{
				Name: "Chris",
				Profile: Profile{
					27,
					"London",
				},
			},
			[]string{"Chris", "London"},
		},
		"slices": {
			[]Profile{
				{33, "London"},
				{34, "Reykjavik"},
			},
			[]string{"London", "Reykjavik"},
		},
		"arrays": {
			[2]Profile{
				{33, "London"},
				{34, "Reykjavik"},
			},
			[]string{"London", "Reykjavik"},
		},
	}
	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			assert.Equal(t, test.ExpectedCalls, got)
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assert.Contains(t, got, "Bar")
		assert.Contains(t, got, "Boz")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{34, "Katowice"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		assert.Equal(t, want, got)
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Katowice"}
		}
		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		assert.Equal(t, want, got)
	})
}
