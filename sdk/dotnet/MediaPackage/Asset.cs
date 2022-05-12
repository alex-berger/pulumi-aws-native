// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaPackage
{
    /// <summary>
    /// Resource schema for AWS::MediaPackage::Asset
    /// </summary>
    [AwsNativeResourceType("aws-native:mediapackage:Asset")]
    public partial class Asset : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the Asset.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The time the Asset was initially submitted for Ingest.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The list of egress endpoints available for the Asset.
        /// </summary>
        [Output("egressEndpoints")]
        public Output<ImmutableArray<Outputs.AssetEgressEndpoint>> EgressEndpoints { get; private set; } = null!;

        /// <summary>
        /// The ID of the PackagingGroup for the Asset.
        /// </summary>
        [Output("packagingGroupId")]
        public Output<string> PackagingGroupId { get; private set; } = null!;

        /// <summary>
        /// The resource ID to include in SPEKE key requests.
        /// </summary>
        [Output("resourceId")]
        public Output<string?> ResourceId { get; private set; } = null!;

        /// <summary>
        /// ARN of the source object in S3.
        /// </summary>
        [Output("sourceArn")]
        public Output<string> SourceArn { get; private set; } = null!;

        /// <summary>
        /// The IAM role_arn used to access the source S3 bucket.
        /// </summary>
        [Output("sourceRoleArn")]
        public Output<string> SourceRoleArn { get; private set; } = null!;

        /// <summary>
        /// A collection of tags associated with a resource
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.AssetTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Asset resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Asset(string name, AssetArgs args, CustomResourceOptions? options = null)
            : base("aws-native:mediapackage:Asset", name, args ?? new AssetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Asset(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:mediapackage:Asset", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Asset resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Asset Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Asset(name, id, options);
        }
    }

    public sealed class AssetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the PackagingGroup for the Asset.
        /// </summary>
        [Input("packagingGroupId", required: true)]
        public Input<string> PackagingGroupId { get; set; } = null!;

        /// <summary>
        /// The resource ID to include in SPEKE key requests.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        /// <summary>
        /// ARN of the source object in S3.
        /// </summary>
        [Input("sourceArn", required: true)]
        public Input<string> SourceArn { get; set; } = null!;

        /// <summary>
        /// The IAM role_arn used to access the source S3 bucket.
        /// </summary>
        [Input("sourceRoleArn", required: true)]
        public Input<string> SourceRoleArn { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.AssetTagArgs>? _tags;

        /// <summary>
        /// A collection of tags associated with a resource
        /// </summary>
        public InputList<Inputs.AssetTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.AssetTagArgs>());
            set => _tags = value;
        }

        public AssetArgs()
        {
        }
    }
}
