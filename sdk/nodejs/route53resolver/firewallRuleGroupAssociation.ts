// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::Route53Resolver::FirewallRuleGroupAssociation.
 */
export class FirewallRuleGroupAssociation extends pulumi.CustomResource {
    /**
     * Get an existing FirewallRuleGroupAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): FirewallRuleGroupAssociation {
        return new FirewallRuleGroupAssociation(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:route53resolver:FirewallRuleGroupAssociation';

    /**
     * Returns true if the given object is an instance of FirewallRuleGroupAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallRuleGroupAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallRuleGroupAssociation.__pulumiType;
    }

    /**
     * Arn
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Rfc3339TimeString
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * The id of the creator request.
     */
    public /*out*/ readonly creatorRequestId!: pulumi.Output<string>;
    /**
     * FirewallRuleGroupId
     */
    public readonly firewallRuleGroupId!: pulumi.Output<string>;
    /**
     * ServicePrincipal
     */
    public /*out*/ readonly managedOwnerName!: pulumi.Output<string>;
    /**
     * Rfc3339TimeString
     */
    public /*out*/ readonly modificationTime!: pulumi.Output<string>;
    /**
     * MutationProtectionStatus
     */
    public readonly mutationProtection!: pulumi.Output<enums.route53resolver.FirewallRuleGroupAssociationMutationProtection | undefined>;
    /**
     * FirewallRuleGroupAssociationName
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * Priority
     */
    public readonly priority!: pulumi.Output<number>;
    /**
     * ResolverFirewallRuleGroupAssociation, possible values are COMPLETE, DELETING, UPDATING, and INACTIVE_OWNER_ACCOUNT_CLOSED.
     */
    public /*out*/ readonly status!: pulumi.Output<enums.route53resolver.FirewallRuleGroupAssociationStatus>;
    /**
     * FirewallDomainListAssociationStatus
     */
    public /*out*/ readonly statusMessage!: pulumi.Output<string>;
    /**
     * Tags
     */
    public readonly tags!: pulumi.Output<outputs.route53resolver.FirewallRuleGroupAssociationTag[] | undefined>;
    /**
     * VpcId
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a FirewallRuleGroupAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallRuleGroupAssociationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.firewallRuleGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'firewallRuleGroupId'");
            }
            if ((!args || args.priority === undefined) && !opts.urn) {
                throw new Error("Missing required property 'priority'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["firewallRuleGroupId"] = args ? args.firewallRuleGroupId : undefined;
            resourceInputs["mutationProtection"] = args ? args.mutationProtection : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["creatorRequestId"] = undefined /*out*/;
            resourceInputs["managedOwnerName"] = undefined /*out*/;
            resourceInputs["modificationTime"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["statusMessage"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["creatorRequestId"] = undefined /*out*/;
            resourceInputs["firewallRuleGroupId"] = undefined /*out*/;
            resourceInputs["managedOwnerName"] = undefined /*out*/;
            resourceInputs["modificationTime"] = undefined /*out*/;
            resourceInputs["mutationProtection"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["priority"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["statusMessage"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["vpcId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallRuleGroupAssociation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a FirewallRuleGroupAssociation resource.
 */
export interface FirewallRuleGroupAssociationArgs {
    /**
     * FirewallRuleGroupId
     */
    firewallRuleGroupId: pulumi.Input<string>;
    /**
     * MutationProtectionStatus
     */
    mutationProtection?: pulumi.Input<enums.route53resolver.FirewallRuleGroupAssociationMutationProtection>;
    /**
     * FirewallRuleGroupAssociationName
     */
    name?: pulumi.Input<string>;
    /**
     * Priority
     */
    priority: pulumi.Input<number>;
    /**
     * Tags
     */
    tags?: pulumi.Input<pulumi.Input<inputs.route53resolver.FirewallRuleGroupAssociationTagArgs>[]>;
    /**
     * VpcId
     */
    vpcId: pulumi.Input<string>;
}
