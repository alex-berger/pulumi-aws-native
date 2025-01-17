// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTWireless.Inputs
{

    public sealed class DeviceProfileLoRaWANDeviceProfileArgs : Pulumi.ResourceArgs
    {
        [Input("classBTimeout")]
        public Input<int>? ClassBTimeout { get; set; }

        [Input("classCTimeout")]
        public Input<int>? ClassCTimeout { get; set; }

        [Input("macVersion")]
        public Input<string>? MacVersion { get; set; }

        [Input("maxDutyCycle")]
        public Input<int>? MaxDutyCycle { get; set; }

        [Input("maxEirp")]
        public Input<int>? MaxEirp { get; set; }

        [Input("pingSlotDr")]
        public Input<int>? PingSlotDr { get; set; }

        [Input("pingSlotFreq")]
        public Input<int>? PingSlotFreq { get; set; }

        [Input("pingSlotPeriod")]
        public Input<int>? PingSlotPeriod { get; set; }

        [Input("regParamsRevision")]
        public Input<string>? RegParamsRevision { get; set; }

        [Input("rfRegion")]
        public Input<string>? RfRegion { get; set; }

        [Input("supports32BitFCnt")]
        public Input<bool>? Supports32BitFCnt { get; set; }

        [Input("supportsClassB")]
        public Input<bool>? SupportsClassB { get; set; }

        [Input("supportsClassC")]
        public Input<bool>? SupportsClassC { get; set; }

        [Input("supportsJoin")]
        public Input<bool>? SupportsJoin { get; set; }

        public DeviceProfileLoRaWANDeviceProfileArgs()
        {
        }
    }
}
