// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaStore.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediastore-container-corsrule.html
    /// </summary>
    [OutputType]
    public sealed class ContainerCorsRule
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediastore-container-corsrule.html#cfn-mediastore-container-corsrule-allowedheaders
        /// </summary>
        public readonly ImmutableArray<string> AllowedHeaders;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediastore-container-corsrule.html#cfn-mediastore-container-corsrule-allowedmethods
        /// </summary>
        public readonly ImmutableArray<string> AllowedMethods;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediastore-container-corsrule.html#cfn-mediastore-container-corsrule-allowedorigins
        /// </summary>
        public readonly ImmutableArray<string> AllowedOrigins;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediastore-container-corsrule.html#cfn-mediastore-container-corsrule-exposeheaders
        /// </summary>
        public readonly ImmutableArray<string> ExposeHeaders;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediastore-container-corsrule.html#cfn-mediastore-container-corsrule-maxageseconds
        /// </summary>
        public readonly int? MaxAgeSeconds;

        [OutputConstructor]
        private ContainerCorsRule(
            ImmutableArray<string> allowedHeaders,

            ImmutableArray<string> allowedMethods,

            ImmutableArray<string> allowedOrigins,

            ImmutableArray<string> exposeHeaders,

            int? maxAgeSeconds)
        {
            AllowedHeaders = allowedHeaders;
            AllowedMethods = allowedMethods;
            AllowedOrigins = allowedOrigins;
            ExposeHeaders = exposeHeaders;
            MaxAgeSeconds = maxAgeSeconds;
        }
    }
}