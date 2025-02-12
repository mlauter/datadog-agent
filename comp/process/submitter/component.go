// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

// Package submitter implements a component to submit collected data in the Process Agent to
// supported Datadog intakes.
package submitter

import (
	"time"

	"go.uber.org/fx"

	"github.com/DataDog/datadog-agent/comp/process/types"
	"github.com/DataDog/datadog-agent/pkg/util/fxutil"
)

// team: processes

// Component is the component type.
type Component interface {
	Submit(start time.Time, checkName string, payload *types.Payload)
	Start() error
	Stop()
}

// Mock implements mock-specific methods.
type Mock interface {
	Component
}

// Module defines the fx options for this component.
var Module = fxutil.Component(
	fx.Provide(newSubmitter),
)

// MockModule defines the fx options for the mock component.
var MockModule = fxutil.Component(
	fx.Provide(newMock),
)
