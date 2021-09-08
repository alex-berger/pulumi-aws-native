// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.FIS.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplatetarget.html
    /// </summary>
    [OutputType]
    public sealed class ExperimentTemplateExperimentTemplateTarget
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplatetarget.html#cfn-fis-experimenttemplate-experimenttemplatetarget-filters
        /// </summary>
        public readonly ImmutableArray<Outputs.ExperimentTemplateExperimentTemplateTargetFilter> Filters;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplatetarget.html#cfn-fis-experimenttemplate-experimenttemplatetarget-resourcearns
        /// </summary>
        public readonly ImmutableArray<string> ResourceArns;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplatetarget.html#cfn-fis-experimenttemplate-experimenttemplatetarget-resourcetags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? ResourceTags;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplatetarget.html#cfn-fis-experimenttemplate-experimenttemplatetarget-resourcetype
        /// </summary>
        public readonly string ResourceType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplatetarget.html#cfn-fis-experimenttemplate-experimenttemplatetarget-selectionmode
        /// </summary>
        public readonly string SelectionMode;

        [OutputConstructor]
        private ExperimentTemplateExperimentTemplateTarget(
            ImmutableArray<Outputs.ExperimentTemplateExperimentTemplateTargetFilter> filters,

            ImmutableArray<string> resourceArns,

            ImmutableDictionary<string, string>? resourceTags,

            string resourceType,

            string selectionMode)
        {
            Filters = filters;
            ResourceArns = resourceArns;
            ResourceTags = resourceTags;
            ResourceType = resourceType;
            SelectionMode = selectionMode;
        }
    }
}