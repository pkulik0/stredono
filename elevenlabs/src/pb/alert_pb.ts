// @generated by protoc-gen-es v1.7.2 with parameter "target=ts"
// @generated from file pb/alert.proto (package stredono, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { EventType } from "./event_pb.js";

/**
 * @generated from enum stredono.AnimationType
 */
export enum AnimationType {
  /**
   * @generated from enum value: PULSE = 0;
   */
  PULSE = 0,

  /**
   * @generated from enum value: HEART_BEAT = 2;
   */
  HEART_BEAT = 2,

  /**
   * @generated from enum value: SHAKE_VERTICALLY = 4;
   */
  SHAKE_VERTICALLY = 4,

  /**
   * @generated from enum value: SHAKE_HORIZONTALLY = 5;
   */
  SHAKE_HORIZONTALLY = 5,

  /**
   * @generated from enum value: TADA = 8;
   */
  TADA = 8,

  /**
   * @generated from enum value: JELLO = 9;
   */
  JELLO = 9,

  /**
   * @generated from enum value: BOUNCE = 10;
   */
  BOUNCE = 10,
}
// Retrieve enum metadata with: proto3.getEnumType(AnimationType)
proto3.util.setEnumType(AnimationType, "stredono.AnimationType", [
  { no: 0, name: "PULSE" },
  { no: 2, name: "HEART_BEAT" },
  { no: 4, name: "SHAKE_VERTICALLY" },
  { no: 5, name: "SHAKE_HORIZONTALLY" },
  { no: 8, name: "TADA" },
  { no: 9, name: "JELLO" },
  { no: 10, name: "BOUNCE" },
]);

/**
 * @generated from enum stredono.Alignment
 */
export enum Alignment {
  /**
   * @generated from enum value: START = 0;
   */
  START = 0,

  /**
   * @generated from enum value: CENTER = 1;
   */
  CENTER = 1,

  /**
   * @generated from enum value: END = 2;
   */
  END = 2,

  /**
   * @generated from enum value: JUSTIFY = 3;
   */
  JUSTIFY = 3,
}
// Retrieve enum metadata with: proto3.getEnumType(Alignment)
proto3.util.setEnumType(Alignment, "stredono.Alignment", [
  { no: 0, name: "START" },
  { no: 1, name: "CENTER" },
  { no: 2, name: "END" },
  { no: 3, name: "JUSTIFY" },
]);

/**
 * @generated from enum stredono.Position
 */
export enum Position {
  /**
   * @generated from enum value: TOP = 0;
   */
  TOP = 0,

  /**
   * @generated from enum value: LEFT = 1;
   */
  LEFT = 1,

  /**
   * @generated from enum value: RIGHT = 2;
   */
  RIGHT = 2,

  /**
   * @generated from enum value: BOTTOM = 3;
   */
  BOTTOM = 3,
}
// Retrieve enum metadata with: proto3.getEnumType(Position)
proto3.util.setEnumType(Position, "stredono.Position", [
  { no: 0, name: "TOP" },
  { no: 1, name: "LEFT" },
  { no: 2, name: "RIGHT" },
  { no: 3, name: "BOTTOM" },
]);

/**
 * @generated from enum stredono.Speed
 */
export enum Speed {
  /**
   * @generated from enum value: OFF = 0;
   */
  OFF = 0,

  /**
   * @generated from enum value: SLOW = 1;
   */
  SLOW = 1,

  /**
   * @generated from enum value: MEDIUM = 2;
   */
  MEDIUM = 2,

  /**
   * @generated from enum value: FAST = 3;
   */
  FAST = 3,

  /**
   * @generated from enum value: FASTER = 4;
   */
  FASTER = 4,
}
// Retrieve enum metadata with: proto3.getEnumType(Speed)
proto3.util.setEnumType(Speed, "stredono.Speed", [
  { no: 0, name: "OFF" },
  { no: 1, name: "SLOW" },
  { no: 2, name: "MEDIUM" },
  { no: 3, name: "FAST" },
  { no: 4, name: "FASTER" },
]);

/**
 * @generated from message stredono.Alert
 */
export class Alert extends Message<Alert> {
  /**
   * @generated from field: string ID = 1;
   */
  ID = "";

  /**
   * @generated from field: stredono.EventType EventType = 2;
   */
  EventType = EventType.TIP;

  /**
   * @generated from field: double Min = 3;
   */
  Min = 0;

  /**
   * @generated from field: optional double Max = 4;
   */
  Max?: number;

  /**
   * @generated from field: string GifUrl = 5;
   */
  GifUrl = "";

  /**
   * @generated from field: string SoundUrl = 6;
   */
  SoundUrl = "";

  /**
   * @generated from field: stredono.AnimationType Animation = 7;
   */
  Animation = AnimationType.PULSE;

  /**
   * @generated from field: stredono.Speed AnimationSpeed = 8;
   */
  AnimationSpeed = Speed.OFF;

  /**
   * @generated from field: string TextColor = 9;
   */
  TextColor = "";

  /**
   * @generated from field: string AccentColor = 10;
   */
  AccentColor = "";

  /**
   * @generated from field: stredono.Alignment Alignment = 11;
   */
  Alignment = Alignment.START;

  /**
   * @generated from field: stredono.Position TextPosition = 12;
   */
  TextPosition = Position.TOP;

  constructor(data?: PartialMessage<Alert>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "stredono.Alert";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "ID", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "EventType", kind: "enum", T: proto3.getEnumType(EventType) },
    { no: 3, name: "Min", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 4, name: "Max", kind: "scalar", T: 1 /* ScalarType.DOUBLE */, opt: true },
    { no: 5, name: "GifUrl", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "SoundUrl", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "Animation", kind: "enum", T: proto3.getEnumType(AnimationType) },
    { no: 8, name: "AnimationSpeed", kind: "enum", T: proto3.getEnumType(Speed) },
    { no: 9, name: "TextColor", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 10, name: "AccentColor", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 11, name: "Alignment", kind: "enum", T: proto3.getEnumType(Alignment) },
    { no: 12, name: "TextPosition", kind: "enum", T: proto3.getEnumType(Position) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Alert {
    return new Alert().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Alert {
    return new Alert().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Alert {
    return new Alert().fromJsonString(jsonString, options);
  }

  static equals(a: Alert | PlainMessage<Alert> | undefined, b: Alert | PlainMessage<Alert> | undefined): boolean {
    return proto3.util.equals(Alert, a, b);
  }
}

/**
 * @generated from message stredono.UsersAlerts
 */
export class UsersAlerts extends Message<UsersAlerts> {
  /**
   * @generated from field: repeated stredono.Alert Alerts = 1;
   */
  Alerts: Alert[] = [];

  constructor(data?: PartialMessage<UsersAlerts>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "stredono.UsersAlerts";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "Alerts", kind: "message", T: Alert, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UsersAlerts {
    return new UsersAlerts().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UsersAlerts {
    return new UsersAlerts().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UsersAlerts {
    return new UsersAlerts().fromJsonString(jsonString, options);
  }

  static equals(a: UsersAlerts | PlainMessage<UsersAlerts> | undefined, b: UsersAlerts | PlainMessage<UsersAlerts> | undefined): boolean {
    return proto3.util.equals(UsersAlerts, a, b);
  }
}

