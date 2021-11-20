// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Lambda::Url
 */
export class Url extends pulumi.CustomResource {
    /**
     * Get an existing Url resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Url {
        return new Url(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:lambda:Url';

    /**
     * Returns true if the given object is an instance of Url.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Url {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Url.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the Function URL.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Can be either AWS_IAM if the requests are authorized via IAM, or NONE if no authorization is configured on the Function URL.
     */
    public readonly authorizationType!: pulumi.Output<enums.lambda.UrlAuthorizationType>;
    public readonly cors!: pulumi.Output<outputs.lambda.UrlCors | undefined>;
    /**
     * The generated url for this resource.
     */
    public /*out*/ readonly functionUrl!: pulumi.Output<string>;
    /**
     * The alias qualifier for the target function. If TargetFunctionArn is unqualified then Qualifier must be passed.
     */
    public readonly qualifier!: pulumi.Output<string | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the function associated with the Function URL.
     */
    public readonly targetFunctionArn!: pulumi.Output<string>;

    /**
     * Create a Url resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UrlArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.authorizationType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authorizationType'");
            }
            if ((!args || args.targetFunctionArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetFunctionArn'");
            }
            inputs["authorizationType"] = args ? args.authorizationType : undefined;
            inputs["cors"] = args ? args.cors : undefined;
            inputs["qualifier"] = args ? args.qualifier : undefined;
            inputs["targetFunctionArn"] = args ? args.targetFunctionArn : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["functionUrl"] = undefined /*out*/;
        } else {
            inputs["arn"] = undefined /*out*/;
            inputs["authorizationType"] = undefined /*out*/;
            inputs["cors"] = undefined /*out*/;
            inputs["functionUrl"] = undefined /*out*/;
            inputs["qualifier"] = undefined /*out*/;
            inputs["targetFunctionArn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Url.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Url resource.
 */
export interface UrlArgs {
    /**
     * Can be either AWS_IAM if the requests are authorized via IAM, or NONE if no authorization is configured on the Function URL.
     */
    authorizationType: pulumi.Input<enums.lambda.UrlAuthorizationType>;
    cors?: pulumi.Input<inputs.lambda.UrlCorsArgs>;
    /**
     * The alias qualifier for the target function. If TargetFunctionArn is unqualified then Qualifier must be passed.
     */
    qualifier?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the function associated with the Function URL.
     */
    targetFunctionArn: pulumi.Input<string>;
}
