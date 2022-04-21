// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redshift

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EventSubscriptionEventCategoriesItem string

const (
	EventSubscriptionEventCategoriesItemConfiguration = EventSubscriptionEventCategoriesItem("configuration")
	EventSubscriptionEventCategoriesItemManagement    = EventSubscriptionEventCategoriesItem("management")
	EventSubscriptionEventCategoriesItemMonitoring    = EventSubscriptionEventCategoriesItem("monitoring")
	EventSubscriptionEventCategoriesItemSecurity      = EventSubscriptionEventCategoriesItem("security")
	EventSubscriptionEventCategoriesItemPending       = EventSubscriptionEventCategoriesItem("pending")
)

func (EventSubscriptionEventCategoriesItem) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSubscriptionEventCategoriesItem)(nil)).Elem()
}

func (e EventSubscriptionEventCategoriesItem) ToEventSubscriptionEventCategoriesItemOutput() EventSubscriptionEventCategoriesItemOutput {
	return pulumi.ToOutput(e).(EventSubscriptionEventCategoriesItemOutput)
}

func (e EventSubscriptionEventCategoriesItem) ToEventSubscriptionEventCategoriesItemOutputWithContext(ctx context.Context) EventSubscriptionEventCategoriesItemOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EventSubscriptionEventCategoriesItemOutput)
}

func (e EventSubscriptionEventCategoriesItem) ToEventSubscriptionEventCategoriesItemPtrOutput() EventSubscriptionEventCategoriesItemPtrOutput {
	return e.ToEventSubscriptionEventCategoriesItemPtrOutputWithContext(context.Background())
}

func (e EventSubscriptionEventCategoriesItem) ToEventSubscriptionEventCategoriesItemPtrOutputWithContext(ctx context.Context) EventSubscriptionEventCategoriesItemPtrOutput {
	return EventSubscriptionEventCategoriesItem(e).ToEventSubscriptionEventCategoriesItemOutputWithContext(ctx).ToEventSubscriptionEventCategoriesItemPtrOutputWithContext(ctx)
}

func (e EventSubscriptionEventCategoriesItem) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EventSubscriptionEventCategoriesItem) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EventSubscriptionEventCategoriesItem) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EventSubscriptionEventCategoriesItem) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EventSubscriptionEventCategoriesItemOutput struct{ *pulumi.OutputState }

func (EventSubscriptionEventCategoriesItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSubscriptionEventCategoriesItem)(nil)).Elem()
}

func (o EventSubscriptionEventCategoriesItemOutput) ToEventSubscriptionEventCategoriesItemOutput() EventSubscriptionEventCategoriesItemOutput {
	return o
}

func (o EventSubscriptionEventCategoriesItemOutput) ToEventSubscriptionEventCategoriesItemOutputWithContext(ctx context.Context) EventSubscriptionEventCategoriesItemOutput {
	return o
}

func (o EventSubscriptionEventCategoriesItemOutput) ToEventSubscriptionEventCategoriesItemPtrOutput() EventSubscriptionEventCategoriesItemPtrOutput {
	return o.ToEventSubscriptionEventCategoriesItemPtrOutputWithContext(context.Background())
}

func (o EventSubscriptionEventCategoriesItemOutput) ToEventSubscriptionEventCategoriesItemPtrOutputWithContext(ctx context.Context) EventSubscriptionEventCategoriesItemPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EventSubscriptionEventCategoriesItem) *EventSubscriptionEventCategoriesItem {
		return &v
	}).(EventSubscriptionEventCategoriesItemPtrOutput)
}

func (o EventSubscriptionEventCategoriesItemOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EventSubscriptionEventCategoriesItemOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EventSubscriptionEventCategoriesItem) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EventSubscriptionEventCategoriesItemOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EventSubscriptionEventCategoriesItemOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EventSubscriptionEventCategoriesItem) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EventSubscriptionEventCategoriesItemPtrOutput struct{ *pulumi.OutputState }

func (EventSubscriptionEventCategoriesItemPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSubscriptionEventCategoriesItem)(nil)).Elem()
}

func (o EventSubscriptionEventCategoriesItemPtrOutput) ToEventSubscriptionEventCategoriesItemPtrOutput() EventSubscriptionEventCategoriesItemPtrOutput {
	return o
}

func (o EventSubscriptionEventCategoriesItemPtrOutput) ToEventSubscriptionEventCategoriesItemPtrOutputWithContext(ctx context.Context) EventSubscriptionEventCategoriesItemPtrOutput {
	return o
}

func (o EventSubscriptionEventCategoriesItemPtrOutput) Elem() EventSubscriptionEventCategoriesItemOutput {
	return o.ApplyT(func(v *EventSubscriptionEventCategoriesItem) EventSubscriptionEventCategoriesItem {
		if v != nil {
			return *v
		}
		var ret EventSubscriptionEventCategoriesItem
		return ret
	}).(EventSubscriptionEventCategoriesItemOutput)
}

func (o EventSubscriptionEventCategoriesItemPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EventSubscriptionEventCategoriesItemPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EventSubscriptionEventCategoriesItem) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// EventSubscriptionEventCategoriesItemInput is an input type that accepts EventSubscriptionEventCategoriesItemArgs and EventSubscriptionEventCategoriesItemOutput values.
// You can construct a concrete instance of `EventSubscriptionEventCategoriesItemInput` via:
//
//          EventSubscriptionEventCategoriesItemArgs{...}
type EventSubscriptionEventCategoriesItemInput interface {
	pulumi.Input

	ToEventSubscriptionEventCategoriesItemOutput() EventSubscriptionEventCategoriesItemOutput
	ToEventSubscriptionEventCategoriesItemOutputWithContext(context.Context) EventSubscriptionEventCategoriesItemOutput
}

var eventSubscriptionEventCategoriesItemPtrType = reflect.TypeOf((**EventSubscriptionEventCategoriesItem)(nil)).Elem()

type EventSubscriptionEventCategoriesItemPtrInput interface {
	pulumi.Input

	ToEventSubscriptionEventCategoriesItemPtrOutput() EventSubscriptionEventCategoriesItemPtrOutput
	ToEventSubscriptionEventCategoriesItemPtrOutputWithContext(context.Context) EventSubscriptionEventCategoriesItemPtrOutput
}

type eventSubscriptionEventCategoriesItemPtr string

func EventSubscriptionEventCategoriesItemPtr(v string) EventSubscriptionEventCategoriesItemPtrInput {
	return (*eventSubscriptionEventCategoriesItemPtr)(&v)
}

func (*eventSubscriptionEventCategoriesItemPtr) ElementType() reflect.Type {
	return eventSubscriptionEventCategoriesItemPtrType
}

func (in *eventSubscriptionEventCategoriesItemPtr) ToEventSubscriptionEventCategoriesItemPtrOutput() EventSubscriptionEventCategoriesItemPtrOutput {
	return pulumi.ToOutput(in).(EventSubscriptionEventCategoriesItemPtrOutput)
}

func (in *eventSubscriptionEventCategoriesItemPtr) ToEventSubscriptionEventCategoriesItemPtrOutputWithContext(ctx context.Context) EventSubscriptionEventCategoriesItemPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EventSubscriptionEventCategoriesItemPtrOutput)
}

// EventSubscriptionEventCategoriesItemArrayInput is an input type that accepts EventSubscriptionEventCategoriesItemArray and EventSubscriptionEventCategoriesItemArrayOutput values.
// You can construct a concrete instance of `EventSubscriptionEventCategoriesItemArrayInput` via:
//
//          EventSubscriptionEventCategoriesItemArray{ EventSubscriptionEventCategoriesItemArgs{...} }
type EventSubscriptionEventCategoriesItemArrayInput interface {
	pulumi.Input

	ToEventSubscriptionEventCategoriesItemArrayOutput() EventSubscriptionEventCategoriesItemArrayOutput
	ToEventSubscriptionEventCategoriesItemArrayOutputWithContext(context.Context) EventSubscriptionEventCategoriesItemArrayOutput
}

type EventSubscriptionEventCategoriesItemArray []EventSubscriptionEventCategoriesItem

func (EventSubscriptionEventCategoriesItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EventSubscriptionEventCategoriesItem)(nil)).Elem()
}

func (i EventSubscriptionEventCategoriesItemArray) ToEventSubscriptionEventCategoriesItemArrayOutput() EventSubscriptionEventCategoriesItemArrayOutput {
	return i.ToEventSubscriptionEventCategoriesItemArrayOutputWithContext(context.Background())
}

func (i EventSubscriptionEventCategoriesItemArray) ToEventSubscriptionEventCategoriesItemArrayOutputWithContext(ctx context.Context) EventSubscriptionEventCategoriesItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionEventCategoriesItemArrayOutput)
}

type EventSubscriptionEventCategoriesItemArrayOutput struct{ *pulumi.OutputState }

func (EventSubscriptionEventCategoriesItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EventSubscriptionEventCategoriesItem)(nil)).Elem()
}

func (o EventSubscriptionEventCategoriesItemArrayOutput) ToEventSubscriptionEventCategoriesItemArrayOutput() EventSubscriptionEventCategoriesItemArrayOutput {
	return o
}

func (o EventSubscriptionEventCategoriesItemArrayOutput) ToEventSubscriptionEventCategoriesItemArrayOutputWithContext(ctx context.Context) EventSubscriptionEventCategoriesItemArrayOutput {
	return o
}

func (o EventSubscriptionEventCategoriesItemArrayOutput) Index(i pulumi.IntInput) EventSubscriptionEventCategoriesItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EventSubscriptionEventCategoriesItem {
		return vs[0].([]EventSubscriptionEventCategoriesItem)[vs[1].(int)]
	}).(EventSubscriptionEventCategoriesItemOutput)
}

// Specifies the Amazon Redshift event severity to be published by the event notification subscription.
type EventSubscriptionSeverity string

const (
	EventSubscriptionSeverityError = EventSubscriptionSeverity("ERROR")
	EventSubscriptionSeverityInfo  = EventSubscriptionSeverity("INFO")
)

func (EventSubscriptionSeverity) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSubscriptionSeverity)(nil)).Elem()
}

func (e EventSubscriptionSeverity) ToEventSubscriptionSeverityOutput() EventSubscriptionSeverityOutput {
	return pulumi.ToOutput(e).(EventSubscriptionSeverityOutput)
}

func (e EventSubscriptionSeverity) ToEventSubscriptionSeverityOutputWithContext(ctx context.Context) EventSubscriptionSeverityOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EventSubscriptionSeverityOutput)
}

func (e EventSubscriptionSeverity) ToEventSubscriptionSeverityPtrOutput() EventSubscriptionSeverityPtrOutput {
	return e.ToEventSubscriptionSeverityPtrOutputWithContext(context.Background())
}

func (e EventSubscriptionSeverity) ToEventSubscriptionSeverityPtrOutputWithContext(ctx context.Context) EventSubscriptionSeverityPtrOutput {
	return EventSubscriptionSeverity(e).ToEventSubscriptionSeverityOutputWithContext(ctx).ToEventSubscriptionSeverityPtrOutputWithContext(ctx)
}

func (e EventSubscriptionSeverity) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EventSubscriptionSeverity) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EventSubscriptionSeverity) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EventSubscriptionSeverity) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EventSubscriptionSeverityOutput struct{ *pulumi.OutputState }

func (EventSubscriptionSeverityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSubscriptionSeverity)(nil)).Elem()
}

func (o EventSubscriptionSeverityOutput) ToEventSubscriptionSeverityOutput() EventSubscriptionSeverityOutput {
	return o
}

func (o EventSubscriptionSeverityOutput) ToEventSubscriptionSeverityOutputWithContext(ctx context.Context) EventSubscriptionSeverityOutput {
	return o
}

func (o EventSubscriptionSeverityOutput) ToEventSubscriptionSeverityPtrOutput() EventSubscriptionSeverityPtrOutput {
	return o.ToEventSubscriptionSeverityPtrOutputWithContext(context.Background())
}

func (o EventSubscriptionSeverityOutput) ToEventSubscriptionSeverityPtrOutputWithContext(ctx context.Context) EventSubscriptionSeverityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EventSubscriptionSeverity) *EventSubscriptionSeverity {
		return &v
	}).(EventSubscriptionSeverityPtrOutput)
}

func (o EventSubscriptionSeverityOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EventSubscriptionSeverityOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EventSubscriptionSeverity) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EventSubscriptionSeverityOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EventSubscriptionSeverityOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EventSubscriptionSeverity) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EventSubscriptionSeverityPtrOutput struct{ *pulumi.OutputState }

func (EventSubscriptionSeverityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSubscriptionSeverity)(nil)).Elem()
}

func (o EventSubscriptionSeverityPtrOutput) ToEventSubscriptionSeverityPtrOutput() EventSubscriptionSeverityPtrOutput {
	return o
}

func (o EventSubscriptionSeverityPtrOutput) ToEventSubscriptionSeverityPtrOutputWithContext(ctx context.Context) EventSubscriptionSeverityPtrOutput {
	return o
}

func (o EventSubscriptionSeverityPtrOutput) Elem() EventSubscriptionSeverityOutput {
	return o.ApplyT(func(v *EventSubscriptionSeverity) EventSubscriptionSeverity {
		if v != nil {
			return *v
		}
		var ret EventSubscriptionSeverity
		return ret
	}).(EventSubscriptionSeverityOutput)
}

func (o EventSubscriptionSeverityPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EventSubscriptionSeverityPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EventSubscriptionSeverity) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// EventSubscriptionSeverityInput is an input type that accepts EventSubscriptionSeverityArgs and EventSubscriptionSeverityOutput values.
// You can construct a concrete instance of `EventSubscriptionSeverityInput` via:
//
//          EventSubscriptionSeverityArgs{...}
type EventSubscriptionSeverityInput interface {
	pulumi.Input

	ToEventSubscriptionSeverityOutput() EventSubscriptionSeverityOutput
	ToEventSubscriptionSeverityOutputWithContext(context.Context) EventSubscriptionSeverityOutput
}

var eventSubscriptionSeverityPtrType = reflect.TypeOf((**EventSubscriptionSeverity)(nil)).Elem()

type EventSubscriptionSeverityPtrInput interface {
	pulumi.Input

	ToEventSubscriptionSeverityPtrOutput() EventSubscriptionSeverityPtrOutput
	ToEventSubscriptionSeverityPtrOutputWithContext(context.Context) EventSubscriptionSeverityPtrOutput
}

type eventSubscriptionSeverityPtr string

func EventSubscriptionSeverityPtr(v string) EventSubscriptionSeverityPtrInput {
	return (*eventSubscriptionSeverityPtr)(&v)
}

func (*eventSubscriptionSeverityPtr) ElementType() reflect.Type {
	return eventSubscriptionSeverityPtrType
}

func (in *eventSubscriptionSeverityPtr) ToEventSubscriptionSeverityPtrOutput() EventSubscriptionSeverityPtrOutput {
	return pulumi.ToOutput(in).(EventSubscriptionSeverityPtrOutput)
}

func (in *eventSubscriptionSeverityPtr) ToEventSubscriptionSeverityPtrOutputWithContext(ctx context.Context) EventSubscriptionSeverityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EventSubscriptionSeverityPtrOutput)
}

// The type of source that will be generating the events.
type EventSubscriptionSourceType string

const (
	EventSubscriptionSourceTypeCluster               = EventSubscriptionSourceType("cluster")
	EventSubscriptionSourceTypeClusterParameterGroup = EventSubscriptionSourceType("cluster-parameter-group")
	EventSubscriptionSourceTypeClusterSecurityGroup  = EventSubscriptionSourceType("cluster-security-group")
	EventSubscriptionSourceTypeClusterSnapshot       = EventSubscriptionSourceType("cluster-snapshot")
	EventSubscriptionSourceTypeScheduledAction       = EventSubscriptionSourceType("scheduled-action")
)

func (EventSubscriptionSourceType) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSubscriptionSourceType)(nil)).Elem()
}

func (e EventSubscriptionSourceType) ToEventSubscriptionSourceTypeOutput() EventSubscriptionSourceTypeOutput {
	return pulumi.ToOutput(e).(EventSubscriptionSourceTypeOutput)
}

func (e EventSubscriptionSourceType) ToEventSubscriptionSourceTypeOutputWithContext(ctx context.Context) EventSubscriptionSourceTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EventSubscriptionSourceTypeOutput)
}

func (e EventSubscriptionSourceType) ToEventSubscriptionSourceTypePtrOutput() EventSubscriptionSourceTypePtrOutput {
	return e.ToEventSubscriptionSourceTypePtrOutputWithContext(context.Background())
}

func (e EventSubscriptionSourceType) ToEventSubscriptionSourceTypePtrOutputWithContext(ctx context.Context) EventSubscriptionSourceTypePtrOutput {
	return EventSubscriptionSourceType(e).ToEventSubscriptionSourceTypeOutputWithContext(ctx).ToEventSubscriptionSourceTypePtrOutputWithContext(ctx)
}

func (e EventSubscriptionSourceType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EventSubscriptionSourceType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EventSubscriptionSourceType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EventSubscriptionSourceType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EventSubscriptionSourceTypeOutput struct{ *pulumi.OutputState }

func (EventSubscriptionSourceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSubscriptionSourceType)(nil)).Elem()
}

func (o EventSubscriptionSourceTypeOutput) ToEventSubscriptionSourceTypeOutput() EventSubscriptionSourceTypeOutput {
	return o
}

func (o EventSubscriptionSourceTypeOutput) ToEventSubscriptionSourceTypeOutputWithContext(ctx context.Context) EventSubscriptionSourceTypeOutput {
	return o
}

func (o EventSubscriptionSourceTypeOutput) ToEventSubscriptionSourceTypePtrOutput() EventSubscriptionSourceTypePtrOutput {
	return o.ToEventSubscriptionSourceTypePtrOutputWithContext(context.Background())
}

func (o EventSubscriptionSourceTypeOutput) ToEventSubscriptionSourceTypePtrOutputWithContext(ctx context.Context) EventSubscriptionSourceTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EventSubscriptionSourceType) *EventSubscriptionSourceType {
		return &v
	}).(EventSubscriptionSourceTypePtrOutput)
}

func (o EventSubscriptionSourceTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EventSubscriptionSourceTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EventSubscriptionSourceType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EventSubscriptionSourceTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EventSubscriptionSourceTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EventSubscriptionSourceType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EventSubscriptionSourceTypePtrOutput struct{ *pulumi.OutputState }

func (EventSubscriptionSourceTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSubscriptionSourceType)(nil)).Elem()
}

func (o EventSubscriptionSourceTypePtrOutput) ToEventSubscriptionSourceTypePtrOutput() EventSubscriptionSourceTypePtrOutput {
	return o
}

func (o EventSubscriptionSourceTypePtrOutput) ToEventSubscriptionSourceTypePtrOutputWithContext(ctx context.Context) EventSubscriptionSourceTypePtrOutput {
	return o
}

func (o EventSubscriptionSourceTypePtrOutput) Elem() EventSubscriptionSourceTypeOutput {
	return o.ApplyT(func(v *EventSubscriptionSourceType) EventSubscriptionSourceType {
		if v != nil {
			return *v
		}
		var ret EventSubscriptionSourceType
		return ret
	}).(EventSubscriptionSourceTypeOutput)
}

func (o EventSubscriptionSourceTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EventSubscriptionSourceTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EventSubscriptionSourceType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// EventSubscriptionSourceTypeInput is an input type that accepts EventSubscriptionSourceTypeArgs and EventSubscriptionSourceTypeOutput values.
// You can construct a concrete instance of `EventSubscriptionSourceTypeInput` via:
//
//          EventSubscriptionSourceTypeArgs{...}
type EventSubscriptionSourceTypeInput interface {
	pulumi.Input

	ToEventSubscriptionSourceTypeOutput() EventSubscriptionSourceTypeOutput
	ToEventSubscriptionSourceTypeOutputWithContext(context.Context) EventSubscriptionSourceTypeOutput
}

var eventSubscriptionSourceTypePtrType = reflect.TypeOf((**EventSubscriptionSourceType)(nil)).Elem()

type EventSubscriptionSourceTypePtrInput interface {
	pulumi.Input

	ToEventSubscriptionSourceTypePtrOutput() EventSubscriptionSourceTypePtrOutput
	ToEventSubscriptionSourceTypePtrOutputWithContext(context.Context) EventSubscriptionSourceTypePtrOutput
}

type eventSubscriptionSourceTypePtr string

func EventSubscriptionSourceTypePtr(v string) EventSubscriptionSourceTypePtrInput {
	return (*eventSubscriptionSourceTypePtr)(&v)
}

func (*eventSubscriptionSourceTypePtr) ElementType() reflect.Type {
	return eventSubscriptionSourceTypePtrType
}

func (in *eventSubscriptionSourceTypePtr) ToEventSubscriptionSourceTypePtrOutput() EventSubscriptionSourceTypePtrOutput {
	return pulumi.ToOutput(in).(EventSubscriptionSourceTypePtrOutput)
}

func (in *eventSubscriptionSourceTypePtr) ToEventSubscriptionSourceTypePtrOutputWithContext(ctx context.Context) EventSubscriptionSourceTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EventSubscriptionSourceTypePtrOutput)
}

// The status of the Amazon Redshift event notification subscription.
type EventSubscriptionStatus string

const (
	EventSubscriptionStatusActive        = EventSubscriptionStatus("active")
	EventSubscriptionStatusNoPermission  = EventSubscriptionStatus("no-permission")
	EventSubscriptionStatusTopicNotExist = EventSubscriptionStatus("topic-not-exist")
)

type EventSubscriptionStatusOutput struct{ *pulumi.OutputState }

func (EventSubscriptionStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSubscriptionStatus)(nil)).Elem()
}

func (o EventSubscriptionStatusOutput) ToEventSubscriptionStatusOutput() EventSubscriptionStatusOutput {
	return o
}

func (o EventSubscriptionStatusOutput) ToEventSubscriptionStatusOutputWithContext(ctx context.Context) EventSubscriptionStatusOutput {
	return o
}

func (o EventSubscriptionStatusOutput) ToEventSubscriptionStatusPtrOutput() EventSubscriptionStatusPtrOutput {
	return o.ToEventSubscriptionStatusPtrOutputWithContext(context.Background())
}

func (o EventSubscriptionStatusOutput) ToEventSubscriptionStatusPtrOutputWithContext(ctx context.Context) EventSubscriptionStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EventSubscriptionStatus) *EventSubscriptionStatus {
		return &v
	}).(EventSubscriptionStatusPtrOutput)
}

func (o EventSubscriptionStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EventSubscriptionStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EventSubscriptionStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EventSubscriptionStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EventSubscriptionStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EventSubscriptionStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EventSubscriptionStatusPtrOutput struct{ *pulumi.OutputState }

func (EventSubscriptionStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSubscriptionStatus)(nil)).Elem()
}

func (o EventSubscriptionStatusPtrOutput) ToEventSubscriptionStatusPtrOutput() EventSubscriptionStatusPtrOutput {
	return o
}

func (o EventSubscriptionStatusPtrOutput) ToEventSubscriptionStatusPtrOutputWithContext(ctx context.Context) EventSubscriptionStatusPtrOutput {
	return o
}

func (o EventSubscriptionStatusPtrOutput) Elem() EventSubscriptionStatusOutput {
	return o.ApplyT(func(v *EventSubscriptionStatus) EventSubscriptionStatus {
		if v != nil {
			return *v
		}
		var ret EventSubscriptionStatus
		return ret
	}).(EventSubscriptionStatusOutput)
}

func (o EventSubscriptionStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EventSubscriptionStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EventSubscriptionStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// The state of the scheduled action.
type ScheduledActionStateEnum string

const (
	ScheduledActionStateEnumActive   = ScheduledActionStateEnum("ACTIVE")
	ScheduledActionStateEnumDisabled = ScheduledActionStateEnum("DISABLED")
)

type ScheduledActionStateEnumOutput struct{ *pulumi.OutputState }

func (ScheduledActionStateEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledActionStateEnum)(nil)).Elem()
}

func (o ScheduledActionStateEnumOutput) ToScheduledActionStateEnumOutput() ScheduledActionStateEnumOutput {
	return o
}

func (o ScheduledActionStateEnumOutput) ToScheduledActionStateEnumOutputWithContext(ctx context.Context) ScheduledActionStateEnumOutput {
	return o
}

func (o ScheduledActionStateEnumOutput) ToScheduledActionStateEnumPtrOutput() ScheduledActionStateEnumPtrOutput {
	return o.ToScheduledActionStateEnumPtrOutputWithContext(context.Background())
}

func (o ScheduledActionStateEnumOutput) ToScheduledActionStateEnumPtrOutputWithContext(ctx context.Context) ScheduledActionStateEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScheduledActionStateEnum) *ScheduledActionStateEnum {
		return &v
	}).(ScheduledActionStateEnumPtrOutput)
}

func (o ScheduledActionStateEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ScheduledActionStateEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScheduledActionStateEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ScheduledActionStateEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScheduledActionStateEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScheduledActionStateEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScheduledActionStateEnumPtrOutput struct{ *pulumi.OutputState }

func (ScheduledActionStateEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledActionStateEnum)(nil)).Elem()
}

func (o ScheduledActionStateEnumPtrOutput) ToScheduledActionStateEnumPtrOutput() ScheduledActionStateEnumPtrOutput {
	return o
}

func (o ScheduledActionStateEnumPtrOutput) ToScheduledActionStateEnumPtrOutputWithContext(ctx context.Context) ScheduledActionStateEnumPtrOutput {
	return o
}

func (o ScheduledActionStateEnumPtrOutput) Elem() ScheduledActionStateEnumOutput {
	return o.ApplyT(func(v *ScheduledActionStateEnum) ScheduledActionStateEnum {
		if v != nil {
			return *v
		}
		var ret ScheduledActionStateEnum
		return ret
	}).(ScheduledActionStateEnumOutput)
}

func (o ScheduledActionStateEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScheduledActionStateEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ScheduledActionStateEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EventSubscriptionEventCategoriesItemInput)(nil)).Elem(), EventSubscriptionEventCategoriesItem("configuration"))
	pulumi.RegisterInputType(reflect.TypeOf((*EventSubscriptionEventCategoriesItemPtrInput)(nil)).Elem(), EventSubscriptionEventCategoriesItem("configuration"))
	pulumi.RegisterInputType(reflect.TypeOf((*EventSubscriptionEventCategoriesItemArrayInput)(nil)).Elem(), EventSubscriptionEventCategoriesItemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventSubscriptionSeverityInput)(nil)).Elem(), EventSubscriptionSeverity("ERROR"))
	pulumi.RegisterInputType(reflect.TypeOf((*EventSubscriptionSeverityPtrInput)(nil)).Elem(), EventSubscriptionSeverity("ERROR"))
	pulumi.RegisterInputType(reflect.TypeOf((*EventSubscriptionSourceTypeInput)(nil)).Elem(), EventSubscriptionSourceType("cluster"))
	pulumi.RegisterInputType(reflect.TypeOf((*EventSubscriptionSourceTypePtrInput)(nil)).Elem(), EventSubscriptionSourceType("cluster"))
	pulumi.RegisterOutputType(EventSubscriptionEventCategoriesItemOutput{})
	pulumi.RegisterOutputType(EventSubscriptionEventCategoriesItemPtrOutput{})
	pulumi.RegisterOutputType(EventSubscriptionEventCategoriesItemArrayOutput{})
	pulumi.RegisterOutputType(EventSubscriptionSeverityOutput{})
	pulumi.RegisterOutputType(EventSubscriptionSeverityPtrOutput{})
	pulumi.RegisterOutputType(EventSubscriptionSourceTypeOutput{})
	pulumi.RegisterOutputType(EventSubscriptionSourceTypePtrOutput{})
	pulumi.RegisterOutputType(EventSubscriptionStatusOutput{})
	pulumi.RegisterOutputType(EventSubscriptionStatusPtrOutput{})
	pulumi.RegisterOutputType(ScheduledActionStateEnumOutput{})
	pulumi.RegisterOutputType(ScheduledActionStateEnumPtrOutput{})
}
