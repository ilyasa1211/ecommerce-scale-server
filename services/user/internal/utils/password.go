package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"runtime"

	"golang.org/x/crypto/argon2"
)

func HashPassword(password string) (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)

	if err != nil {
		return "", fmt.Errorf("failed to generate salt: %w", err)
	}

	return hashPasswordWithSalt(password, salt)
}
func hashPasswordWithParams(password string, version int, salt []byte, time, memory uint32, threads uint8, keyLen uint32) (string, error) {
	buffer := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLen)
	encodedPass := base64.RawStdEncoding.EncodeToString(buffer)
	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)

	result := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, memory, time, threads, encodedSalt, encodedPass)

	return result, nil
}
func hashPasswordWithSalt(password string, salt []byte) (string, error) {
	time := uint32(1)
	memory := uint32(64 * 1024)
	threads := uint8(runtime.NumCPU())
	keyLen := uint32(64)

	buffer := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLen)
	encodedPass := base64.RawStdEncoding.EncodeToString(buffer)
	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)

	result := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, memory, time, threads, encodedSalt, encodedPass)

	return result, nil
}

func VerifyPassword(password, hashedPassword string) (bool, error) {
	version, memory, time, threads, salt, keyLen, err := extractHashedPassword(hashedPassword)

	if err != nil {
		return false, fmt.Errorf("failed to extract hashed password: %w", err)
	}

	result, err := hashPasswordWithParams(password, version, salt, time, memory, threads, keyLen)

	if err != nil {
		return false, fmt.Errorf("failed to hash password with params: %w", err)
	}

	if result == hashedPassword {
		return true, nil
	}

	return false, nil
}
func extractHashedPassword(hashedPassword string) (version int, memory, time uint32, threads uint8, salt []byte, keyLen uint32, err error) {
	var encodedPass, encodedSalt string

	fmt.Sscanf(hashedPassword, "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		&version, &memory, &time, &threads, &encodedSalt, &encodedPass)

	salt, err = base64.RawStdEncoding.DecodeString(encodedSalt)

	if err != nil {
		return 0, 0, 0, 0, nil, 0, fmt.Errorf("failed to decode salt: %w", err)
	}

	bufferPass, err := base64.RawStdEncoding.DecodeString(encodedPass)

	if err != nil {
		return 0, 0, 0, 0, nil, 0, fmt.Errorf("failed to decode password: %w", err)
	}

	keyLen = uint32(len(bufferPass))

	return
}
