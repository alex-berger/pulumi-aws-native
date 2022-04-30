// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Inputs
{

    public sealed class ChannelAutomaticInputFailoverSettingsArgs : Pulumi.ResourceArgs
    {
        [Input("errorClearTimeMsec")]
        public Input<int>? ErrorClearTimeMsec { get; set; }

        [Input("failoverConditions")]
        private InputList<Inputs.ChannelFailoverConditionArgs>? _failoverConditions;
        public InputList<Inputs.ChannelFailoverConditionArgs> FailoverConditions
        {
            get => _failoverConditions ?? (_failoverConditions = new InputList<Inputs.ChannelFailoverConditionArgs>());
            set => _failoverConditions = value;
        }

        [Input("inputPreference")]
        public Input<string>? InputPreference { get; set; }

        [Input("secondaryInputId")]
        public Input<string>? SecondaryInputId { get; set; }

        public ChannelAutomaticInputFailoverSettingsArgs()
        {
        }
    }
}
