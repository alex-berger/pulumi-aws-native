// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Route53
{
    /// <summary>
    /// Resource Type definition for AWS::Route53::RecordSetGroup
    /// </summary>
    [Obsolete(@"RecordSetGroup is not yet supported by AWS Cloud Control API, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:route53:RecordSetGroup")]
    public partial class RecordSetGroup : Pulumi.CustomResource
    {
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        [Output("hostedZoneId")]
        public Output<string?> HostedZoneId { get; private set; } = null!;

        [Output("hostedZoneName")]
        public Output<string?> HostedZoneName { get; private set; } = null!;

        [Output("recordSets")]
        public Output<ImmutableArray<Outputs.RecordSetGroupRecordSet>> RecordSets { get; private set; } = null!;


        /// <summary>
        /// Create a RecordSetGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RecordSetGroup(string name, RecordSetGroupArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:route53:RecordSetGroup", name, args ?? new RecordSetGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RecordSetGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:route53:RecordSetGroup", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing RecordSetGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RecordSetGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RecordSetGroup(name, id, options);
        }
    }

    public sealed class RecordSetGroupArgs : Pulumi.ResourceArgs
    {
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("hostedZoneId")]
        public Input<string>? HostedZoneId { get; set; }

        [Input("hostedZoneName")]
        public Input<string>? HostedZoneName { get; set; }

        [Input("recordSets")]
        private InputList<Inputs.RecordSetGroupRecordSetArgs>? _recordSets;
        public InputList<Inputs.RecordSetGroupRecordSetArgs> RecordSets
        {
            get => _recordSets ?? (_recordSets = new InputList<Inputs.RecordSetGroupRecordSetArgs>());
            set => _recordSets = value;
        }

        public RecordSetGroupArgs()
        {
        }
    }
}