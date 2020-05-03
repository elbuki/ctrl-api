package bcrypt_test

import (
	"testing"

	"github.com/elbuki/ctrl-api/src/bcrypt"
)

type scenario struct {
	encrytor   bcrypt.Encryptor
	passphrase []byte
	expectErr  bool
}

var scenarios = []scenario{
	scenario{bcrypt.NewEncryptor(6), []byte("foo"), false},
	scenario{bcrypt.NewEncryptor(40), nil, true},
}

func TestEncryption(t *testing.T) {
	for i, s := range scenarios {
		b, err := s.encrytor.Generate(s.passphrase)
		if !s.expectErr && err != nil {
			t.Errorf("#%d: not expecting error but there was one: %v", i, err)
		}

		if s.expectErr && err == nil {
			t.Errorf("#%d: expecting an error but generate went fine", i)
		}

		if err := s.encrytor.Compare(b, []byte("test")); err == nil {
			t.Errorf("#%d: expecting an error on compare but it went ok", i)
		}

		err = s.encrytor.Compare(b, s.passphrase)

		if !s.expectErr && err != nil {
			t.Errorf("#%d: unexpected error on compare: %v", i, err)
		}

		if s.expectErr && err == nil {
			t.Errorf("#%d: hash is too short but the compare went fine", i)
		}
	}
}
