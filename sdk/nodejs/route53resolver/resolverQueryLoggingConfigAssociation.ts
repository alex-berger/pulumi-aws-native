// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverqueryloggingconfigassociation.html
 */
export class ResolverQueryLoggingConfigAssociation extends pulumi.CustomResource {
    /**
     * Get an existing ResolverQueryLoggingConfigAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ResolverQueryLoggingConfigAssociation {
        return new ResolverQueryLoggingConfigAssociation(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:Route53Resolver:ResolverQueryLoggingConfigAssociation';

    /**
     * Returns true if the given object is an instance of ResolverQueryLoggingConfigAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ResolverQueryLoggingConfigAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ResolverQueryLoggingConfigAssociation.__pulumiType;
    }

    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    public /*out*/ readonly error!: pulumi.Output<string>;
    public /*out*/ readonly errorMessage!: pulumi.Output<string>;
    public /*out*/ readonly id!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverqueryloggingconfigassociation.html#cfn-route53resolver-resolverqueryloggingconfigassociation-resolverquerylogconfigid
     */
    public readonly resolverQueryLogConfigId!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverqueryloggingconfigassociation.html#cfn-route53resolver-resolverqueryloggingconfigassociation-resourceid
     */
    public readonly resourceId!: pulumi.Output<string | undefined>;
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a ResolverQueryLoggingConfigAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ResolverQueryLoggingConfigAssociationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            inputs["resolverQueryLogConfigId"] = args ? args.resolverQueryLogConfigId : undefined;
            inputs["resourceId"] = args ? args.resourceId : undefined;
            inputs["creationTime"] = undefined /*out*/;
            inputs["error"] = undefined /*out*/;
            inputs["errorMessage"] = undefined /*out*/;
            inputs["id"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        } else {
            inputs["creationTime"] = undefined /*out*/;
            inputs["error"] = undefined /*out*/;
            inputs["errorMessage"] = undefined /*out*/;
            inputs["id"] = undefined /*out*/;
            inputs["resolverQueryLogConfigId"] = undefined /*out*/;
            inputs["resourceId"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ResolverQueryLoggingConfigAssociation.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ResolverQueryLoggingConfigAssociation resource.
 */
export interface ResolverQueryLoggingConfigAssociationArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverqueryloggingconfigassociation.html#cfn-route53resolver-resolverqueryloggingconfigassociation-resolverquerylogconfigid
     */
    readonly resolverQueryLogConfigId?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverqueryloggingconfigassociation.html#cfn-route53resolver-resolverqueryloggingconfigassociation-resourceid
     */
    readonly resourceId?: pulumi.Input<string>;
}
