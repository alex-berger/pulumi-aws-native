// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type schema for AWS::Batch::SchedulingPolicy
 */
export class SchedulingPolicy extends pulumi.CustomResource {
    /**
     * Get an existing SchedulingPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SchedulingPolicy {
        return new SchedulingPolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:batch:SchedulingPolicy';

    /**
     * Returns true if the given object is an instance of SchedulingPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SchedulingPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SchedulingPolicy.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly fairsharePolicy!: pulumi.Output<outputs.batch.SchedulingPolicyFairsharePolicy | undefined>;
    /**
     * Name of Scheduling Policy.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * A key-value pair to associate with a resource.
     */
    public readonly tags!: pulumi.Output<any | undefined>;

    /**
     * Create a SchedulingPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SchedulingPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["fairsharePolicy"] = args ? args.fairsharePolicy : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["fairsharePolicy"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SchedulingPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a SchedulingPolicy resource.
 */
export interface SchedulingPolicyArgs {
    fairsharePolicy?: pulumi.Input<inputs.batch.SchedulingPolicyFairsharePolicyArgs>;
    /**
     * Name of Scheduling Policy.
     */
    name?: pulumi.Input<string>;
    /**
     * A key-value pair to associate with a resource.
     */
    tags?: any;
}