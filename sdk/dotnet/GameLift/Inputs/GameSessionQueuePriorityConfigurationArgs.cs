// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GameLift.Inputs
{

    public sealed class GameSessionQueuePriorityConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("locationOrder")]
        private InputList<string>? _locationOrder;
        public InputList<string> LocationOrder
        {
            get => _locationOrder ?? (_locationOrder = new InputList<string>());
            set => _locationOrder = value;
        }

        [Input("priorityOrder")]
        private InputList<string>? _priorityOrder;
        public InputList<string> PriorityOrder
        {
            get => _priorityOrder ?? (_priorityOrder = new InputList<string>());
            set => _priorityOrder = value;
        }

        public GameSessionQueuePriorityConfigurationArgs()
        {
        }
    }
}
