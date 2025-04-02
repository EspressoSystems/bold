// Copyright 2023-2024, Offchain Labs, Inc.
// For license information, see:
// https://github.com/EspressoSystems/bold/blob/main/LICENSE.md

package solimpl

import protocol "github.com/EspressoSystems/bold/chain-abstraction"

var (
	_ = protocol.SpecEdge(&specEdge{})
	_ = protocol.SpecChallengeManager(&specChallengeManager{})
)
