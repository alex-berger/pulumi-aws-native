// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ApiGateway::RequestValidator
 */
export class RequestValidator extends pulumi.CustomResource {
    /**
     * Get an existing RequestValidator resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RequestValidator {
        return new RequestValidator(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:apigateway:RequestValidator';

    /**
     * Returns true if the given object is an instance of RequestValidator.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RequestValidator {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RequestValidator.__pulumiType;
    }

    /**
     * Name of the request validator.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * ID of the request validator.
     */
    public /*out*/ readonly requestValidatorId!: pulumi.Output<string>;
    /**
     * The identifier of the targeted API entity.
     */
    public readonly restApiId!: pulumi.Output<string>;
    /**
     * Indicates whether to validate the request body according to the configured schema for the targeted API and method. 
     */
    public readonly validateRequestBody!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates whether to validate request parameters.
     */
    public readonly validateRequestParameters!: pulumi.Output<boolean | undefined>;

    /**
     * Create a RequestValidator resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RequestValidatorArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.restApiId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'restApiId'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["restApiId"] = args ? args.restApiId : undefined;
            resourceInputs["validateRequestBody"] = args ? args.validateRequestBody : undefined;
            resourceInputs["validateRequestParameters"] = args ? args.validateRequestParameters : undefined;
            resourceInputs["requestValidatorId"] = undefined /*out*/;
        } else {
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["requestValidatorId"] = undefined /*out*/;
            resourceInputs["restApiId"] = undefined /*out*/;
            resourceInputs["validateRequestBody"] = undefined /*out*/;
            resourceInputs["validateRequestParameters"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RequestValidator.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a RequestValidator resource.
 */
export interface RequestValidatorArgs {
    /**
     * Name of the request validator.
     */
    name?: pulumi.Input<string>;
    /**
     * The identifier of the targeted API entity.
     */
    restApiId: pulumi.Input<string>;
    /**
     * Indicates whether to validate the request body according to the configured schema for the targeted API and method. 
     */
    validateRequestBody?: pulumi.Input<boolean>;
    /**
     * Indicates whether to validate request parameters.
     */
    validateRequestParameters?: pulumi.Input<boolean>;
}
