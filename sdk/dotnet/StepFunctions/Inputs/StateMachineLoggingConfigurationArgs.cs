// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.StepFunctions.Inputs
{

    public sealed class StateMachineLoggingConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("destinations")]
        private InputList<Inputs.StateMachineLogDestinationArgs>? _destinations;
        public InputList<Inputs.StateMachineLogDestinationArgs> Destinations
        {
            get => _destinations ?? (_destinations = new InputList<Inputs.StateMachineLogDestinationArgs>());
            set => _destinations = value;
        }

        [Input("includeExecutionData")]
        public Input<bool>? IncludeExecutionData { get; set; }

        [Input("level")]
        public Input<Pulumi.AwsNative.StepFunctions.StateMachineLoggingConfigurationLevel>? Level { get; set; }

        public StateMachineLoggingConfigurationArgs()
        {
        }
    }
}
