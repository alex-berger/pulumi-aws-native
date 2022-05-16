// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource-specific logging allows you to specify a logging level for a specific thing group.
 */
export class ResourceSpecificLogging extends pulumi.CustomResource {
    /**
     * Get an existing ResourceSpecificLogging resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ResourceSpecificLogging {
        return new ResourceSpecificLogging(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:iot:ResourceSpecificLogging';

    /**
     * Returns true if the given object is an instance of ResourceSpecificLogging.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ResourceSpecificLogging {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ResourceSpecificLogging.__pulumiType;
    }

    /**
     * The log level for a specific target. Valid values are: ERROR, WARN, INFO, DEBUG, or DISABLED.
     */
    public readonly logLevel!: pulumi.Output<enums.iot.ResourceSpecificLoggingLogLevel>;
    /**
     * Unique Id for a Target (TargetType:TargetName), this will be internally built to serve as primary identifier for a log target.
     */
    public /*out*/ readonly targetId!: pulumi.Output<string>;
    /**
     * The target name.
     */
    public readonly targetName!: pulumi.Output<string>;
    /**
     * The target type. Value must be THING_GROUP, CLIENT_ID, SOURCE_IP, PRINCIPAL_ID.
     */
    public readonly targetType!: pulumi.Output<enums.iot.ResourceSpecificLoggingTargetType>;

    /**
     * Create a ResourceSpecificLogging resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ResourceSpecificLoggingArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.logLevel === undefined) && !opts.urn) {
                throw new Error("Missing required property 'logLevel'");
            }
            if ((!args || args.targetName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetName'");
            }
            if ((!args || args.targetType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetType'");
            }
            resourceInputs["logLevel"] = args ? args.logLevel : undefined;
            resourceInputs["targetName"] = args ? args.targetName : undefined;
            resourceInputs["targetType"] = args ? args.targetType : undefined;
            resourceInputs["targetId"] = undefined /*out*/;
        } else {
            resourceInputs["logLevel"] = undefined /*out*/;
            resourceInputs["targetId"] = undefined /*out*/;
            resourceInputs["targetName"] = undefined /*out*/;
            resourceInputs["targetType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ResourceSpecificLogging.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ResourceSpecificLogging resource.
 */
export interface ResourceSpecificLoggingArgs {
    /**
     * The log level for a specific target. Valid values are: ERROR, WARN, INFO, DEBUG, or DISABLED.
     */
    logLevel: pulumi.Input<enums.iot.ResourceSpecificLoggingLogLevel>;
    /**
     * The target name.
     */
    targetName: pulumi.Input<string>;
    /**
     * The target type. Value must be THING_GROUP, CLIENT_ID, SOURCE_IP, PRINCIPAL_ID.
     */
    targetType: pulumi.Input<enums.iot.ResourceSpecificLoggingTargetType>;
}
