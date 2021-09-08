// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.PinpointEmail.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpointemail-configurationseteventdestination-cloudwatchdestination.html
    /// </summary>
    [OutputType]
    public sealed class ConfigurationSetEventDestinationCloudWatchDestination
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpointemail-configurationseteventdestination-cloudwatchdestination.html#cfn-pinpointemail-configurationseteventdestination-cloudwatchdestination-dimensionconfigurations
        /// </summary>
        public readonly ImmutableArray<Outputs.ConfigurationSetEventDestinationDimensionConfiguration> DimensionConfigurations;

        [OutputConstructor]
        private ConfigurationSetEventDestinationCloudWatchDestination(ImmutableArray<Outputs.ConfigurationSetEventDestinationDimensionConfiguration> dimensionConfigurations)
        {
            DimensionConfigurations = dimensionConfigurations;
        }
    }
}
