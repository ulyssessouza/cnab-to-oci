// Package cnabtooci provides converters to transform CNAB to OCI format and vice versa
// In `cmd/cnab-to-oci` you can find the `main` package to create the `cnab-to-oci` binary
// see https://github.com/docker/cnab-to-oci for more information about it.
//
// It can also be used as a library. For more, please refer to http://godoc.org/github.com/docker/cnab-to-oci
// The main functions are located in the remotes package, and they are:
//
// remotes.Pull in https://github.com/docker/cnab-to-oci/blob/master/remotes/pull.go
//
// remotes.Push in https://github.com/docker/cnab-to-oci/blob/master/remotes/push.go
//
// remotes.Fixup in https://github.com/docker/cnab-to-oci/blob/master/remotes/fixup.go
package cnabtooci
