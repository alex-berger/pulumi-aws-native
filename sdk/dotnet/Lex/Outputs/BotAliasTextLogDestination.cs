// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex.Outputs
{

    /// <summary>
    /// Defines the Amazon CloudWatch Logs destination log group for conversation text logs.
    /// </summary>
    [OutputType]
    public sealed class BotAliasTextLogDestination
    {
        public readonly Outputs.BotAliasCloudWatchLogGroupLogDestination CloudWatch;

        [OutputConstructor]
        private BotAliasTextLogDestination(Outputs.BotAliasCloudWatchLogGroupLogDestination cloudWatch)
        {
            CloudWatch = cloudWatch;
        }
    }
}
