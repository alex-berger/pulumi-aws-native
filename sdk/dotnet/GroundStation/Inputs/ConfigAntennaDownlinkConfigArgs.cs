// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GroundStation.Inputs
{

    public sealed class ConfigAntennaDownlinkConfigArgs : Pulumi.ResourceArgs
    {
        [Input("spectrumConfig")]
        public Input<Inputs.ConfigSpectrumConfigArgs>? SpectrumConfig { get; set; }

        public ConfigAntennaDownlinkConfigArgs()
        {
        }
    }
}