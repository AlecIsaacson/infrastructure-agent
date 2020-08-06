// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
package agent

import (
	"github.com/newrelic/infrastructure-agent/pkg/backend/identityapi"
	"github.com/newrelic/infrastructure-agent/pkg/backend/state"
)

type identityRegisterService struct {
	client identityapi.RegisterClient
	state  state.RegisterSM
}

func NewIdentityRegisterService(
	client identityapi.RegisterClient,
	sm state.RegisterSM,
) *identityRegisterService {
	return &identityRegisterService{
		client: client,
		state:  sm,
	}
}
