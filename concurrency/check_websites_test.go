package concurrency

import "testing"

var StubWebsiteCheckerFalse = func(_ string) bool {
	return false
}

var StubWebsiteCheckerTrue = func(_ string) bool {
	return true
}

func TestCheckWebsites(t *testing.T) {
	equalMaps := func(t *testing.T, want map[string]bool, got map[string]bool) {
		t.Helper()
		equalMaps := func(a map[string]bool, b map[string]bool) bool {
			for k, _ := range a {
				valueInB, ok := b[k]
				if !ok || valueInB != a[k] {
					return false
				}
			}
			return true
		}

		if len(want) != len(got) || !equalMaps(want, got) {
			t.Errorf("Expected %v, got %v", want, got)
		}
	}

	t.Run("when WebsiteChecker returns false", func(t *testing.T) {
		t.Run("it adds the url as false in the results", func(t *testing.T) {
			urls := []string{"https://www.google.com"}
			got := CheckWebsites(StubWebsiteCheckerFalse, urls)
			want := map[string]bool{"https://www.google.com": false}
			equalMaps(t, want, got)
		})
	})

	t.Run("when WebsiteChecker returns true", func(t *testing.T) {
		t.Run("it adds the url as true in the results", func(t *testing.T) {
			urls := []string{"https://www.google.com"}
			got := CheckWebsites(StubWebsiteCheckerTrue, urls)
			want := map[string]bool{"https://www.google.com": true}
			equalMaps(t, want, got)
		})
	})
}
