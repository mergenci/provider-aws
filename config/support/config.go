/*
Copyright 2021 Upbound Inc.
*/

package support

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure adds configurations for the "support" group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_support_case", func(r *config.Resource) {
		r.Kind = "SupportCase" // "case" is a keyword in Go, so we rename to "supportcase".
	})
}
