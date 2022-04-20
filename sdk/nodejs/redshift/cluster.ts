// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * An example resource schema demonstrating some basic constructs and validation rules.
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:redshift:Cluster';

    /**
     * Returns true if the given object is an instance of Cluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cluster.__pulumiType;
    }

    /**
     * Major version upgrades can be applied during the maintenance window to the Amazon Redshift engine that is running on the cluster. Default value is True
     */
    public readonly allowVersionUpgrade!: pulumi.Output<boolean | undefined>;
    /**
     * The value represents how the cluster is configured to use AQUA (Advanced Query Accelerator) after the cluster is restored. Possible values include the following.
     *
     * enabled - Use AQUA if it is available for the current Region and Amazon Redshift node type.
     * disabled - Don't use AQUA.
     * auto - Amazon Redshift determines whether to use AQUA.
     */
    public readonly aquaConfigurationStatus!: pulumi.Output<string | undefined>;
    /**
     * The number of days that automated snapshots are retained. If the value is 0, automated snapshots are disabled. Default value is 1
     */
    public readonly automatedSnapshotRetentionPeriod!: pulumi.Output<number | undefined>;
    /**
     * The EC2 Availability Zone (AZ) in which you want Amazon Redshift to provision the cluster. Default: A random, system-chosen Availability Zone in the region that is specified by the endpoint
     */
    public readonly availabilityZone!: pulumi.Output<string | undefined>;
    /**
     * The option to enable relocation for an Amazon Redshift cluster between Availability Zones after the cluster modification is complete.
     */
    public readonly availabilityZoneRelocation!: pulumi.Output<boolean | undefined>;
    /**
     * The availability zone relocation status of the cluster
     */
    public readonly availabilityZoneRelocationStatus!: pulumi.Output<string | undefined>;
    /**
     * A boolean value indicating whether the resize operation is using the classic resize process. If you don't provide this parameter or set the value to false , the resize type is elastic.
     */
    public readonly classic!: pulumi.Output<boolean | undefined>;
    /**
     * A unique identifier for the cluster. You use this identifier to refer to the cluster for any subsequent cluster operations such as deleting or modifying. All alphabetical characters must be lower case, no hypens at the end, no two consecutive hyphens. Cluster name should be unique for all clusters within an AWS account
     */
    public readonly clusterIdentifier!: pulumi.Output<string | undefined>;
    /**
     * The name of the parameter group to be associated with this cluster.
     */
    public readonly clusterParameterGroupName!: pulumi.Output<string | undefined>;
    /**
     * A list of security groups to be associated with this cluster.
     */
    public readonly clusterSecurityGroups!: pulumi.Output<string[] | undefined>;
    /**
     * The name of a cluster subnet group to be associated with this cluster.
     */
    public readonly clusterSubnetGroupName!: pulumi.Output<string | undefined>;
    /**
     * The type of the cluster. When cluster type is specified as single-node, the NumberOfNodes parameter is not required and if multi-node, the NumberOfNodes parameter is required
     */
    public readonly clusterType!: pulumi.Output<string>;
    /**
     * The version of the Amazon Redshift engine software that you want to deploy on the cluster.The version selected runs on all the nodes in the cluster.
     */
    public readonly clusterVersion!: pulumi.Output<string | undefined>;
    /**
     * The name of the first database to be created when the cluster is created. To create additional databases after the cluster is created, connect to the cluster with a SQL client and use SQL commands to create a database.
     */
    public readonly dBName!: pulumi.Output<string>;
    /**
     * A boolean indicating whether to enable the deferred maintenance window.
     */
    public readonly deferMaintenance!: pulumi.Output<boolean | undefined>;
    /**
     * An integer indicating the duration of the maintenance window in days. If you specify a duration, you can't specify an end time. The duration must be 45 days or less.
     */
    public readonly deferMaintenanceDuration!: pulumi.Output<number | undefined>;
    /**
     * A timestamp indicating end time for the deferred maintenance window. If you specify an end time, you can't specify a duration.
     */
    public readonly deferMaintenanceEndTime!: pulumi.Output<string | undefined>;
    /**
     * A unique identifier for the deferred maintenance window.
     */
    public /*out*/ readonly deferMaintenanceIdentifier!: pulumi.Output<string>;
    /**
     * A timestamp indicating the start time for the deferred maintenance window.
     */
    public readonly deferMaintenanceStartTime!: pulumi.Output<string | undefined>;
    /**
     * The destination AWS Region that you want to copy snapshots to. Constraints: Must be the name of a valid AWS Region. For more information, see Regions and Endpoints in the Amazon Web Services [https://docs.aws.amazon.com/general/latest/gr/rande.html#redshift_region] General Reference
     */
    public readonly destinationRegion!: pulumi.Output<string | undefined>;
    /**
     * The Elastic IP (EIP) address for the cluster.
     */
    public readonly elasticIp!: pulumi.Output<string | undefined>;
    /**
     * If true, the data in the cluster is encrypted at rest.
     */
    public readonly encrypted!: pulumi.Output<boolean | undefined>;
    public readonly endpoint!: pulumi.Output<outputs.redshift.ClusterEndpoint | undefined>;
    /**
     * An option that specifies whether to create the cluster with enhanced VPC routing enabled. To create a cluster that uses enhanced VPC routing, the cluster must be in a VPC. For more information, see Enhanced VPC Routing in the Amazon Redshift Cluster Management Guide.
     *
     * If this option is true , enhanced VPC routing is enabled.
     *
     * Default: false
     */
    public readonly enhancedVpcRouting!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the name of the HSM client certificate the Amazon Redshift cluster uses to retrieve the data encryption keys stored in an HSM
     */
    public readonly hsmClientCertificateIdentifier!: pulumi.Output<string | undefined>;
    /**
     * Specifies the name of the HSM configuration that contains the information the Amazon Redshift cluster can use to retrieve and store keys in an HSM.
     */
    public readonly hsmConfigurationIdentifier!: pulumi.Output<string | undefined>;
    /**
     * A list of AWS Identity and Access Management (IAM) roles that can be used by the cluster to access other AWS services. You must supply the IAM roles in their Amazon Resource Name (ARN) format. You can supply up to 10 IAM roles in a single request
     */
    public readonly iamRoles!: pulumi.Output<string[] | undefined>;
    /**
     * The AWS Key Management Service (KMS) key ID of the encryption key that you want to use to encrypt data in the cluster.
     */
    public readonly kmsKeyId!: pulumi.Output<string | undefined>;
    public readonly loggingProperties!: pulumi.Output<outputs.redshift.ClusterLoggingProperties | undefined>;
    /**
     * The name for the maintenance track that you want to assign for the cluster. This name change is asynchronous. The new track name stays in the PendingModifiedValues for the cluster until the next maintenance window. When the maintenance track changes, the cluster is switched to the latest cluster release available for the maintenance track. At this point, the maintenance track name is applied.
     */
    public readonly maintenanceTrackName!: pulumi.Output<string | undefined>;
    /**
     * The number of days to retain newly copied snapshots in the destination AWS Region after they are copied from the source AWS Region. If the value is -1, the manual snapshot is retained indefinitely.
     *
     * The value must be either -1 or an integer between 1 and 3,653.
     */
    public readonly manualSnapshotRetentionPeriod!: pulumi.Output<number | undefined>;
    /**
     * The password associated with the master user account for the cluster that is being created. Password must be between 8 and 64 characters in length, should have at least one uppercase letter.Must contain at least one lowercase letter.Must contain one number.Can be any printable ASCII character.
     */
    public readonly masterUserPassword!: pulumi.Output<string>;
    /**
     * The user name associated with the master user account for the cluster that is being created. The user name can't be PUBLIC and first character must be a letter.
     */
    public readonly masterUsername!: pulumi.Output<string>;
    /**
     * The node type to be provisioned for the cluster.Valid Values: ds2.xlarge | ds2.8xlarge | dc1.large | dc1.8xlarge | dc2.large | dc2.8xlarge | ra3.4xlarge | ra3.16xlarge
     */
    public readonly nodeType!: pulumi.Output<string>;
    /**
     * The number of compute nodes in the cluster. This parameter is required when the ClusterType parameter is specified as multi-node.
     */
    public readonly numberOfNodes!: pulumi.Output<number | undefined>;
    public readonly ownerAccount!: pulumi.Output<string | undefined>;
    /**
     * The port number on which the cluster accepts incoming connections. The cluster is accessible only via the JDBC and ODBC connection strings
     */
    public readonly port!: pulumi.Output<number | undefined>;
    /**
     * The weekly time range (in UTC) during which automated cluster maintenance can occur.
     */
    public readonly preferredMaintenanceWindow!: pulumi.Output<string | undefined>;
    /**
     * If true, the cluster can be accessed from a public network.
     */
    public readonly publiclyAccessible!: pulumi.Output<boolean | undefined>;
    /**
     * The Redshift operation to be performed. Resource Action supports pause-cluster, resume-cluster APIs
     */
    public readonly resourceAction!: pulumi.Output<string | undefined>;
    /**
     * The identifier of the database revision. You can retrieve this value from the response to the DescribeClusterDbRevisions request.
     */
    public readonly revisionTarget!: pulumi.Output<string | undefined>;
    /**
     * A boolean indicating if we want to rotate Encryption Keys.
     */
    public readonly rotateEncryptionKey!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the cluster the source snapshot was created from. This parameter is required if your IAM user has a policy containing a snapshot resource element that specifies anything other than * for the cluster name.
     */
    public readonly snapshotClusterIdentifier!: pulumi.Output<string | undefined>;
    /**
     * The name of the snapshot copy grant to use when snapshots of an AWS KMS-encrypted cluster are copied to the destination region.
     */
    public readonly snapshotCopyGrantName!: pulumi.Output<string | undefined>;
    /**
     * Indicates whether to apply the snapshot retention period to newly copied manual snapshots instead of automated snapshots.
     */
    public readonly snapshotCopyManual!: pulumi.Output<boolean | undefined>;
    /**
     * The number of days to retain automated snapshots in the destination region after they are copied from the source region. 
     *
     *  Default is 7. 
     *
     *  Constraints: Must be at least 1 and no more than 35.
     */
    public readonly snapshotCopyRetentionPeriod!: pulumi.Output<number | undefined>;
    /**
     * The name of the snapshot from which to create the new cluster. This parameter isn't case sensitive.
     */
    public readonly snapshotIdentifier!: pulumi.Output<string | undefined>;
    /**
     * The list of tags for the cluster parameter group.
     */
    public readonly tags!: pulumi.Output<outputs.redshift.ClusterTag[] | undefined>;
    /**
     * A list of Virtual Private Cloud (VPC) security groups to be associated with the cluster.
     */
    public readonly vpcSecurityGroupIds!: pulumi.Output<string[] | undefined>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.clusterType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterType'");
            }
            if ((!args || args.dBName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dBName'");
            }
            if ((!args || args.masterUserPassword === undefined) && !opts.urn) {
                throw new Error("Missing required property 'masterUserPassword'");
            }
            if ((!args || args.masterUsername === undefined) && !opts.urn) {
                throw new Error("Missing required property 'masterUsername'");
            }
            if ((!args || args.nodeType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeType'");
            }
            resourceInputs["allowVersionUpgrade"] = args ? args.allowVersionUpgrade : undefined;
            resourceInputs["aquaConfigurationStatus"] = args ? args.aquaConfigurationStatus : undefined;
            resourceInputs["automatedSnapshotRetentionPeriod"] = args ? args.automatedSnapshotRetentionPeriod : undefined;
            resourceInputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            resourceInputs["availabilityZoneRelocation"] = args ? args.availabilityZoneRelocation : undefined;
            resourceInputs["availabilityZoneRelocationStatus"] = args ? args.availabilityZoneRelocationStatus : undefined;
            resourceInputs["classic"] = args ? args.classic : undefined;
            resourceInputs["clusterIdentifier"] = args ? args.clusterIdentifier : undefined;
            resourceInputs["clusterParameterGroupName"] = args ? args.clusterParameterGroupName : undefined;
            resourceInputs["clusterSecurityGroups"] = args ? args.clusterSecurityGroups : undefined;
            resourceInputs["clusterSubnetGroupName"] = args ? args.clusterSubnetGroupName : undefined;
            resourceInputs["clusterType"] = args ? args.clusterType : undefined;
            resourceInputs["clusterVersion"] = args ? args.clusterVersion : undefined;
            resourceInputs["dBName"] = args ? args.dBName : undefined;
            resourceInputs["deferMaintenance"] = args ? args.deferMaintenance : undefined;
            resourceInputs["deferMaintenanceDuration"] = args ? args.deferMaintenanceDuration : undefined;
            resourceInputs["deferMaintenanceEndTime"] = args ? args.deferMaintenanceEndTime : undefined;
            resourceInputs["deferMaintenanceStartTime"] = args ? args.deferMaintenanceStartTime : undefined;
            resourceInputs["destinationRegion"] = args ? args.destinationRegion : undefined;
            resourceInputs["elasticIp"] = args ? args.elasticIp : undefined;
            resourceInputs["encrypted"] = args ? args.encrypted : undefined;
            resourceInputs["endpoint"] = args ? args.endpoint : undefined;
            resourceInputs["enhancedVpcRouting"] = args ? args.enhancedVpcRouting : undefined;
            resourceInputs["hsmClientCertificateIdentifier"] = args ? args.hsmClientCertificateIdentifier : undefined;
            resourceInputs["hsmConfigurationIdentifier"] = args ? args.hsmConfigurationIdentifier : undefined;
            resourceInputs["iamRoles"] = args ? args.iamRoles : undefined;
            resourceInputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            resourceInputs["loggingProperties"] = args ? args.loggingProperties : undefined;
            resourceInputs["maintenanceTrackName"] = args ? args.maintenanceTrackName : undefined;
            resourceInputs["manualSnapshotRetentionPeriod"] = args ? args.manualSnapshotRetentionPeriod : undefined;
            resourceInputs["masterUserPassword"] = args ? args.masterUserPassword : undefined;
            resourceInputs["masterUsername"] = args ? args.masterUsername : undefined;
            resourceInputs["nodeType"] = args ? args.nodeType : undefined;
            resourceInputs["numberOfNodes"] = args ? args.numberOfNodes : undefined;
            resourceInputs["ownerAccount"] = args ? args.ownerAccount : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["preferredMaintenanceWindow"] = args ? args.preferredMaintenanceWindow : undefined;
            resourceInputs["publiclyAccessible"] = args ? args.publiclyAccessible : undefined;
            resourceInputs["resourceAction"] = args ? args.resourceAction : undefined;
            resourceInputs["revisionTarget"] = args ? args.revisionTarget : undefined;
            resourceInputs["rotateEncryptionKey"] = args ? args.rotateEncryptionKey : undefined;
            resourceInputs["snapshotClusterIdentifier"] = args ? args.snapshotClusterIdentifier : undefined;
            resourceInputs["snapshotCopyGrantName"] = args ? args.snapshotCopyGrantName : undefined;
            resourceInputs["snapshotCopyManual"] = args ? args.snapshotCopyManual : undefined;
            resourceInputs["snapshotCopyRetentionPeriod"] = args ? args.snapshotCopyRetentionPeriod : undefined;
            resourceInputs["snapshotIdentifier"] = args ? args.snapshotIdentifier : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcSecurityGroupIds"] = args ? args.vpcSecurityGroupIds : undefined;
            resourceInputs["deferMaintenanceIdentifier"] = undefined /*out*/;
        } else {
            resourceInputs["allowVersionUpgrade"] = undefined /*out*/;
            resourceInputs["aquaConfigurationStatus"] = undefined /*out*/;
            resourceInputs["automatedSnapshotRetentionPeriod"] = undefined /*out*/;
            resourceInputs["availabilityZone"] = undefined /*out*/;
            resourceInputs["availabilityZoneRelocation"] = undefined /*out*/;
            resourceInputs["availabilityZoneRelocationStatus"] = undefined /*out*/;
            resourceInputs["classic"] = undefined /*out*/;
            resourceInputs["clusterIdentifier"] = undefined /*out*/;
            resourceInputs["clusterParameterGroupName"] = undefined /*out*/;
            resourceInputs["clusterSecurityGroups"] = undefined /*out*/;
            resourceInputs["clusterSubnetGroupName"] = undefined /*out*/;
            resourceInputs["clusterType"] = undefined /*out*/;
            resourceInputs["clusterVersion"] = undefined /*out*/;
            resourceInputs["dBName"] = undefined /*out*/;
            resourceInputs["deferMaintenance"] = undefined /*out*/;
            resourceInputs["deferMaintenanceDuration"] = undefined /*out*/;
            resourceInputs["deferMaintenanceEndTime"] = undefined /*out*/;
            resourceInputs["deferMaintenanceIdentifier"] = undefined /*out*/;
            resourceInputs["deferMaintenanceStartTime"] = undefined /*out*/;
            resourceInputs["destinationRegion"] = undefined /*out*/;
            resourceInputs["elasticIp"] = undefined /*out*/;
            resourceInputs["encrypted"] = undefined /*out*/;
            resourceInputs["endpoint"] = undefined /*out*/;
            resourceInputs["enhancedVpcRouting"] = undefined /*out*/;
            resourceInputs["hsmClientCertificateIdentifier"] = undefined /*out*/;
            resourceInputs["hsmConfigurationIdentifier"] = undefined /*out*/;
            resourceInputs["iamRoles"] = undefined /*out*/;
            resourceInputs["kmsKeyId"] = undefined /*out*/;
            resourceInputs["loggingProperties"] = undefined /*out*/;
            resourceInputs["maintenanceTrackName"] = undefined /*out*/;
            resourceInputs["manualSnapshotRetentionPeriod"] = undefined /*out*/;
            resourceInputs["masterUserPassword"] = undefined /*out*/;
            resourceInputs["masterUsername"] = undefined /*out*/;
            resourceInputs["nodeType"] = undefined /*out*/;
            resourceInputs["numberOfNodes"] = undefined /*out*/;
            resourceInputs["ownerAccount"] = undefined /*out*/;
            resourceInputs["port"] = undefined /*out*/;
            resourceInputs["preferredMaintenanceWindow"] = undefined /*out*/;
            resourceInputs["publiclyAccessible"] = undefined /*out*/;
            resourceInputs["resourceAction"] = undefined /*out*/;
            resourceInputs["revisionTarget"] = undefined /*out*/;
            resourceInputs["rotateEncryptionKey"] = undefined /*out*/;
            resourceInputs["snapshotClusterIdentifier"] = undefined /*out*/;
            resourceInputs["snapshotCopyGrantName"] = undefined /*out*/;
            resourceInputs["snapshotCopyManual"] = undefined /*out*/;
            resourceInputs["snapshotCopyRetentionPeriod"] = undefined /*out*/;
            resourceInputs["snapshotIdentifier"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["vpcSecurityGroupIds"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Cluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * Major version upgrades can be applied during the maintenance window to the Amazon Redshift engine that is running on the cluster. Default value is True
     */
    allowVersionUpgrade?: pulumi.Input<boolean>;
    /**
     * The value represents how the cluster is configured to use AQUA (Advanced Query Accelerator) after the cluster is restored. Possible values include the following.
     *
     * enabled - Use AQUA if it is available for the current Region and Amazon Redshift node type.
     * disabled - Don't use AQUA.
     * auto - Amazon Redshift determines whether to use AQUA.
     */
    aquaConfigurationStatus?: pulumi.Input<string>;
    /**
     * The number of days that automated snapshots are retained. If the value is 0, automated snapshots are disabled. Default value is 1
     */
    automatedSnapshotRetentionPeriod?: pulumi.Input<number>;
    /**
     * The EC2 Availability Zone (AZ) in which you want Amazon Redshift to provision the cluster. Default: A random, system-chosen Availability Zone in the region that is specified by the endpoint
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * The option to enable relocation for an Amazon Redshift cluster between Availability Zones after the cluster modification is complete.
     */
    availabilityZoneRelocation?: pulumi.Input<boolean>;
    /**
     * The availability zone relocation status of the cluster
     */
    availabilityZoneRelocationStatus?: pulumi.Input<string>;
    /**
     * A boolean value indicating whether the resize operation is using the classic resize process. If you don't provide this parameter or set the value to false , the resize type is elastic.
     */
    classic?: pulumi.Input<boolean>;
    /**
     * A unique identifier for the cluster. You use this identifier to refer to the cluster for any subsequent cluster operations such as deleting or modifying. All alphabetical characters must be lower case, no hypens at the end, no two consecutive hyphens. Cluster name should be unique for all clusters within an AWS account
     */
    clusterIdentifier?: pulumi.Input<string>;
    /**
     * The name of the parameter group to be associated with this cluster.
     */
    clusterParameterGroupName?: pulumi.Input<string>;
    /**
     * A list of security groups to be associated with this cluster.
     */
    clusterSecurityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of a cluster subnet group to be associated with this cluster.
     */
    clusterSubnetGroupName?: pulumi.Input<string>;
    /**
     * The type of the cluster. When cluster type is specified as single-node, the NumberOfNodes parameter is not required and if multi-node, the NumberOfNodes parameter is required
     */
    clusterType: pulumi.Input<string>;
    /**
     * The version of the Amazon Redshift engine software that you want to deploy on the cluster.The version selected runs on all the nodes in the cluster.
     */
    clusterVersion?: pulumi.Input<string>;
    /**
     * The name of the first database to be created when the cluster is created. To create additional databases after the cluster is created, connect to the cluster with a SQL client and use SQL commands to create a database.
     */
    dBName: pulumi.Input<string>;
    /**
     * A boolean indicating whether to enable the deferred maintenance window.
     */
    deferMaintenance?: pulumi.Input<boolean>;
    /**
     * An integer indicating the duration of the maintenance window in days. If you specify a duration, you can't specify an end time. The duration must be 45 days or less.
     */
    deferMaintenanceDuration?: pulumi.Input<number>;
    /**
     * A timestamp indicating end time for the deferred maintenance window. If you specify an end time, you can't specify a duration.
     */
    deferMaintenanceEndTime?: pulumi.Input<string>;
    /**
     * A timestamp indicating the start time for the deferred maintenance window.
     */
    deferMaintenanceStartTime?: pulumi.Input<string>;
    /**
     * The destination AWS Region that you want to copy snapshots to. Constraints: Must be the name of a valid AWS Region. For more information, see Regions and Endpoints in the Amazon Web Services [https://docs.aws.amazon.com/general/latest/gr/rande.html#redshift_region] General Reference
     */
    destinationRegion?: pulumi.Input<string>;
    /**
     * The Elastic IP (EIP) address for the cluster.
     */
    elasticIp?: pulumi.Input<string>;
    /**
     * If true, the data in the cluster is encrypted at rest.
     */
    encrypted?: pulumi.Input<boolean>;
    endpoint?: pulumi.Input<inputs.redshift.ClusterEndpointArgs>;
    /**
     * An option that specifies whether to create the cluster with enhanced VPC routing enabled. To create a cluster that uses enhanced VPC routing, the cluster must be in a VPC. For more information, see Enhanced VPC Routing in the Amazon Redshift Cluster Management Guide.
     *
     * If this option is true , enhanced VPC routing is enabled.
     *
     * Default: false
     */
    enhancedVpcRouting?: pulumi.Input<boolean>;
    /**
     * Specifies the name of the HSM client certificate the Amazon Redshift cluster uses to retrieve the data encryption keys stored in an HSM
     */
    hsmClientCertificateIdentifier?: pulumi.Input<string>;
    /**
     * Specifies the name of the HSM configuration that contains the information the Amazon Redshift cluster can use to retrieve and store keys in an HSM.
     */
    hsmConfigurationIdentifier?: pulumi.Input<string>;
    /**
     * A list of AWS Identity and Access Management (IAM) roles that can be used by the cluster to access other AWS services. You must supply the IAM roles in their Amazon Resource Name (ARN) format. You can supply up to 10 IAM roles in a single request
     */
    iamRoles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The AWS Key Management Service (KMS) key ID of the encryption key that you want to use to encrypt data in the cluster.
     */
    kmsKeyId?: pulumi.Input<string>;
    loggingProperties?: pulumi.Input<inputs.redshift.ClusterLoggingPropertiesArgs>;
    /**
     * The name for the maintenance track that you want to assign for the cluster. This name change is asynchronous. The new track name stays in the PendingModifiedValues for the cluster until the next maintenance window. When the maintenance track changes, the cluster is switched to the latest cluster release available for the maintenance track. At this point, the maintenance track name is applied.
     */
    maintenanceTrackName?: pulumi.Input<string>;
    /**
     * The number of days to retain newly copied snapshots in the destination AWS Region after they are copied from the source AWS Region. If the value is -1, the manual snapshot is retained indefinitely.
     *
     * The value must be either -1 or an integer between 1 and 3,653.
     */
    manualSnapshotRetentionPeriod?: pulumi.Input<number>;
    /**
     * The password associated with the master user account for the cluster that is being created. Password must be between 8 and 64 characters in length, should have at least one uppercase letter.Must contain at least one lowercase letter.Must contain one number.Can be any printable ASCII character.
     */
    masterUserPassword: pulumi.Input<string>;
    /**
     * The user name associated with the master user account for the cluster that is being created. The user name can't be PUBLIC and first character must be a letter.
     */
    masterUsername: pulumi.Input<string>;
    /**
     * The node type to be provisioned for the cluster.Valid Values: ds2.xlarge | ds2.8xlarge | dc1.large | dc1.8xlarge | dc2.large | dc2.8xlarge | ra3.4xlarge | ra3.16xlarge
     */
    nodeType: pulumi.Input<string>;
    /**
     * The number of compute nodes in the cluster. This parameter is required when the ClusterType parameter is specified as multi-node.
     */
    numberOfNodes?: pulumi.Input<number>;
    ownerAccount?: pulumi.Input<string>;
    /**
     * The port number on which the cluster accepts incoming connections. The cluster is accessible only via the JDBC and ODBC connection strings
     */
    port?: pulumi.Input<number>;
    /**
     * The weekly time range (in UTC) during which automated cluster maintenance can occur.
     */
    preferredMaintenanceWindow?: pulumi.Input<string>;
    /**
     * If true, the cluster can be accessed from a public network.
     */
    publiclyAccessible?: pulumi.Input<boolean>;
    /**
     * The Redshift operation to be performed. Resource Action supports pause-cluster, resume-cluster APIs
     */
    resourceAction?: pulumi.Input<string>;
    /**
     * The identifier of the database revision. You can retrieve this value from the response to the DescribeClusterDbRevisions request.
     */
    revisionTarget?: pulumi.Input<string>;
    /**
     * A boolean indicating if we want to rotate Encryption Keys.
     */
    rotateEncryptionKey?: pulumi.Input<boolean>;
    /**
     * The name of the cluster the source snapshot was created from. This parameter is required if your IAM user has a policy containing a snapshot resource element that specifies anything other than * for the cluster name.
     */
    snapshotClusterIdentifier?: pulumi.Input<string>;
    /**
     * The name of the snapshot copy grant to use when snapshots of an AWS KMS-encrypted cluster are copied to the destination region.
     */
    snapshotCopyGrantName?: pulumi.Input<string>;
    /**
     * Indicates whether to apply the snapshot retention period to newly copied manual snapshots instead of automated snapshots.
     */
    snapshotCopyManual?: pulumi.Input<boolean>;
    /**
     * The number of days to retain automated snapshots in the destination region after they are copied from the source region. 
     *
     *  Default is 7. 
     *
     *  Constraints: Must be at least 1 and no more than 35.
     */
    snapshotCopyRetentionPeriod?: pulumi.Input<number>;
    /**
     * The name of the snapshot from which to create the new cluster. This parameter isn't case sensitive.
     */
    snapshotIdentifier?: pulumi.Input<string>;
    /**
     * The list of tags for the cluster parameter group.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.redshift.ClusterTagArgs>[]>;
    /**
     * A list of Virtual Private Cloud (VPC) security groups to be associated with the cluster.
     */
    vpcSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
}
