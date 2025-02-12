// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

// Package processcheck implements a component to handle Process data collection in the Process Agent.
package processcheck

import (
	"go.uber.org/fx"

	"github.com/DataDog/datadog-agent/pkg/util/fxutil"
)

// team: processes

// Component is the component type.
type Component interface {
}

// Module defines the fx options for this component.
var Module = fxutil.Component(
	fx.Provide(newCheck),
)
