// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ServiceDiscovery::Service
 *
 * @deprecated Service is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class Service extends pulumi.CustomResource {
    /**
     * Get an existing Service resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Service {
        pulumi.log.warn("Service is deprecated: Service is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new Service(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:servicediscovery:Service';

    /**
     * Returns true if the given object is an instance of Service.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Service {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Service.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly dnsConfig!: pulumi.Output<outputs.servicediscovery.ServiceDnsConfig | undefined>;
    public readonly healthCheckConfig!: pulumi.Output<outputs.servicediscovery.ServiceHealthCheckConfig | undefined>;
    public readonly healthCheckCustomConfig!: pulumi.Output<outputs.servicediscovery.ServiceHealthCheckCustomConfig | undefined>;
    public readonly name!: pulumi.Output<string | undefined>;
    public readonly namespaceId!: pulumi.Output<string | undefined>;
    public readonly tags!: pulumi.Output<outputs.servicediscovery.ServiceTag[] | undefined>;
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a Service resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated Service is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args?: ServiceArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Service is deprecated: Service is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dnsConfig"] = args ? args.dnsConfig : undefined;
            resourceInputs["healthCheckConfig"] = args ? args.healthCheckConfig : undefined;
            resourceInputs["healthCheckCustomConfig"] = args ? args.healthCheckCustomConfig : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namespaceId"] = args ? args.namespaceId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["arn"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["dnsConfig"] = undefined /*out*/;
            resourceInputs["healthCheckConfig"] = undefined /*out*/;
            resourceInputs["healthCheckCustomConfig"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["namespaceId"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Service.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Service resource.
 */
export interface ServiceArgs {
    description?: pulumi.Input<string>;
    dnsConfig?: pulumi.Input<inputs.servicediscovery.ServiceDnsConfigArgs>;
    healthCheckConfig?: pulumi.Input<inputs.servicediscovery.ServiceHealthCheckConfigArgs>;
    healthCheckCustomConfig?: pulumi.Input<inputs.servicediscovery.ServiceHealthCheckCustomConfigArgs>;
    name?: pulumi.Input<string>;
    namespaceId?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.servicediscovery.ServiceTagArgs>[]>;
    type?: pulumi.Input<string>;
}
