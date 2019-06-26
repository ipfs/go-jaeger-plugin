package jaeger

import (
	"os"

	plugin "github.com/ipfs/go-ipfs/plugin"
	opentracing "github.com/opentracing/opentracing-go"
	config "github.com/uber/jaeger-client-go/config"
)

// Plugins is exported list of plugins that will be loaded
var Plugins = []plugin.Plugin{
	&jaegerPlugin{},
}

type jaegerPlugin struct{}

var _ plugin.PluginTracer = (*jaegerPlugin)(nil)

var tracerName = "#TRACER-NAME-NOT-SET"
var tracerEnv = "IPFS_TRACER_NAME"

func (*jaegerPlugin) Name() string {
	return "jaeger"
}

func (*jaegerPlugin) Version() string {
	return "0.0.1"
}

func (*jaegerPlugin) Init() error {
	maybeName := os.Getenv(tracerEnv)
	if maybeName != "" {
		tracerName = maybeName
	}
	return nil
}

//Initalize a Jaeger tracer and set it as the global tracer in opentracing api
func (*jaegerPlugin) InitTracer() (opentracing.Tracer, error) {
	tracerCfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}
	//we are ignoring the closer for now
	tracer, _, err := tracerCfg.New(tracerName)
	if err != nil {
		return nil, err
	}
	return tracer, nil
}
