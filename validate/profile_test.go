package validate

import "testing"

func TestKnownProfiles_AllRecognizeThemselves(t *testing.T) {
	for _, p := range KnownProfiles() {
		if !p.IsKnown() {
			t.Errorf("KnownProfiles() member %q is not IsKnown()", p)
		}
	}
}

func TestProfile_IsKnown_Unknown(t *testing.T) {
	if Profile("nonexistent").IsKnown() {
		t.Error("expected unknown profile to report IsKnown() == false")
	}
}
