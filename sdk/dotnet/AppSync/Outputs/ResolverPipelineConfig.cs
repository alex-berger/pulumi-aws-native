// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppSync.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-resolver-pipelineconfig.html
    /// </summary>
    [OutputType]
    public sealed class ResolverPipelineConfig
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-resolver-pipelineconfig.html#cfn-appsync-resolver-pipelineconfig-functions
        /// </summary>
        public readonly ImmutableArray<string> Functions;

        [OutputConstructor]
        private ResolverPipelineConfig(ImmutableArray<string> functions)
        {
            Functions = functions;
        }
    }
}
