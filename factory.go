package customdialroutingextension

import (
	"context"

	"github.com/atoulme/customdialroutingextension/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

func NewFactory() extension.Factory {
	return extension.NewFactory(metadata.Type, createDefaultConfig, createExtension, metadata.ExtensionStability)
}

func createDefaultConfig() component.Config {
	return &Config{}
}

func createExtension(_ context.Context, _ extension.Settings, config component.Config) (extension.Extension, error) {
	return &customDialer{cfg: config.(*Config)}, nil
}
