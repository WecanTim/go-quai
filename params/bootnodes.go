// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/dominant-strategies/go-quai/common"

// ColosseumBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the Colosseum test network.
var ColosseumBootnodes = []string{
	"enode://aa76ab25bd9657efc8371103649ad9fc099388f1b303b18ab9b5fbdd43c8bae1e039357eb6a5a3aa86679f58ed0231c88bd207393448a7226d63fdc4651698cb@35.202.3.128",  // us-central1-a
	"enode://7a80da93d6ba5299fccf4dc195c0a9efaee5ef3ef9ae00c7ae4db72f10cd26750b7acaf956a381f3101860b58c931b56a4524051440382f1be5e33dc247a9b52@35.184.84.143", // us-central1-a
	"enode://1ba881701345af5491ec9ced31d2c325bdd296219015856109e8405bfa51665c0dc91b85a93541ba2ae6294d0dee4ef9c8cab7364594234447b30237631e90c6@35.224.101.46", // us-central1-a
}

// GardenBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Garden test network.
var GardenBootnodes = []string{
	"enode://f353567741755a4b35afb4641618b34c7e3c1666b49c9e519e8336c124ba8c2c556f00e9edcd87599994e9674ebdea3503a5ff1948a257c1540fe9f8c0fedf62@34.22.72.242",  // asia-northeast3-a
	"enode://aa3d4daa8ea958c8a4fe3c56a7c5c0509754eec9b2e4eb45ac5fe9eb16bea442da0b50a0ac09e5fdf7711eb6b64df76dcd22e48d69f2f938d19fc60b9a8fd579@34.175.158.49", // europe-southwest1-a
	"enode://3ecc23f02242be9249ac8b3aa2217514151106c152aa34b92f4287551d85300f9708740cb28d5ee39df3aec08315d624d8311fffc4c5e7ea24bc7422fc801972@35.198.7.119",  // southamerica-east1-b
	"enode://13ed279b8013b61ef41466d4f07679fe17c407af4f6ae09b34042c72635eef314e9020f732a96a739abaa88f51108b76698f6b7335397be99418c371e5c3bcce@34.29.49.205",  // us-central1-a
}

// OrchardBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Orchard test network
var OrchardBootnodes = []string{
	"enode://f99403abcfbee37e3232e6bb4d7fda4d70496585a53857ccb4aff5ec32d0f27186b5097430d5806f20f2003f35cfec5d778598a3945d530f212b7072caab9b8a@35.188.17.207", // us-central1-b
	"enode://142e48e3c36e5fe21aebf2941f2e63eb4720febe67de17dd84baf010e33c76275567ede53674007ab2848eec53022cd0cb94bcbea10960ae93edb5497c8caa2a@104.198.69.162", // us-central1-a
	"enode://d6d27b273682f8abc7ffff04dc9006bd40f0a079a8ba24da351b714506bb82c1f106ff073fa04983345aef15c876c602209b48b37d8ee10dad581fd1d9db9263@34.23.150.43", // us-east1-c
}

// GalenaBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Galena test network
var GalenaBootnodes = []string{
	"enode://12fa0822887c7c8829c1d30e7aaf430587cc71827f5ab2028ee39acf02c113eb01a4c130258074d3c7868f68280f854036a8a51beb4df96af314f94c7f45966e@34.175.138.4", // europe-southwest1-a
	"enode://1dab5f7739dfd6fb037a676607fc86ecc76eff3906abaa8f8662e236eb01937b27db1309c89e0c3af807a64b491f401f52ff75daca49689c99cfe949c359623c@35.199.83.89", // southamerica-east1-b
	"enode://402a7cc416700f8717ff35416dfcb3c286921bc4261a9e8cc9948cfac9e3589c8db25fffcf7d6877ec0e7ad93c83e6bde93d4918923dfb9c24c1635a57c5aba2@34.64.122.86", // asia-northeast3-a
}

var V5Bootnodes = []string{}

const dnsPrefix = "enrtree://ALE24Z2TEZV2XK46RXVB6IIN5HB5WTI4F4SMAVLYCAQIUPU53RSUU@"

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	var net string
	switch genesis {
	case ColosseumGenesisHash:
		net = "colosseum"
	case GardenGenesisHash:
		net = "garden"
	case OrchardGenesisHash:
		net = "orchard"
	case GalenaGenesisHash:
		net = "galena"
	default:
		return ""
	}
	return dnsPrefix + common.NodeLocation.Name() + "." + net + ".quainodes.io"
}
