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
    /// Resource Type definition for AWS::EC2::EIP
    /// </summary>
    [Obsolete(@"EIP is not yet supported by AWS Cloud Control API, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:ec2:EIP")]
    public partial class EIP : Pulumi.CustomResource
    {
        [Output("allocationId")]
        public Output<string> AllocationId { get; private set; } = null!;

        [Output("domain")]
        public Output<string?> Domain { get; private set; } = null!;

        [Output("instanceId")]
        public Output<string?> InstanceId { get; private set; } = null!;

        [Output("publicIpv4Pool")]
        public Output<string?> PublicIpv4Pool { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.EIPTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a EIP resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EIP(string name, EIPArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:ec2:EIP", name, args ?? new EIPArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EIP(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ec2:EIP", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing EIP resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EIP Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new EIP(name, id, options);
        }
    }

    public sealed class EIPArgs : Pulumi.ResourceArgs
    {
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        [Input("publicIpv4Pool")]
        public Input<string>? PublicIpv4Pool { get; set; }

        [Input("tags")]
        private InputList<Inputs.EIPTagArgs>? _tags;
        public InputList<Inputs.EIPTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.EIPTagArgs>());
            set => _tags = value;
        }

        public EIPArgs()
        {
        }
    }
}