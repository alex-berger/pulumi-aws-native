// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dms

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
	case "aws-native:dms:Certificate":
		r = &Certificate{}
	case "aws-native:dms:Endpoint":
		r = &Endpoint{}
	case "aws-native:dms:EventSubscription":
		r = &EventSubscription{}
	case "aws-native:dms:ReplicationInstance":
		r = &ReplicationInstance{}
	case "aws-native:dms:ReplicationSubnetGroup":
		r = &ReplicationSubnetGroup{}
	case "aws-native:dms:ReplicationTask":
		r = &ReplicationTask{}
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
		"dms",
		&module{version},
	)
}
