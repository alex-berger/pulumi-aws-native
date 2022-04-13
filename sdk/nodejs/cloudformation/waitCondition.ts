// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::CloudFormation::WaitCondition
 *
 * @deprecated WaitCondition is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class WaitCondition extends pulumi.CustomResource {
    /**
     * Get an existing WaitCondition resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): WaitCondition {
        pulumi.log.warn("WaitCondition is deprecated: WaitCondition is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new WaitCondition(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:cloudformation:WaitCondition';

    /**
     * Returns true if the given object is an instance of WaitCondition.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WaitCondition {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WaitCondition.__pulumiType;
    }

    public readonly count!: pulumi.Output<number | undefined>;
    public /*out*/ readonly data!: pulumi.Output<any>;
    public readonly handle!: pulumi.Output<string | undefined>;
    public readonly timeout!: pulumi.Output<string | undefined>;

    /**
     * Create a WaitCondition resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated WaitCondition is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args?: WaitConditionArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("WaitCondition is deprecated: WaitCondition is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["count"] = args ? args.count : undefined;
            resourceInputs["handle"] = args ? args.handle : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
            resourceInputs["data"] = undefined /*out*/;
        } else {
            resourceInputs["count"] = undefined /*out*/;
            resourceInputs["data"] = undefined /*out*/;
            resourceInputs["handle"] = undefined /*out*/;
            resourceInputs["timeout"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WaitCondition.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a WaitCondition resource.
 */
export interface WaitConditionArgs {
    count?: pulumi.Input<number>;
    handle?: pulumi.Input<string>;
    timeout?: pulumi.Input<string>;
}
