// Code generated by "stringer -type=NmDeviceType"; DO NOT EDIT

package gonetworkmanager

import "fmt"

const _NmDeviceType_name = "NmDeviceTypeUnknownNmDeviceTypeEthernetNmDeviceTypeWifiNmDeviceTypeUnused1NmDeviceTypeUnused2NmDeviceTypeBtNmDeviceTypeOlpcMeshNmDeviceTypeWimaxNmDeviceTypeModemNmDeviceTypeInfinibandNmDeviceTypeBondNmDeviceTypeVlanNmDeviceTypeAdslNmDeviceTypeBridgeNmDeviceTypeGenericNmDeviceTypeTeam"

var _NmDeviceType_index = [...]uint16{0, 19, 39, 55, 74, 93, 107, 127, 144, 161, 183, 199, 215, 231, 249, 268, 284}

func (i NmDeviceType) String() string {
	if i >= NmDeviceType(len(_NmDeviceType_index)-1) {
		return fmt.Sprintf("NmDeviceType(%d)", i)
	}
	return _NmDeviceType_name[_NmDeviceType_index[i]:_NmDeviceType_index[i+1]]
}