// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Budgets.Outputs
{

    [OutputType]
    public sealed class BudgetsActionSsmActionDefinition
    {
        public readonly ImmutableArray<string> InstanceIds;
        public readonly string Region;
        public readonly Pulumi.AwsNative.Budgets.BudgetsActionSsmActionDefinitionSubtype Subtype;

        [OutputConstructor]
        private BudgetsActionSsmActionDefinition(
            ImmutableArray<string> instanceIds,

            string region,

            Pulumi.AwsNative.Budgets.BudgetsActionSsmActionDefinitionSubtype subtype)
        {
            InstanceIds = instanceIds;
            Region = region;
            Subtype = subtype;
        }
    }
}
