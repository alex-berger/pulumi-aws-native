// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Lambda::Version
 *
 * @deprecated Version is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class Version extends pulumi.CustomResource {
    /**
     * Get an existing Version resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Version {
        pulumi.log.warn("Version is deprecated: Version is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new Version(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:lambda:Version';

    /**
     * Returns true if the given object is an instance of Version.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Version {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Version.__pulumiType;
    }

    public readonly codeSha256!: pulumi.Output<string | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly functionName!: pulumi.Output<string>;
    public readonly provisionedConcurrencyConfig!: pulumi.Output<outputs.lambda.VersionProvisionedConcurrencyConfiguration | undefined>;
    public /*out*/ readonly version!: pulumi.Output<string>;

    /**
     * Create a Version resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated Version is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: VersionArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Version is deprecated: Version is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.functionName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'functionName'");
            }
            resourceInputs["codeSha256"] = args ? args.codeSha256 : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["functionName"] = args ? args.functionName : undefined;
            resourceInputs["provisionedConcurrencyConfig"] = args ? args.provisionedConcurrencyConfig : undefined;
            resourceInputs["version"] = undefined /*out*/;
        } else {
            resourceInputs["codeSha256"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["functionName"] = undefined /*out*/;
            resourceInputs["provisionedConcurrencyConfig"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Version.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Version resource.
 */
export interface VersionArgs {
    codeSha256?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    functionName: pulumi.Input<string>;
    provisionedConcurrencyConfig?: pulumi.Input<inputs.lambda.VersionProvisionedConcurrencyConfigurationArgs>;
}