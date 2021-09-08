// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-automaticinputfailoversettings.html
    /// </summary>
    [OutputType]
    public sealed class ChannelAutomaticInputFailoverSettings
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-automaticinputfailoversettings.html#cfn-medialive-channel-automaticinputfailoversettings-errorcleartimemsec
        /// </summary>
        public readonly int? ErrorClearTimeMsec;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-automaticinputfailoversettings.html#cfn-medialive-channel-automaticinputfailoversettings-failoverconditions
        /// </summary>
        public readonly ImmutableArray<Outputs.ChannelFailoverCondition> FailoverConditions;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-automaticinputfailoversettings.html#cfn-medialive-channel-automaticinputfailoversettings-inputpreference
        /// </summary>
        public readonly string? InputPreference;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-automaticinputfailoversettings.html#cfn-medialive-channel-automaticinputfailoversettings-secondaryinputid
        /// </summary>
        public readonly string? SecondaryInputId;

        [OutputConstructor]
        private ChannelAutomaticInputFailoverSettings(
            int? errorClearTimeMsec,

            ImmutableArray<Outputs.ChannelFailoverCondition> failoverConditions,

            string? inputPreference,

            string? secondaryInputId)
        {
            ErrorClearTimeMsec = errorClearTimeMsec;
            FailoverConditions = failoverConditions;
            InputPreference = inputPreference;
            SecondaryInputId = secondaryInputId;
        }
    }
}