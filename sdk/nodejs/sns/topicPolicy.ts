// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SNS::TopicPolicy
 *
 * @deprecated TopicPolicy is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class TopicPolicy extends pulumi.CustomResource {
    /**
     * Get an existing TopicPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): TopicPolicy {
        pulumi.log.warn("TopicPolicy is deprecated: TopicPolicy is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new TopicPolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:sns:TopicPolicy';

    /**
     * Returns true if the given object is an instance of TopicPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TopicPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TopicPolicy.__pulumiType;
    }

    public readonly policyDocument!: pulumi.Output<any>;
    public readonly topics!: pulumi.Output<string[]>;

    /**
     * Create a TopicPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated TopicPolicy is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: TopicPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("TopicPolicy is deprecated: TopicPolicy is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.policyDocument === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyDocument'");
            }
            if ((!args || args.topics === undefined) && !opts.urn) {
                throw new Error("Missing required property 'topics'");
            }
            resourceInputs["policyDocument"] = args ? args.policyDocument : undefined;
            resourceInputs["topics"] = args ? args.topics : undefined;
        } else {
            resourceInputs["policyDocument"] = undefined /*out*/;
            resourceInputs["topics"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TopicPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a TopicPolicy resource.
 */
export interface TopicPolicyArgs {
    policyDocument: any;
    topics: pulumi.Input<pulumi.Input<string>[]>;
}
