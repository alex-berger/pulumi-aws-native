// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFormation
{
    /// <summary>
    /// Resource Type definition for AWS::CloudFormation::WaitConditionHandle
    /// </summary>
    [Obsolete(@"WaitConditionHandle is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:cloudformation:WaitConditionHandle")]
    public partial class WaitConditionHandle : Pulumi.CustomResource
    {
        /// <summary>
        /// Create a WaitConditionHandle resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WaitConditionHandle(string name, WaitConditionHandleArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:cloudformation:WaitConditionHandle", name, args ?? new WaitConditionHandleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WaitConditionHandle(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:cloudformation:WaitConditionHandle", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing WaitConditionHandle resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WaitConditionHandle Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new WaitConditionHandle(name, id, options);
        }
    }

    public sealed class WaitConditionHandleArgs : Pulumi.ResourceArgs
    {
        public WaitConditionHandleArgs()
        {
        }
    }
}
