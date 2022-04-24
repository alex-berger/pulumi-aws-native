// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::RDS::DBSubnetGroup
 *
 * @deprecated DBSubnetGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class DBSubnetGroup extends pulumi.CustomResource {
    /**
     * Get an existing DBSubnetGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DBSubnetGroup {
        pulumi.log.warn("DBSubnetGroup is deprecated: DBSubnetGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new DBSubnetGroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:rds:DBSubnetGroup';

    /**
     * Returns true if the given object is an instance of DBSubnetGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DBSubnetGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DBSubnetGroup.__pulumiType;
    }

    public readonly dBSubnetGroupDescription!: pulumi.Output<string>;
    public readonly dBSubnetGroupName!: pulumi.Output<string | undefined>;
    public readonly subnetIds!: pulumi.Output<string[]>;
    public readonly tags!: pulumi.Output<outputs.rds.DBSubnetGroupTag[] | undefined>;

    /**
     * Create a DBSubnetGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated DBSubnetGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: DBSubnetGroupArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("DBSubnetGroup is deprecated: DBSubnetGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.dBSubnetGroupDescription === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dBSubnetGroupDescription'");
            }
            if ((!args || args.subnetIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetIds'");
            }
            resourceInputs["dBSubnetGroupDescription"] = args ? args.dBSubnetGroupDescription : undefined;
            resourceInputs["dBSubnetGroupName"] = args ? args.dBSubnetGroupName : undefined;
            resourceInputs["subnetIds"] = args ? args.subnetIds : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
        } else {
            resourceInputs["dBSubnetGroupDescription"] = undefined /*out*/;
            resourceInputs["dBSubnetGroupName"] = undefined /*out*/;
            resourceInputs["subnetIds"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DBSubnetGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DBSubnetGroup resource.
 */
export interface DBSubnetGroupArgs {
    dBSubnetGroupDescription: pulumi.Input<string>;
    dBSubnetGroupName?: pulumi.Input<string>;
    subnetIds: pulumi.Input<pulumi.Input<string>[]>;
    tags?: pulumi.Input<pulumi.Input<inputs.rds.DBSubnetGroupTagArgs>[]>;
}
