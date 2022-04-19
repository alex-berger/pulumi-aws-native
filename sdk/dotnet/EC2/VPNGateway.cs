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
    /// Resource Type definition for AWS::EC2::VPNGateway
    /// </summary>
    [Obsolete(@"VPNGateway is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:ec2:VPNGateway")]
    public partial class VPNGateway : Pulumi.CustomResource
    {
        [Output("amazonSideAsn")]
        public Output<int?> AmazonSideAsn { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.VPNGatewayTag>> Tags { get; private set; } = null!;

        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a VPNGateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VPNGateway(string name, VPNGatewayArgs args, CustomResourceOptions? options = null)
            : base("aws-native:ec2:VPNGateway", name, args ?? new VPNGatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VPNGateway(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ec2:VPNGateway", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing VPNGateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VPNGateway Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VPNGateway(name, id, options);
        }
    }

    public sealed class VPNGatewayArgs : Pulumi.ResourceArgs
    {
        [Input("amazonSideAsn")]
        public Input<int>? AmazonSideAsn { get; set; }

        [Input("tags")]
        private InputList<Inputs.VPNGatewayTagArgs>? _tags;
        public InputList<Inputs.VPNGatewayTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.VPNGatewayTagArgs>());
            set => _tags = value;
        }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public VPNGatewayArgs()
        {
        }
    }
}
