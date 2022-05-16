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
    /// Resource Type definition for AWS::EC2::VPCPeeringConnection
    /// </summary>
    [Obsolete(@"VPCPeeringConnection is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:ec2:VPCPeeringConnection")]
    public partial class VPCPeeringConnection : Pulumi.CustomResource
    {
        [Output("peerOwnerId")]
        public Output<string?> PeerOwnerId { get; private set; } = null!;

        [Output("peerRegion")]
        public Output<string?> PeerRegion { get; private set; } = null!;

        [Output("peerRoleArn")]
        public Output<string?> PeerRoleArn { get; private set; } = null!;

        [Output("peerVpcId")]
        public Output<string> PeerVpcId { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.VPCPeeringConnectionTag>> Tags { get; private set; } = null!;

        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a VPCPeeringConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VPCPeeringConnection(string name, VPCPeeringConnectionArgs args, CustomResourceOptions? options = null)
            : base("aws-native:ec2:VPCPeeringConnection", name, args ?? new VPCPeeringConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VPCPeeringConnection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ec2:VPCPeeringConnection", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing VPCPeeringConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VPCPeeringConnection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VPCPeeringConnection(name, id, options);
        }
    }

    public sealed class VPCPeeringConnectionArgs : Pulumi.ResourceArgs
    {
        [Input("peerOwnerId")]
        public Input<string>? PeerOwnerId { get; set; }

        [Input("peerRegion")]
        public Input<string>? PeerRegion { get; set; }

        [Input("peerRoleArn")]
        public Input<string>? PeerRoleArn { get; set; }

        [Input("peerVpcId", required: true)]
        public Input<string> PeerVpcId { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.VPCPeeringConnectionTagArgs>? _tags;
        public InputList<Inputs.VPCPeeringConnectionTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.VPCPeeringConnectionTagArgs>());
            set => _tags = value;
        }

        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public VPCPeeringConnectionArgs()
        {
        }
    }
}