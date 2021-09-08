// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Batch.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties-logconfiguration.html
    /// </summary>
    public sealed class JobDefinitionLogConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties-logconfiguration.html#cfn-batch-jobdefinition-containerproperties-logconfiguration-logdriver
        /// </summary>
        [Input("logDriver", required: true)]
        public Input<string> LogDriver { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties-logconfiguration.html#cfn-batch-jobdefinition-containerproperties-logconfiguration-options
        /// </summary>
        [Input("options")]
        public InputUnion<System.Text.Json.JsonElement, string>? Options { get; set; }

        [Input("secretOptions")]
        private InputList<Inputs.JobDefinitionSecretArgs>? _secretOptions;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties-logconfiguration.html#cfn-batch-jobdefinition-containerproperties-logconfiguration-secretoptions
        /// </summary>
        public InputList<Inputs.JobDefinitionSecretArgs> SecretOptions
        {
            get => _secretOptions ?? (_secretOptions = new InputList<Inputs.JobDefinitionSecretArgs>());
            set => _secretOptions = value;
        }

        public JobDefinitionLogConfigurationArgs()
        {
        }
    }
}