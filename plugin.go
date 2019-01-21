package main

import (
	"github.com/jaegertracing/jaeger/plugin"
	"github.com/jaegertracing/jaeger/storage"
)

var factory = NewFactory()

// Export the needed symbols
var Configurable plugin.Configurable = factory
var StorageFactory storage.Factory = factory;
