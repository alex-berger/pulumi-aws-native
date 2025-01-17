// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudwatch

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "aws-native:cloudwatch:Alarm":
		r = &Alarm{}
	case "aws-native:cloudwatch:AnomalyDetector":
		r = &AnomalyDetector{}
	case "aws-native:cloudwatch:CompositeAlarm":
		r = &CompositeAlarm{}
	case "aws-native:cloudwatch:Dashboard":
		r = &Dashboard{}
	case "aws-native:cloudwatch:InsightRule":
		r = &InsightRule{}
	case "aws-native:cloudwatch:MetricStream":
		r = &MetricStream{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := aws.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"aws-native",
		"cloudwatch",
		&module{version},
	)
}
