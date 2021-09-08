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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-applicationsettings-campaignhook.html
    /// </summary>
    [OutputType]
    public sealed class ApplicationSettingsCampaignHook
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-applicationsettings-campaignhook.html#cfn-pinpoint-applicationsettings-campaignhook-lambdafunctionname
        /// </summary>
        public readonly string? LambdaFunctionName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-applicationsettings-campaignhook.html#cfn-pinpoint-applicationsettings-campaignhook-mode
        /// </summary>
        public readonly string? Mode;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-applicationsettings-campaignhook.html#cfn-pinpoint-applicationsettings-campaignhook-weburl
        /// </summary>
        public readonly string? WebUrl;

        [OutputConstructor]
        private ApplicationSettingsCampaignHook(
            string? lambdaFunctionName,

            string? mode,

            string? webUrl)
        {
            LambdaFunctionName = lambdaFunctionName;
            Mode = mode;
            WebUrl = webUrl;
        }
    }
}
