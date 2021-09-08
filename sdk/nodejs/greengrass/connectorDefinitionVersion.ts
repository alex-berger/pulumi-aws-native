// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-connectordefinitionversion.html
 */
export class ConnectorDefinitionVersion extends pulumi.CustomResource {
    /**
     * Get an existing ConnectorDefinitionVersion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ConnectorDefinitionVersion {
        return new ConnectorDefinitionVersion(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:greengrass:ConnectorDefinitionVersion';

    /**
     * Returns true if the given object is an instance of ConnectorDefinitionVersion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConnectorDefinitionVersion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConnectorDefinitionVersion.__pulumiType;
    }

    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-connectordefinitionversion.html#cfn-greengrass-connectordefinitionversion-connectordefinitionid
     */
    public readonly connectorDefinitionId!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-connectordefinitionversion.html#cfn-greengrass-connectordefinitionversion-connectors
     */
    public readonly connectors!: pulumi.Output<outputs.greengrass.ConnectorDefinitionVersionConnector[]>;

    /**
     * Create a ConnectorDefinitionVersion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectorDefinitionVersionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.connectorDefinitionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectorDefinitionId'");
            }
            if ((!args || args.connectors === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectors'");
            }
            inputs["connectorDefinitionId"] = args ? args.connectorDefinitionId : undefined;
            inputs["connectors"] = args ? args.connectors : undefined;
        } else {
            inputs["connectorDefinitionId"] = undefined /*out*/;
            inputs["connectors"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ConnectorDefinitionVersion.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ConnectorDefinitionVersion resource.
 */
export interface ConnectorDefinitionVersionArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-connectordefinitionversion.html#cfn-greengrass-connectordefinitionversion-connectordefinitionid
     */
    connectorDefinitionId: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-connectordefinitionversion.html#cfn-greengrass-connectordefinitionversion-connectors
     */
    connectors: pulumi.Input<pulumi.Input<inputs.greengrass.ConnectorDefinitionVersionConnectorArgs>[]>;
}