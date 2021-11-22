package config

import (
	"crypto/rsa"

	jwt "github.com/golang-jwt/jwt/v4"
)

const privKey = `
-----BEGIN RSA PRIVATE KEY-----
MIICXgIBAAKBgQDN+p9a9oMyqRzkae8yLdJcEK0O0WesH6JiMz+KDrpUwAoAM/KP
DnxFnROJDSBHyHEmPVn5x8GqV5lQ9+6l97jdEEcPo6wkshycM82fgcxOmvtAy4Uo
xq/AeplYqplhcUTGVuo4ZldOLmN8ksGmzhWpsOdT0bkYipHCn5sWZxd21QIDAQAB
AoGBAMJ0++KVXXEDZMpjFDWsOq898xNNMHG3/8ZzmWXN161RC1/7qt/RjhLuYtX9
NV9vZRrzyrDcHAKj5pMhLgUzpColKzvdG2vKCldUs2b0c8HEGmjsmpmgoI1Tdf9D
G1QK+q9pKHlbj/MLr4vZPX6xEwAFeqRKlzL30JPD+O6mOXs1AkEA8UDzfadH1Y+H
bcNN2COvCqzqJMwLNRMXHDmUsjHfR2gtzk6D5dDyEaL+O4FLiQCaNXGWWoDTy/HJ
Clh1Z0+KYwJBANqRtJ+RvdgHMq0Yd45MMyy0ODGr1B3PoRbUK8EdXpyUNMi1g3iJ
tXMbLywNkTfcEXZTlbbkVYwrEl6P2N1r42cCQQDb9UQLBEFSTRJE2RRYQ/CL4yt3
cTGmqkkfyr/v19ii2jEpMBzBo8eQnPL+fdvIhWwT3gQfb+WqxD9v10bzcmnRAkEA
mzTgeHd7wg3KdJRtQYTmyhXn2Y3VAJ5SG+3qbCW466NqoCQVCeFwEh75rmSr/Giv
lcDhDZCzFuf3EWNAcmuMfQJARsWfM6q7v2p6vkYLLJ7+VvIwookkr6wymF5Zgb9d
E6oTM2EeUPSyyrj5IdsU2JCNBH1m3JnUflz8p8/NYCoOZg==
-----END RSA PRIVATE KEY-----
`

const hmacSecret = `SECRET_2gtzk`

type Config struct {
	HostAddr string
	TenantID string
	CertFile string
}

func Load() *Config {

	return &Config{}
}

func PrivateKey() *rsa.PrivateKey {

	p, _ := jwt.ParseRSAPrivateKeyFromPEM([]byte(privKey))
	return p
}

func PublicKey() *rsa.PublicKey {
	return &PrivateKey().PublicKey
}

func HMACKey() []byte {
	return []byte(hmacSecret)
}
