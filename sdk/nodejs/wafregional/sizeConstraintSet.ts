// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::WAFRegional::SizeConstraintSet
 *
 * @deprecated SizeConstraintSet is not yet supported by AWS Cloud Control API, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class SizeConstraintSet extends pulumi.CustomResource {
    /**
     * Get an existing SizeConstraintSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SizeConstraintSet {
        pulumi.log.warn("SizeConstraintSet is deprecated: SizeConstraintSet is not yet supported by AWS Cloud Control API, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new SizeConstraintSet(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:wafregional:SizeConstraintSet';

    /**
     * Returns true if the given object is an instance of SizeConstraintSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SizeConstraintSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SizeConstraintSet.__pulumiType;
    }

    public readonly name!: pulumi.Output<string>;
    public readonly sizeConstraints!: pulumi.Output<outputs.wafregional.SizeConstraintSetSizeConstraint[] | undefined>;

    /**
     * Create a SizeConstraintSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated SizeConstraintSet is not yet supported by AWS Cloud Control API, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: SizeConstraintSetArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("SizeConstraintSet is deprecated: SizeConstraintSet is not yet supported by AWS Cloud Control API, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["sizeConstraints"] = args ? args.sizeConstraints : undefined;
        } else {
            inputs["name"] = undefined /*out*/;
            inputs["sizeConstraints"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SizeConstraintSet.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a SizeConstraintSet resource.
 */
export interface SizeConstraintSetArgs {
    name: pulumi.Input<string>;
    sizeConstraints?: pulumi.Input<pulumi.Input<inputs.wafregional.SizeConstraintSetSizeConstraintArgs>[]>;
}