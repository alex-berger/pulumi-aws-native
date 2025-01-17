// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Backup
{
    public static class GetBackupVault
    {
        /// <summary>
        /// Resource Type definition for AWS::Backup::BackupVault
        /// </summary>
        public static Task<GetBackupVaultResult> InvokeAsync(GetBackupVaultArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBackupVaultResult>("aws-native:backup:getBackupVault", args ?? new GetBackupVaultArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Backup::BackupVault
        /// </summary>
        public static Output<GetBackupVaultResult> Invoke(GetBackupVaultInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetBackupVaultResult>("aws-native:backup:getBackupVault", args ?? new GetBackupVaultInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBackupVaultArgs : Pulumi.InvokeArgs
    {
        [Input("backupVaultName", required: true)]
        public string BackupVaultName { get; set; } = null!;

        public GetBackupVaultArgs()
        {
        }
    }

    public sealed class GetBackupVaultInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("backupVaultName", required: true)]
        public Input<string> BackupVaultName { get; set; } = null!;

        public GetBackupVaultInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBackupVaultResult
    {
        public readonly object? AccessPolicy;
        public readonly string? BackupVaultArn;
        public readonly object? BackupVaultTags;
        public readonly Outputs.BackupVaultLockConfigurationType? LockConfiguration;
        public readonly Outputs.BackupVaultNotificationObjectType? Notifications;

        [OutputConstructor]
        private GetBackupVaultResult(
            object? accessPolicy,

            string? backupVaultArn,

            object? backupVaultTags,

            Outputs.BackupVaultLockConfigurationType? lockConfiguration,

            Outputs.BackupVaultNotificationObjectType? notifications)
        {
            AccessPolicy = accessPolicy;
            BackupVaultArn = backupVaultArn;
            BackupVaultTags = backupVaultTags;
            LockConfiguration = lockConfiguration;
            Notifications = notifications;
        }
    }
}
