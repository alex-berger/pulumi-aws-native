// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Configuration
{
    /// <summary>
    /// Resource Type definition for AWS::Config::DeliveryChannel
    /// </summary>
    [Obsolete(@"DeliveryChannel is not yet supported by AWS Cloud Control API, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:configuration:DeliveryChannel")]
    public partial class DeliveryChannel : Pulumi.CustomResource
    {
        [Output("configSnapshotDeliveryProperties")]
        public Output<Outputs.DeliveryChannelConfigSnapshotDeliveryProperties?> ConfigSnapshotDeliveryProperties { get; private set; } = null!;

        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        [Output("s3BucketName")]
        public Output<string> S3BucketName { get; private set; } = null!;

        [Output("s3KeyPrefix")]
        public Output<string?> S3KeyPrefix { get; private set; } = null!;

        [Output("s3KmsKeyArn")]
        public Output<string?> S3KmsKeyArn { get; private set; } = null!;

        [Output("snsTopicARN")]
        public Output<string?> SnsTopicARN { get; private set; } = null!;


        /// <summary>
        /// Create a DeliveryChannel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DeliveryChannel(string name, DeliveryChannelArgs args, CustomResourceOptions? options = null)
            : base("aws-native:configuration:DeliveryChannel", name, args ?? new DeliveryChannelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DeliveryChannel(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:configuration:DeliveryChannel", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing DeliveryChannel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DeliveryChannel Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DeliveryChannel(name, id, options);
        }
    }

    public sealed class DeliveryChannelArgs : Pulumi.ResourceArgs
    {
        [Input("configSnapshotDeliveryProperties")]
        public Input<Inputs.DeliveryChannelConfigSnapshotDeliveryPropertiesArgs>? ConfigSnapshotDeliveryProperties { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("s3BucketName", required: true)]
        public Input<string> S3BucketName { get; set; } = null!;

        [Input("s3KeyPrefix")]
        public Input<string>? S3KeyPrefix { get; set; }

        [Input("s3KmsKeyArn")]
        public Input<string>? S3KmsKeyArn { get; set; }

        [Input("snsTopicARN")]
        public Input<string>? SnsTopicARN { get; set; }

        public DeliveryChannelArgs()
        {
        }
    }
}