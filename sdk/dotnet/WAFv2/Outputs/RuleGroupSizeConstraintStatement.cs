// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFv2.Outputs
{

    /// <summary>
    /// Size Constraint statement.
    /// </summary>
    [OutputType]
    public sealed class RuleGroupSizeConstraintStatement
    {
        public readonly Pulumi.AwsNative.WAFv2.RuleGroupSizeConstraintStatementComparisonOperator ComparisonOperator;
        public readonly Outputs.RuleGroupFieldToMatch FieldToMatch;
        public readonly double Size;
        public readonly ImmutableArray<Outputs.RuleGroupTextTransformation> TextTransformations;

        [OutputConstructor]
        private RuleGroupSizeConstraintStatement(
            Pulumi.AwsNative.WAFv2.RuleGroupSizeConstraintStatementComparisonOperator comparisonOperator,

            Outputs.RuleGroupFieldToMatch fieldToMatch,

            double size,

            ImmutableArray<Outputs.RuleGroupTextTransformation> textTransformations)
        {
            ComparisonOperator = comparisonOperator;
            FieldToMatch = fieldToMatch;
            Size = size;
            TextTransformations = textTransformations;
        }
    }
}
