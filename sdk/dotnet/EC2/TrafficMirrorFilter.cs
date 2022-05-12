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
    /// Resource Type definition for AWS::EC2::TrafficMirrorFilter
    /// </summary>
    [Obsolete(@"TrafficMirrorFilter is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:ec2:TrafficMirrorFilter")]
    public partial class TrafficMirrorFilter : Pulumi.CustomResource
    {
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("networkServices")]
        public Output<ImmutableArray<string>> NetworkServices { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.TrafficMirrorFilterTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a TrafficMirrorFilter resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TrafficMirrorFilter(string name, TrafficMirrorFilterArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:ec2:TrafficMirrorFilter", name, args ?? new TrafficMirrorFilterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TrafficMirrorFilter(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ec2:TrafficMirrorFilter", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing TrafficMirrorFilter resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TrafficMirrorFilter Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new TrafficMirrorFilter(name, id, options);
        }
    }

    public sealed class TrafficMirrorFilterArgs : Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("networkServices")]
        private InputList<string>? _networkServices;
        public InputList<string> NetworkServices
        {
            get => _networkServices ?? (_networkServices = new InputList<string>());
            set => _networkServices = value;
        }

        [Input("tags")]
        private InputList<Inputs.TrafficMirrorFilterTagArgs>? _tags;
        public InputList<Inputs.TrafficMirrorFilterTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.TrafficMirrorFilterTagArgs>());
            set => _tags = value;
        }

        public TrafficMirrorFilterArgs()
        {
        }
    }
}
