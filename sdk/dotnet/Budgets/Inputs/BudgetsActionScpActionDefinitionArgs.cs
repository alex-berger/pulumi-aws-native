// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Budgets.Inputs
{

    public sealed class BudgetsActionScpActionDefinitionArgs : Pulumi.ResourceArgs
    {
        [Input("policyId", required: true)]
        public Input<string> PolicyId { get; set; } = null!;

        [Input("targetIds", required: true)]
        private InputList<string>? _targetIds;
        public InputList<string> TargetIds
        {
            get => _targetIds ?? (_targetIds = new InputList<string>());
            set => _targetIds = value;
        }

        public BudgetsActionScpActionDefinitionArgs()
        {
        }
    }
}