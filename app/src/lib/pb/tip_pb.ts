// @generated by protoc-gen-es v1.7.2 with parameter "target=ts"
// @generated from file pb/tip.proto (package stredono, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";
import { Currency } from "./enums_pb.js";

/**
 * @generated from enum stredono.TipStatus
 */
export enum TipStatus {
  /**
   * @generated from enum value: INITIATED = 0;
   */
  INITIATED = 0,

  /**
   * @generated from enum value: PAYMENT_PENDING = 1;
   */
  PAYMENT_PENDING = 1,

  /**
   * @generated from enum value: PAYMENT_SUCCESS = 2;
   */
  PAYMENT_SUCCESS = 2,
}
// Retrieve enum metadata with: proto3.getEnumType(TipStatus)
proto3.util.setEnumType(TipStatus, "stredono.TipStatus", [
  { no: 0, name: "INITIATED" },
  { no: 1, name: "PAYMENT_PENDING" },
  { no: 2, name: "PAYMENT_SUCCESS" },
]);

/**
 * @generated from message stredono.Tip
 */
export class Tip extends Message<Tip> {
  /**
   * @generated from field: string SenderId = 1;
   */
  SenderId = "";

  /**
   * @generated from field: string DisplayName = 2;
   */
  DisplayName = "";

  /**
   * @generated from field: string Email = 3;
   */
  Email = "";

  /**
   * @generated from field: string Message = 4;
   */
  Message = "";

  /**
   * @generated from field: double Amount = 5;
   */
  Amount = 0;

  /**
   * @generated from field: stredono.Currency Currency = 6;
   */
  Currency = Currency.UNKNOWN;

  /**
   * @generated from field: string RecipientId = 7;
   */
  RecipientId = "";

  /**
   * @generated from field: stredono.TipStatus Status = 8;
   */
  Status = TipStatus.INITIATED;

  /**
   * @generated from field: int64 Timestamp = 9;
   */
  Timestamp = protoInt64.zero;

  constructor(data?: PartialMessage<Tip>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "stredono.Tip";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "SenderId", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "DisplayName", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "Email", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "Message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "Amount", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 6, name: "Currency", kind: "enum", T: proto3.getEnumType(Currency) },
    { no: 7, name: "RecipientId", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 8, name: "Status", kind: "enum", T: proto3.getEnumType(TipStatus) },
    { no: 9, name: "Timestamp", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Tip {
    return new Tip().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Tip {
    return new Tip().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Tip {
    return new Tip().fromJsonString(jsonString, options);
  }

  static equals(a: Tip | PlainMessage<Tip> | undefined, b: Tip | PlainMessage<Tip> | undefined): boolean {
    return proto3.util.equals(Tip, a, b);
  }
}

