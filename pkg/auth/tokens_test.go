package auth

import "testing"

func TestGenerateNewTokens(t *testing.T) {
	at, rt, err := GenerateNewTokens("jopa")
	t.Logf(`
		AccessToken:  %s
		RefreshToken: %s
	`, at, rt)

	if err != nil {
		t.Fatalf("Unable to generate tokens: %s", err.Error())
	}

	if at == "" {
		t.Fatal("Access token is empty")
	}

	if rt == "" {
		t.Fatal("Refresh token is empty")
	}
}

func TestParseToken(t *testing.T) {
	at, _, err := GenerateNewTokens("Jaba")
	if err != nil {
		t.Fatalf("Unable to generate tokens: %s", err.Error())
	}

	userid, _, err := ParseToken(at, "ACCESS_KEY")
	if err != nil {
		t.Fatalf("Unable to parse AccessToken: %s", err.Error())
	}

	if userid == "" {
		t.Fatal("Userid is empty")
	}
}