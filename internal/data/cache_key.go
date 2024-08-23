package data

import "time"

type CacheKey string

func (k CacheKey) EmailLogin() (string, time.Duration) {
	return "email_login_code:" + k.String(), 5 * time.Minute
}
func (k CacheKey) EmailLoginLimit() (string, time.Duration) {
	return "email_login_limit:" + k.String(), 1 * time.Minute
}
func (k CacheKey) String() string {
	return string(k)
}
