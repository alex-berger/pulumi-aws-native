// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkmanager

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
	case "aws-native:networkmanager:ConnectAttachment":
		r = &ConnectAttachment{}
	case "aws-native:networkmanager:ConnectPeer":
		r = &ConnectPeer{}
	case "aws-native:networkmanager:CoreNetwork":
		r = &CoreNetwork{}
	case "aws-native:networkmanager:CustomerGatewayAssociation":
		r = &CustomerGatewayAssociation{}
	case "aws-native:networkmanager:Device":
		r = &Device{}
	case "aws-native:networkmanager:GlobalNetwork":
		r = &GlobalNetwork{}
	case "aws-native:networkmanager:Link":
		r = &Link{}
	case "aws-native:networkmanager:LinkAssociation":
		r = &LinkAssociation{}
	case "aws-native:networkmanager:Site":
		r = &Site{}
	case "aws-native:networkmanager:SiteToSiteVpnAttachment":
		r = &SiteToSiteVpnAttachment{}
	case "aws-native:networkmanager:TransitGatewayRegistration":
		r = &TransitGatewayRegistration{}
	case "aws-native:networkmanager:VpcAttachment":
		r = &VpcAttachment{}
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
		"networkmanager",
		&module{version},
	)
}
