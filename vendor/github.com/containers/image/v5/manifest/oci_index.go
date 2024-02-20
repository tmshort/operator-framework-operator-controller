package manifest

import (
	"github.com/containers/image/v5/internal/manifest"
	imgspecv1 "github.com/opencontainers/image-spec/specs-go/v1"
)

// OCI1Index is just an alias for the OCI index type, but one which we can
// provide methods for.
type OCI1Index = manifest.OCI1IndexPublic

// OCI1IndexFromComponents creates an OCI1 image index instance from the
// supplied data.
func OCI1IndexFromComponents(components []imgspecv1.Descriptor, annotations map[string]string) *OCI1Index {
	return manifest.OCI1IndexPublicFromComponents(components, annotations)
}

// OCI1IndexClone creates a deep copy of the passed-in index.
func OCI1IndexClone(index *OCI1Index) *OCI1Index {
	return manifest.OCI1IndexPublicClone(index)
}

// OCI1IndexFromManifest creates an OCI1 manifest index instance from marshalled
// JSON, presumably generated by encoding a OCI1 manifest index.
func OCI1IndexFromManifest(manifestBlob []byte) (*OCI1Index, error) {
	return manifest.OCI1IndexPublicFromManifest(manifestBlob)
}