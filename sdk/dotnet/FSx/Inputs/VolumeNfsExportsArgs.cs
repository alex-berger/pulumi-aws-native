// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.FSx.Inputs
{

    public sealed class VolumeNfsExportsArgs : Pulumi.ResourceArgs
    {
        [Input("clientConfigurations", required: true)]
        private InputList<Inputs.VolumeClientConfigurationsArgs>? _clientConfigurations;
        public InputList<Inputs.VolumeClientConfigurationsArgs> ClientConfigurations
        {
            get => _clientConfigurations ?? (_clientConfigurations = new InputList<Inputs.VolumeClientConfigurationsArgs>());
            set => _clientConfigurations = value;
        }

        public VolumeNfsExportsArgs()
        {
        }
    }
}
