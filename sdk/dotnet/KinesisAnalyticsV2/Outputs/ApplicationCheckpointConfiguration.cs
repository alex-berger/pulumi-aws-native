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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-checkpointconfiguration.html
    /// </summary>
    [OutputType]
    public sealed class ApplicationCheckpointConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-checkpointconfiguration.html#cfn-kinesisanalyticsv2-application-checkpointconfiguration-checkpointinterval
        /// </summary>
        public readonly int? CheckpointInterval;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-checkpointconfiguration.html#cfn-kinesisanalyticsv2-application-checkpointconfiguration-checkpointingenabled
        /// </summary>
        public readonly bool? CheckpointingEnabled;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-checkpointconfiguration.html#cfn-kinesisanalyticsv2-application-checkpointconfiguration-configurationtype
        /// </summary>
        public readonly string ConfigurationType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-checkpointconfiguration.html#cfn-kinesisanalyticsv2-application-checkpointconfiguration-minpausebetweencheckpoints
        /// </summary>
        public readonly int? MinPauseBetweenCheckpoints;

        [OutputConstructor]
        private ApplicationCheckpointConfiguration(
            int? checkpointInterval,

            bool? checkpointingEnabled,

            string configurationType,

            int? minPauseBetweenCheckpoints)
        {
            CheckpointInterval = checkpointInterval;
            CheckpointingEnabled = checkpointingEnabled;
            ConfigurationType = configurationType;
            MinPauseBetweenCheckpoints = minPauseBetweenCheckpoints;
        }
    }
}
