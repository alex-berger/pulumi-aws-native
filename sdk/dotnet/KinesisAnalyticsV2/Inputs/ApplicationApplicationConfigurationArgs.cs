// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisAnalyticsV2.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationconfiguration.html
    /// </summary>
    public sealed class ApplicationApplicationConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationconfiguration.html#cfn-kinesisanalyticsv2-application-applicationconfiguration-applicationcodeconfiguration
        /// </summary>
        [Input("applicationCodeConfiguration")]
        public Input<Inputs.ApplicationApplicationCodeConfigurationArgs>? ApplicationCodeConfiguration { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationconfiguration.html#cfn-kinesisanalyticsv2-application-applicationconfiguration-applicationsnapshotconfiguration
        /// </summary>
        [Input("applicationSnapshotConfiguration")]
        public Input<Inputs.ApplicationApplicationSnapshotConfigurationArgs>? ApplicationSnapshotConfiguration { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationconfiguration.html#cfn-kinesisanalyticsv2-application-applicationconfiguration-environmentproperties
        /// </summary>
        [Input("environmentProperties")]
        public Input<Inputs.ApplicationEnvironmentPropertiesArgs>? EnvironmentProperties { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationconfiguration.html#cfn-kinesisanalyticsv2-application-applicationconfiguration-flinkapplicationconfiguration
        /// </summary>
        [Input("flinkApplicationConfiguration")]
        public Input<Inputs.ApplicationFlinkApplicationConfigurationArgs>? FlinkApplicationConfiguration { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationconfiguration.html#cfn-kinesisanalyticsv2-application-applicationconfiguration-sqlapplicationconfiguration
        /// </summary>
        [Input("sqlApplicationConfiguration")]
        public Input<Inputs.ApplicationSqlApplicationConfigurationArgs>? SqlApplicationConfiguration { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationconfiguration.html#cfn-kinesisanalyticsv2-application-applicationconfiguration-zeppelinapplicationconfiguration
        /// </summary>
        [Input("zeppelinApplicationConfiguration")]
        public Input<Inputs.ApplicationZeppelinApplicationConfigurationArgs>? ZeppelinApplicationConfiguration { get; set; }

        public ApplicationApplicationConfigurationArgs()
        {
        }
    }
}
