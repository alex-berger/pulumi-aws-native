// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html
 */
export class BackupVault extends pulumi.CustomResource {
    /**
     * Get an existing BackupVault resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): BackupVault {
        return new BackupVault(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:Backup:BackupVault';

    /**
     * Returns true if the given object is an instance of BackupVault.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BackupVault {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BackupVault.__pulumiType;
    }

    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html#cfn-backup-backupvault-accesspolicy
     */
    public readonly accessPolicy!: pulumi.Output<any | string | undefined>;
    public /*out*/ readonly backupVaultArn!: pulumi.Output<string>;
    public readonly backupVaultName!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html#cfn-backup-backupvault-backupvaulttags
     */
    public readonly backupVaultTags!: pulumi.Output<any | string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html#cfn-backup-backupvault-encryptionkeyarn
     */
    public readonly encryptionKeyArn!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html#cfn-backup-backupvault-notifications
     */
    public readonly notifications!: pulumi.Output<outputs.Backup.BackupVaultNotificationObjectType | undefined>;

    /**
     * Create a BackupVault resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BackupVaultArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.backupVaultName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backupVaultName'");
            }
            inputs["accessPolicy"] = args ? args.accessPolicy : undefined;
            inputs["backupVaultName"] = args ? args.backupVaultName : undefined;
            inputs["backupVaultTags"] = args ? args.backupVaultTags : undefined;
            inputs["encryptionKeyArn"] = args ? args.encryptionKeyArn : undefined;
            inputs["notifications"] = args ? args.notifications : undefined;
            inputs["backupVaultArn"] = undefined /*out*/;
        } else {
            inputs["accessPolicy"] = undefined /*out*/;
            inputs["backupVaultArn"] = undefined /*out*/;
            inputs["backupVaultName"] = undefined /*out*/;
            inputs["backupVaultTags"] = undefined /*out*/;
            inputs["encryptionKeyArn"] = undefined /*out*/;
            inputs["notifications"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(BackupVault.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a BackupVault resource.
 */
export interface BackupVaultArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html#cfn-backup-backupvault-accesspolicy
     */
    readonly accessPolicy?: pulumi.Input<any | string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html#cfn-backup-backupvault-backupvaultname
     */
    readonly backupVaultName: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html#cfn-backup-backupvault-backupvaulttags
     */
    readonly backupVaultTags?: pulumi.Input<any | string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html#cfn-backup-backupvault-encryptionkeyarn
     */
    readonly encryptionKeyArn?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html#cfn-backup-backupvault-notifications
     */
    readonly notifications?: pulumi.Input<inputs.Backup.BackupVaultNotificationObjectTypeArgs>;
}
