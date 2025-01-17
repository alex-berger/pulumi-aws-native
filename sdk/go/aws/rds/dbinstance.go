// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::RDS::DBInstance
//
// Deprecated: DBInstance is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type DBInstance struct {
	pulumi.CustomResourceState

	AllocatedStorage                   pulumi.StringPtrOutput                `pulumi:"allocatedStorage"`
	AllowMajorVersionUpgrade           pulumi.BoolPtrOutput                  `pulumi:"allowMajorVersionUpgrade"`
	AssociatedRoles                    DBInstanceRoleArrayOutput             `pulumi:"associatedRoles"`
	AutoMinorVersionUpgrade            pulumi.BoolPtrOutput                  `pulumi:"autoMinorVersionUpgrade"`
	AvailabilityZone                   pulumi.StringPtrOutput                `pulumi:"availabilityZone"`
	BackupRetentionPeriod              pulumi.IntPtrOutput                   `pulumi:"backupRetentionPeriod"`
	CACertificateIdentifier            pulumi.StringPtrOutput                `pulumi:"cACertificateIdentifier"`
	CharacterSetName                   pulumi.StringPtrOutput                `pulumi:"characterSetName"`
	CopyTagsToSnapshot                 pulumi.BoolPtrOutput                  `pulumi:"copyTagsToSnapshot"`
	DBClusterIdentifier                pulumi.StringPtrOutput                `pulumi:"dBClusterIdentifier"`
	DBInstanceClass                    pulumi.StringOutput                   `pulumi:"dBInstanceClass"`
	DBInstanceIdentifier               pulumi.StringPtrOutput                `pulumi:"dBInstanceIdentifier"`
	DBName                             pulumi.StringPtrOutput                `pulumi:"dBName"`
	DBParameterGroupName               pulumi.StringPtrOutput                `pulumi:"dBParameterGroupName"`
	DBSecurityGroups                   pulumi.StringArrayOutput              `pulumi:"dBSecurityGroups"`
	DBSnapshotIdentifier               pulumi.StringPtrOutput                `pulumi:"dBSnapshotIdentifier"`
	DBSubnetGroupName                  pulumi.StringPtrOutput                `pulumi:"dBSubnetGroupName"`
	DeleteAutomatedBackups             pulumi.BoolPtrOutput                  `pulumi:"deleteAutomatedBackups"`
	DeletionProtection                 pulumi.BoolPtrOutput                  `pulumi:"deletionProtection"`
	Domain                             pulumi.StringPtrOutput                `pulumi:"domain"`
	DomainIAMRoleName                  pulumi.StringPtrOutput                `pulumi:"domainIAMRoleName"`
	EnableCloudwatchLogsExports        pulumi.StringArrayOutput              `pulumi:"enableCloudwatchLogsExports"`
	EnableIAMDatabaseAuthentication    pulumi.BoolPtrOutput                  `pulumi:"enableIAMDatabaseAuthentication"`
	EnablePerformanceInsights          pulumi.BoolPtrOutput                  `pulumi:"enablePerformanceInsights"`
	EndpointAddress                    pulumi.StringPtrOutput                `pulumi:"endpointAddress"`
	EndpointPort                       pulumi.StringPtrOutput                `pulumi:"endpointPort"`
	Engine                             pulumi.StringPtrOutput                `pulumi:"engine"`
	EngineVersion                      pulumi.StringPtrOutput                `pulumi:"engineVersion"`
	Iops                               pulumi.IntPtrOutput                   `pulumi:"iops"`
	KmsKeyId                           pulumi.StringPtrOutput                `pulumi:"kmsKeyId"`
	LicenseModel                       pulumi.StringPtrOutput                `pulumi:"licenseModel"`
	MasterUserPassword                 pulumi.StringPtrOutput                `pulumi:"masterUserPassword"`
	MasterUsername                     pulumi.StringPtrOutput                `pulumi:"masterUsername"`
	MaxAllocatedStorage                pulumi.IntPtrOutput                   `pulumi:"maxAllocatedStorage"`
	MonitoringInterval                 pulumi.IntPtrOutput                   `pulumi:"monitoringInterval"`
	MonitoringRoleArn                  pulumi.StringPtrOutput                `pulumi:"monitoringRoleArn"`
	MultiAZ                            pulumi.BoolPtrOutput                  `pulumi:"multiAZ"`
	OptionGroupName                    pulumi.StringPtrOutput                `pulumi:"optionGroupName"`
	PerformanceInsightsKMSKeyId        pulumi.StringPtrOutput                `pulumi:"performanceInsightsKMSKeyId"`
	PerformanceInsightsRetentionPeriod pulumi.IntPtrOutput                   `pulumi:"performanceInsightsRetentionPeriod"`
	Port                               pulumi.StringPtrOutput                `pulumi:"port"`
	PreferredBackupWindow              pulumi.StringPtrOutput                `pulumi:"preferredBackupWindow"`
	PreferredMaintenanceWindow         pulumi.StringPtrOutput                `pulumi:"preferredMaintenanceWindow"`
	ProcessorFeatures                  DBInstanceProcessorFeatureArrayOutput `pulumi:"processorFeatures"`
	PromotionTier                      pulumi.IntPtrOutput                   `pulumi:"promotionTier"`
	PubliclyAccessible                 pulumi.BoolPtrOutput                  `pulumi:"publiclyAccessible"`
	SourceDBInstanceIdentifier         pulumi.StringPtrOutput                `pulumi:"sourceDBInstanceIdentifier"`
	SourceRegion                       pulumi.StringPtrOutput                `pulumi:"sourceRegion"`
	StorageEncrypted                   pulumi.BoolPtrOutput                  `pulumi:"storageEncrypted"`
	StorageType                        pulumi.StringPtrOutput                `pulumi:"storageType"`
	Tags                               DBInstanceTagArrayOutput              `pulumi:"tags"`
	Timezone                           pulumi.StringPtrOutput                `pulumi:"timezone"`
	UseDefaultProcessorFeatures        pulumi.BoolPtrOutput                  `pulumi:"useDefaultProcessorFeatures"`
	VPCSecurityGroups                  pulumi.StringArrayOutput              `pulumi:"vPCSecurityGroups"`
}

// NewDBInstance registers a new resource with the given unique name, arguments, and options.
func NewDBInstance(ctx *pulumi.Context,
	name string, args *DBInstanceArgs, opts ...pulumi.ResourceOption) (*DBInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DBInstanceClass == nil {
		return nil, errors.New("invalid value for required argument 'DBInstanceClass'")
	}
	var resource DBInstance
	err := ctx.RegisterResource("aws-native:rds:DBInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDBInstance gets an existing DBInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDBInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DBInstanceState, opts ...pulumi.ResourceOption) (*DBInstance, error) {
	var resource DBInstance
	err := ctx.ReadResource("aws-native:rds:DBInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DBInstance resources.
type dbinstanceState struct {
}

type DBInstanceState struct {
}

func (DBInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*dbinstanceState)(nil)).Elem()
}

type dbinstanceArgs struct {
	AllocatedStorage                   *string                      `pulumi:"allocatedStorage"`
	AllowMajorVersionUpgrade           *bool                        `pulumi:"allowMajorVersionUpgrade"`
	AssociatedRoles                    []DBInstanceRole             `pulumi:"associatedRoles"`
	AutoMinorVersionUpgrade            *bool                        `pulumi:"autoMinorVersionUpgrade"`
	AvailabilityZone                   *string                      `pulumi:"availabilityZone"`
	BackupRetentionPeriod              *int                         `pulumi:"backupRetentionPeriod"`
	CACertificateIdentifier            *string                      `pulumi:"cACertificateIdentifier"`
	CharacterSetName                   *string                      `pulumi:"characterSetName"`
	CopyTagsToSnapshot                 *bool                        `pulumi:"copyTagsToSnapshot"`
	DBClusterIdentifier                *string                      `pulumi:"dBClusterIdentifier"`
	DBInstanceClass                    string                       `pulumi:"dBInstanceClass"`
	DBInstanceIdentifier               *string                      `pulumi:"dBInstanceIdentifier"`
	DBName                             *string                      `pulumi:"dBName"`
	DBParameterGroupName               *string                      `pulumi:"dBParameterGroupName"`
	DBSecurityGroups                   []string                     `pulumi:"dBSecurityGroups"`
	DBSnapshotIdentifier               *string                      `pulumi:"dBSnapshotIdentifier"`
	DBSubnetGroupName                  *string                      `pulumi:"dBSubnetGroupName"`
	DeleteAutomatedBackups             *bool                        `pulumi:"deleteAutomatedBackups"`
	DeletionProtection                 *bool                        `pulumi:"deletionProtection"`
	Domain                             *string                      `pulumi:"domain"`
	DomainIAMRoleName                  *string                      `pulumi:"domainIAMRoleName"`
	EnableCloudwatchLogsExports        []string                     `pulumi:"enableCloudwatchLogsExports"`
	EnableIAMDatabaseAuthentication    *bool                        `pulumi:"enableIAMDatabaseAuthentication"`
	EnablePerformanceInsights          *bool                        `pulumi:"enablePerformanceInsights"`
	EndpointAddress                    *string                      `pulumi:"endpointAddress"`
	EndpointPort                       *string                      `pulumi:"endpointPort"`
	Engine                             *string                      `pulumi:"engine"`
	EngineVersion                      *string                      `pulumi:"engineVersion"`
	Iops                               *int                         `pulumi:"iops"`
	KmsKeyId                           *string                      `pulumi:"kmsKeyId"`
	LicenseModel                       *string                      `pulumi:"licenseModel"`
	MasterUserPassword                 *string                      `pulumi:"masterUserPassword"`
	MasterUsername                     *string                      `pulumi:"masterUsername"`
	MaxAllocatedStorage                *int                         `pulumi:"maxAllocatedStorage"`
	MonitoringInterval                 *int                         `pulumi:"monitoringInterval"`
	MonitoringRoleArn                  *string                      `pulumi:"monitoringRoleArn"`
	MultiAZ                            *bool                        `pulumi:"multiAZ"`
	OptionGroupName                    *string                      `pulumi:"optionGroupName"`
	PerformanceInsightsKMSKeyId        *string                      `pulumi:"performanceInsightsKMSKeyId"`
	PerformanceInsightsRetentionPeriod *int                         `pulumi:"performanceInsightsRetentionPeriod"`
	Port                               *string                      `pulumi:"port"`
	PreferredBackupWindow              *string                      `pulumi:"preferredBackupWindow"`
	PreferredMaintenanceWindow         *string                      `pulumi:"preferredMaintenanceWindow"`
	ProcessorFeatures                  []DBInstanceProcessorFeature `pulumi:"processorFeatures"`
	PromotionTier                      *int                         `pulumi:"promotionTier"`
	PubliclyAccessible                 *bool                        `pulumi:"publiclyAccessible"`
	SourceDBInstanceIdentifier         *string                      `pulumi:"sourceDBInstanceIdentifier"`
	SourceRegion                       *string                      `pulumi:"sourceRegion"`
	StorageEncrypted                   *bool                        `pulumi:"storageEncrypted"`
	StorageType                        *string                      `pulumi:"storageType"`
	Tags                               []DBInstanceTag              `pulumi:"tags"`
	Timezone                           *string                      `pulumi:"timezone"`
	UseDefaultProcessorFeatures        *bool                        `pulumi:"useDefaultProcessorFeatures"`
	VPCSecurityGroups                  []string                     `pulumi:"vPCSecurityGroups"`
}

// The set of arguments for constructing a DBInstance resource.
type DBInstanceArgs struct {
	AllocatedStorage                   pulumi.StringPtrInput
	AllowMajorVersionUpgrade           pulumi.BoolPtrInput
	AssociatedRoles                    DBInstanceRoleArrayInput
	AutoMinorVersionUpgrade            pulumi.BoolPtrInput
	AvailabilityZone                   pulumi.StringPtrInput
	BackupRetentionPeriod              pulumi.IntPtrInput
	CACertificateIdentifier            pulumi.StringPtrInput
	CharacterSetName                   pulumi.StringPtrInput
	CopyTagsToSnapshot                 pulumi.BoolPtrInput
	DBClusterIdentifier                pulumi.StringPtrInput
	DBInstanceClass                    pulumi.StringInput
	DBInstanceIdentifier               pulumi.StringPtrInput
	DBName                             pulumi.StringPtrInput
	DBParameterGroupName               pulumi.StringPtrInput
	DBSecurityGroups                   pulumi.StringArrayInput
	DBSnapshotIdentifier               pulumi.StringPtrInput
	DBSubnetGroupName                  pulumi.StringPtrInput
	DeleteAutomatedBackups             pulumi.BoolPtrInput
	DeletionProtection                 pulumi.BoolPtrInput
	Domain                             pulumi.StringPtrInput
	DomainIAMRoleName                  pulumi.StringPtrInput
	EnableCloudwatchLogsExports        pulumi.StringArrayInput
	EnableIAMDatabaseAuthentication    pulumi.BoolPtrInput
	EnablePerformanceInsights          pulumi.BoolPtrInput
	EndpointAddress                    pulumi.StringPtrInput
	EndpointPort                       pulumi.StringPtrInput
	Engine                             pulumi.StringPtrInput
	EngineVersion                      pulumi.StringPtrInput
	Iops                               pulumi.IntPtrInput
	KmsKeyId                           pulumi.StringPtrInput
	LicenseModel                       pulumi.StringPtrInput
	MasterUserPassword                 pulumi.StringPtrInput
	MasterUsername                     pulumi.StringPtrInput
	MaxAllocatedStorage                pulumi.IntPtrInput
	MonitoringInterval                 pulumi.IntPtrInput
	MonitoringRoleArn                  pulumi.StringPtrInput
	MultiAZ                            pulumi.BoolPtrInput
	OptionGroupName                    pulumi.StringPtrInput
	PerformanceInsightsKMSKeyId        pulumi.StringPtrInput
	PerformanceInsightsRetentionPeriod pulumi.IntPtrInput
	Port                               pulumi.StringPtrInput
	PreferredBackupWindow              pulumi.StringPtrInput
	PreferredMaintenanceWindow         pulumi.StringPtrInput
	ProcessorFeatures                  DBInstanceProcessorFeatureArrayInput
	PromotionTier                      pulumi.IntPtrInput
	PubliclyAccessible                 pulumi.BoolPtrInput
	SourceDBInstanceIdentifier         pulumi.StringPtrInput
	SourceRegion                       pulumi.StringPtrInput
	StorageEncrypted                   pulumi.BoolPtrInput
	StorageType                        pulumi.StringPtrInput
	Tags                               DBInstanceTagArrayInput
	Timezone                           pulumi.StringPtrInput
	UseDefaultProcessorFeatures        pulumi.BoolPtrInput
	VPCSecurityGroups                  pulumi.StringArrayInput
}

func (DBInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dbinstanceArgs)(nil)).Elem()
}

type DBInstanceInput interface {
	pulumi.Input

	ToDBInstanceOutput() DBInstanceOutput
	ToDBInstanceOutputWithContext(ctx context.Context) DBInstanceOutput
}

func (*DBInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**DBInstance)(nil)).Elem()
}

func (i *DBInstance) ToDBInstanceOutput() DBInstanceOutput {
	return i.ToDBInstanceOutputWithContext(context.Background())
}

func (i *DBInstance) ToDBInstanceOutputWithContext(ctx context.Context) DBInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBInstanceOutput)
}

type DBInstanceOutput struct{ *pulumi.OutputState }

func (DBInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DBInstance)(nil)).Elem()
}

func (o DBInstanceOutput) ToDBInstanceOutput() DBInstanceOutput {
	return o
}

func (o DBInstanceOutput) ToDBInstanceOutputWithContext(ctx context.Context) DBInstanceOutput {
	return o
}

func (o DBInstanceOutput) AllocatedStorage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringPtrOutput { return v.AllocatedStorage }).(pulumi.StringPtrOutput)
}

func (o DBInstanceOutput) AllowMajorVersionUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.BoolPtrOutput { return v.AllowMajorVersionUpgrade }).(pulumi.BoolPtrOutput)
}

func (o DBInstanceOutput) AssociatedRoles() DBInstanceRoleArrayOutput {
	return o.ApplyT(func(v *DBInstance) DBInstanceRoleArrayOutput { return v.AssociatedRoles }).(DBInstanceRoleArrayOutput)
}

func (o DBInstanceOutput) AutoMinorVersionUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.BoolPtrOutput { return v.AutoMinorVersionUpgrade }).(pulumi.BoolPtrOutput)
}

func (o DBInstanceOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringPtrOutput { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

func (o DBInstanceOutput) BackupRetentionPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.IntPtrOutput { return v.BackupRetentionPeriod }).(pulumi.IntPtrOutput)
}

func (o DBInstanceOutput) CACertificateIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringPtrOutput { return v.CACertificateIdentifier }).(pulumi.StringPtrOutput)
}

func (o DBInstanceOutput) CharacterSetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringPtrOutput { return v.CharacterSetName }).(pulumi.StringPtrOutput)
}

func (o DBInstanceOutput) CopyTagsToSnapshot() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.BoolPtrOutput { return v.CopyTagsToSnapshot }).(pulumi.BoolPtrOutput)
}

func (o DBInstanceOutput) DBClusterIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringPtrOutput { return v.DBClusterIdentifier }).(pulumi.StringPtrOutput)
}

func (o DBInstanceOutput) DBInstanceClass() pulumi.StringOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringOutput { return v.DBInstanceClass }).(pulumi.StringOutput)
}

func (o DBInstanceOutput) DBInstanceIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringPtrOutput { return v.DBInstanceIdentifier }).(pulumi.StringPtrOutput)
}

func (o DBInstanceOutput) DBName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringPtrOutput { return v.DBName }).(pulumi.StringPtrOutput)
}

func (o DBInstanceOutput) DBParameterGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringPtrOutput { return v.DBParameterGroupName }).(pulumi.StringPtrOutput)
}

func (o DBInstanceOutput) DBSecurityGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringArrayOutput { return v.DBSecurityGroups }).(pulumi.StringArrayOutput)
}

func (o DBInstanceOutput) DBSnapshotIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringPtrOutput { return v.DBSnapshotIdentifier }).(pulumi.StringPtrOutput)
}

func (o DBInstanceOutput) DBSubnetGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringPtrOutput { return v.DBSubnetGroupName }).(pulumi.StringPtrOutput)
}

func (o DBInstanceOutput) DeleteAutomatedBackups() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.BoolPtrOutput { return v.DeleteAutomatedBackups }).(pulumi.BoolPtrOutput)
}

func (o DBInstanceOutput) DeletionProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.BoolPtrOutput { return v.DeletionProtection }).(pulumi.BoolPtrOutput)
}

func (o DBInstanceOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringPtrOutput { return v.Domain }).(pulumi.StringPtrOutput)
}

func (o DBInstanceOutput) DomainIAMRoleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringPtrOutput { return v.DomainIAMRoleName }).(pulumi.StringPtrOutput)
}

func (o DBInstanceOutput) EnableCloudwatchLogsExports() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringArrayOutput { return v.EnableCloudwatchLogsExports }).(pulumi.StringArrayOutput)
}

func (o DBInstanceOutput) EnableIAMDatabaseAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.BoolPtrOutput { return v.EnableIAMDatabaseAuthentication }).(pulumi.BoolPtrOutput)
}

func (o DBInstanceOutput) EnablePerformanceInsights() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.BoolPtrOutput { return v.EnablePerformanceInsights }).(pulumi.BoolPtrOutput)
}

func (o DBInstanceOutput) EndpointAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringPtrOutput { return v.EndpointAddress }).(pulumi.StringPtrOutput)
}

func (o DBInstanceOutput) EndpointPort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringPtrOutput { return v.EndpointPort }).(pulumi.StringPtrOutput)
}

func (o DBInstanceOutput) Engine() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringPtrOutput { return v.Engine }).(pulumi.StringPtrOutput)
}

func (o DBInstanceOutput) EngineVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringPtrOutput { return v.EngineVersion }).(pulumi.StringPtrOutput)
}

func (o DBInstanceOutput) Iops() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.IntPtrOutput { return v.Iops }).(pulumi.IntPtrOutput)
}

func (o DBInstanceOutput) KmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringPtrOutput { return v.KmsKeyId }).(pulumi.StringPtrOutput)
}

func (o DBInstanceOutput) LicenseModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringPtrOutput { return v.LicenseModel }).(pulumi.StringPtrOutput)
}

func (o DBInstanceOutput) MasterUserPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringPtrOutput { return v.MasterUserPassword }).(pulumi.StringPtrOutput)
}

func (o DBInstanceOutput) MasterUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringPtrOutput { return v.MasterUsername }).(pulumi.StringPtrOutput)
}

func (o DBInstanceOutput) MaxAllocatedStorage() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.IntPtrOutput { return v.MaxAllocatedStorage }).(pulumi.IntPtrOutput)
}

func (o DBInstanceOutput) MonitoringInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.IntPtrOutput { return v.MonitoringInterval }).(pulumi.IntPtrOutput)
}

func (o DBInstanceOutput) MonitoringRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringPtrOutput { return v.MonitoringRoleArn }).(pulumi.StringPtrOutput)
}

func (o DBInstanceOutput) MultiAZ() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.BoolPtrOutput { return v.MultiAZ }).(pulumi.BoolPtrOutput)
}

func (o DBInstanceOutput) OptionGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringPtrOutput { return v.OptionGroupName }).(pulumi.StringPtrOutput)
}

func (o DBInstanceOutput) PerformanceInsightsKMSKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringPtrOutput { return v.PerformanceInsightsKMSKeyId }).(pulumi.StringPtrOutput)
}

func (o DBInstanceOutput) PerformanceInsightsRetentionPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.IntPtrOutput { return v.PerformanceInsightsRetentionPeriod }).(pulumi.IntPtrOutput)
}

func (o DBInstanceOutput) Port() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringPtrOutput { return v.Port }).(pulumi.StringPtrOutput)
}

func (o DBInstanceOutput) PreferredBackupWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringPtrOutput { return v.PreferredBackupWindow }).(pulumi.StringPtrOutput)
}

func (o DBInstanceOutput) PreferredMaintenanceWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringPtrOutput { return v.PreferredMaintenanceWindow }).(pulumi.StringPtrOutput)
}

func (o DBInstanceOutput) ProcessorFeatures() DBInstanceProcessorFeatureArrayOutput {
	return o.ApplyT(func(v *DBInstance) DBInstanceProcessorFeatureArrayOutput { return v.ProcessorFeatures }).(DBInstanceProcessorFeatureArrayOutput)
}

func (o DBInstanceOutput) PromotionTier() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.IntPtrOutput { return v.PromotionTier }).(pulumi.IntPtrOutput)
}

func (o DBInstanceOutput) PubliclyAccessible() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.BoolPtrOutput { return v.PubliclyAccessible }).(pulumi.BoolPtrOutput)
}

func (o DBInstanceOutput) SourceDBInstanceIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringPtrOutput { return v.SourceDBInstanceIdentifier }).(pulumi.StringPtrOutput)
}

func (o DBInstanceOutput) SourceRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringPtrOutput { return v.SourceRegion }).(pulumi.StringPtrOutput)
}

func (o DBInstanceOutput) StorageEncrypted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.BoolPtrOutput { return v.StorageEncrypted }).(pulumi.BoolPtrOutput)
}

func (o DBInstanceOutput) StorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringPtrOutput { return v.StorageType }).(pulumi.StringPtrOutput)
}

func (o DBInstanceOutput) Tags() DBInstanceTagArrayOutput {
	return o.ApplyT(func(v *DBInstance) DBInstanceTagArrayOutput { return v.Tags }).(DBInstanceTagArrayOutput)
}

func (o DBInstanceOutput) Timezone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringPtrOutput { return v.Timezone }).(pulumi.StringPtrOutput)
}

func (o DBInstanceOutput) UseDefaultProcessorFeatures() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.BoolPtrOutput { return v.UseDefaultProcessorFeatures }).(pulumi.BoolPtrOutput)
}

func (o DBInstanceOutput) VPCSecurityGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringArrayOutput { return v.VPCSecurityGroups }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DBInstanceInput)(nil)).Elem(), &DBInstance{})
	pulumi.RegisterOutputType(DBInstanceOutput{})
}
