// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lightsail

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Lightsail::Alarm
type Alarm struct {
	pulumi.CustomResourceState

	AlarmArn pulumi.StringOutput `pulumi:"alarmArn"`
	// The name for the alarm. Specify the name of an existing alarm to update, and overwrite the previous configuration of the alarm.
	AlarmName pulumi.StringOutput `pulumi:"alarmName"`
	// The arithmetic operation to use when comparing the specified statistic to the threshold. The specified statistic value is used as the first operand.
	ComparisonOperator pulumi.StringOutput `pulumi:"comparisonOperator"`
	// The contact protocols to use for the alarm, such as Email, SMS (text messaging), or both.
	ContactProtocols pulumi.StringArrayOutput `pulumi:"contactProtocols"`
	// The number of data points that must be not within the specified threshold to trigger the alarm. If you are setting an "M out of N" alarm, this value (datapointsToAlarm) is the M.
	DatapointsToAlarm pulumi.IntPtrOutput `pulumi:"datapointsToAlarm"`
	// The number of most recent periods over which data is compared to the specified threshold. If you are setting an "M out of N" alarm, this value (evaluationPeriods) is the N.
	EvaluationPeriods pulumi.IntOutput `pulumi:"evaluationPeriods"`
	// The name of the metric to associate with the alarm.
	MetricName pulumi.StringOutput `pulumi:"metricName"`
	// The validation status of the SSL/TLS certificate.
	MonitoredResourceName pulumi.StringOutput `pulumi:"monitoredResourceName"`
	// Indicates whether the alarm is enabled. Notifications are enabled by default if you don't specify this parameter.
	NotificationEnabled pulumi.BoolPtrOutput `pulumi:"notificationEnabled"`
	// The alarm states that trigger a notification.
	NotificationTriggers pulumi.StringArrayOutput `pulumi:"notificationTriggers"`
	// The current state of the alarm.
	State pulumi.StringOutput `pulumi:"state"`
	// The value against which the specified statistic is compared.
	Threshold pulumi.Float64Output `pulumi:"threshold"`
	// Sets how this alarm will handle missing data points.
	TreatMissingData pulumi.StringPtrOutput `pulumi:"treatMissingData"`
}

// NewAlarm registers a new resource with the given unique name, arguments, and options.
func NewAlarm(ctx *pulumi.Context,
	name string, args *AlarmArgs, opts ...pulumi.ResourceOption) (*Alarm, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ComparisonOperator == nil {
		return nil, errors.New("invalid value for required argument 'ComparisonOperator'")
	}
	if args.EvaluationPeriods == nil {
		return nil, errors.New("invalid value for required argument 'EvaluationPeriods'")
	}
	if args.MetricName == nil {
		return nil, errors.New("invalid value for required argument 'MetricName'")
	}
	if args.MonitoredResourceName == nil {
		return nil, errors.New("invalid value for required argument 'MonitoredResourceName'")
	}
	if args.Threshold == nil {
		return nil, errors.New("invalid value for required argument 'Threshold'")
	}
	var resource Alarm
	err := ctx.RegisterResource("aws-native:lightsail:Alarm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlarm gets an existing Alarm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlarm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlarmState, opts ...pulumi.ResourceOption) (*Alarm, error) {
	var resource Alarm
	err := ctx.ReadResource("aws-native:lightsail:Alarm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Alarm resources.
type alarmState struct {
}

type AlarmState struct {
}

func (AlarmState) ElementType() reflect.Type {
	return reflect.TypeOf((*alarmState)(nil)).Elem()
}

type alarmArgs struct {
	// The name for the alarm. Specify the name of an existing alarm to update, and overwrite the previous configuration of the alarm.
	AlarmName *string `pulumi:"alarmName"`
	// The arithmetic operation to use when comparing the specified statistic to the threshold. The specified statistic value is used as the first operand.
	ComparisonOperator string `pulumi:"comparisonOperator"`
	// The contact protocols to use for the alarm, such as Email, SMS (text messaging), or both.
	ContactProtocols []string `pulumi:"contactProtocols"`
	// The number of data points that must be not within the specified threshold to trigger the alarm. If you are setting an "M out of N" alarm, this value (datapointsToAlarm) is the M.
	DatapointsToAlarm *int `pulumi:"datapointsToAlarm"`
	// The number of most recent periods over which data is compared to the specified threshold. If you are setting an "M out of N" alarm, this value (evaluationPeriods) is the N.
	EvaluationPeriods int `pulumi:"evaluationPeriods"`
	// The name of the metric to associate with the alarm.
	MetricName string `pulumi:"metricName"`
	// The validation status of the SSL/TLS certificate.
	MonitoredResourceName string `pulumi:"monitoredResourceName"`
	// Indicates whether the alarm is enabled. Notifications are enabled by default if you don't specify this parameter.
	NotificationEnabled *bool `pulumi:"notificationEnabled"`
	// The alarm states that trigger a notification.
	NotificationTriggers []string `pulumi:"notificationTriggers"`
	// The value against which the specified statistic is compared.
	Threshold float64 `pulumi:"threshold"`
	// Sets how this alarm will handle missing data points.
	TreatMissingData *string `pulumi:"treatMissingData"`
}

// The set of arguments for constructing a Alarm resource.
type AlarmArgs struct {
	// The name for the alarm. Specify the name of an existing alarm to update, and overwrite the previous configuration of the alarm.
	AlarmName pulumi.StringPtrInput
	// The arithmetic operation to use when comparing the specified statistic to the threshold. The specified statistic value is used as the first operand.
	ComparisonOperator pulumi.StringInput
	// The contact protocols to use for the alarm, such as Email, SMS (text messaging), or both.
	ContactProtocols pulumi.StringArrayInput
	// The number of data points that must be not within the specified threshold to trigger the alarm. If you are setting an "M out of N" alarm, this value (datapointsToAlarm) is the M.
	DatapointsToAlarm pulumi.IntPtrInput
	// The number of most recent periods over which data is compared to the specified threshold. If you are setting an "M out of N" alarm, this value (evaluationPeriods) is the N.
	EvaluationPeriods pulumi.IntInput
	// The name of the metric to associate with the alarm.
	MetricName pulumi.StringInput
	// The validation status of the SSL/TLS certificate.
	MonitoredResourceName pulumi.StringInput
	// Indicates whether the alarm is enabled. Notifications are enabled by default if you don't specify this parameter.
	NotificationEnabled pulumi.BoolPtrInput
	// The alarm states that trigger a notification.
	NotificationTriggers pulumi.StringArrayInput
	// The value against which the specified statistic is compared.
	Threshold pulumi.Float64Input
	// Sets how this alarm will handle missing data points.
	TreatMissingData pulumi.StringPtrInput
}

func (AlarmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alarmArgs)(nil)).Elem()
}

type AlarmInput interface {
	pulumi.Input

	ToAlarmOutput() AlarmOutput
	ToAlarmOutputWithContext(ctx context.Context) AlarmOutput
}

func (*Alarm) ElementType() reflect.Type {
	return reflect.TypeOf((**Alarm)(nil)).Elem()
}

func (i *Alarm) ToAlarmOutput() AlarmOutput {
	return i.ToAlarmOutputWithContext(context.Background())
}

func (i *Alarm) ToAlarmOutputWithContext(ctx context.Context) AlarmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlarmOutput)
}

type AlarmOutput struct{ *pulumi.OutputState }

func (AlarmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Alarm)(nil)).Elem()
}

func (o AlarmOutput) ToAlarmOutput() AlarmOutput {
	return o
}

func (o AlarmOutput) ToAlarmOutputWithContext(ctx context.Context) AlarmOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AlarmInput)(nil)).Elem(), &Alarm{})
	pulumi.RegisterOutputType(AlarmOutput{})
}
