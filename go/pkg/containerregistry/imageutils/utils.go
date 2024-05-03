package imageutils

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/google/go-containerregistry/pkg/name"
	"github.com/pkg/errors"
)

var (
	// Based on https://docs.docker.com/registry/spec/api/#content-digests
	digestRegex   = regexp.MustCompile("([A-Za-z0-9_+.-]+:[A-Fa-f0-9]+)$")
	protocolRegex = regexp.MustCompile("^(https?://)")
)

func NormalizeRegistryUrl(regUrl string) (bool, string) {
	protocol := protocolRegex.FindString(regUrl)
	s := protocolRegex.ReplaceAllString(regUrl, "")
	s = strings.TrimRight(s, "/")
	return protocol == "http://", s
}

func NormalizeRepositoryName(repository string) string {
	return strings.TrimLeft(repository, "/")
}

var ErrNoIdentifier = errors.Errorf("could not parse a valid tag or digest from image reference")

type Identifier interface {
	fmt.Stringer
	Name() string
	IsTag() bool
	IsDigest() bool
}

type Tag string

func (t Tag) Name() string   { return string(t) }
func (Tag) IsTag() bool      { return true }
func (Tag) IsDigest() bool   { return false }
func (t Tag) String() string { return "tag<" + string(t) + ">" }

type Digest string

func (d Digest) Name() string   { return string(d) }
func (Digest) IsTag() bool      { return false }
func (Digest) IsDigest() bool   { return true }
func (d Digest) String() string { return "digest<" + string(d) + ">" }

type ImageRef struct {
	Registry   string
	Repository string
	Identifier Identifier
}

func ParseReference(ref string) (*ImageRef, error) {
	parsed, err := name.ParseReference(ref)
	if err != nil {
		return nil, err
	}

	var id Identifier
	if tag, ok := parsed.(name.Tag); ok {
		id = Tag(tag.Identifier())
	} else if digest, ok := parsed.(name.Digest); ok {
		id = Digest(digest.Identifier())
	} else {
		return nil, ErrNoIdentifier
	}

	return &ImageRef{
		Registry:   parsed.Context().RegistryStr(),
		Repository: parsed.Context().RepositoryStr(),
		Identifier: id,
	}, nil
}

func IsDigest(d string) bool {
	return digestRegex.MatchString(d)
}

func Join(registry, repository, identifier string) string {
	delim := ":"
	if IsDigest(identifier) {
		delim = "@"
	}
	_, regUrl := NormalizeRegistryUrl(registry)
	return fmt.Sprintf(
		"%[1]s/%[2]s%[3]s%[4]s",
		regUrl,                              // 1
		NormalizeRepositoryName(repository), // 2
		delim,                               // 3
		identifier,                          //4
	)
}
