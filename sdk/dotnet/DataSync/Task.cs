// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataSync
{
    /// <summary>
    /// Resource schema for AWS::DataSync::Task.
    /// </summary>
    [AwsNativeResourceType("aws-native:datasync:Task")]
    public partial class Task : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the Amazon CloudWatch log group that is used to monitor and log events in the task.
        /// </summary>
        [Output("cloudWatchLogGroupArn")]
        public Output<string?> CloudWatchLogGroupArn { get; private set; } = null!;

        /// <summary>
        /// The ARN of an AWS storage resource's location.
        /// </summary>
        [Output("destinationLocationArn")]
        public Output<string> DestinationLocationArn { get; private set; } = null!;

        [Output("destinationNetworkInterfaceArns")]
        public Output<ImmutableArray<string>> DestinationNetworkInterfaceArns { get; private set; } = null!;

        /// <summary>
        /// Errors that AWS DataSync encountered during execution of the task. You can use this error code to help troubleshoot issues.
        /// </summary>
        [Output("errorCode")]
        public Output<string> ErrorCode { get; private set; } = null!;

        /// <summary>
        /// Detailed description of an error that was encountered during the task execution.
        /// </summary>
        [Output("errorDetail")]
        public Output<string> ErrorDetail { get; private set; } = null!;

        [Output("excludes")]
        public Output<ImmutableArray<Outputs.TaskFilterRule>> Excludes { get; private set; } = null!;

        [Output("includes")]
        public Output<ImmutableArray<Outputs.TaskFilterRule>> Includes { get; private set; } = null!;

        /// <summary>
        /// The name of a task. This value is a text reference that is used to identify the task in the console.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        [Output("options")]
        public Output<Outputs.TaskOptions?> Options { get; private set; } = null!;

        [Output("schedule")]
        public Output<Outputs.TaskSchedule?> Schedule { get; private set; } = null!;

        /// <summary>
        /// The ARN of the source location for the task.
        /// </summary>
        [Output("sourceLocationArn")]
        public Output<string> SourceLocationArn { get; private set; } = null!;

        [Output("sourceNetworkInterfaceArns")]
        public Output<ImmutableArray<string>> SourceNetworkInterfaceArns { get; private set; } = null!;

        /// <summary>
        /// The status of the task that was described.
        /// </summary>
        [Output("status")]
        public Output<Pulumi.AwsNative.DataSync.TaskStatus> Status { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.TaskTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// The ARN of the task.
        /// </summary>
        [Output("taskArn")]
        public Output<string> TaskArn { get; private set; } = null!;


        /// <summary>
        /// Create a Task resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Task(string name, TaskArgs args, CustomResourceOptions? options = null)
            : base("aws-native:datasync:Task", name, args ?? new TaskArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Task(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:datasync:Task", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Task resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Task Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Task(name, id, options);
        }
    }

    public sealed class TaskArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the Amazon CloudWatch log group that is used to monitor and log events in the task.
        /// </summary>
        [Input("cloudWatchLogGroupArn")]
        public Input<string>? CloudWatchLogGroupArn { get; set; }

        /// <summary>
        /// The ARN of an AWS storage resource's location.
        /// </summary>
        [Input("destinationLocationArn", required: true)]
        public Input<string> DestinationLocationArn { get; set; } = null!;

        [Input("excludes")]
        private InputList<Inputs.TaskFilterRuleArgs>? _excludes;
        public InputList<Inputs.TaskFilterRuleArgs> Excludes
        {
            get => _excludes ?? (_excludes = new InputList<Inputs.TaskFilterRuleArgs>());
            set => _excludes = value;
        }

        [Input("includes")]
        private InputList<Inputs.TaskFilterRuleArgs>? _includes;
        public InputList<Inputs.TaskFilterRuleArgs> Includes
        {
            get => _includes ?? (_includes = new InputList<Inputs.TaskFilterRuleArgs>());
            set => _includes = value;
        }

        /// <summary>
        /// The name of a task. This value is a text reference that is used to identify the task in the console.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("options")]
        public Input<Inputs.TaskOptionsArgs>? Options { get; set; }

        [Input("schedule")]
        public Input<Inputs.TaskScheduleArgs>? Schedule { get; set; }

        /// <summary>
        /// The ARN of the source location for the task.
        /// </summary>
        [Input("sourceLocationArn", required: true)]
        public Input<string> SourceLocationArn { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.TaskTagArgs>? _tags;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public InputList<Inputs.TaskTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.TaskTagArgs>());
            set => _tags = value;
        }

        public TaskArgs()
        {
        }
    }
}
