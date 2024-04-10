// this package is useful if you are trying to write a CLI that uses the same auth session as pvnctl
package session

import (
	"os"
	"sync"

	auth_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/auth"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

var (
	homeDir, _      = os.UserHomeDir()
	sessionDir      = homeDir + "/.config/pvn/"
	sessionPath     = sessionDir + "session.pb"
	session         *auth_pb.Session
	loadSessionOnce sync.Once
	// these are separated out from session object so that the process can choose to
	// override it without writing to session
	InProcessContext      string
	InProcessAdminContext string
)

func GetSession() *auth_pb.Session {
	loadSessionOnce.Do(func() {
		session = &auth_pb.Session{
			Contexts:      map[string]*auth_pb.AuthContext{},
			AdminContexts: map[string]*auth_pb.AuthContext{},
		}

		_, err := os.Stat(sessionPath)
		if os.IsNotExist(err) {
			return
		}
		if err != nil {
			panic(err)
		}

		sessionBytes, err := os.ReadFile(sessionPath)
		if err != nil {
			panic(err)
		}

		err = proto.Unmarshal(sessionBytes, session)
		if err != nil {
			panic(err)
		}

		if session.Contexts == nil {
			session.Contexts = map[string]*auth_pb.AuthContext{}
		}
		if session.AdminContexts == nil {
			session.AdminContexts = map[string]*auth_pb.AuthContext{}
		}

		InProcessContext = session.CurrentContext
		InProcessAdminContext = session.CurrentAdminContext
	})

	return session
}

func SaveSession(s *auth_pb.Session) error {
	bytes, err := proto.Marshal(s)
	if err != nil {
		return errors.Wrapf(err, "error marshaling session")
	}

	err = os.MkdirAll(sessionDir, os.FileMode(0755))
	if err != nil {
		return errors.Wrapf(err, "error making directories for session")
	}

	err = os.WriteFile(sessionPath, bytes, os.FileMode(0600))
	if err != nil {
		return errors.Wrapf(err, "error writing session")
	}
	return err
}
