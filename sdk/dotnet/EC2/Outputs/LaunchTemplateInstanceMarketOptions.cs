// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-instancemarketoptions.html
    /// </summary>
    [OutputType]
    public sealed class LaunchTemplateInstanceMarketOptions
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-instancemarketoptions.html#cfn-ec2-launchtemplate-launchtemplatedata-instancemarketoptions-markettype
        /// </summary>
        public readonly string? MarketType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-instancemarketoptions.html#cfn-ec2-launchtemplate-launchtemplatedata-instancemarketoptions-spotoptions
        /// </summary>
        public readonly Outputs.LaunchTemplateSpotOptions? SpotOptions;

        [OutputConstructor]
        private LaunchTemplateInstanceMarketOptions(
            string? marketType,

            Outputs.LaunchTemplateSpotOptions? spotOptions)
        {
            MarketType = marketType;
            SpotOptions = spotOptions;
        }
    }
}
