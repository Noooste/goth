package twitter_test

import (
	"testing"

	"github.com/Noooste/goth"
	"github.com/Noooste/goth/providers/twitter"
	"github.com/stretchr/testify/assert"
)

func Test_Implements_Session(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	s := &twitter.Session{}

	a.Implements((*goth.Session)(nil), s)
}

func Test_GetAuthURL(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	s := &twitter.Session{}

	_, err := s.GetAuthURL()
	a.Error(err)

	s.AuthURL = "/foo"

	url, _ := s.GetAuthURL()
	a.Equal(url, "/foo")
}

func Test_ToJSON(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	s := &twitter.Session{}

	data := s.Marshal()
	a.Equal(data, `{"AuthURL":"","AccessToken":null,"RequestToken":null}`)
}

func Test_String(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	s := &twitter.Session{}

	a.Equal(s.String(), s.Marshal())
}
