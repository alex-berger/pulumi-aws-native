// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTEvents.Inputs
{

    /// <summary>
    /// Sends an AWS IoT Events input, passing in information about the detector model instance and the event that triggered the action.
    /// </summary>
    public sealed class DetectorModelIotEventsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the AWS IoT Events input where the data is sent.
        /// </summary>
        [Input("inputName", required: true)]
        public Input<string> InputName { get; set; } = null!;

        [Input("payload")]
        public Input<Inputs.DetectorModelPayloadArgs>? Payload { get; set; }

        public DetectorModelIotEventsArgs()
        {
        }
    }
}