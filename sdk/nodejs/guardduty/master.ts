// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::GuardDuty::Master
 *
 * @deprecated Master is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class Master extends pulumi.CustomResource {
    /**
     * Get an existing Master resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Master {
        pulumi.log.warn("Master is deprecated: Master is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new Master(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:guardduty:Master';

    /**
     * Returns true if the given object is an instance of Master.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Master {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Master.__pulumiType;
    }

    public readonly detectorId!: pulumi.Output<string>;
    public readonly invitationId!: pulumi.Output<string | undefined>;
    public /*out*/ readonly masterId!: pulumi.Output<string>;

    /**
     * Create a Master resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated Master is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: MasterArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Master is deprecated: Master is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.detectorId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'detectorId'");
            }
            resourceInputs["detectorId"] = args ? args.detectorId : undefined;
            resourceInputs["invitationId"] = args ? args.invitationId : undefined;
            resourceInputs["masterId"] = undefined /*out*/;
        } else {
            resourceInputs["detectorId"] = undefined /*out*/;
            resourceInputs["invitationId"] = undefined /*out*/;
            resourceInputs["masterId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Master.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Master resource.
 */
export interface MasterArgs {
    detectorId: pulumi.Input<string>;
    invitationId?: pulumi.Input<string>;
}
