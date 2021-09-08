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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputattachment.html
    /// </summary>
    [OutputType]
    public sealed class ChannelInputAttachment
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputattachment.html#cfn-medialive-channel-inputattachment-automaticinputfailoversettings
        /// </summary>
        public readonly Outputs.ChannelAutomaticInputFailoverSettings? AutomaticInputFailoverSettings;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputattachment.html#cfn-medialive-channel-inputattachment-inputattachmentname
        /// </summary>
        public readonly string? InputAttachmentName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputattachment.html#cfn-medialive-channel-inputattachment-inputid
        /// </summary>
        public readonly string? InputId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputattachment.html#cfn-medialive-channel-inputattachment-inputsettings
        /// </summary>
        public readonly Outputs.ChannelInputSettings? InputSettings;

        [OutputConstructor]
        private ChannelInputAttachment(
            Outputs.ChannelAutomaticInputFailoverSettings? automaticInputFailoverSettings,

            string? inputAttachmentName,

            string? inputId,

            Outputs.ChannelInputSettings? inputSettings)
        {
            AutomaticInputFailoverSettings = automaticInputFailoverSettings;
            InputAttachmentName = inputAttachmentName;
            InputId = inputId;
            InputSettings = inputSettings;
        }
    }
}
