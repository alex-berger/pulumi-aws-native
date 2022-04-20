// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticache

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
	case "aws-native:elasticache:CacheCluster":
		r = &CacheCluster{}
	case "aws-native:elasticache:GlobalReplicationGroup":
		r = &GlobalReplicationGroup{}
	case "aws-native:elasticache:ParameterGroup":
		r = &ParameterGroup{}
	case "aws-native:elasticache:ReplicationGroup":
		r = &ReplicationGroup{}
	case "aws-native:elasticache:SecurityGroup":
		r = &SecurityGroup{}
	case "aws-native:elasticache:SecurityGroupIngress":
		r = &SecurityGroupIngress{}
	case "aws-native:elasticache:SubnetGroup":
		r = &SubnetGroup{}
	case "aws-native:elasticache:User":
		r = &User{}
	case "aws-native:elasticache:UserGroup":
		r = &UserGroup{}
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
		"elasticache",
		&module{version},
	)
}
