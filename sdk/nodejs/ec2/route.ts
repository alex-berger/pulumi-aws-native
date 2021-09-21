// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EC2::Route
 *
 * @deprecated Route is not yet supported by AWS Cloud Control API, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class Route extends pulumi.CustomResource {
    /**
     * Get an existing Route resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Route {
        pulumi.log.warn("Route is deprecated: Route is not yet supported by AWS Cloud Control API, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new Route(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ec2:Route';

    /**
     * Returns true if the given object is an instance of Route.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Route {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Route.__pulumiType;
    }

    public readonly carrierGatewayId!: pulumi.Output<string | undefined>;
    public readonly destinationCidrBlock!: pulumi.Output<string | undefined>;
    public readonly destinationIpv6CidrBlock!: pulumi.Output<string | undefined>;
    public readonly egressOnlyInternetGatewayId!: pulumi.Output<string | undefined>;
    public readonly gatewayId!: pulumi.Output<string | undefined>;
    public readonly instanceId!: pulumi.Output<string | undefined>;
    public readonly localGatewayId!: pulumi.Output<string | undefined>;
    public readonly natGatewayId!: pulumi.Output<string | undefined>;
    public readonly networkInterfaceId!: pulumi.Output<string | undefined>;
    public readonly routeTableId!: pulumi.Output<string>;
    public readonly transitGatewayId!: pulumi.Output<string | undefined>;
    public readonly vpcEndpointId!: pulumi.Output<string | undefined>;
    public readonly vpcPeeringConnectionId!: pulumi.Output<string | undefined>;

    /**
     * Create a Route resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated Route is not yet supported by AWS Cloud Control API, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: RouteArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Route is deprecated: Route is not yet supported by AWS Cloud Control API, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.routeTableId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'routeTableId'");
            }
            inputs["carrierGatewayId"] = args ? args.carrierGatewayId : undefined;
            inputs["destinationCidrBlock"] = args ? args.destinationCidrBlock : undefined;
            inputs["destinationIpv6CidrBlock"] = args ? args.destinationIpv6CidrBlock : undefined;
            inputs["egressOnlyInternetGatewayId"] = args ? args.egressOnlyInternetGatewayId : undefined;
            inputs["gatewayId"] = args ? args.gatewayId : undefined;
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["localGatewayId"] = args ? args.localGatewayId : undefined;
            inputs["natGatewayId"] = args ? args.natGatewayId : undefined;
            inputs["networkInterfaceId"] = args ? args.networkInterfaceId : undefined;
            inputs["routeTableId"] = args ? args.routeTableId : undefined;
            inputs["transitGatewayId"] = args ? args.transitGatewayId : undefined;
            inputs["vpcEndpointId"] = args ? args.vpcEndpointId : undefined;
            inputs["vpcPeeringConnectionId"] = args ? args.vpcPeeringConnectionId : undefined;
        } else {
            inputs["carrierGatewayId"] = undefined /*out*/;
            inputs["destinationCidrBlock"] = undefined /*out*/;
            inputs["destinationIpv6CidrBlock"] = undefined /*out*/;
            inputs["egressOnlyInternetGatewayId"] = undefined /*out*/;
            inputs["gatewayId"] = undefined /*out*/;
            inputs["instanceId"] = undefined /*out*/;
            inputs["localGatewayId"] = undefined /*out*/;
            inputs["natGatewayId"] = undefined /*out*/;
            inputs["networkInterfaceId"] = undefined /*out*/;
            inputs["routeTableId"] = undefined /*out*/;
            inputs["transitGatewayId"] = undefined /*out*/;
            inputs["vpcEndpointId"] = undefined /*out*/;
            inputs["vpcPeeringConnectionId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Route.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Route resource.
 */
export interface RouteArgs {
    carrierGatewayId?: pulumi.Input<string>;
    destinationCidrBlock?: pulumi.Input<string>;
    destinationIpv6CidrBlock?: pulumi.Input<string>;
    egressOnlyInternetGatewayId?: pulumi.Input<string>;
    gatewayId?: pulumi.Input<string>;
    instanceId?: pulumi.Input<string>;
    localGatewayId?: pulumi.Input<string>;
    natGatewayId?: pulumi.Input<string>;
    networkInterfaceId?: pulumi.Input<string>;
    routeTableId: pulumi.Input<string>;
    transitGatewayId?: pulumi.Input<string>;
    vpcEndpointId?: pulumi.Input<string>;
    vpcPeeringConnectionId?: pulumi.Input<string>;
}