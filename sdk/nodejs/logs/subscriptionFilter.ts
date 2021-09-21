// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Logs::SubscriptionFilter
 *
 * @deprecated SubscriptionFilter is not yet supported by AWS Cloud Control API, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class SubscriptionFilter extends pulumi.CustomResource {
    /**
     * Get an existing SubscriptionFilter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SubscriptionFilter {
        pulumi.log.warn("SubscriptionFilter is deprecated: SubscriptionFilter is not yet supported by AWS Cloud Control API, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new SubscriptionFilter(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:logs:SubscriptionFilter';

    /**
     * Returns true if the given object is an instance of SubscriptionFilter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SubscriptionFilter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SubscriptionFilter.__pulumiType;
    }

    public readonly destinationArn!: pulumi.Output<string>;
    public readonly filterPattern!: pulumi.Output<string>;
    public readonly logGroupName!: pulumi.Output<string>;
    public readonly roleArn!: pulumi.Output<string | undefined>;

    /**
     * Create a SubscriptionFilter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated SubscriptionFilter is not yet supported by AWS Cloud Control API, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: SubscriptionFilterArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("SubscriptionFilter is deprecated: SubscriptionFilter is not yet supported by AWS Cloud Control API, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.destinationArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destinationArn'");
            }
            if ((!args || args.filterPattern === undefined) && !opts.urn) {
                throw new Error("Missing required property 'filterPattern'");
            }
            if ((!args || args.logGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'logGroupName'");
            }
            inputs["destinationArn"] = args ? args.destinationArn : undefined;
            inputs["filterPattern"] = args ? args.filterPattern : undefined;
            inputs["logGroupName"] = args ? args.logGroupName : undefined;
            inputs["roleArn"] = args ? args.roleArn : undefined;
        } else {
            inputs["destinationArn"] = undefined /*out*/;
            inputs["filterPattern"] = undefined /*out*/;
            inputs["logGroupName"] = undefined /*out*/;
            inputs["roleArn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SubscriptionFilter.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a SubscriptionFilter resource.
 */
export interface SubscriptionFilterArgs {
    destinationArn: pulumi.Input<string>;
    filterPattern: pulumi.Input<string>;
    logGroupName: pulumi.Input<string>;
    roleArn?: pulumi.Input<string>;
}