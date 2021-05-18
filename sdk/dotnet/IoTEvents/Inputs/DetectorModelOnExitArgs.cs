// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTEvents.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-onexit.html
    /// </summary>
    public sealed class DetectorModelOnExitArgs : Pulumi.ResourceArgs
    {
        [Input("events")]
        private InputList<Inputs.DetectorModelEventArgs>? _events;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-onexit.html#cfn-iotevents-detectormodel-onexit-events
        /// </summary>
        public InputList<Inputs.DetectorModelEventArgs> Events
        {
            get => _events ?? (_events = new InputList<Inputs.DetectorModelEventArgs>());
            set => _events = value;
        }

        public DetectorModelOnExitArgs()
        {
        }
    }
}
