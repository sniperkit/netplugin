package drivers

import (
	"encoding/json"
	"fmt"

	"github.com/contiv/netplugin/core"
)

// implements the State interface for a network implemented using
// vlans with ovs. The state is stored as Json objects.
const (
	BASE_PATH           = "/contiv.io/"
	CFG_PATH            = BASE_PATH + "config/"
	NW_CFG_PATH_PREFIX  = CFG_PATH + "nets/"
	NW_CFG_PATH         = NW_CFG_PATH_PREFIX + "%s"
	EP_CFG_PATH_PREFIX  = CFG_PATH + "eps/"
	EP_CFG_PATH         = EP_CFG_PATH_PREFIX + "%s"
	OPER_PATH           = BASE_PATH + "oper/"
	NW_OPER_PATH_PREFIX = OPER_PATH + "nets/"
	NW_OPER_PATH        = NW_OPER_PATH_PREFIX + "%s"
	EP_OPER_PATH_PREFIX = OPER_PATH + "eps/"
	EP_OPER_PATH        = EP_OPER_PATH_PREFIX + "%s"
)

type OvsCfgNetworkState struct {
	StateDriver core.StateDriver `json:"-"`
	Id          string           `json:"id"`
}

func (s *OvsCfgNetworkState) Write() error {
	key := fmt.Sprintf(NW_CFG_PATH, s.Id)
	return s.StateDriver.WriteState(key, s, json.Marshal)
}

func (s *OvsCfgNetworkState) Read(id string) error {
	key := fmt.Sprintf(NW_CFG_PATH, id)
	return s.StateDriver.ReadState(key, s, json.Unmarshal)
}

func (s *OvsCfgNetworkState) Clear() error {
	key := fmt.Sprintf(NW_CFG_PATH, s.Id)
	return s.StateDriver.ClearState(key)
}

type OvsOperNetworkState struct {
	StateDriver core.StateDriver `json:"-"`
	Id          string           `json:"id"`
	EpCount     int              `json:"epCount"`
}

func (s *OvsOperNetworkState) Write() error {
	key := fmt.Sprintf(NW_OPER_PATH, s.Id)
	return s.StateDriver.WriteState(key, s, json.Marshal)
}

func (s *OvsOperNetworkState) Read(id string) error {
	key := fmt.Sprintf(NW_OPER_PATH, id)
	return s.StateDriver.ReadState(key, s, json.Unmarshal)
}

func (s *OvsOperNetworkState) Clear() error {
	key := fmt.Sprintf(NW_OPER_PATH, s.Id)
	return s.StateDriver.ClearState(key)
}