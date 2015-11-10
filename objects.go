package models

import("encoding/json"
		  "fmt"
      )

type ConfigObj interface {
	     UnmarshalObject(data []byte) (ConfigObj, error)
}

//
// This file is handcoded for now. Eventually this would be generated by yang compiler
//
type IPV4Route struct {
	DestinationNw string
	NetworkMask string
	Cost int
	NextHopIp string
	OutgoingInterface  string
	Protocol string
}

func (obj IPV4Route) UnmarshalObject(body []byte) (ConfigObj, error) {
	var v4Route IPV4Route
	 var err  error
	if err = json.Unmarshal(body, &v4Route); err != nil  {
		fmt.Println("### IPV4Route Create is called", v4Route)
	}
	 fmt.Println("### IPV4Route Create is Unmarshal Object", err, v4Route)
	 return v4Route, err
}

type GlobalConfig struct {
	AS uint32
	RouterId string
}

type GlobalState struct {
	AS uint32
	RouterId string
	TotalPaths uint32
	TotalPrefixes uint32
}

type PeerType int

const (
	PeerTypeInternal PeerType = iota
	PeerTypeExternal
)

type BgpCounters struct {
	Update uint64
	Notification uint64
}

type Messages struct {
	Sent BgpCounters
	Received BgpCounters
}

type Queues struct {
	Input uint32
	Output uint32
}

type NeighborConfig struct {
	PeerAS uint32
	LocalAS uint32
	AuthPassword string
	Description string
	NeighborAddress string
}

type NeighborState struct {
	PeerAS uint32
	LocalAS uint32
	PeerType PeerType
	AuthPassword string
	Description string
	NeighborAddress string
	SessionState uint32
	Messages Messages
	Queues Queues
}

/* Start - Asicd objects */
type Vlan struct {
    VlanId    int32
    Ports     string
    PortTagType string
}

/* FIXME : RouterIf needs to be replaced by generic
 * layer 2 object name e.x Port-21 or Vlan-5 etc.
 * Internally this l2 object name can be translated
 * into appropriate key.
 */
type IPv4Intf struct {
    IpAddr   string
    RouterIf int32
}

type IPv4Neighbor struct {
    IpAddr   string
    MacAddr  string
    VlanId   int32
    RouterIf int32
}

func (obj Vlan) UnmarshalObject(body []byte) (ConfigObj, error) {
    var vlanObj Vlan
    var err error
    if err = json.Unmarshal(body, &vlanObj); err != nil  {
        fmt.Println("### Vlan create called, unmarshal failed", vlanObj)
    }
    return vlanObj, err
}

func (obj IPv4Intf) UnmarshalObject(body []byte) (ConfigObj, error) {
    var v4Intf IPv4Intf
    var err error
    if err = json.Unmarshal(body, &v4Intf); err != nil  {
        fmt.Println("### IPv4Intf create called, unmarshal failed", v4Intf)
    }
    return v4Intf, err
}

func (obj IPv4Neighbor) UnmarshalObject(body []byte) (ConfigObj, error) {
    var v4Nbr IPv4Neighbor
    var err error
    if err = json.Unmarshal(body, &v4Nbr); err != nil  {
        fmt.Println("### IPv4Neighbor create called, unmarshal failed", v4Nbr)
    }
    return v4Nbr, err
}
/* End - Asicd objects*/
