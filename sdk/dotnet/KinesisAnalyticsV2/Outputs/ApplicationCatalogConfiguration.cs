// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisAnalyticsV2.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-catalogconfiguration.html
    /// </summary>
    [OutputType]
    public sealed class ApplicationCatalogConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-catalogconfiguration.html#cfn-kinesisanalyticsv2-application-catalogconfiguration-gluedatacatalogconfiguration
        /// </summary>
        public readonly Outputs.ApplicationGlueDataCatalogConfiguration? GlueDataCatalogConfiguration;

        [OutputConstructor]
        private ApplicationCatalogConfiguration(Outputs.ApplicationGlueDataCatalogConfiguration? glueDataCatalogConfiguration)
        {
            GlueDataCatalogConfiguration = glueDataCatalogConfiguration;
        }
    }
}
