// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package docdb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::DocDB::DBCluster
//
// Deprecated: DBCluster is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type DBCluster struct {
	pulumi.CustomResourceState

	AvailabilityZones           pulumi.StringArrayOutput `pulumi:"availabilityZones"`
	BackupRetentionPeriod       pulumi.IntPtrOutput      `pulumi:"backupRetentionPeriod"`
	ClusterResourceId           pulumi.StringOutput      `pulumi:"clusterResourceId"`
	CopyTagsToSnapshot          pulumi.BoolPtrOutput     `pulumi:"copyTagsToSnapshot"`
	DBClusterIdentifier         pulumi.StringPtrOutput   `pulumi:"dBClusterIdentifier"`
	DBClusterParameterGroupName pulumi.StringPtrOutput   `pulumi:"dBClusterParameterGroupName"`
	DBSubnetGroupName           pulumi.StringPtrOutput   `pulumi:"dBSubnetGroupName"`
	DeletionProtection          pulumi.BoolPtrOutput     `pulumi:"deletionProtection"`
	EnableCloudwatchLogsExports pulumi.StringArrayOutput `pulumi:"enableCloudwatchLogsExports"`
	Endpoint                    pulumi.StringOutput      `pulumi:"endpoint"`
	EngineVersion               pulumi.StringPtrOutput   `pulumi:"engineVersion"`
	KmsKeyId                    pulumi.StringPtrOutput   `pulumi:"kmsKeyId"`
	MasterUserPassword          pulumi.StringPtrOutput   `pulumi:"masterUserPassword"`
	MasterUsername              pulumi.StringPtrOutput   `pulumi:"masterUsername"`
	Port                        pulumi.IntPtrOutput      `pulumi:"port"`
	PreferredBackupWindow       pulumi.StringPtrOutput   `pulumi:"preferredBackupWindow"`
	PreferredMaintenanceWindow  pulumi.StringPtrOutput   `pulumi:"preferredMaintenanceWindow"`
	ReadEndpoint                pulumi.StringOutput      `pulumi:"readEndpoint"`
	SnapshotIdentifier          pulumi.StringPtrOutput   `pulumi:"snapshotIdentifier"`
	StorageEncrypted            pulumi.BoolPtrOutput     `pulumi:"storageEncrypted"`
	Tags                        DBClusterTagArrayOutput  `pulumi:"tags"`
	VpcSecurityGroupIds         pulumi.StringArrayOutput `pulumi:"vpcSecurityGroupIds"`
}

// NewDBCluster registers a new resource with the given unique name, arguments, and options.
func NewDBCluster(ctx *pulumi.Context,
	name string, args *DBClusterArgs, opts ...pulumi.ResourceOption) (*DBCluster, error) {
	if args == nil {
		args = &DBClusterArgs{}
	}

	var resource DBCluster
	err := ctx.RegisterResource("aws-native:docdb:DBCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDBCluster gets an existing DBCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDBCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DBClusterState, opts ...pulumi.ResourceOption) (*DBCluster, error) {
	var resource DBCluster
	err := ctx.ReadResource("aws-native:docdb:DBCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DBCluster resources.
type dbclusterState struct {
}

type DBClusterState struct {
}

func (DBClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*dbclusterState)(nil)).Elem()
}

type dbclusterArgs struct {
	AvailabilityZones           []string       `pulumi:"availabilityZones"`
	BackupRetentionPeriod       *int           `pulumi:"backupRetentionPeriod"`
	CopyTagsToSnapshot          *bool          `pulumi:"copyTagsToSnapshot"`
	DBClusterIdentifier         *string        `pulumi:"dBClusterIdentifier"`
	DBClusterParameterGroupName *string        `pulumi:"dBClusterParameterGroupName"`
	DBSubnetGroupName           *string        `pulumi:"dBSubnetGroupName"`
	DeletionProtection          *bool          `pulumi:"deletionProtection"`
	EnableCloudwatchLogsExports []string       `pulumi:"enableCloudwatchLogsExports"`
	EngineVersion               *string        `pulumi:"engineVersion"`
	KmsKeyId                    *string        `pulumi:"kmsKeyId"`
	MasterUserPassword          *string        `pulumi:"masterUserPassword"`
	MasterUsername              *string        `pulumi:"masterUsername"`
	Port                        *int           `pulumi:"port"`
	PreferredBackupWindow       *string        `pulumi:"preferredBackupWindow"`
	PreferredMaintenanceWindow  *string        `pulumi:"preferredMaintenanceWindow"`
	SnapshotIdentifier          *string        `pulumi:"snapshotIdentifier"`
	StorageEncrypted            *bool          `pulumi:"storageEncrypted"`
	Tags                        []DBClusterTag `pulumi:"tags"`
	VpcSecurityGroupIds         []string       `pulumi:"vpcSecurityGroupIds"`
}

// The set of arguments for constructing a DBCluster resource.
type DBClusterArgs struct {
	AvailabilityZones           pulumi.StringArrayInput
	BackupRetentionPeriod       pulumi.IntPtrInput
	CopyTagsToSnapshot          pulumi.BoolPtrInput
	DBClusterIdentifier         pulumi.StringPtrInput
	DBClusterParameterGroupName pulumi.StringPtrInput
	DBSubnetGroupName           pulumi.StringPtrInput
	DeletionProtection          pulumi.BoolPtrInput
	EnableCloudwatchLogsExports pulumi.StringArrayInput
	EngineVersion               pulumi.StringPtrInput
	KmsKeyId                    pulumi.StringPtrInput
	MasterUserPassword          pulumi.StringPtrInput
	MasterUsername              pulumi.StringPtrInput
	Port                        pulumi.IntPtrInput
	PreferredBackupWindow       pulumi.StringPtrInput
	PreferredMaintenanceWindow  pulumi.StringPtrInput
	SnapshotIdentifier          pulumi.StringPtrInput
	StorageEncrypted            pulumi.BoolPtrInput
	Tags                        DBClusterTagArrayInput
	VpcSecurityGroupIds         pulumi.StringArrayInput
}

func (DBClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dbclusterArgs)(nil)).Elem()
}

type DBClusterInput interface {
	pulumi.Input

	ToDBClusterOutput() DBClusterOutput
	ToDBClusterOutputWithContext(ctx context.Context) DBClusterOutput
}

func (*DBCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**DBCluster)(nil)).Elem()
}

func (i *DBCluster) ToDBClusterOutput() DBClusterOutput {
	return i.ToDBClusterOutputWithContext(context.Background())
}

func (i *DBCluster) ToDBClusterOutputWithContext(ctx context.Context) DBClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBClusterOutput)
}

type DBClusterOutput struct{ *pulumi.OutputState }

func (DBClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DBCluster)(nil)).Elem()
}

func (o DBClusterOutput) ToDBClusterOutput() DBClusterOutput {
	return o
}

func (o DBClusterOutput) ToDBClusterOutputWithContext(ctx context.Context) DBClusterOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DBClusterInput)(nil)).Elem(), &DBCluster{})
	pulumi.RegisterOutputType(DBClusterOutput{})
}
