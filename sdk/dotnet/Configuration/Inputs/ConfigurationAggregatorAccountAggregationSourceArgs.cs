// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Configuration.Inputs
{

    public sealed class ConfigurationAggregatorAccountAggregationSourceArgs : Pulumi.ResourceArgs
    {
        [Input("accountIds", required: true)]
        private InputList<string>? _accountIds;
        public InputList<string> AccountIds
        {
            get => _accountIds ?? (_accountIds = new InputList<string>());
            set => _accountIds = value;
        }

        [Input("allAwsRegions")]
        public Input<bool>? AllAwsRegions { get; set; }

        [Input("awsRegions")]
        private InputList<string>? _awsRegions;
        public InputList<string> AwsRegions
        {
            get => _awsRegions ?? (_awsRegions = new InputList<string>());
            set => _awsRegions = value;
        }

        public ConfigurationAggregatorAccountAggregationSourceArgs()
        {
        }
    }
}
