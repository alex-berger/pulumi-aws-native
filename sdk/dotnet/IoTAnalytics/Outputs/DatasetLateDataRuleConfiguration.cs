// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTAnalytics.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-latedataruleconfiguration.html
    /// </summary>
    [OutputType]
    public sealed class DatasetLateDataRuleConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-latedataruleconfiguration.html#cfn-iotanalytics-dataset-latedataruleconfiguration-deltatimesessionwindowconfiguration
        /// </summary>
        public readonly Outputs.DatasetDeltaTimeSessionWindowConfiguration? DeltaTimeSessionWindowConfiguration;

        [OutputConstructor]
        private DatasetLateDataRuleConfiguration(Outputs.DatasetDeltaTimeSessionWindowConfiguration? deltaTimeSessionWindowConfiguration)
        {
            DeltaTimeSessionWindowConfiguration = deltaTimeSessionWindowConfiguration;
        }
    }
}