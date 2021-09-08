// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-failovercondition.html
    /// </summary>
    public sealed class ChannelFailoverConditionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-failovercondition.html#cfn-medialive-channel-failovercondition-failoverconditionsettings
        /// </summary>
        [Input("failoverConditionSettings")]
        public Input<Inputs.ChannelFailoverConditionSettingsArgs>? FailoverConditionSettings { get; set; }

        public ChannelFailoverConditionArgs()
        {
        }
    }
}