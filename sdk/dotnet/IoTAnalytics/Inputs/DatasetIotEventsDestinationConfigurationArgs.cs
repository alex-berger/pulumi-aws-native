// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTAnalytics.Inputs
{

    public sealed class DatasetIotEventsDestinationConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("inputName", required: true)]
        public Input<string> InputName { get; set; } = null!;

        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        public DatasetIotEventsDestinationConfigurationArgs()
        {
        }
    }
}
