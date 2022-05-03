// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redshift

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
	case "aws-native:redshift:Cluster":
		r = &Cluster{}
	case "aws-native:redshift:ClusterParameterGroup":
		r = &ClusterParameterGroup{}
	case "aws-native:redshift:ClusterSecurityGroup":
		r = &ClusterSecurityGroup{}
	case "aws-native:redshift:ClusterSecurityGroupIngress":
		r = &ClusterSecurityGroupIngress{}
	case "aws-native:redshift:ClusterSubnetGroup":
		r = &ClusterSubnetGroup{}
	case "aws-native:redshift:EndpointAccess":
		r = &EndpointAccess{}
	case "aws-native:redshift:EndpointAuthorization":
		r = &EndpointAuthorization{}
	case "aws-native:redshift:EventSubscription":
		r = &EventSubscription{}
	case "aws-native:redshift:ScheduledAction":
		r = &ScheduledAction{}
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
		"redshift",
		&module{version},
	)
}
