package services

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/plattyp/addon/resources"
)

// MOCK ACCESSOR
type testUserAccessor struct {
	invoked int
}

func (t *testUserAccessor) CreateUser(planID int64, region string, herokuID string) (*resources.User, error) {
	t.invoked++
	return nil, nil
}

func (t *testUserAccessor) FetchUser(id int64) (*resources.User, error) {
	t.invoked++
	// Simulate not found user
	if id == 2 {
		return nil, errors.New("Not Found")
	}
	r := &resources.User{}
	r.ID = id
	return r, nil
}

func (t *testUserAccessor) UpdatePlan(id int64, planID int64) error {
	t.invoked++
	return nil
}

func (t *testUserAccessor) DeleteUser(id int64) error {
	t.invoked++
	return nil
}

func TestValidateSSOTokenReturns403IfTokenIsInvalid(t *testing.T) {
	u := testUserAccessor{}
	_, statusCode, err := ValidateSSOToken(&u, 1, "1234567", time.Now().Add(time.Duration(time.Duration(-2))*time.Hour).Unix())
	if statusCode != 403 {
		t.Errorf("Should have returned a 403, but returned: %d", statusCode)
	}
	if err != nil && err.Error() != "Token is invalid" {
		t.Errorf("Should have an error `Token is invalid` but returned: %s", err.Error())
	}
	if u.invoked != 0 {
		t.Errorf("User accessor should not have been invoked, but was: %d", u.invoked)
	}
}

func TestValidateSSOTokenReturns403IfCurrentTimeIsAheadOfTimestamp(t *testing.T) {
	currentSalt := os.Getenv("HEROKU_SSO_SALT")
	timestamp := time.Now().Add(time.Duration(time.Duration(-1)) * time.Hour).Unix()
	preToken := fmt.Sprintf("%d:%s:%d", 1, currentSalt, timestamp)
	h := sha1.New()
	h.Write([]byte(preToken))
	testToken := fmt.Sprintf("%x", h.Sum(nil))

	u := testUserAccessor{}
	_, statusCode, err := ValidateSSOToken(&u, 1, testToken, timestamp)
	if statusCode != 403 {
		t.Errorf("Should have returned a 403, but returned: %d", statusCode)
	}
	if err != nil && err.Error() != "Token has expired" {
		t.Errorf("Should have an error `Token has expired` but returned: %s", err.Error())
	}
	if u.invoked != 0 {
		t.Errorf("User accessor should not have been invoked, but was: %d", u.invoked)
	}
}

func TestValidateSSOTokenReturns404IfUserDoesntExist(t *testing.T) {
	currentSalt := os.Getenv("HEROKU_SSO_SALT")
	timestamp := time.Now().Add(time.Duration(time.Duration(1)) * time.Hour).Unix()
	preToken := fmt.Sprintf("%d:%s:%d", 2, currentSalt, timestamp)
	h := sha1.New()
	h.Write([]byte(preToken))
	testToken := fmt.Sprintf("%x", h.Sum(nil))

	u := testUserAccessor{}
	_, statusCode, err := ValidateSSOToken(&u, 2, testToken, timestamp)
	if statusCode != 404 {
		t.Errorf("Should have returned a 404, but returned: %d", statusCode)
	}
	if err != nil && err.Error() != "Not Found" {
		t.Errorf("Should have an error `Token has expired` but returned: %s", err.Error())
	}
	if u.invoked != 1 {
		t.Errorf("User accessor should have been invoked once, but was: %d", u.invoked)
	}
}

func TestValidateSSOTokenReturns200AndUserIfValidTokenAndUser(t *testing.T) {
	currentSalt := os.Getenv("HEROKU_SSO_SALT")
	timestamp := time.Now().Add(time.Duration(time.Duration(1)) * time.Hour).Unix()
	preToken := fmt.Sprintf("%d:%s:%d", 1, currentSalt, timestamp)
	h := sha1.New()
	h.Write([]byte(preToken))
	testToken := fmt.Sprintf("%x", h.Sum(nil))

	u := testUserAccessor{}
	user, statusCode, err := ValidateSSOToken(&u, 1, testToken, timestamp)
	if statusCode != 200 {
		t.Errorf("Should have returned a 200, but returned: %d", statusCode)
	}
	if err != nil {
		t.Errorf("Should have not returned an error, but did: %s", err.Error())
	}
	if u.invoked != 1 {
		t.Errorf("User accessor should have been invoked once, but was: %d", u.invoked)
	}
	if user == nil || user.ID != 1 {
		t.Errorf("Should have returned the User resource but did not")
	}
}
