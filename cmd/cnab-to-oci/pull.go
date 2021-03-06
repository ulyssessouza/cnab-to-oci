package main

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/docker/cnab-to-oci/remotes"
	"github.com/docker/distribution/reference"
	"github.com/spf13/cobra"
)

type pullOptions struct {
	output    string
	targetRef string
}

func pullCmd() *cobra.Command {
	var opts pullOptions
	cmd := &cobra.Command{
		Use:  "pull <ref> [options]",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.targetRef = args[0]
			return runPull(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.output, "output", "o", "pulled.json", "output file")
	return cmd
}

func runPull(opts pullOptions) error {
	resolver := createResolver()
	ref, err := reference.ParseNormalizedNamed(opts.targetRef)
	if err != nil {
		return err
	}
	b, err := remotes.Pull(context.Background(), ref, resolver)
	if err != nil {
		return err
	}
	bytes, err := json.MarshalIndent(b, "", " ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(opts.output, bytes, 0644)
}
