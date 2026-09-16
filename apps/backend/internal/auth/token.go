package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"sync"
	"time"
)

var ErrInvalidAccessToken = errors.New("invalid access token")

type Principal struct {
	UserID      uint64
	PublicID    string
	Email       string
	Permissions map[string]struct{}
}

type accessEntry struct {
	principal Principal
	expiresAt time.Time
}

type TokenManager struct {
	mu      sync.RWMutex
	entries map[[32]byte]accessEntry
	ttl     time.Duration
}

func NewTokenManager(ttl time.Duration) *TokenManager {
	return &TokenManager{entries: make(map[[32]byte]accessEntry), ttl: ttl}
}

func (m *TokenManager) Issue(principal Principal) (string, time.Time, error) {
	raw := make([]byte, 32)
	if _, err := rand.Read(raw); err != nil {
		return "", time.Time{}, err
	}
	token := base64.RawURLEncoding.EncodeToString(raw)
	hash := sha256.Sum256([]byte(token))
	expiresAt := time.Now().UTC().Add(m.ttl)
	m.mu.Lock()
	m.entries[hash] = accessEntry{principal: principal, expiresAt: expiresAt}
	m.mu.Unlock()
	return token, expiresAt, nil
}

func (m *TokenManager) Verify(token string) (Principal, error) {
	hash := sha256.Sum256([]byte(token))
	m.mu.RLock()
	entry, ok := m.entries[hash]
	m.mu.RUnlock()
	if !ok || time.Now().UTC().After(entry.expiresAt) {
		if ok {
			m.Revoke(token)
		}
		return Principal{}, ErrInvalidAccessToken
	}
	return entry.principal, nil
}

func (m *TokenManager) Revoke(token string) {
	hash := sha256.Sum256([]byte(token))
	m.mu.Lock()
	delete(m.entries, hash)
	m.mu.Unlock()
}
