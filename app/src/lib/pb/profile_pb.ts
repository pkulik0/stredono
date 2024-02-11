// @generated by protoc-gen-es v1.7.2 with parameter "target=ts"
// @generated from file pb/profile.proto (package stredono, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message stredono.Profile
 */
export class Profile extends Message<Profile> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  /**
   * @generated from field: string uid = 2;
   */
  uid = "";

  /**
   * @generated from field: string url = 3;
   */
  url = "";

  /**
   * @generated from field: string description = 4;
   */
  description = "";

  /**
   * @generated from field: string avatarUrl = 5;
   */
  avatarUrl = "";

  /**
   * @generated from field: float minimumAmount = 6;
   */
  minimumAmount = 0;

  constructor(data?: PartialMessage<Profile>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "stredono.Profile";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "uid", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "avatarUrl", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "minimumAmount", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Profile {
    return new Profile().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Profile {
    return new Profile().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Profile {
    return new Profile().fromJsonString(jsonString, options);
  }

  static equals(a: Profile | PlainMessage<Profile> | undefined, b: Profile | PlainMessage<Profile> | undefined): boolean {
    return proto3.util.equals(Profile, a, b);
  }
}

