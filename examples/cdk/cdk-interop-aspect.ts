import * as cdk from 'aws-cdk-lib/lib/core';
import * as pulumi from "@pulumi/pulumi";
import {CdkResource, normalize, firstToLower} from "./cdk-interop";
import { Stack, CfnElement, IResolvable, Aspects, Token } from "aws-cdk-lib";
import { Construct, ConstructOrder } from "constructs";

export class CdkStackComponent extends pulumi.ComponentResource {
  constructor(name: string, args: (scope: Construct) => cdk.Stack, opts?: pulumi.CustomResourceOptions) {
    super("aws-native:cdk:StackComponent", name, args, opts);
    const app = new cdk.App();
    const stack = args(app);
    app.synth();
  }
}


export class AwsPulumiAdapter extends Stack {
    constructor(scope: Construct, id: string) {
      super(undefined, id);
  
      const host = new PulumiCDKBridge(scope, id, this);
  
      Aspects.of(scope).add({
        visit: (node) => {
          if (node === scope) {
            host.convert();
          }
        },
      });
    }
  }

class PulumiCDKBridge extends Construct {
    constructor(
      scope: Construct,
      id: string,
      private readonly host: AwsPulumiAdapter
    ) {
      super(scope, id);
      
    }
  
    convert() {
      for (const r of this.host.node.findAll()) {
        if (r instanceof CfnElement) {
          const cfn = this.host.resolve(
            (r as any)._toCloudFormation()
          ) as CloudFormationTemplate;
          for (const [logical, value] of Object.entries(cfn.Resources || {})) {
            const typeName = value.Type;
            let props = this.processIntrinsics(value.Properties);
            props = normalize(props);
            new CdkResource(logical, typeName, props);
          }
          for (const [conditionId, condition] of Object.entries(
            cfn.Conditions || {}
          )) {
            // Do something with the condition
          }
          for (const [outputId, args] of Object.entries(cfn.Outputs || {})) {
            // Do something with the outputs
          }
        }
      }
    }

  private processIntrinsics(obj: any): any {
    if (typeof obj === "string") {
      if (Token.isUnresolved(obj)) {
        console.debug(`Unresolved: ${JSON.stringify(obj)}`);
        return this.host.resolve(obj);
      }
      return obj;
    }

    if (typeof obj !== "object") {
      return obj;
    }

    if (Array.isArray(obj)) {
      return obj.map((x) => this.processIntrinsics(x));
    }

    const ref = obj.Ref;
    if (ref) {
      return this.resolveRef(ref);
    }

    const intrinsic = Object.keys(obj)[0];
    if (intrinsic?.startsWith("Fn::") && Object.keys(obj).length === 1) {
      return this.resolveIntrinsic(intrinsic, obj[intrinsic]);
    } else if (intrinsic?.startsWith("Fn:") && !intrinsic?.startsWith("Fn::")) {
      console.warn(
        'Found possible intrinsic function starting with "Fn:" instead of "Fn::". Typo?'
      );
    }

    const result: any = {};
    for (const [k, v] of Object.entries(obj)) {
      result[k] = this.processIntrinsics(v);
    }

    return result;
  }

  private resolveIntrinsic(fn: string, params: any) {
    switch (fn) {
      case "Fn::GetAtt": {
        console.debug(`Fn::GetAtt(${params[0]}, ${firstToLower(params[1])})`);
        return this.resolveAtt(params[0], firstToLower(params[1]));
      }

      case "Fn::Join": {
        const [delim, strings] = params;
        return this.processIntrinsics(delim) + this.processIntrinsics(strings);
      }

      case "Fn::Transform": {
        // https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-macros.html
        throw new Error(
          "Fn::Transform is not supported – Cfn Template Macros are not supported yet"
        );
      }

      case "Fn::ImportValue": {
        // TODO: support cross cfn stack references?
        // This is related to the Export Name from outputs https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/outputs-section-structure.html
        // We might revisit this once the CDKTF supports cross stack references
        throw new Error(`Fn::ImportValue is not yet supported.`);
      }

      default:
        throw new Error(
          `unsupported intrinsic function ${fn} (params: ${JSON.stringify(
            params
          )})`
        );
    }
  }

  private resolveRef(ref: string) {
    if (ref?.startsWith("AWS::")) {
      throw new Error(`don't support pseudo parameters ${ref}`);
    }
  }

  private resolveAtt(logicalId: string, attribute: string) {
    const child = this.host.node.tryFindChild(logicalId);
    if (!(child instanceof CfnElement)) {
      for (const c of this.host.node.findAll(ConstructOrder.POSTORDER)) {
        console.debug(`Id: ${c.node.id}`)
      }
      throw new Error(
        `unable to resolve a "Ref" to a resource with the logical ID ${logicalId}: ${JSON.stringify(child)}`
      );
    }

    for (const [id, value] of Object.entries(((child as CfnElement)._toCloudFormation() as CloudFormationTemplate).Resources || {})) {
      const resolvedId = this.host.resolve(id);
      if (!value.Properties) {
        console.debug(`No value for id: ${resolvedId}`)
        continue
      }
      const att = value.Properties[attribute];
      if (!att) {
        throw new Error(
          `no "${attribute}" attribute mapping for resource ${logicalId}`
        );
      }
    }

    console.debug(`resolveAtt for ${logicalId}`)
      throw new Error(
        `no "${attribute}" attribute mapping for resource ${logicalId}`
      );
  }
}

export interface CloudFormationResource {
  readonly Type: string;
  readonly Properties: any;
  readonly Condition?: string;
}

export interface CloudFormationTemplate {
  Resources?: { [id: string]: CloudFormationResource };
  Conditions?: { [id: string]: any };
  Outputs?: { [id: string]: any };
}
