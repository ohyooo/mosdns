/*
 * Copyright (C) 2020-2022, IrineSistiana
 *
 * This file is part of mosdns.
 *
 * mosdns is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * mosdns is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package main

import (
	"fmt"
	"github.com/IrineSistiana/mosdns/v4/coremain"
	"github.com/IrineSistiana/mosdns/v4/mlog"
	_ "github.com/IrineSistiana/mosdns/v4/plugin"
	_ "github.com/IrineSistiana/mosdns/v4/tools"
	"github.com/spf13/cobra"
	_ "net/http/pprof"
)

var (
	version = "v4.5.3-maintenance"
)

func init() {
	coremain.AddSubCmd(&cobra.Command{
		Use:   "version",
		Short: "Print out version info and exit.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version)
		},
	})
}

func main() {
	if err := coremain.Run(); err != nil {
		mlog.S().Fatal(err)
	}
}
