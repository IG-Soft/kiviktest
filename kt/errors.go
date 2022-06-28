package kt

import (
	"net/http"

	kivik "github.com/IG-Soft/kivik/v3"
)

// CheckError compares the error's status code with that expected.
func (c *Context) CheckError(err error) (match bool, success bool) {
	c.T.Helper()
	status := c.Int("status")
	if status == 0 && err == nil {
		return true, true
	}
	switch actualStatus := kivik.StatusCode(err); actualStatus {
	case status:
		// This is expected
		return true, status == 0
	case 0:
		c.Errorf("Expected failure %d/%s, got success", status, http.StatusText(status))
		return false, true
	default:
		if status == 0 {
			c.Errorf("Unexpected failure: %d/%s", kivik.StatusCode(err), err)
			return false, false
		}
		c.Errorf("Unexpected failure state.\nExpected: %d/%s\n  Actual: %d/%s", status, http.StatusText(status), actualStatus, err)
		return false, false
	}
}

// IsExpected checks the error against the expected status, and returns true
// if they match.
func (c *Context) IsExpected(err error) bool {
	c.T.Helper()
	m, _ := c.CheckError(err)
	return m
}

// IsSuccess is similar to IsExpected, except for its return value. This method
// returns true if the expected status == 0, regardless of the error.
func (c *Context) IsSuccess(err error) bool {
	c.T.Helper()
	_, s := c.CheckError(err)
	return s
}

// IsExpectedSuccess combines IsExpected() and IsSuccess(), returning true only
// if there is no error, and no error was expected.
func (c *Context) IsExpectedSuccess(err error) bool {
	c.T.Helper()
	m, s := c.CheckError(err)
	return m && s
}
