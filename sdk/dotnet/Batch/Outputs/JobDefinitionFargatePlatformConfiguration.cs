// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Batch.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties-fargateplatformconfiguration.html
    /// </summary>
    [OutputType]
    public sealed class JobDefinitionFargatePlatformConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties-fargateplatformconfiguration.html#cfn-batch-jobdefinition-containerproperties-fargateplatformconfiguration-platformversion
        /// </summary>
        public readonly string? PlatformVersion;

        [OutputConstructor]
        private JobDefinitionFargatePlatformConfiguration(string? platformVersion)
        {
            PlatformVersion = platformVersion;
        }
    }
}