/*
Real-time Online/Offline Charging System (OCS) for Telecom & ISP environments
Copyright (C) ITsysCOM GmbH

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>
*/

package ers

import (
	"fmt"

	"github.com/cgrates/cgrates/config"
	"github.com/cgrates/cgrates/utils"
)

type EventReader interface {
	Config() *config.EventReaderCfg // return it's configuration
	Init() error                    // subscribe the reader on the path
	Read() (*utils.CGREvent, error) // produce a single record in the events file
}

// NewEventReader instantiates the event reader based on configuration at index
func NewEventReader(rdrCfg *config.EventReaderCfg,
	rdrExit chan struct{}, appExit chan bool) (er EventReader, err error) {
	switch rdrCfg.Type {
	default:
		err = fmt.Errorf("unsupported reader type: <%s>", rdrCfg.Type)
	case utils.MetaFileCSV:
		return NewCSVFileER(rdrCfg, rdrExit, appExit)
	}
	return
}