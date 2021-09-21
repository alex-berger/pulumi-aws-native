// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EC2::ClientVpnAuthorizationRule
 *
 * @deprecated ClientVpnAuthorizationRule is not yet supported by AWS Cloud Control API, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class ClientVpnAuthorizationRule extends pulumi.CustomResource {
    /**
     * Get an existing ClientVpnAuthorizationRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ClientVpnAuthorizationRule {
        pulumi.log.warn("ClientVpnAuthorizationRule is deprecated: ClientVpnAuthorizationRule is not yet supported by AWS Cloud Control API, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new ClientVpnAuthorizationRule(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ec2:ClientVpnAuthorizationRule';

    /**
     * Returns true if the given object is an instance of ClientVpnAuthorizationRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClientVpnAuthorizationRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClientVpnAuthorizationRule.__pulumiType;
    }

    public readonly accessGroupId!: pulumi.Output<string | undefined>;
    public readonly authorizeAllGroups!: pulumi.Output<boolean | undefined>;
    public readonly clientVpnEndpointId!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly targetNetworkCidr!: pulumi.Output<string>;

    /**
     * Create a ClientVpnAuthorizationRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated ClientVpnAuthorizationRule is not yet supported by AWS Cloud Control API, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: ClientVpnAuthorizationRuleArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("ClientVpnAuthorizationRule is deprecated: ClientVpnAuthorizationRule is not yet supported by AWS Cloud Control API, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.clientVpnEndpointId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientVpnEndpointId'");
            }
            if ((!args || args.targetNetworkCidr === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetNetworkCidr'");
            }
            inputs["accessGroupId"] = args ? args.accessGroupId : undefined;
            inputs["authorizeAllGroups"] = args ? args.authorizeAllGroups : undefined;
            inputs["clientVpnEndpointId"] = args ? args.clientVpnEndpointId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["targetNetworkCidr"] = args ? args.targetNetworkCidr : undefined;
        } else {
            inputs["accessGroupId"] = undefined /*out*/;
            inputs["authorizeAllGroups"] = undefined /*out*/;
            inputs["clientVpnEndpointId"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["targetNetworkCidr"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ClientVpnAuthorizationRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ClientVpnAuthorizationRule resource.
 */
export interface ClientVpnAuthorizationRuleArgs {
    accessGroupId?: pulumi.Input<string>;
    authorizeAllGroups?: pulumi.Input<boolean>;
    clientVpnEndpointId: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    targetNetworkCidr: pulumi.Input<string>;
}