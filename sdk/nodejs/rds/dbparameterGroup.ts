// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::RDS::DBParameterGroup
 *
 * @deprecated DBParameterGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class DBParameterGroup extends pulumi.CustomResource {
    /**
     * Get an existing DBParameterGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DBParameterGroup {
        pulumi.log.warn("DBParameterGroup is deprecated: DBParameterGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new DBParameterGroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:rds:DBParameterGroup';

    /**
     * Returns true if the given object is an instance of DBParameterGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DBParameterGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DBParameterGroup.__pulumiType;
    }

    public readonly description!: pulumi.Output<string>;
    public readonly family!: pulumi.Output<string>;
    public readonly parameters!: pulumi.Output<any | undefined>;
    public readonly tags!: pulumi.Output<outputs.rds.DBParameterGroupTag[] | undefined>;

    /**
     * Create a DBParameterGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated DBParameterGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: DBParameterGroupArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("DBParameterGroup is deprecated: DBParameterGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.family === undefined) && !opts.urn) {
                throw new Error("Missing required property 'family'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["family"] = args ? args.family : undefined;
            resourceInputs["parameters"] = args ? args.parameters : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
        } else {
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["family"] = undefined /*out*/;
            resourceInputs["parameters"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DBParameterGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DBParameterGroup resource.
 */
export interface DBParameterGroupArgs {
    description: pulumi.Input<string>;
    family: pulumi.Input<string>;
    parameters?: any;
    tags?: pulumi.Input<pulumi.Input<inputs.rds.DBParameterGroupTagArgs>[]>;
}