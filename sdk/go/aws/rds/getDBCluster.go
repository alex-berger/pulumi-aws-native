// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::RDS::DBCluster
func LookupDBCluster(ctx *pulumi.Context, args *LookupDBClusterArgs, opts ...pulumi.InvokeOption) (*LookupDBClusterResult, error) {
	var rv LookupDBClusterResult
	err := ctx.Invoke("aws-native:rds:getDBCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDBClusterArgs struct {
	Id string `pulumi:"id"`
}

type LookupDBClusterResult struct {
	AssociatedRoles                 []DBClusterRole                `pulumi:"associatedRoles"`
	BacktrackWindow                 *int                           `pulumi:"backtrackWindow"`
	BackupRetentionPeriod           *int                           `pulumi:"backupRetentionPeriod"`
	CopyTagsToSnapshot              *bool                          `pulumi:"copyTagsToSnapshot"`
	DBClusterParameterGroupName     *string                        `pulumi:"dBClusterParameterGroupName"`
	DeletionProtection              *bool                          `pulumi:"deletionProtection"`
	EnableCloudwatchLogsExports     []string                       `pulumi:"enableCloudwatchLogsExports"`
	EnableHttpEndpoint              *bool                          `pulumi:"enableHttpEndpoint"`
	EnableIAMDatabaseAuthentication *bool                          `pulumi:"enableIAMDatabaseAuthentication"`
	EndpointAddress                 *string                        `pulumi:"endpointAddress"`
	EndpointPort                    *string                        `pulumi:"endpointPort"`
	Engine                          *string                        `pulumi:"engine"`
	EngineVersion                   *string                        `pulumi:"engineVersion"`
	GlobalClusterIdentifier         *string                        `pulumi:"globalClusterIdentifier"`
	Id                              *string                        `pulumi:"id"`
	MasterUserPassword              *string                        `pulumi:"masterUserPassword"`
	Port                            *int                           `pulumi:"port"`
	PreferredBackupWindow           *string                        `pulumi:"preferredBackupWindow"`
	PreferredMaintenanceWindow      *string                        `pulumi:"preferredMaintenanceWindow"`
	ReadEndpointAddress             *string                        `pulumi:"readEndpointAddress"`
	ReplicationSourceIdentifier     *string                        `pulumi:"replicationSourceIdentifier"`
	ScalingConfiguration            *DBClusterScalingConfiguration `pulumi:"scalingConfiguration"`
	Tags                            []DBClusterTag                 `pulumi:"tags"`
	VpcSecurityGroupIds             []string                       `pulumi:"vpcSecurityGroupIds"`
}

func LookupDBClusterOutput(ctx *pulumi.Context, args LookupDBClusterOutputArgs, opts ...pulumi.InvokeOption) LookupDBClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDBClusterResult, error) {
			args := v.(LookupDBClusterArgs)
			r, err := LookupDBCluster(ctx, &args, opts...)
			var s LookupDBClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDBClusterResultOutput)
}

type LookupDBClusterOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupDBClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDBClusterArgs)(nil)).Elem()
}

type LookupDBClusterResultOutput struct{ *pulumi.OutputState }

func (LookupDBClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDBClusterResult)(nil)).Elem()
}

func (o LookupDBClusterResultOutput) ToLookupDBClusterResultOutput() LookupDBClusterResultOutput {
	return o
}

func (o LookupDBClusterResultOutput) ToLookupDBClusterResultOutputWithContext(ctx context.Context) LookupDBClusterResultOutput {
	return o
}

func (o LookupDBClusterResultOutput) AssociatedRoles() DBClusterRoleArrayOutput {
	return o.ApplyT(func(v LookupDBClusterResult) []DBClusterRole { return v.AssociatedRoles }).(DBClusterRoleArrayOutput)
}

func (o LookupDBClusterResultOutput) BacktrackWindow() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDBClusterResult) *int { return v.BacktrackWindow }).(pulumi.IntPtrOutput)
}

func (o LookupDBClusterResultOutput) BackupRetentionPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDBClusterResult) *int { return v.BackupRetentionPeriod }).(pulumi.IntPtrOutput)
}

func (o LookupDBClusterResultOutput) CopyTagsToSnapshot() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDBClusterResult) *bool { return v.CopyTagsToSnapshot }).(pulumi.BoolPtrOutput)
}

func (o LookupDBClusterResultOutput) DBClusterParameterGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDBClusterResult) *string { return v.DBClusterParameterGroupName }).(pulumi.StringPtrOutput)
}

func (o LookupDBClusterResultOutput) DeletionProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDBClusterResult) *bool { return v.DeletionProtection }).(pulumi.BoolPtrOutput)
}

func (o LookupDBClusterResultOutput) EnableCloudwatchLogsExports() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDBClusterResult) []string { return v.EnableCloudwatchLogsExports }).(pulumi.StringArrayOutput)
}

func (o LookupDBClusterResultOutput) EnableHttpEndpoint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDBClusterResult) *bool { return v.EnableHttpEndpoint }).(pulumi.BoolPtrOutput)
}

func (o LookupDBClusterResultOutput) EnableIAMDatabaseAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDBClusterResult) *bool { return v.EnableIAMDatabaseAuthentication }).(pulumi.BoolPtrOutput)
}

func (o LookupDBClusterResultOutput) EndpointAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDBClusterResult) *string { return v.EndpointAddress }).(pulumi.StringPtrOutput)
}

func (o LookupDBClusterResultOutput) EndpointPort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDBClusterResult) *string { return v.EndpointPort }).(pulumi.StringPtrOutput)
}

func (o LookupDBClusterResultOutput) Engine() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDBClusterResult) *string { return v.Engine }).(pulumi.StringPtrOutput)
}

func (o LookupDBClusterResultOutput) EngineVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDBClusterResult) *string { return v.EngineVersion }).(pulumi.StringPtrOutput)
}

func (o LookupDBClusterResultOutput) GlobalClusterIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDBClusterResult) *string { return v.GlobalClusterIdentifier }).(pulumi.StringPtrOutput)
}

func (o LookupDBClusterResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDBClusterResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupDBClusterResultOutput) MasterUserPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDBClusterResult) *string { return v.MasterUserPassword }).(pulumi.StringPtrOutput)
}

func (o LookupDBClusterResultOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDBClusterResult) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o LookupDBClusterResultOutput) PreferredBackupWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDBClusterResult) *string { return v.PreferredBackupWindow }).(pulumi.StringPtrOutput)
}

func (o LookupDBClusterResultOutput) PreferredMaintenanceWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDBClusterResult) *string { return v.PreferredMaintenanceWindow }).(pulumi.StringPtrOutput)
}

func (o LookupDBClusterResultOutput) ReadEndpointAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDBClusterResult) *string { return v.ReadEndpointAddress }).(pulumi.StringPtrOutput)
}

func (o LookupDBClusterResultOutput) ReplicationSourceIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDBClusterResult) *string { return v.ReplicationSourceIdentifier }).(pulumi.StringPtrOutput)
}

func (o LookupDBClusterResultOutput) ScalingConfiguration() DBClusterScalingConfigurationPtrOutput {
	return o.ApplyT(func(v LookupDBClusterResult) *DBClusterScalingConfiguration { return v.ScalingConfiguration }).(DBClusterScalingConfigurationPtrOutput)
}

func (o LookupDBClusterResultOutput) Tags() DBClusterTagArrayOutput {
	return o.ApplyT(func(v LookupDBClusterResult) []DBClusterTag { return v.Tags }).(DBClusterTagArrayOutput)
}

func (o LookupDBClusterResultOutput) VpcSecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDBClusterResult) []string { return v.VpcSecurityGroupIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDBClusterResultOutput{})
}
