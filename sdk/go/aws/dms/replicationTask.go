// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dms

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html
type ReplicationTask struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-cdcstartposition
	CdcStartPosition pulumi.StringPtrOutput `pulumi:"cdcStartPosition"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-cdcstarttime
	CdcStartTime pulumi.Float64PtrOutput `pulumi:"cdcStartTime"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-cdcstopposition
	CdcStopPosition pulumi.StringPtrOutput `pulumi:"cdcStopPosition"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-migrationtype
	MigrationType pulumi.StringOutput `pulumi:"migrationType"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-replicationinstancearn
	ReplicationInstanceArn pulumi.StringOutput `pulumi:"replicationInstanceArn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-replicationtaskidentifier
	ReplicationTaskIdentifier pulumi.StringPtrOutput `pulumi:"replicationTaskIdentifier"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-replicationtasksettings
	ReplicationTaskSettings pulumi.StringPtrOutput `pulumi:"replicationTaskSettings"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-resourceidentifier
	ResourceIdentifier pulumi.StringPtrOutput `pulumi:"resourceIdentifier"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-sourceendpointarn
	SourceEndpointArn pulumi.StringOutput `pulumi:"sourceEndpointArn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-tablemappings
	TableMappings pulumi.StringOutput `pulumi:"tableMappings"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-tags
	Tags aws.TagArrayOutput `pulumi:"tags"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-targetendpointarn
	TargetEndpointArn pulumi.StringOutput `pulumi:"targetEndpointArn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-taskdata
	TaskData pulumi.StringPtrOutput `pulumi:"taskData"`
}

// NewReplicationTask registers a new resource with the given unique name, arguments, and options.
func NewReplicationTask(ctx *pulumi.Context,
	name string, args *ReplicationTaskArgs, opts ...pulumi.ResourceOption) (*ReplicationTask, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MigrationType == nil {
		return nil, errors.New("invalid value for required argument 'MigrationType'")
	}
	if args.ReplicationInstanceArn == nil {
		return nil, errors.New("invalid value for required argument 'ReplicationInstanceArn'")
	}
	if args.SourceEndpointArn == nil {
		return nil, errors.New("invalid value for required argument 'SourceEndpointArn'")
	}
	if args.TableMappings == nil {
		return nil, errors.New("invalid value for required argument 'TableMappings'")
	}
	if args.TargetEndpointArn == nil {
		return nil, errors.New("invalid value for required argument 'TargetEndpointArn'")
	}
	var resource ReplicationTask
	err := ctx.RegisterResource("aws-native:dms:ReplicationTask", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationTask gets an existing ReplicationTask resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationTask(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationTaskState, opts ...pulumi.ResourceOption) (*ReplicationTask, error) {
	var resource ReplicationTask
	err := ctx.ReadResource("aws-native:dms:ReplicationTask", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicationTask resources.
type replicationTaskState struct {
}

type ReplicationTaskState struct {
}

func (ReplicationTaskState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationTaskState)(nil)).Elem()
}

type replicationTaskArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-cdcstartposition
	CdcStartPosition *string `pulumi:"cdcStartPosition"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-cdcstarttime
	CdcStartTime *float64 `pulumi:"cdcStartTime"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-cdcstopposition
	CdcStopPosition *string `pulumi:"cdcStopPosition"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-migrationtype
	MigrationType string `pulumi:"migrationType"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-replicationinstancearn
	ReplicationInstanceArn string `pulumi:"replicationInstanceArn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-replicationtaskidentifier
	ReplicationTaskIdentifier *string `pulumi:"replicationTaskIdentifier"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-replicationtasksettings
	ReplicationTaskSettings *string `pulumi:"replicationTaskSettings"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-resourceidentifier
	ResourceIdentifier *string `pulumi:"resourceIdentifier"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-sourceendpointarn
	SourceEndpointArn string `pulumi:"sourceEndpointArn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-tablemappings
	TableMappings string `pulumi:"tableMappings"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-tags
	Tags []aws.Tag `pulumi:"tags"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-targetendpointarn
	TargetEndpointArn string `pulumi:"targetEndpointArn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-taskdata
	TaskData *string `pulumi:"taskData"`
}

// The set of arguments for constructing a ReplicationTask resource.
type ReplicationTaskArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-cdcstartposition
	CdcStartPosition pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-cdcstarttime
	CdcStartTime pulumi.Float64PtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-cdcstopposition
	CdcStopPosition pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-migrationtype
	MigrationType pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-replicationinstancearn
	ReplicationInstanceArn pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-replicationtaskidentifier
	ReplicationTaskIdentifier pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-replicationtasksettings
	ReplicationTaskSettings pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-resourceidentifier
	ResourceIdentifier pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-sourceendpointarn
	SourceEndpointArn pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-tablemappings
	TableMappings pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-tags
	Tags aws.TagArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-targetendpointarn
	TargetEndpointArn pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html#cfn-dms-replicationtask-taskdata
	TaskData pulumi.StringPtrInput
}

func (ReplicationTaskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationTaskArgs)(nil)).Elem()
}

type ReplicationTaskInput interface {
	pulumi.Input

	ToReplicationTaskOutput() ReplicationTaskOutput
	ToReplicationTaskOutputWithContext(ctx context.Context) ReplicationTaskOutput
}

func (*ReplicationTask) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationTask)(nil))
}

func (i *ReplicationTask) ToReplicationTaskOutput() ReplicationTaskOutput {
	return i.ToReplicationTaskOutputWithContext(context.Background())
}

func (i *ReplicationTask) ToReplicationTaskOutputWithContext(ctx context.Context) ReplicationTaskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationTaskOutput)
}

type ReplicationTaskOutput struct{ *pulumi.OutputState }

func (ReplicationTaskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationTask)(nil))
}

func (o ReplicationTaskOutput) ToReplicationTaskOutput() ReplicationTaskOutput {
	return o
}

func (o ReplicationTaskOutput) ToReplicationTaskOutputWithContext(ctx context.Context) ReplicationTaskOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ReplicationTaskOutput{})
}