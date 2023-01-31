// Copyright 2019 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"runtime"

	"github.com/google/go-containerregistry/pkg/gcrane"
	"github.com/spf13/cobra"
)

// NewCmdCopy creates a new cobra.Command for the copy subcommand.
func NewCmdCopy() *cobra.Command {
	recursive := false
	jobs := 1
	year := 0
	month := 0
	day := 0
	hour := 0
	minute := 0
	second := 0
	year2 := 0
	month2 := 0
	day2 := 0
	hour2 := 0
	minute2 := 0
	second2 := 0
	cmd := &cobra.Command{
		Use:     "copy SRC DST",
		Aliases: []string{"cp"},
		Short:   "Efficiently copy a remote image from src to dst",
		Args:    cobra.ExactArgs(2),
		RunE: func(cc *cobra.Command, args []string) error {
			src, dst := args[0], args[1]
			ctx := cc.Context()
			if recursive {
				return gcrane.CopyRepository(ctx, src, dst, year, month, day, hour, minute, second, year2, month2, day2, hour2, minute2, second2, gcrane.WithJobs(jobs), gcrane.WithUserAgent(userAgent()), gcrane.WithContext(ctx))
			}
			return gcrane.Copy(src, dst, gcrane.WithUserAgent(userAgent()), gcrane.WithContext(ctx))
		},
	}

	cmd.Flags().BoolVarP(&recursive, "recursive", "r", false, "Whether to recurse through repos")
	cmd.Flags().IntVarP(&jobs, "jobs", "j", runtime.GOMAXPROCS(0), "The maximum number of concurrent copies")
	cmd.Flags().IntVarP(&year, "year", "y", 0, "year")
	cmd.Flags().IntVarP(&month, "month", "o", 0, "month")
	cmd.Flags().IntVarP(&day, "day", "d", 0, "day")
	cmd.Flags().IntVarP(&hour, "hour", "g", 0, "hour")
	cmd.Flags().IntVarP(&minute, "minute", "m", 0, "minute")
	cmd.Flags().IntVarP(&second, "second", "s", 0, "second")
	cmd.Flags().IntVarP(&year2, "year2", "z", 0, "year")
	cmd.Flags().IntVarP(&month2, "month2", "p", 0, "month")
	cmd.Flags().IntVarP(&day2, "day2", "e", 0, "day")
	cmd.Flags().IntVarP(&hour2, "hour2", "i", 0, "hour")
	cmd.Flags().IntVarP(&minute2, "minute2", "n", 0, "minute")
	cmd.Flags().IntVarP(&second2, "second2", "t", 0, "second")

	return cmd
}
