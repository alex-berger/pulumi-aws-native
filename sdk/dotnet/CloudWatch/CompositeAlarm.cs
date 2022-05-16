// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudWatch
{
    /// <summary>
    /// The AWS::CloudWatch::CompositeAlarm type specifies an alarm which aggregates the states of other Alarms (Metric or Composite Alarms) as defined by the AlarmRule expression
    /// </summary>
    [AwsNativeResourceType("aws-native:cloudwatch:CompositeAlarm")]
    public partial class CompositeAlarm : Pulumi.CustomResource
    {
        /// <summary>
        /// Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE.
        /// </summary>
        [Output("actionsEnabled")]
        public Output<bool?> ActionsEnabled { get; private set; } = null!;

        /// <summary>
        /// The list of actions to execute when this alarm transitions into an ALARM state from any other state. Specify each action as an Amazon Resource Name (ARN).
        /// </summary>
        [Output("alarmActions")]
        public Output<ImmutableArray<string>> AlarmActions { get; private set; } = null!;

        /// <summary>
        /// The description of the alarm
        /// </summary>
        [Output("alarmDescription")]
        public Output<string?> AlarmDescription { get; private set; } = null!;

        /// <summary>
        /// The name of the Composite Alarm
        /// </summary>
        [Output("alarmName")]
        public Output<string> AlarmName { get; private set; } = null!;

        /// <summary>
        /// Expression which aggregates the state of other Alarms (Metric or Composite Alarms)
        /// </summary>
        [Output("alarmRule")]
        public Output<string> AlarmRule { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) of the alarm
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN).
        /// </summary>
        [Output("insufficientDataActions")]
        public Output<ImmutableArray<string>> InsufficientDataActions { get; private set; } = null!;

        /// <summary>
        /// The actions to execute when this alarm transitions to the OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).
        /// </summary>
        [Output("oKActions")]
        public Output<ImmutableArray<string>> OKActions { get; private set; } = null!;


        /// <summary>
        /// Create a CompositeAlarm resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CompositeAlarm(string name, CompositeAlarmArgs args, CustomResourceOptions? options = null)
            : base("aws-native:cloudwatch:CompositeAlarm", name, args ?? new CompositeAlarmArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CompositeAlarm(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:cloudwatch:CompositeAlarm", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing CompositeAlarm resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CompositeAlarm Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CompositeAlarm(name, id, options);
        }
    }

    public sealed class CompositeAlarmArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE.
        /// </summary>
        [Input("actionsEnabled")]
        public Input<bool>? ActionsEnabled { get; set; }

        [Input("alarmActions")]
        private InputList<string>? _alarmActions;

        /// <summary>
        /// The list of actions to execute when this alarm transitions into an ALARM state from any other state. Specify each action as an Amazon Resource Name (ARN).
        /// </summary>
        public InputList<string> AlarmActions
        {
            get => _alarmActions ?? (_alarmActions = new InputList<string>());
            set => _alarmActions = value;
        }

        /// <summary>
        /// The description of the alarm
        /// </summary>
        [Input("alarmDescription")]
        public Input<string>? AlarmDescription { get; set; }

        /// <summary>
        /// The name of the Composite Alarm
        /// </summary>
        [Input("alarmName", required: true)]
        public Input<string> AlarmName { get; set; } = null!;

        /// <summary>
        /// Expression which aggregates the state of other Alarms (Metric or Composite Alarms)
        /// </summary>
        [Input("alarmRule", required: true)]
        public Input<string> AlarmRule { get; set; } = null!;

        [Input("insufficientDataActions")]
        private InputList<string>? _insufficientDataActions;

        /// <summary>
        /// The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN).
        /// </summary>
        public InputList<string> InsufficientDataActions
        {
            get => _insufficientDataActions ?? (_insufficientDataActions = new InputList<string>());
            set => _insufficientDataActions = value;
        }

        [Input("oKActions")]
        private InputList<string>? _oKActions;

        /// <summary>
        /// The actions to execute when this alarm transitions to the OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).
        /// </summary>
        public InputList<string> OKActions
        {
            get => _oKActions ?? (_oKActions = new InputList<string>());
            set => _oKActions = value;
        }

        public CompositeAlarmArgs()
        {
        }
    }
}