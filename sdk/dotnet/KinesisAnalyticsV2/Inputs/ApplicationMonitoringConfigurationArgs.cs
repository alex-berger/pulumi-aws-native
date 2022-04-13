// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisAnalyticsV2.Inputs
{

    public sealed class ApplicationMonitoringConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("configurationType", required: true)]
        public Input<string> ConfigurationType { get; set; } = null!;

        [Input("logLevel")]
        public Input<string>? LogLevel { get; set; }

        [Input("metricsLevel")]
        public Input<string>? MetricsLevel { get; set; }

        public ApplicationMonitoringConfigurationArgs()
        {
        }
    }
}
