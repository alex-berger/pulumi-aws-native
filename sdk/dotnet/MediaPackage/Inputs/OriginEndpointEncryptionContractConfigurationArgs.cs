// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaPackage.Inputs
{

    /// <summary>
    /// The configuration to use for encrypting one or more content tracks separately for endpoints that use SPEKE 2.0.
    /// </summary>
    public sealed class OriginEndpointEncryptionContractConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A collection of audio encryption presets.
        /// </summary>
        [Input("presetSpeke20Audio", required: true)]
        public Input<Pulumi.AwsNative.MediaPackage.OriginEndpointEncryptionContractConfigurationPresetSpeke20Audio> PresetSpeke20Audio { get; set; } = null!;

        /// <summary>
        /// A collection of video encryption presets.
        /// </summary>
        [Input("presetSpeke20Video", required: true)]
        public Input<Pulumi.AwsNative.MediaPackage.OriginEndpointEncryptionContractConfigurationPresetSpeke20Video> PresetSpeke20Video { get; set; } = null!;

        public OriginEndpointEncryptionContractConfigurationArgs()
        {
        }
    }
}
