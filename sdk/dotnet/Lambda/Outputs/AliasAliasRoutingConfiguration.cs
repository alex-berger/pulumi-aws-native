// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lambda.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-alias-aliasroutingconfiguration.html
    /// </summary>
    [OutputType]
    public sealed class AliasAliasRoutingConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-alias-aliasroutingconfiguration.html#cfn-lambda-alias-aliasroutingconfiguration-additionalversionweights
        /// </summary>
        public readonly ImmutableArray<Outputs.AliasVersionWeight> AdditionalVersionWeights;

        [OutputConstructor]
        private AliasAliasRoutingConfiguration(ImmutableArray<Outputs.AliasVersionWeight> additionalVersionWeights)
        {
            AdditionalVersionWeights = additionalVersionWeights;
        }
    }
}
