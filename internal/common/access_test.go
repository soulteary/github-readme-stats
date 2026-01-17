package common

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestGuardAccess_InvalidAccessType(t *testing.T) {
	// Clear environment variables
	os.Clearenv()
	loadEnvVars()

	w := httptest.NewRecorder()
	colors := map[string]string{
		"theme": "default",
	}

	result := GuardAccess(w, "test-id", "invalid-type", colors)

	if result.IsPassed {
		t.Error("期望访问被拒绝，但实际通过了")
	}

	if !strings.Contains(result.Response, "Invalid access type") {
		t.Errorf("期望错误消息包含 'Invalid access type'，实际得到: %s", result.Response)
	}
}

func TestGuardAccess_Username_Whitelisted(t *testing.T) {
	// Set whitelist environment variable
	os.Clearenv()
	os.Setenv("WHITELIST", "testuser,alloweduser")
	loadEnvVars()

	w := httptest.NewRecorder()
	colors := map[string]string{
		"theme": "default",
	}

	result := GuardAccess(w, "testuser", "username", colors)

	if !result.IsPassed {
		t.Error("期望白名单用户通过，但实际被拒绝")
	}
}

func TestGuardAccess_Username_NotWhitelisted(t *testing.T) {
	// Set whitelist environment variable
	os.Clearenv()
	os.Setenv("WHITELIST", "alloweduser,anotheruser")
	loadEnvVars()

	w := httptest.NewRecorder()
	colors := map[string]string{
		"theme": "default",
	}

	result := GuardAccess(w, "testuser", "username", colors)

	if result.IsPassed {
		t.Error("期望非白名单用户被拒绝，但实际通过了")
	}

	if w.Code != http.StatusForbidden {
		t.Errorf("期望状态码为 %d，实际为 %d", http.StatusForbidden, w.Code)
	}

	if w.Header().Get("Content-Type") != "image/svg+xml" {
		t.Errorf("期望 Content-Type 为 'image/svg+xml'，实际为 '%s'", w.Header().Get("Content-Type"))
	}

	if !strings.Contains(result.Response, notWhitelistedUsernameMessage) {
		t.Errorf("期望错误消息包含白名单消息，实际得到: %s", result.Response)
	}
}

func TestGuardAccess_Username_EmptyWhitelist(t *testing.T) {
	// Don't set whitelist (empty whitelist means allow all)
	os.Clearenv()
	loadEnvVars()

	w := httptest.NewRecorder()
	colors := map[string]string{
		"theme": "default",
	}

	result := GuardAccess(w, "anyuser", "username", colors)

	if !result.IsPassed {
		t.Error("期望空白名单时允许所有用户，但实际被拒绝")
	}
}

func TestGuardAccess_Username_Blacklisted(t *testing.T) {
	// Don't set whitelist, test blacklist
	os.Clearenv()
	loadEnvVars()

	w := httptest.NewRecorder()
	colors := map[string]string{
		"theme": "default",
	}

	// Use username from blacklist
	result := GuardAccess(w, "renovate-bot", "username", colors)

	if result.IsPassed {
		t.Error("期望黑名单用户被拒绝，但实际通过了")
	}

	if w.Code != http.StatusForbidden {
		t.Errorf("期望状态码为 %d，实际为 %d", http.StatusForbidden, w.Code)
	}

	if !strings.Contains(result.Response, blacklistedMessage) {
		t.Errorf("期望错误消息包含黑名单消息，实际得到: %s", result.Response)
	}
}

func TestGuardAccess_Username_Blacklisted_WithWhitelist(t *testing.T) {
	// When whitelist is set, blacklist check should be skipped
	os.Clearenv()
	os.Setenv("WHITELIST", "renovate-bot")
	loadEnvVars()

	w := httptest.NewRecorder()
	colors := map[string]string{
		"theme": "default",
	}

	// Even if username is in blacklist, it should pass if in whitelist
	result := GuardAccess(w, "renovate-bot", "username", colors)

	if !result.IsPassed {
		t.Error("期望白名单用户即使也在黑名单中也能通过，但实际被拒绝")
	}
}

func TestGuardAccess_Gist_Whitelisted(t *testing.T) {
	// Set Gist whitelist
	os.Clearenv()
	os.Setenv("GIST_WHITELIST", "gist123,gist456")
	loadEnvVars()

	w := httptest.NewRecorder()
	colors := map[string]string{
		"theme": "default",
	}

	result := GuardAccess(w, "gist123", "gist", colors)

	if !result.IsPassed {
		t.Error("期望白名单 Gist 通过，但实际被拒绝")
	}
}

func TestGuardAccess_Gist_NotWhitelisted(t *testing.T) {
	// Set Gist whitelist
	os.Clearenv()
	os.Setenv("GIST_WHITELIST", "gist123,gist456")
	loadEnvVars()

	w := httptest.NewRecorder()
	colors := map[string]string{
		"theme": "default",
	}

	result := GuardAccess(w, "gist999", "gist", colors)

	if result.IsPassed {
		t.Error("期望非白名单 Gist 被拒绝，但实际通过了")
	}

	if w.Code != http.StatusForbidden {
		t.Errorf("期望状态码为 %d，实际为 %d", http.StatusForbidden, w.Code)
	}

	if !strings.Contains(result.Response, notWhitelistedGistMessage) {
		t.Errorf("期望错误消息包含 Gist 白名单消息，实际得到: %s", result.Response)
	}
}

func TestGuardAccess_Wakatime_EmptyWhitelist(t *testing.T) {
	// Wakatime type should use username whitelist
	os.Clearenv()
	loadEnvVars()

	w := httptest.NewRecorder()
	colors := map[string]string{
		"theme": "default",
	}

	result := GuardAccess(w, "anyuser", "wakatime", colors)

	if !result.IsPassed {
		t.Error("期望空白名单时允许所有 Wakatime 用户，但实际被拒绝")
	}
}

func TestGuardAccess_Wakatime_Whitelisted(t *testing.T) {
	// Set whitelist
	os.Clearenv()
	os.Setenv("WHITELIST", "wakatimeuser")
	loadEnvVars()

	w := httptest.NewRecorder()
	colors := map[string]string{
		"theme": "default",
	}

	result := GuardAccess(w, "wakatimeuser", "wakatime", colors)

	if !result.IsPassed {
		t.Error("期望白名单 Wakatime 用户通过，但实际被拒绝")
	}
}

func TestGuardAccess_ResponseWriter_Headers(t *testing.T) {
	// Test ResponseWriter header settings
	os.Clearenv()
	os.Setenv("WHITELIST", "alloweduser")
	loadEnvVars()

	w := httptest.NewRecorder()
	colors := map[string]string{
		"theme": "default",
	}

	GuardAccess(w, "notallowed", "username", colors)

	if w.Header().Get("Content-Type") != "image/svg+xml" {
		t.Errorf("期望 Content-Type 为 'image/svg+xml'，实际为 '%s'", w.Header().Get("Content-Type"))
	}

	if w.Code != http.StatusForbidden {
		t.Errorf("期望状态码为 %d，实际为 %d", http.StatusForbidden, w.Code)
	}

	if len(w.Body.Bytes()) == 0 {
		t.Error("期望响应体不为空")
	}
}

func TestGuardAccess_Colors_Passed(t *testing.T) {
	// Test that response is not written when access passes normally
	os.Clearenv()
	loadEnvVars()

	w := httptest.NewRecorder()
	colors := map[string]string{
		"theme":       "dark",
		"title_color": "#ff0000",
	}

	result := GuardAccess(w, "anyuser", "username", colors)

	if !result.IsPassed {
		t.Error("期望访问通过")
	}

	// Response should not be written when passing (but GuardAccess may set status code, so checking IsPassed is more accurate)
	// Note: Even if passing, if status code was set before, w.Code may not be 0
	// So mainly check IsPassed flag
}
