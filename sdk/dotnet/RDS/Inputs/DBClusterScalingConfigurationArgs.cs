// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RDS.Inputs
{

    public sealed class DBClusterScalingConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("autoPause")]
        public Input<bool>? AutoPause { get; set; }

        [Input("maxCapacity")]
        public Input<int>? MaxCapacity { get; set; }

        [Input("minCapacity")]
        public Input<int>? MinCapacity { get; set; }

        [Input("secondsUntilAutoPause")]
        public Input<int>? SecondsUntilAutoPause { get; set; }

        public DBClusterScalingConfigurationArgs()
        {
        }
    }
}
