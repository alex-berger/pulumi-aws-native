// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EC2::TransitGatewayRoute
 *
 * @deprecated TransitGatewayRoute is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class TransitGatewayRoute extends pulumi.CustomResource {
    /**
     * Get an existing TransitGatewayRoute resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): TransitGatewayRoute {
        pulumi.log.warn("TransitGatewayRoute is deprecated: TransitGatewayRoute is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new TransitGatewayRoute(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ec2:TransitGatewayRoute';

    /**
     * Returns true if the given object is an instance of TransitGatewayRoute.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TransitGatewayRoute {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TransitGatewayRoute.__pulumiType;
    }

    public readonly blackhole!: pulumi.Output<boolean | undefined>;
    public readonly destinationCidrBlock!: pulumi.Output<string | undefined>;
    public readonly transitGatewayAttachmentId!: pulumi.Output<string | undefined>;
    public readonly transitGatewayRouteTableId!: pulumi.Output<string>;

    /**
     * Create a TransitGatewayRoute resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated TransitGatewayRoute is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: TransitGatewayRouteArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("TransitGatewayRoute is deprecated: TransitGatewayRoute is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.transitGatewayRouteTableId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'transitGatewayRouteTableId'");
            }
            resourceInputs["blackhole"] = args ? args.blackhole : undefined;
            resourceInputs["destinationCidrBlock"] = args ? args.destinationCidrBlock : undefined;
            resourceInputs["transitGatewayAttachmentId"] = args ? args.transitGatewayAttachmentId : undefined;
            resourceInputs["transitGatewayRouteTableId"] = args ? args.transitGatewayRouteTableId : undefined;
        } else {
            resourceInputs["blackhole"] = undefined /*out*/;
            resourceInputs["destinationCidrBlock"] = undefined /*out*/;
            resourceInputs["transitGatewayAttachmentId"] = undefined /*out*/;
            resourceInputs["transitGatewayRouteTableId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TransitGatewayRoute.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a TransitGatewayRoute resource.
 */
export interface TransitGatewayRouteArgs {
    blackhole?: pulumi.Input<boolean>;
    destinationCidrBlock?: pulumi.Input<string>;
    transitGatewayAttachmentId?: pulumi.Input<string>;
    transitGatewayRouteTableId: pulumi.Input<string>;
}
