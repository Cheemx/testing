package main

import (
	"reflect"
	"testing"
)

type Profile struct {
	Age  int
	City string
}

type Person struct {
	Name    string
	Profile Profile
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Cheems"},
			[]string{"Cheems"},
		},
		{
			"struct with  two string fields",
			struct {
				Name string
				City string
			}{
				"Cheems",
				"Pune",
			},
			[]string{"Cheems", "Pune"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{
				"Cheems",
				22,
			},
			[]string{"Cheems"},
		},
		{
			"struct with nested structs",
			Person{
				"Cheems",
				Profile{22, "Pune"},
			},
			[]string{"Cheems", "Pune"},
		},
		{
			"pointers to things",
			&Person{
				"Cheems",
				Profile{22, "Pune"},
			},
			[]string{"Cheems", "Pune"},
		},
		{
			"slices",
			[]Profile{
				{22, "Pune"},
				{22, "Beed"},
			},
			[]string{"Pune", "Beed"},
		},
		{
			"arrays",
			[2]Profile{
				{22, "Pune"},
				{22, "Beed"},
			},
			[]string{"Pune", "Beed"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
	t.Run("map", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string

		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{22, "Pune"}
			aChannel <- Profile{22, "Beed"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Pune", "Beed"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{22, "Pune"}, Profile{22, "Beed"}
		}

		var got []string
		want := []string{"Pune", "Beed"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}

// First draft of tests!

// func TestWalk(t *testing.T) {
// 	expected := "Cheems"
// 	var got []string

// 	x := struct {
// 		Name string
// 	}{expected}

// 	walk(x, func(input string) {
// 		got = append(got, input)
// 	})

// 	if len(got) != 1 {
// 		t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
// 	}

// 	if got[0] != expected {
// 		t.Errorf("got %q, want %q", got[0], expected)
// 	}
// }
