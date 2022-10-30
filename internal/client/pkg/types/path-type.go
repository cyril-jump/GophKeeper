package types

// PathType string
type PathType string

const (
	BlobPath     PathType = "/api/materials/blob"
	CardPath     PathType = "/api/materials/card"
	CredPath     PathType = "/api/materials/cred"
	TextPath     PathType = "/api/materials/text"
	LoginPath    PathType = "/api/user/login"
	RegisterPath PathType = "/api/user/register"
)

func (c PathType) String() string {
	return string(c)
}
