// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html
 */
export class DBProxyTargetGroup extends pulumi.CustomResource {
    /**
     * Get an existing DBProxyTargetGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DBProxyTargetGroup {
        return new DBProxyTargetGroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:RDS:DBProxyTargetGroup';

    /**
     * Returns true if the given object is an instance of DBProxyTargetGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DBProxyTargetGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DBProxyTargetGroup.__pulumiType;
    }

    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-connectionpoolconfigurationinfo
     */
    public readonly connectionPoolConfigurationInfo!: pulumi.Output<outputs.RDS.DBProxyTargetGroupConnectionPoolConfigurationInfoFormat | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-dbclusteridentifiers
     */
    public readonly dBClusterIdentifiers!: pulumi.Output<string[] | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-dbinstanceidentifiers
     */
    public readonly dBInstanceIdentifiers!: pulumi.Output<string[] | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-dbproxyname
     */
    public readonly dBProxyName!: pulumi.Output<string>;
    public /*out*/ readonly targetGroupArn!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-targetgroupname
     */
    public readonly targetGroupName!: pulumi.Output<string>;

    /**
     * Create a DBProxyTargetGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DBProxyTargetGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.dBProxyName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dBProxyName'");
            }
            if ((!args || args.targetGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetGroupName'");
            }
            inputs["connectionPoolConfigurationInfo"] = args ? args.connectionPoolConfigurationInfo : undefined;
            inputs["dBClusterIdentifiers"] = args ? args.dBClusterIdentifiers : undefined;
            inputs["dBInstanceIdentifiers"] = args ? args.dBInstanceIdentifiers : undefined;
            inputs["dBProxyName"] = args ? args.dBProxyName : undefined;
            inputs["targetGroupName"] = args ? args.targetGroupName : undefined;
            inputs["targetGroupArn"] = undefined /*out*/;
        } else {
            inputs["connectionPoolConfigurationInfo"] = undefined /*out*/;
            inputs["dBClusterIdentifiers"] = undefined /*out*/;
            inputs["dBInstanceIdentifiers"] = undefined /*out*/;
            inputs["dBProxyName"] = undefined /*out*/;
            inputs["targetGroupArn"] = undefined /*out*/;
            inputs["targetGroupName"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DBProxyTargetGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DBProxyTargetGroup resource.
 */
export interface DBProxyTargetGroupArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-connectionpoolconfigurationinfo
     */
    readonly connectionPoolConfigurationInfo?: pulumi.Input<inputs.RDS.DBProxyTargetGroupConnectionPoolConfigurationInfoFormatArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-dbclusteridentifiers
     */
    readonly dBClusterIdentifiers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-dbinstanceidentifiers
     */
    readonly dBInstanceIdentifiers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-dbproxyname
     */
    readonly dBProxyName: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-targetgroupname
     */
    readonly targetGroupName: pulumi.Input<string>;
}
