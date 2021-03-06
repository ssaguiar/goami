// Copyright 2014 Helton Marques
//
//	Use of this source code is governed by a LGPL
//	license that can be found in the LICENSE file.
//

package ami

import (
	"errors"
)

//	SIPnotify
//		Send a SIP notify
//
func SIPnotify(socket *Socket, actionID string, channel string, variable string) (map[string]string, error) {
	// verify channel and variable and action ID
	if len(channel) == 0 || len(actionID) == 0 || len(variable) == 0 {
		return nil, errors.New("Invalid parameters")
	}

	command := []string{
		"Action: SIPnotify",
		"\r\nActionID: ",
		actionID,
		"\r\nChannel: ",
		channel,
		"\r\nVariable: ",
		variable,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//  SIPPeers
//      Lists SIP peers in text format with details on current status.
//      Peerlist will follow as separate events, followed by a final event called PeerlistComplete
//
func SIPpeers(socket *Socket, actionID string) ([]map[string]string, error) {
	command, err := getCommand("SIPpeers", actionID)
	if err != nil {
		return nil, err
	}
	return getMessageList(socket, command, actionID,
		"PeerEntry", "PeerlistComplete")
}

//	SIPqualifypeer
//		Qualify SIP peers.
//
func SIPqualifypeer(socket *Socket, actionID string, peer string) (map[string]string, error) {
	// verify peer and variable and action ID
	if len(peer) == 0 || len(actionID) == 0 {
		return nil, errors.New("Invalid parameters")
	}

	command := []string{
		"Action: SIPqualifypeer",
		"\r\nActionID: ",
		actionID,
		"\r\nPeer: ",
		peer,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//  SIPshowpeer
//      Show one SIP peer with details on current status.
//
func SIPshowpeer(socket *Socket, actionID string, peer string) (map[string]string, error) {
	// verify peer and action ID
	if len(peer) == 0 || len(actionID) == 0 {
		return nil, errors.New("Invalid parameters")
	}

	command := []string{
		"Action: SIPshowpeer",
		"\r\nActionID: ",
		actionID,
		"\r\nPeer: ",
		peer,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	SIPshowregistry
//		Show SIP registrations (text format).
//
func SIPshowregistry(socket *Socket, actionID string) (map[string]string, error) {
	// action ID
	if len(actionID) == 0 {
		return nil, errors.New("Invalid parameters")
	}

	command := []string{
		"Action: SIPshowregistry",
		"\r\nActionID: ",
		actionID,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}
