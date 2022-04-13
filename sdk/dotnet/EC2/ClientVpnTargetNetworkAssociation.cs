// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2
{
    /// <summary>
    /// Resource Type definition for AWS::EC2::ClientVpnTargetNetworkAssociation
    /// </summary>
    [Obsolete(@"ClientVpnTargetNetworkAssociation is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:ec2:ClientVpnTargetNetworkAssociation")]
    public partial class ClientVpnTargetNetworkAssociation : Pulumi.CustomResource
    {
        [Output("clientVpnEndpointId")]
        public Output<string> ClientVpnEndpointId { get; private set; } = null!;

        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;


        /// <summary>
        /// Create a ClientVpnTargetNetworkAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClientVpnTargetNetworkAssociation(string name, ClientVpnTargetNetworkAssociationArgs args, CustomResourceOptions? options = null)
            : base("aws-native:ec2:ClientVpnTargetNetworkAssociation", name, args ?? new ClientVpnTargetNetworkAssociationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClientVpnTargetNetworkAssociation(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ec2:ClientVpnTargetNetworkAssociation", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ClientVpnTargetNetworkAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClientVpnTargetNetworkAssociation Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ClientVpnTargetNetworkAssociation(name, id, options);
        }
    }

    public sealed class ClientVpnTargetNetworkAssociationArgs : Pulumi.ResourceArgs
    {
        [Input("clientVpnEndpointId", required: true)]
        public Input<string> ClientVpnEndpointId { get; set; } = null!;

        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        public ClientVpnTargetNetworkAssociationArgs()
        {
        }
    }
}
