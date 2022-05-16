// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkFirewall.Inputs
{

    public sealed class RuleGroupTCPFlagFieldArgs : Pulumi.ResourceArgs
    {
        [Input("flags", required: true)]
        private InputList<Pulumi.AwsNative.NetworkFirewall.RuleGroupTCPFlag>? _flags;
        public InputList<Pulumi.AwsNative.NetworkFirewall.RuleGroupTCPFlag> Flags
        {
            get => _flags ?? (_flags = new InputList<Pulumi.AwsNative.NetworkFirewall.RuleGroupTCPFlag>());
            set => _flags = value;
        }

        [Input("masks")]
        private InputList<Pulumi.AwsNative.NetworkFirewall.RuleGroupTCPFlag>? _masks;
        public InputList<Pulumi.AwsNative.NetworkFirewall.RuleGroupTCPFlag> Masks
        {
            get => _masks ?? (_masks = new InputList<Pulumi.AwsNative.NetworkFirewall.RuleGroupTCPFlag>());
            set => _masks = value;
        }

        public RuleGroupTCPFlagFieldArgs()
        {
        }
    }
}