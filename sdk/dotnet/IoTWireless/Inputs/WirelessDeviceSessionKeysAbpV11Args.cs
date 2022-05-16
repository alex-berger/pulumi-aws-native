// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTWireless.Inputs
{

    public sealed class WirelessDeviceSessionKeysAbpV11Args : Pulumi.ResourceArgs
    {
        [Input("appSKey", required: true)]
        public Input<string> AppSKey { get; set; } = null!;

        [Input("fNwkSIntKey", required: true)]
        public Input<string> FNwkSIntKey { get; set; } = null!;

        [Input("nwkSEncKey", required: true)]
        public Input<string> NwkSEncKey { get; set; } = null!;

        [Input("sNwkSIntKey", required: true)]
        public Input<string> SNwkSIntKey { get; set; } = null!;

        public WirelessDeviceSessionKeysAbpV11Args()
        {
        }
    }
}