// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::RDS::DBCluster
 *
 * @deprecated DBCluster is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class DBCluster extends pulumi.CustomResource {
    /**
     * Get an existing DBCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DBCluster {
        pulumi.log.warn("DBCluster is deprecated: DBCluster is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new DBCluster(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:rds:DBCluster';

    /**
     * Returns true if the given object is an instance of DBCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DBCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DBCluster.__pulumiType;
    }

    public readonly associatedRoles!: pulumi.Output<outputs.rds.DBClusterRole[] | undefined>;
    public readonly availabilityZones!: pulumi.Output<string[] | undefined>;
    public readonly backtrackWindow!: pulumi.Output<number | undefined>;
    public readonly backupRetentionPeriod!: pulumi.Output<number | undefined>;
    public readonly copyTagsToSnapshot!: pulumi.Output<boolean | undefined>;
    public readonly dBClusterIdentifier!: pulumi.Output<string | undefined>;
    public readonly dBClusterParameterGroupName!: pulumi.Output<string | undefined>;
    public readonly dBSubnetGroupName!: pulumi.Output<string | undefined>;
    public readonly databaseName!: pulumi.Output<string | undefined>;
    public readonly deletionProtection!: pulumi.Output<boolean | undefined>;
    public readonly enableCloudwatchLogsExports!: pulumi.Output<string[] | undefined>;
    public readonly enableHttpEndpoint!: pulumi.Output<boolean | undefined>;
    public readonly enableIAMDatabaseAuthentication!: pulumi.Output<boolean | undefined>;
    public readonly endpointAddress!: pulumi.Output<string | undefined>;
    public readonly endpointPort!: pulumi.Output<string | undefined>;
    public readonly engine!: pulumi.Output<string>;
    public readonly engineMode!: pulumi.Output<string | undefined>;
    public readonly engineVersion!: pulumi.Output<string | undefined>;
    public readonly globalClusterIdentifier!: pulumi.Output<string | undefined>;
    public readonly kmsKeyId!: pulumi.Output<string | undefined>;
    public readonly masterUserPassword!: pulumi.Output<string | undefined>;
    public readonly masterUsername!: pulumi.Output<string | undefined>;
    public readonly port!: pulumi.Output<number | undefined>;
    public readonly preferredBackupWindow!: pulumi.Output<string | undefined>;
    public readonly preferredMaintenanceWindow!: pulumi.Output<string | undefined>;
    public readonly readEndpointAddress!: pulumi.Output<string | undefined>;
    public readonly replicationSourceIdentifier!: pulumi.Output<string | undefined>;
    public readonly restoreType!: pulumi.Output<string | undefined>;
    public readonly scalingConfiguration!: pulumi.Output<outputs.rds.DBClusterScalingConfiguration | undefined>;
    public readonly snapshotIdentifier!: pulumi.Output<string | undefined>;
    public readonly sourceDBClusterIdentifier!: pulumi.Output<string | undefined>;
    public readonly sourceRegion!: pulumi.Output<string | undefined>;
    public readonly storageEncrypted!: pulumi.Output<boolean | undefined>;
    public readonly tags!: pulumi.Output<outputs.rds.DBClusterTag[] | undefined>;
    public readonly useLatestRestorableTime!: pulumi.Output<boolean | undefined>;
    public readonly vpcSecurityGroupIds!: pulumi.Output<string[] | undefined>;

    /**
     * Create a DBCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated DBCluster is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: DBClusterArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("DBCluster is deprecated: DBCluster is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.engine === undefined) && !opts.urn) {
                throw new Error("Missing required property 'engine'");
            }
            resourceInputs["associatedRoles"] = args ? args.associatedRoles : undefined;
            resourceInputs["availabilityZones"] = args ? args.availabilityZones : undefined;
            resourceInputs["backtrackWindow"] = args ? args.backtrackWindow : undefined;
            resourceInputs["backupRetentionPeriod"] = args ? args.backupRetentionPeriod : undefined;
            resourceInputs["copyTagsToSnapshot"] = args ? args.copyTagsToSnapshot : undefined;
            resourceInputs["dBClusterIdentifier"] = args ? args.dBClusterIdentifier : undefined;
            resourceInputs["dBClusterParameterGroupName"] = args ? args.dBClusterParameterGroupName : undefined;
            resourceInputs["dBSubnetGroupName"] = args ? args.dBSubnetGroupName : undefined;
            resourceInputs["databaseName"] = args ? args.databaseName : undefined;
            resourceInputs["deletionProtection"] = args ? args.deletionProtection : undefined;
            resourceInputs["enableCloudwatchLogsExports"] = args ? args.enableCloudwatchLogsExports : undefined;
            resourceInputs["enableHttpEndpoint"] = args ? args.enableHttpEndpoint : undefined;
            resourceInputs["enableIAMDatabaseAuthentication"] = args ? args.enableIAMDatabaseAuthentication : undefined;
            resourceInputs["endpointAddress"] = args ? args.endpointAddress : undefined;
            resourceInputs["endpointPort"] = args ? args.endpointPort : undefined;
            resourceInputs["engine"] = args ? args.engine : undefined;
            resourceInputs["engineMode"] = args ? args.engineMode : undefined;
            resourceInputs["engineVersion"] = args ? args.engineVersion : undefined;
            resourceInputs["globalClusterIdentifier"] = args ? args.globalClusterIdentifier : undefined;
            resourceInputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            resourceInputs["masterUserPassword"] = args ? args.masterUserPassword : undefined;
            resourceInputs["masterUsername"] = args ? args.masterUsername : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["preferredBackupWindow"] = args ? args.preferredBackupWindow : undefined;
            resourceInputs["preferredMaintenanceWindow"] = args ? args.preferredMaintenanceWindow : undefined;
            resourceInputs["readEndpointAddress"] = args ? args.readEndpointAddress : undefined;
            resourceInputs["replicationSourceIdentifier"] = args ? args.replicationSourceIdentifier : undefined;
            resourceInputs["restoreType"] = args ? args.restoreType : undefined;
            resourceInputs["scalingConfiguration"] = args ? args.scalingConfiguration : undefined;
            resourceInputs["snapshotIdentifier"] = args ? args.snapshotIdentifier : undefined;
            resourceInputs["sourceDBClusterIdentifier"] = args ? args.sourceDBClusterIdentifier : undefined;
            resourceInputs["sourceRegion"] = args ? args.sourceRegion : undefined;
            resourceInputs["storageEncrypted"] = args ? args.storageEncrypted : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["useLatestRestorableTime"] = args ? args.useLatestRestorableTime : undefined;
            resourceInputs["vpcSecurityGroupIds"] = args ? args.vpcSecurityGroupIds : undefined;
        } else {
            resourceInputs["associatedRoles"] = undefined /*out*/;
            resourceInputs["availabilityZones"] = undefined /*out*/;
            resourceInputs["backtrackWindow"] = undefined /*out*/;
            resourceInputs["backupRetentionPeriod"] = undefined /*out*/;
            resourceInputs["copyTagsToSnapshot"] = undefined /*out*/;
            resourceInputs["dBClusterIdentifier"] = undefined /*out*/;
            resourceInputs["dBClusterParameterGroupName"] = undefined /*out*/;
            resourceInputs["dBSubnetGroupName"] = undefined /*out*/;
            resourceInputs["databaseName"] = undefined /*out*/;
            resourceInputs["deletionProtection"] = undefined /*out*/;
            resourceInputs["enableCloudwatchLogsExports"] = undefined /*out*/;
            resourceInputs["enableHttpEndpoint"] = undefined /*out*/;
            resourceInputs["enableIAMDatabaseAuthentication"] = undefined /*out*/;
            resourceInputs["endpointAddress"] = undefined /*out*/;
            resourceInputs["endpointPort"] = undefined /*out*/;
            resourceInputs["engine"] = undefined /*out*/;
            resourceInputs["engineMode"] = undefined /*out*/;
            resourceInputs["engineVersion"] = undefined /*out*/;
            resourceInputs["globalClusterIdentifier"] = undefined /*out*/;
            resourceInputs["kmsKeyId"] = undefined /*out*/;
            resourceInputs["masterUserPassword"] = undefined /*out*/;
            resourceInputs["masterUsername"] = undefined /*out*/;
            resourceInputs["port"] = undefined /*out*/;
            resourceInputs["preferredBackupWindow"] = undefined /*out*/;
            resourceInputs["preferredMaintenanceWindow"] = undefined /*out*/;
            resourceInputs["readEndpointAddress"] = undefined /*out*/;
            resourceInputs["replicationSourceIdentifier"] = undefined /*out*/;
            resourceInputs["restoreType"] = undefined /*out*/;
            resourceInputs["scalingConfiguration"] = undefined /*out*/;
            resourceInputs["snapshotIdentifier"] = undefined /*out*/;
            resourceInputs["sourceDBClusterIdentifier"] = undefined /*out*/;
            resourceInputs["sourceRegion"] = undefined /*out*/;
            resourceInputs["storageEncrypted"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["useLatestRestorableTime"] = undefined /*out*/;
            resourceInputs["vpcSecurityGroupIds"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DBCluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DBCluster resource.
 */
export interface DBClusterArgs {
    associatedRoles?: pulumi.Input<pulumi.Input<inputs.rds.DBClusterRoleArgs>[]>;
    availabilityZones?: pulumi.Input<pulumi.Input<string>[]>;
    backtrackWindow?: pulumi.Input<number>;
    backupRetentionPeriod?: pulumi.Input<number>;
    copyTagsToSnapshot?: pulumi.Input<boolean>;
    dBClusterIdentifier?: pulumi.Input<string>;
    dBClusterParameterGroupName?: pulumi.Input<string>;
    dBSubnetGroupName?: pulumi.Input<string>;
    databaseName?: pulumi.Input<string>;
    deletionProtection?: pulumi.Input<boolean>;
    enableCloudwatchLogsExports?: pulumi.Input<pulumi.Input<string>[]>;
    enableHttpEndpoint?: pulumi.Input<boolean>;
    enableIAMDatabaseAuthentication?: pulumi.Input<boolean>;
    endpointAddress?: pulumi.Input<string>;
    endpointPort?: pulumi.Input<string>;
    engine: pulumi.Input<string>;
    engineMode?: pulumi.Input<string>;
    engineVersion?: pulumi.Input<string>;
    globalClusterIdentifier?: pulumi.Input<string>;
    kmsKeyId?: pulumi.Input<string>;
    masterUserPassword?: pulumi.Input<string>;
    masterUsername?: pulumi.Input<string>;
    port?: pulumi.Input<number>;
    preferredBackupWindow?: pulumi.Input<string>;
    preferredMaintenanceWindow?: pulumi.Input<string>;
    readEndpointAddress?: pulumi.Input<string>;
    replicationSourceIdentifier?: pulumi.Input<string>;
    restoreType?: pulumi.Input<string>;
    scalingConfiguration?: pulumi.Input<inputs.rds.DBClusterScalingConfigurationArgs>;
    snapshotIdentifier?: pulumi.Input<string>;
    sourceDBClusterIdentifier?: pulumi.Input<string>;
    sourceRegion?: pulumi.Input<string>;
    storageEncrypted?: pulumi.Input<boolean>;
    tags?: pulumi.Input<pulumi.Input<inputs.rds.DBClusterTagArgs>[]>;
    useLatestRestorableTime?: pulumi.Input<boolean>;
    vpcSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
}
