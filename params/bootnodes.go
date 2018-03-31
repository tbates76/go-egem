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

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
// Egem Go Bootnodes
"enode://24c98339ee08f63f158cd579fbb02c540b4805ab1990439362dfd4ef7a19445d94215c476bbc0451dded04f2d1e95adceb8d08ea6d68b3bdda507e3220398edb@[154.20.195.166]:30666", //N0
"enode://5c27e3a28d776f456496e1d93f01004fcdce36ff2e88dbfe6748e1c1ab2f837119587bb5d8ef594e70b45c0af258010f3784603ea6708d4942a84b6316107c14@[154.20.195.166]:30666",
"enode://49c45b59d5a1b40a9bdc8dc4c2dc382f4926147e814dbd5d91af61e39dc9f1e5694d1a24d4ec332e06a678847fe88d1b84169af83d68dbb883fd455252d96d53@[45.32.43.8]:30666", //N2
"enode://1c61e7921907e96aeea73d2baa1e93dd0b21acc606ae6bfe53ce0cdda3f3c657afec753d4fce812ab2a5f60281792c82c9bb65cf78e9eb1fd4d7a9745bcd544a@[199.247.0.245]:30666", //N1
"enode://d41c47d11708490583305c57fb57acfbebb52c96e0c04b7f41456119a1bce40f1474cd6f43457b3d2bc7baf5de2de194c3d369045707290fa37c49d698050db9@[45.77.231.167]:30666", //N3
"enode://e2f2471622b308217e513951ce75bfce412fe6a541ba49f4fa87784bad2602201f0bd32838089d9469bc0325237261b45b7e47a59fbf570d915186d03e891de9@[46.105.132.9]:30303",
"enode://316d2c8c55dd6100ede975696dd7a5fe94fc3310b88fecb98bf74385ac5447c55d199322d9db74f1cee9d301cf97715e7f3bc0642309db07f52f6b3edd95b549@[45.76.63.206]:30666",
"enode://69ebb51eb3ec7ca212be8a75a98ebf6cd8728e07e149d55e3ef48adb68c3179cdaf18457aa82556fd68f4cacd51b6b717df2c4c49f03466ef3f52c904088a8d3@[207.148.117.17]:30666",
"enode://d92ddb5d7200c706d6c865b8c5f1e78e0810fb748750d858ed178dafc65155de983f22f23ac9b01c8e4574da5ecc7ae3d8d8c93d6c33e37c6a2e348770c0ab3e@[104.207.149.147]:30666",
"enode://46fd233e086dd6af9a63ee5ec7f09c3ce3dfc0c6a5022baae8c995ef44f5bf3e98959fb47e451778dead1fe46ad396771e0f3009e8d5ac534bfc7453185400b9@[45.63.43.108]:30666",
"enode://6646b3abd9826b7749a2a238dc35807707a62fccdfab806afb3214f67e41d5b514a8def552e9a017f11411980dc05e8f8845265937316dae381d453f39f39d60@[173.199.119.64]:30666",
"enode://26a0f2623c0aa463e3bff15ea0d444901fb754bb89ed2b41a0e41222abf7af3eb27e3e5c1eb336ccc2485291be8c8729e797ca79938d3bd943bf681398553865@[45.63.27.79]:30666",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// RoGem test network.
var TestnetBootnodes = []string{

}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{

}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{

}
