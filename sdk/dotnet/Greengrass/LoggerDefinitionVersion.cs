// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Greengrass
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-loggerdefinitionversion.html
    /// </summary>
    [AwsNativeResourceType("aws-native:greengrass:LoggerDefinitionVersion")]
    public partial class LoggerDefinitionVersion : Pulumi.CustomResource
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-loggerdefinitionversion.html#cfn-greengrass-loggerdefinitionversion-loggerdefinitionid
        /// </summary>
        [Output("loggerDefinitionId")]
        public Output<string> LoggerDefinitionId { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-loggerdefinitionversion.html#cfn-greengrass-loggerdefinitionversion-loggers
        /// </summary>
        [Output("loggers")]
        public Output<ImmutableArray<Outputs.LoggerDefinitionVersionLogger>> Loggers { get; private set; } = null!;


        /// <summary>
        /// Create a LoggerDefinitionVersion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LoggerDefinitionVersion(string name, LoggerDefinitionVersionArgs args, CustomResourceOptions? options = null)
            : base("aws-native:greengrass:LoggerDefinitionVersion", name, args ?? new LoggerDefinitionVersionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LoggerDefinitionVersion(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:greengrass:LoggerDefinitionVersion", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LoggerDefinitionVersion resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LoggerDefinitionVersion Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new LoggerDefinitionVersion(name, id, options);
        }
    }

    public sealed class LoggerDefinitionVersionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-loggerdefinitionversion.html#cfn-greengrass-loggerdefinitionversion-loggerdefinitionid
        /// </summary>
        [Input("loggerDefinitionId", required: true)]
        public Input<string> LoggerDefinitionId { get; set; } = null!;

        [Input("loggers", required: true)]
        private InputList<Inputs.LoggerDefinitionVersionLoggerArgs>? _loggers;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-loggerdefinitionversion.html#cfn-greengrass-loggerdefinitionversion-loggers
        /// </summary>
        public InputList<Inputs.LoggerDefinitionVersionLoggerArgs> Loggers
        {
            get => _loggers ?? (_loggers = new InputList<Inputs.LoggerDefinitionVersionLoggerArgs>());
            set => _loggers = value;
        }

        public LoggerDefinitionVersionArgs()
        {
        }
    }
}