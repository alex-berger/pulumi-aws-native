// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Backup
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupselection.html
    /// </summary>
    [AwsNativeResourceType("aws-native:Backup:BackupSelection")]
    public partial class BackupSelection : Pulumi.CustomResource
    {
        [Output("backupPlanId")]
        public Output<string> BackupPlanId { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupselection.html#cfn-backup-backupselection-backupselection
        /// </summary>
        [Output("backupSelection")]
        public Output<Outputs.BackupSelectionBackupSelectionResourceType> BackupSelectionValue { get; private set; } = null!;

        [Output("selectionId")]
        public Output<string> SelectionId { get; private set; } = null!;


        /// <summary>
        /// Create a BackupSelection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BackupSelection(string name, BackupSelectionArgs args, CustomResourceOptions? options = null)
            : base("aws-native:Backup:BackupSelection", name, args ?? new BackupSelectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BackupSelection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:Backup:BackupSelection", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing BackupSelection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BackupSelection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new BackupSelection(name, id, options);
        }
    }

    public sealed class BackupSelectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupselection.html#cfn-backup-backupselection-backupplanid
        /// </summary>
        [Input("backupPlanId", required: true)]
        public Input<string> BackupPlanId { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupselection.html#cfn-backup-backupselection-backupselection
        /// </summary>
        [Input("backupSelection", required: true)]
        public Input<Inputs.BackupSelectionBackupSelectionResourceTypeArgs> BackupSelectionValue { get; set; } = null!;

        public BackupSelectionArgs()
        {
        }
    }
}
