// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::MediaLive::Input
 *
 * @deprecated Input is not yet supported by AWS Cloud Control API, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class Input extends pulumi.CustomResource {
    /**
     * Get an existing Input resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Input {
        pulumi.log.warn("Input is deprecated: Input is not yet supported by AWS Cloud Control API, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new Input(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:medialive:Input';

    /**
     * Returns true if the given object is an instance of Input.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Input {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Input.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly destinations!: pulumi.Output<outputs.medialive.InputInputDestinationRequest[] | undefined>;
    public readonly inputDevices!: pulumi.Output<outputs.medialive.InputInputDeviceSettings[] | undefined>;
    public readonly inputSecurityGroups!: pulumi.Output<string[] | undefined>;
    public readonly mediaConnectFlows!: pulumi.Output<outputs.medialive.InputMediaConnectFlowRequest[] | undefined>;
    public readonly name!: pulumi.Output<string | undefined>;
    public readonly roleArn!: pulumi.Output<string | undefined>;
    public readonly sources!: pulumi.Output<outputs.medialive.InputInputSourceRequest[] | undefined>;
    public readonly tags!: pulumi.Output<any | undefined>;
    public readonly type!: pulumi.Output<string | undefined>;
    public readonly vpc!: pulumi.Output<outputs.medialive.InputInputVpcRequest | undefined>;

    /**
     * Create a Input resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated Input is not yet supported by AWS Cloud Control API, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args?: InputArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Input is deprecated: Input is not yet supported by AWS Cloud Control API, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            inputs["destinations"] = args ? args.destinations : undefined;
            inputs["inputDevices"] = args ? args.inputDevices : undefined;
            inputs["inputSecurityGroups"] = args ? args.inputSecurityGroups : undefined;
            inputs["mediaConnectFlows"] = args ? args.mediaConnectFlows : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["roleArn"] = args ? args.roleArn : undefined;
            inputs["sources"] = args ? args.sources : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["vpc"] = args ? args.vpc : undefined;
            inputs["arn"] = undefined /*out*/;
        } else {
            inputs["arn"] = undefined /*out*/;
            inputs["destinations"] = undefined /*out*/;
            inputs["inputDevices"] = undefined /*out*/;
            inputs["inputSecurityGroups"] = undefined /*out*/;
            inputs["mediaConnectFlows"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["roleArn"] = undefined /*out*/;
            inputs["sources"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["vpc"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Input.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Input resource.
 */
export interface InputArgs {
    destinations?: pulumi.Input<pulumi.Input<inputs.medialive.InputInputDestinationRequestArgs>[]>;
    inputDevices?: pulumi.Input<pulumi.Input<inputs.medialive.InputInputDeviceSettingsArgs>[]>;
    inputSecurityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    mediaConnectFlows?: pulumi.Input<pulumi.Input<inputs.medialive.InputMediaConnectFlowRequestArgs>[]>;
    name?: pulumi.Input<string>;
    roleArn?: pulumi.Input<string>;
    sources?: pulumi.Input<pulumi.Input<inputs.medialive.InputInputSourceRequestArgs>[]>;
    tags?: any;
    type?: pulumi.Input<string>;
    vpc?: pulumi.Input<inputs.medialive.InputInputVpcRequestArgs>;
}