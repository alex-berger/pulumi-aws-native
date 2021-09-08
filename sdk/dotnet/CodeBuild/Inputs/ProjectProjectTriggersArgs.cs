// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodeBuild.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-projecttriggers.html
    /// </summary>
    public sealed class ProjectProjectTriggersArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-projecttriggers.html#cfn-codebuild-project-projecttriggers-buildtype
        /// </summary>
        [Input("buildType")]
        public Input<string>? BuildType { get; set; }

        [Input("filterGroups")]
        private InputList<Inputs.ProjectFilterGroupArgs>? _filterGroups;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-projecttriggers.html#cfn-codebuild-project-projecttriggers-filtergroups
        /// </summary>
        public InputList<Inputs.ProjectFilterGroupArgs> FilterGroups
        {
            get => _filterGroups ?? (_filterGroups = new InputList<Inputs.ProjectFilterGroupArgs>());
            set => _filterGroups = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-projecttriggers.html#cfn-codebuild-project-projecttriggers-webhook
        /// </summary>
        [Input("webhook")]
        public Input<bool>? Webhook { get; set; }

        public ProjectProjectTriggersArgs()
        {
        }
    }
}