package remotes

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/containerd/containerd/errdefs"
	"github.com/containerd/containerd/remotes"
	"github.com/deislabs/duffle/pkg/bundle"
	"github.com/docker/cnab-to-oci"
	"github.com/docker/distribution/reference"
	"github.com/opencontainers/go-digest"
	ocischemav1 "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/pkg/errors"
)

// ManifestOption is a callback used to customize a manifest before pushing it
type ManifestOption func(*ocischemav1.Index) error

// Push pushes a bundle as an OCI Image Index manifest
func Push(ctx context.Context, b *bundle.Bundle, ref reference.Named, resolver remotes.Resolver, options ...ManifestOption) (ocischemav1.Descriptor, error) {
	confDescriptor, confPayload, err := prepareConfig(b)
	if err != nil {
		return ocischemav1.Descriptor{}, err
	}
	indexDescriptor, indexPayload, err := prepareIndex(b, ref, confDescriptor, options...)
	if err != nil {
		return ocischemav1.Descriptor{}, err
	}
	// Push the bundle config
	if err := pushPayload(ctx, resolver, ref.Name(), confDescriptor, confPayload); err != nil {
		return ocischemav1.Descriptor{}, fmt.Errorf("error while pushing bundle config: %s", err)
	}
	// Push the bundle index
	if err := pushPayload(ctx, resolver, ref.String(), indexDescriptor, indexPayload); err != nil {
		return ocischemav1.Descriptor{}, fmt.Errorf("error while pushing bundle index: %s", err)
	}
	return indexDescriptor, nil
}

func prepareConfig(b *bundle.Bundle) (ocischemav1.Descriptor, []byte, error) {
	conf := oci.CreateBundleConfig(b)
	confPayload, err := json.Marshal(conf)
	if err != nil {
		return ocischemav1.Descriptor{}, nil, fmt.Errorf("invalid bundle config %q: %s", b.Name, err)
	}
	confDescriptor := ocischemav1.Descriptor{
		Digest:    digest.FromBytes(confPayload),
		MediaType: oci.BundleConfigMediaType,
		Size:      int64(len(confPayload)),
	}
	return confDescriptor, confPayload, nil
}

func prepareIndex(b *bundle.Bundle, ref reference.Named, confDescriptor ocischemav1.Descriptor, options ...ManifestOption) (ocischemav1.Descriptor, []byte, error) {
	ix, err := oci.ConvertBundleToOCIIndex(b, ref, confDescriptor)
	if err != nil {
		return ocischemav1.Descriptor{}, nil, err
	}
	for _, opts := range options {
		if err := opts(ix); err != nil {
			return ocischemav1.Descriptor{}, nil, fmt.Errorf("failed to prepare bundle manifest %q: %s", ref, err)
		}
	}

	indexPayload, err := json.Marshal(ix)
	if err != nil {
		return ocischemav1.Descriptor{}, nil, fmt.Errorf("invalid bundle manifest %q: %s", ref, err)
	}
	indexDescriptor := ocischemav1.Descriptor{
		Digest:    digest.FromBytes(indexPayload),
		MediaType: ocischemav1.MediaTypeImageIndex,
		Size:      int64(len(indexPayload)),
	}
	return indexDescriptor, indexPayload, nil
}

func pushPayload(ctx context.Context, resolver remotes.Resolver, reference string, descriptor ocischemav1.Descriptor, payload []byte) error {
	pusher, err := resolver.Pusher(ctx, reference)
	if err != nil {
		return err
	}
	writer, err := pusher.Push(ctx, descriptor)
	if err != nil {
		if errors.Cause(err) == errdefs.ErrAlreadyExists {
			return nil
		}
		return err
	}
	defer writer.Close()
	if _, err := writer.Write(payload); err != nil {
		if errors.Cause(err) == errdefs.ErrAlreadyExists {
			return nil
		}
		return err
	}
	err = writer.Commit(ctx, descriptor.Size, descriptor.Digest)
	if errors.Cause(err) == errdefs.ErrAlreadyExists {
		return nil
	}
	return err

}