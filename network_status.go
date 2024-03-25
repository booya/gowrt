package gowrt

import (
	"encoding/json"
	"fmt"
)

type FlowControl struct {
	Autoneg                bool     `json:"autoneg"`
	Supported              []string `json:"supported"`
	LinkAdvertising        []string `json:"link-advertising"`
	LinkPartnerAdvertising []string `json:"link-partner-advertising"`
	Negotiated             []string `json:"negotiated"`
}

type NetworkStatistics struct {
	Collisions        int `json:"collisions"`
	RxFrameErrors     int `json:"rx_frame_errors"`
	TxCompressed      int `json:"tx_compressed"`
	Multicast         int `json:"multicast"`
	RxLengthErrors    int `json:"rx_length_errors"`
	TxDropped         int `json:"tx_dropped"`
	RxBytes           int `json:"rx_bytes"`
	RxMissedErrors    int `json:"rx_missed_errors"`
	TxErrors          int `json:"tx_errors"`
	RxCompressed      int `json:"rx_compressed"`
	RxOverErrors      int `json:"rx_over_errors"`
	TxFifoErrors      int `json:"tx_fifo_errors"`
	RxCrcErrors       int `json:"rx_crc_errors"`
	RxPackets         int `json:"rx_packets"`
	TxHeartbeatErrors int `json:"tx_heartbeat_errors"`
	RxDropped         int `json:"rx_dropped"`
	TxAbortedErrors   int `json:"tx_aborted_errors"`
	TxPackets         int `json:"tx_packets"`
	RxErrors          int `json:"rx_errors"`
	TxBytes           int `json:"tx_bytes"`
	TxWindowErrors    int `json:"tx_window_errors"`
	RxFifoErrors      int `json:"rx_fifo_errors"`
	TxCarrierErrors   int `json:"tx_carrier_errors"`
}

type NetworkStatus struct {
	External                   bool              `json:"external"`
	Present                    bool              `json:"present"`
	Type                       string            `json:"type"`
	Up                         bool              `json:"up"`
	Carrier                    bool              `json:"carrier"`
	AuthStatus                 bool              `json:"auth_status"`
	LinkAdvertising            []string          `json:"link-advertising"`
	LinkPartnerAdvertising     []string          `json:"link-partner-advertising"`
	LinkSupported              []string          `json:"link-supported"`
	Speed                      string            `json:"speed"`
	Autoneg                    bool              `json:"autoneg"`
	FlowControl                FlowControl       `json:"flow-control"`
	HwTcOffload                bool              `json:"hw-tc-offload"`
	Devtype                    string            `json:"devtype"`
	Mtu                        int               `json:"mtu"`
	Mtu6                       int               `json:"mtu6"`
	Macaddr                    string            `json:"macaddr"`
	Txqueuelen                 int               `json:"txqueuelen"`
	Ipv6                       bool              `json:"ipv6"`
	Ip6segmentrouting          bool              `json:"ip6segmentrouting"`
	Promisc                    bool              `json:"promisc"`
	Rpfilter                   int               `json:"rpfilter"`
	Acceptlocal                bool              `json:"acceptlocal"`
	Igmpversion                int               `json:"igmpversion"`
	Mldversion                 int               `json:"mldversion"`
	Neigh4reachabletime        int               `json:"neigh4reachabletime"`
	Neigh6reachabletime        int               `json:"neigh6reachabletime"`
	Neigh4gcstaletime          int               `json:"neigh4gcstaletime"`
	Neigh6gcstaletime          int               `json:"neigh6gcstaletime"`
	Neigh4locktime             int               `json:"neigh4locktime"`
	Dadtransmits               int               `json:"dadtransmits"`
	Multicast                  bool              `json:"multicast"`
	Sendredirects              bool              `json:"sendredirects"`
	DropV4UnicastInL2Multicast bool              `json:"drop_v4_unicast_in_l2_multicast"`
	DropV6UnicastInL2Multicast bool              `json:"drop_v6_unicast_in_l2_multicast"`
	DropGratuitousArp          bool              `json:"drop_gratuitous_arp"`
	DropUnsolicitedNa          bool              `json:"drop_unsolicited_na"`
	ArpAccept                  bool              `json:"arp_accept"`
	Gro                        bool              `json:"gro"`
	Statistics                 NetworkStatistics `json:"statistics"`
}

func (c *Client) GetNetworkStatus(id, iface string) (NetworkStatus, error) {
	var networkStatus NetworkStatus
	body := map[string]interface{}{"name": iface}
	call := NewRpcCall(id, "call", "network.device", "status", body)
	response, err := c.ApiCall(call)
	if err != nil {
		return networkStatus, fmt.Errorf("get network status: %s", err)
	}
	err = json.Unmarshal(response, &networkStatus)
	if err != nil {
		return networkStatus, fmt.Errorf("unmarshal network status: %s", err)
	}
	return networkStatus, nil
}
