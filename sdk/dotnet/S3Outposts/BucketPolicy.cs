// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3Outposts
{
    /// <summary>
    /// Resource Type Definition for AWS::S3Outposts::BucketPolicy
    /// </summary>
    [AwsNativeResourceType("aws-native:s3outposts:BucketPolicy")]
    public partial class BucketPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the specified bucket.
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// A policy document containing permissions to add to the specified bucket.
        /// </summary>
        [Output("policyDocument")]
        public Output<object> PolicyDocument { get; private set; } = null!;


        /// <summary>
        /// Create a BucketPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BucketPolicy(string name, BucketPolicyArgs args, CustomResourceOptions? options = null)
            : base("aws-native:s3outposts:BucketPolicy", name, args ?? new BucketPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BucketPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:s3outposts:BucketPolicy", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing BucketPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BucketPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new BucketPolicy(name, id, options);
        }
    }

    public sealed class BucketPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the specified bucket.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// A policy document containing permissions to add to the specified bucket.
        /// </summary>
        [Input("policyDocument", required: true)]
        public Input<object> PolicyDocument { get; set; } = null!;

        public BucketPolicyArgs()
        {
        }
    }
}
