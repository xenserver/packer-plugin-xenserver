// Copyright (c) 2026. Citrix Systems, Inc.

package common

import "os"

type envFallback struct {
	field *string
	env   string
}

// applyEnvFallbacks treats an explicit "" as unset: these are plain strings, and an empty credential is unusable.
func applyEnvFallbacks(fallbacks []envFallback) {
	for _, f := range fallbacks {
		if *f.field == "" {
			*f.field = os.Getenv(f.env)
		}
	}
}
