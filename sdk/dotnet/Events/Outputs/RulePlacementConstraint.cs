// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Events.Outputs
{

    [OutputType]
    public sealed class RulePlacementConstraint
    {
        public readonly string? Expression;
        public readonly string? Type;

        [OutputConstructor]
        private RulePlacementConstraint(
            string? expression,

            string? type)
        {
            Expression = expression;
            Type = type;
        }
    }
}
