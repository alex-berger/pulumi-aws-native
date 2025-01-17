// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::OpsWorks::Instance
 *
 * @deprecated Instance is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Instance {
        pulumi.log.warn("Instance is deprecated: Instance is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new Instance(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:opsworks:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    public readonly agentVersion!: pulumi.Output<string | undefined>;
    public readonly amiId!: pulumi.Output<string | undefined>;
    public readonly architecture!: pulumi.Output<string | undefined>;
    public readonly autoScalingType!: pulumi.Output<string | undefined>;
    public readonly availabilityZone!: pulumi.Output<string | undefined>;
    public readonly blockDeviceMappings!: pulumi.Output<outputs.opsworks.InstanceBlockDeviceMapping[] | undefined>;
    public readonly ebsOptimized!: pulumi.Output<boolean | undefined>;
    public readonly elasticIps!: pulumi.Output<string[] | undefined>;
    public readonly hostname!: pulumi.Output<string | undefined>;
    public readonly installUpdatesOnBoot!: pulumi.Output<boolean | undefined>;
    public readonly instanceType!: pulumi.Output<string>;
    public readonly layerIds!: pulumi.Output<string[]>;
    public readonly os!: pulumi.Output<string | undefined>;
    public /*out*/ readonly privateDnsName!: pulumi.Output<string>;
    public /*out*/ readonly privateIp!: pulumi.Output<string>;
    public /*out*/ readonly publicDnsName!: pulumi.Output<string>;
    public /*out*/ readonly publicIp!: pulumi.Output<string>;
    public readonly rootDeviceType!: pulumi.Output<string | undefined>;
    public readonly sshKeyName!: pulumi.Output<string | undefined>;
    public readonly stackId!: pulumi.Output<string>;
    public readonly subnetId!: pulumi.Output<string | undefined>;
    public readonly tenancy!: pulumi.Output<string | undefined>;
    public readonly timeBasedAutoScaling!: pulumi.Output<outputs.opsworks.InstanceTimeBasedAutoScaling | undefined>;
    public readonly virtualizationType!: pulumi.Output<string | undefined>;
    public readonly volumes!: pulumi.Output<string[] | undefined>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated Instance is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Instance is deprecated: Instance is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.instanceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceType'");
            }
            if ((!args || args.layerIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'layerIds'");
            }
            if ((!args || args.stackId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'stackId'");
            }
            resourceInputs["agentVersion"] = args ? args.agentVersion : undefined;
            resourceInputs["amiId"] = args ? args.amiId : undefined;
            resourceInputs["architecture"] = args ? args.architecture : undefined;
            resourceInputs["autoScalingType"] = args ? args.autoScalingType : undefined;
            resourceInputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            resourceInputs["blockDeviceMappings"] = args ? args.blockDeviceMappings : undefined;
            resourceInputs["ebsOptimized"] = args ? args.ebsOptimized : undefined;
            resourceInputs["elasticIps"] = args ? args.elasticIps : undefined;
            resourceInputs["hostname"] = args ? args.hostname : undefined;
            resourceInputs["installUpdatesOnBoot"] = args ? args.installUpdatesOnBoot : undefined;
            resourceInputs["instanceType"] = args ? args.instanceType : undefined;
            resourceInputs["layerIds"] = args ? args.layerIds : undefined;
            resourceInputs["os"] = args ? args.os : undefined;
            resourceInputs["rootDeviceType"] = args ? args.rootDeviceType : undefined;
            resourceInputs["sshKeyName"] = args ? args.sshKeyName : undefined;
            resourceInputs["stackId"] = args ? args.stackId : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["tenancy"] = args ? args.tenancy : undefined;
            resourceInputs["timeBasedAutoScaling"] = args ? args.timeBasedAutoScaling : undefined;
            resourceInputs["virtualizationType"] = args ? args.virtualizationType : undefined;
            resourceInputs["volumes"] = args ? args.volumes : undefined;
            resourceInputs["privateDnsName"] = undefined /*out*/;
            resourceInputs["privateIp"] = undefined /*out*/;
            resourceInputs["publicDnsName"] = undefined /*out*/;
            resourceInputs["publicIp"] = undefined /*out*/;
        } else {
            resourceInputs["agentVersion"] = undefined /*out*/;
            resourceInputs["amiId"] = undefined /*out*/;
            resourceInputs["architecture"] = undefined /*out*/;
            resourceInputs["autoScalingType"] = undefined /*out*/;
            resourceInputs["availabilityZone"] = undefined /*out*/;
            resourceInputs["blockDeviceMappings"] = undefined /*out*/;
            resourceInputs["ebsOptimized"] = undefined /*out*/;
            resourceInputs["elasticIps"] = undefined /*out*/;
            resourceInputs["hostname"] = undefined /*out*/;
            resourceInputs["installUpdatesOnBoot"] = undefined /*out*/;
            resourceInputs["instanceType"] = undefined /*out*/;
            resourceInputs["layerIds"] = undefined /*out*/;
            resourceInputs["os"] = undefined /*out*/;
            resourceInputs["privateDnsName"] = undefined /*out*/;
            resourceInputs["privateIp"] = undefined /*out*/;
            resourceInputs["publicDnsName"] = undefined /*out*/;
            resourceInputs["publicIp"] = undefined /*out*/;
            resourceInputs["rootDeviceType"] = undefined /*out*/;
            resourceInputs["sshKeyName"] = undefined /*out*/;
            resourceInputs["stackId"] = undefined /*out*/;
            resourceInputs["subnetId"] = undefined /*out*/;
            resourceInputs["tenancy"] = undefined /*out*/;
            resourceInputs["timeBasedAutoScaling"] = undefined /*out*/;
            resourceInputs["virtualizationType"] = undefined /*out*/;
            resourceInputs["volumes"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Instance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    agentVersion?: pulumi.Input<string>;
    amiId?: pulumi.Input<string>;
    architecture?: pulumi.Input<string>;
    autoScalingType?: pulumi.Input<string>;
    availabilityZone?: pulumi.Input<string>;
    blockDeviceMappings?: pulumi.Input<pulumi.Input<inputs.opsworks.InstanceBlockDeviceMappingArgs>[]>;
    ebsOptimized?: pulumi.Input<boolean>;
    elasticIps?: pulumi.Input<pulumi.Input<string>[]>;
    hostname?: pulumi.Input<string>;
    installUpdatesOnBoot?: pulumi.Input<boolean>;
    instanceType: pulumi.Input<string>;
    layerIds: pulumi.Input<pulumi.Input<string>[]>;
    os?: pulumi.Input<string>;
    rootDeviceType?: pulumi.Input<string>;
    sshKeyName?: pulumi.Input<string>;
    stackId: pulumi.Input<string>;
    subnetId?: pulumi.Input<string>;
    tenancy?: pulumi.Input<string>;
    timeBasedAutoScaling?: pulumi.Input<inputs.opsworks.InstanceTimeBasedAutoScalingArgs>;
    virtualizationType?: pulumi.Input<string>;
    volumes?: pulumi.Input<pulumi.Input<string>[]>;
}
