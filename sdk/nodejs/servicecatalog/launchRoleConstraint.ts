// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ServiceCatalog::LaunchRoleConstraint
 *
 * @deprecated LaunchRoleConstraint is not yet supported by AWS Cloud Control API, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class LaunchRoleConstraint extends pulumi.CustomResource {
    /**
     * Get an existing LaunchRoleConstraint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): LaunchRoleConstraint {
        pulumi.log.warn("LaunchRoleConstraint is deprecated: LaunchRoleConstraint is not yet supported by AWS Cloud Control API, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new LaunchRoleConstraint(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:servicecatalog:LaunchRoleConstraint';

    /**
     * Returns true if the given object is an instance of LaunchRoleConstraint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LaunchRoleConstraint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LaunchRoleConstraint.__pulumiType;
    }

    public readonly acceptLanguage!: pulumi.Output<string | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly localRoleName!: pulumi.Output<string | undefined>;
    public readonly portfolioId!: pulumi.Output<string>;
    public readonly productId!: pulumi.Output<string>;
    public readonly roleArn!: pulumi.Output<string | undefined>;

    /**
     * Create a LaunchRoleConstraint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated LaunchRoleConstraint is not yet supported by AWS Cloud Control API, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: LaunchRoleConstraintArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("LaunchRoleConstraint is deprecated: LaunchRoleConstraint is not yet supported by AWS Cloud Control API, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.portfolioId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'portfolioId'");
            }
            if ((!args || args.productId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'productId'");
            }
            inputs["acceptLanguage"] = args ? args.acceptLanguage : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["localRoleName"] = args ? args.localRoleName : undefined;
            inputs["portfolioId"] = args ? args.portfolioId : undefined;
            inputs["productId"] = args ? args.productId : undefined;
            inputs["roleArn"] = args ? args.roleArn : undefined;
        } else {
            inputs["acceptLanguage"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["localRoleName"] = undefined /*out*/;
            inputs["portfolioId"] = undefined /*out*/;
            inputs["productId"] = undefined /*out*/;
            inputs["roleArn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(LaunchRoleConstraint.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a LaunchRoleConstraint resource.
 */
export interface LaunchRoleConstraintArgs {
    acceptLanguage?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    localRoleName?: pulumi.Input<string>;
    portfolioId: pulumi.Input<string>;
    productId: pulumi.Input<string>;
    roleArn?: pulumi.Input<string>;
}