// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::EC2::NetworkInterface resource creates network interface
type NetworkInterface struct {
	pulumi.CustomResourceState

	// A description for the network interface.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A list of security group IDs associated with this network interface.
	GroupSet pulumi.StringArrayOutput `pulumi:"groupSet"`
	// Indicates the type of network interface.
	InterfaceType pulumi.StringPtrOutput `pulumi:"interfaceType"`
	// The number of IPv6 addresses to assign to a network interface. Amazon EC2 automatically selects the IPv6 addresses from the subnet range. To specify specific IPv6 addresses, use the Ipv6Addresses property and don't specify this property.
	Ipv6AddressCount pulumi.IntPtrOutput `pulumi:"ipv6AddressCount"`
	// One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet to associate with the network interface. If you're specifying a number of IPv6 addresses, use the Ipv6AddressCount property and don't specify this property.
	Ipv6Addresses NetworkInterfaceInstanceIpv6AddressArrayOutput `pulumi:"ipv6Addresses"`
	// Returns the primary private IP address of the network interface.
	PrimaryPrivateIpAddress pulumi.StringOutput `pulumi:"primaryPrivateIpAddress"`
	// Assigns a single private IP address to the network interface, which is used as the primary private IP address. If you want to specify multiple private IP address, use the PrivateIpAddresses property.
	PrivateIpAddress pulumi.StringPtrOutput `pulumi:"privateIpAddress"`
	// Assigns a list of private IP addresses to the network interface. You can specify a primary private IP address by setting the value of the Primary property to true in the PrivateIpAddressSpecification property. If you want EC2 to automatically assign private IP addresses, use the SecondaryPrivateIpAddressCount property and do not specify this property.
	PrivateIpAddresses NetworkInterfacePrivateIpAddressSpecificationArrayOutput `pulumi:"privateIpAddresses"`
	// The number of secondary private IPv4 addresses to assign to a network interface. When you specify a number of secondary IPv4 addresses, Amazon EC2 selects these IP addresses within the subnet's IPv4 CIDR range. You can't specify this option and specify more than one private IP address using privateIpAddresses
	SecondaryPrivateIpAddressCount pulumi.IntPtrOutput `pulumi:"secondaryPrivateIpAddressCount"`
	// Returns the secondary private IP addresses of the network interface.
	SecondaryPrivateIpAddresses pulumi.StringArrayOutput `pulumi:"secondaryPrivateIpAddresses"`
	// Indicates whether traffic to or from the instance is validated.
	SourceDestCheck pulumi.BoolPtrOutput `pulumi:"sourceDestCheck"`
	// The ID of the subnet to associate with the network interface.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// An arbitrary set of tags (key-value pairs) for this network interface.
	Tags NetworkInterfaceTagArrayOutput `pulumi:"tags"`
}

// NewNetworkInterface registers a new resource with the given unique name, arguments, and options.
func NewNetworkInterface(ctx *pulumi.Context,
	name string, args *NetworkInterfaceArgs, opts ...pulumi.ResourceOption) (*NetworkInterface, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	var resource NetworkInterface
	err := ctx.RegisterResource("aws-native:ec2:NetworkInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkInterface gets an existing NetworkInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkInterfaceState, opts ...pulumi.ResourceOption) (*NetworkInterface, error) {
	var resource NetworkInterface
	err := ctx.ReadResource("aws-native:ec2:NetworkInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkInterface resources.
type networkInterfaceState struct {
}

type NetworkInterfaceState struct {
}

func (NetworkInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceState)(nil)).Elem()
}

type networkInterfaceArgs struct {
	// A description for the network interface.
	Description *string `pulumi:"description"`
	// A list of security group IDs associated with this network interface.
	GroupSet []string `pulumi:"groupSet"`
	// Indicates the type of network interface.
	InterfaceType *string `pulumi:"interfaceType"`
	// The number of IPv6 addresses to assign to a network interface. Amazon EC2 automatically selects the IPv6 addresses from the subnet range. To specify specific IPv6 addresses, use the Ipv6Addresses property and don't specify this property.
	Ipv6AddressCount *int `pulumi:"ipv6AddressCount"`
	// One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet to associate with the network interface. If you're specifying a number of IPv6 addresses, use the Ipv6AddressCount property and don't specify this property.
	Ipv6Addresses []NetworkInterfaceInstanceIpv6Address `pulumi:"ipv6Addresses"`
	// Assigns a single private IP address to the network interface, which is used as the primary private IP address. If you want to specify multiple private IP address, use the PrivateIpAddresses property.
	PrivateIpAddress *string `pulumi:"privateIpAddress"`
	// Assigns a list of private IP addresses to the network interface. You can specify a primary private IP address by setting the value of the Primary property to true in the PrivateIpAddressSpecification property. If you want EC2 to automatically assign private IP addresses, use the SecondaryPrivateIpAddressCount property and do not specify this property.
	PrivateIpAddresses []NetworkInterfacePrivateIpAddressSpecification `pulumi:"privateIpAddresses"`
	// The number of secondary private IPv4 addresses to assign to a network interface. When you specify a number of secondary IPv4 addresses, Amazon EC2 selects these IP addresses within the subnet's IPv4 CIDR range. You can't specify this option and specify more than one private IP address using privateIpAddresses
	SecondaryPrivateIpAddressCount *int `pulumi:"secondaryPrivateIpAddressCount"`
	// Indicates whether traffic to or from the instance is validated.
	SourceDestCheck *bool `pulumi:"sourceDestCheck"`
	// The ID of the subnet to associate with the network interface.
	SubnetId string `pulumi:"subnetId"`
	// An arbitrary set of tags (key-value pairs) for this network interface.
	Tags []NetworkInterfaceTag `pulumi:"tags"`
}

// The set of arguments for constructing a NetworkInterface resource.
type NetworkInterfaceArgs struct {
	// A description for the network interface.
	Description pulumi.StringPtrInput
	// A list of security group IDs associated with this network interface.
	GroupSet pulumi.StringArrayInput
	// Indicates the type of network interface.
	InterfaceType pulumi.StringPtrInput
	// The number of IPv6 addresses to assign to a network interface. Amazon EC2 automatically selects the IPv6 addresses from the subnet range. To specify specific IPv6 addresses, use the Ipv6Addresses property and don't specify this property.
	Ipv6AddressCount pulumi.IntPtrInput
	// One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet to associate with the network interface. If you're specifying a number of IPv6 addresses, use the Ipv6AddressCount property and don't specify this property.
	Ipv6Addresses NetworkInterfaceInstanceIpv6AddressArrayInput
	// Assigns a single private IP address to the network interface, which is used as the primary private IP address. If you want to specify multiple private IP address, use the PrivateIpAddresses property.
	PrivateIpAddress pulumi.StringPtrInput
	// Assigns a list of private IP addresses to the network interface. You can specify a primary private IP address by setting the value of the Primary property to true in the PrivateIpAddressSpecification property. If you want EC2 to automatically assign private IP addresses, use the SecondaryPrivateIpAddressCount property and do not specify this property.
	PrivateIpAddresses NetworkInterfacePrivateIpAddressSpecificationArrayInput
	// The number of secondary private IPv4 addresses to assign to a network interface. When you specify a number of secondary IPv4 addresses, Amazon EC2 selects these IP addresses within the subnet's IPv4 CIDR range. You can't specify this option and specify more than one private IP address using privateIpAddresses
	SecondaryPrivateIpAddressCount pulumi.IntPtrInput
	// Indicates whether traffic to or from the instance is validated.
	SourceDestCheck pulumi.BoolPtrInput
	// The ID of the subnet to associate with the network interface.
	SubnetId pulumi.StringInput
	// An arbitrary set of tags (key-value pairs) for this network interface.
	Tags NetworkInterfaceTagArrayInput
}

func (NetworkInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceArgs)(nil)).Elem()
}

type NetworkInterfaceInput interface {
	pulumi.Input

	ToNetworkInterfaceOutput() NetworkInterfaceOutput
	ToNetworkInterfaceOutputWithContext(ctx context.Context) NetworkInterfaceOutput
}

func (*NetworkInterface) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterface)(nil)).Elem()
}

func (i *NetworkInterface) ToNetworkInterfaceOutput() NetworkInterfaceOutput {
	return i.ToNetworkInterfaceOutputWithContext(context.Background())
}

func (i *NetworkInterface) ToNetworkInterfaceOutputWithContext(ctx context.Context) NetworkInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceOutput)
}

type NetworkInterfaceOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterface)(nil)).Elem()
}

func (o NetworkInterfaceOutput) ToNetworkInterfaceOutput() NetworkInterfaceOutput {
	return o
}

func (o NetworkInterfaceOutput) ToNetworkInterfaceOutputWithContext(ctx context.Context) NetworkInterfaceOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInterfaceInput)(nil)).Elem(), &NetworkInterface{})
	pulumi.RegisterOutputType(NetworkInterfaceOutput{})
}
