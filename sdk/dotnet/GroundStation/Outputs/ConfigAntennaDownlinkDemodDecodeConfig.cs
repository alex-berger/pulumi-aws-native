// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GroundStation.Outputs
{

    [OutputType]
    public sealed class ConfigAntennaDownlinkDemodDecodeConfig
    {
        public readonly Outputs.ConfigDecodeConfig? DecodeConfig;
        public readonly Outputs.ConfigDemodulationConfig? DemodulationConfig;
        public readonly Outputs.ConfigSpectrumConfig? SpectrumConfig;

        [OutputConstructor]
        private ConfigAntennaDownlinkDemodDecodeConfig(
            Outputs.ConfigDecodeConfig? decodeConfig,

            Outputs.ConfigDemodulationConfig? demodulationConfig,

            Outputs.ConfigSpectrumConfig? spectrumConfig)
        {
            DecodeConfig = decodeConfig;
            DemodulationConfig = demodulationConfig;
            SpectrumConfig = spectrumConfig;
        }
    }
}
