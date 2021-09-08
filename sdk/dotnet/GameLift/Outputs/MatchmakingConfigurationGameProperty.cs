// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GameLift.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-matchmakingconfiguration-gameproperty.html
    /// </summary>
    [OutputType]
    public sealed class MatchmakingConfigurationGameProperty
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-matchmakingconfiguration-gameproperty.html#cfn-gamelift-matchmakingconfiguration-gameproperty-key
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-matchmakingconfiguration-gameproperty.html#cfn-gamelift-matchmakingconfiguration-gameproperty-value
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private MatchmakingConfigurationGameProperty(
            string key,

            string value)
        {
            Key = key;
            Value = value;
        }
    }
}
