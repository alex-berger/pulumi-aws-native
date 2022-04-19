// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lambda.Inputs
{

    public sealed class AliasRoutingConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("additionalVersionWeights", required: true)]
        private InputList<Inputs.AliasVersionWeightArgs>? _additionalVersionWeights;
        public InputList<Inputs.AliasVersionWeightArgs> AdditionalVersionWeights
        {
            get => _additionalVersionWeights ?? (_additionalVersionWeights = new InputList<Inputs.AliasVersionWeightArgs>());
            set => _additionalVersionWeights = value;
        }

        public AliasRoutingConfigurationArgs()
        {
        }
    }
}
