package services

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/plattyp/addon/accessor"
	"github.com/plattyp/addon/resources"
)

// ValidateSSOToken is used to determine if an SSO token is valid
func ValidateSSOToken(userAccessor accessor.UserAccessor, userID int64, token string, timestamp int64) (*resources.User, int, error) {
	herokuSSOSalt := os.Getenv("HEROKU_SSO_SALT")
	preToken := fmt.Sprintf("%d:%s:%d", userID, herokuSSOSalt, timestamp)
	h := sha1.New()
	h.Write([]byte(preToken))
	encodedToken := fmt.Sprintf("%x", h.Sum(nil))

	// Validate Token
	if token != encodedToken {
		return nil, 403, errors.New("Token is invalid")
	}

	// Validate Timestamp
	currentTime := time.Now().UTC().Unix()
	if currentTime > timestamp {
		return nil, 403, errors.New("Token has expired")
	}

	// Lookup User
	user, err := userAccessor.FetchUser(userID)
	if err != nil {
		return nil, 404, err
	}

	// Successful Authentication
	return user, 200, nil
}
