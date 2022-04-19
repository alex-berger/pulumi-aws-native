// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ECR
{
    /// <summary>
    /// The AWS::ECR::RegistryPolicy is used to specify permissions for another AWS account and is used when configuring cross-account replication. For more information, see Registry permissions in the Amazon Elastic Container Registry User Guide: https://docs.aws.amazon.com/AmazonECR/latest/userguide/registry-permissions.html
    /// </summary>
    [AwsNativeResourceType("aws-native:ecr:RegistryPolicy")]
    public partial class RegistryPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// The JSON policy text to apply to your registry. The policy text follows the same format as IAM policy text. For more information, see Registry permissions (https://docs.aws.amazon.com/AmazonECR/latest/userguide/registry-permissions.html) in the Amazon Elastic Container Registry User Guide.
        /// </summary>
        [Output("policyText")]
        public Output<object> PolicyText { get; private set; } = null!;

        [Output("registryId")]
        public Output<string> RegistryId { get; private set; } = null!;


        /// <summary>
        /// Create a RegistryPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RegistryPolicy(string name, RegistryPolicyArgs args, CustomResourceOptions? options = null)
            : base("aws-native:ecr:RegistryPolicy", name, args ?? new RegistryPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RegistryPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ecr:RegistryPolicy", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing RegistryPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RegistryPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RegistryPolicy(name, id, options);
        }
    }

    public sealed class RegistryPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The JSON policy text to apply to your registry. The policy text follows the same format as IAM policy text. For more information, see Registry permissions (https://docs.aws.amazon.com/AmazonECR/latest/userguide/registry-permissions.html) in the Amazon Elastic Container Registry User Guide.
        /// </summary>
        [Input("policyText", required: true)]
        public Input<object> PolicyText { get; set; } = null!;

        public RegistryPolicyArgs()
        {
        }
    }
}
