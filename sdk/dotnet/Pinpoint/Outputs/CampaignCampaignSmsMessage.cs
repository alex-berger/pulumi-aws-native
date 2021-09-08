// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pinpoint.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-campaignsmsmessage.html
    /// </summary>
    [OutputType]
    public sealed class CampaignCampaignSmsMessage
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-campaignsmsmessage.html#cfn-pinpoint-campaign-campaignsmsmessage-body
        /// </summary>
        public readonly string? Body;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-campaignsmsmessage.html#cfn-pinpoint-campaign-campaignsmsmessage-entityid
        /// </summary>
        public readonly string? EntityId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-campaignsmsmessage.html#cfn-pinpoint-campaign-campaignsmsmessage-messagetype
        /// </summary>
        public readonly string? MessageType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-campaignsmsmessage.html#cfn-pinpoint-campaign-campaignsmsmessage-originationnumber
        /// </summary>
        public readonly string? OriginationNumber;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-campaignsmsmessage.html#cfn-pinpoint-campaign-campaignsmsmessage-senderid
        /// </summary>
        public readonly string? SenderId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-campaignsmsmessage.html#cfn-pinpoint-campaign-campaignsmsmessage-templateid
        /// </summary>
        public readonly string? TemplateId;

        [OutputConstructor]
        private CampaignCampaignSmsMessage(
            string? body,

            string? entityId,

            string? messageType,

            string? originationNumber,

            string? senderId,

            string? templateId)
        {
            Body = body;
            EntityId = entityId;
            MessageType = messageType;
            OriginationNumber = originationNumber;
            SenderId = senderId;
            TemplateId = templateId;
        }
    }
}
