// @generated by protoc-gen-es v1.7.2 with parameter "target=ts"
// @generated from file pb/twitch.proto (package stredono, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";

/**
 * @generated from message stredono.TokenEntry
 */
export class TokenEntry extends Message<TokenEntry> {
  /**
   * @generated from field: string Uid = 1;
   */
  Uid = "";

  /**
   * @generated from field: bytes EncryptedToken = 2;
   */
  EncryptedToken = new Uint8Array(0);

  constructor(data?: PartialMessage<TokenEntry>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "stredono.TokenEntry";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "Uid", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "EncryptedToken", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TokenEntry {
    return new TokenEntry().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TokenEntry {
    return new TokenEntry().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TokenEntry {
    return new TokenEntry().fromJsonString(jsonString, options);
  }

  static equals(a: TokenEntry | PlainMessage<TokenEntry> | undefined, b: TokenEntry | PlainMessage<TokenEntry> | undefined): boolean {
    return proto3.util.equals(TokenEntry, a, b);
  }
}

/**
 * @generated from message stredono.Token
 */
export class Token extends Message<Token> {
  /**
   * @generated from field: string AccessToken = 1;
   */
  AccessToken = "";

  /**
   * @generated from field: string RefreshToken = 2;
   */
  RefreshToken = "";

  constructor(data?: PartialMessage<Token>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "stredono.Token";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "AccessToken", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "RefreshToken", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Token {
    return new Token().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Token {
    return new Token().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Token {
    return new Token().fromJsonString(jsonString, options);
  }

  static equals(a: Token | PlainMessage<Token> | undefined, b: Token | PlainMessage<Token> | undefined): boolean {
    return proto3.util.equals(Token, a, b);
  }
}

/**
 * @generated from message stredono.TwitchReward
 */
export class TwitchReward extends Message<TwitchReward> {
  /**
   * @generated from field: string Id = 1;
   */
  Id = "";

  /**
   * @generated from field: string Name = 2;
   */
  Name = "";

  /**
   * @generated from field: int64 Cost = 3;
   */
  Cost = protoInt64.zero;

  /**
   * @generated from field: bool IsEnabled = 4;
   */
  IsEnabled = false;

  constructor(data?: PartialMessage<TwitchReward>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "stredono.TwitchReward";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "Id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "Name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "Cost", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 4, name: "IsEnabled", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TwitchReward {
    return new TwitchReward().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TwitchReward {
    return new TwitchReward().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TwitchReward {
    return new TwitchReward().fromJsonString(jsonString, options);
  }

  static equals(a: TwitchReward | PlainMessage<TwitchReward> | undefined, b: TwitchReward | PlainMessage<TwitchReward> | undefined): boolean {
    return proto3.util.equals(TwitchReward, a, b);
  }
}

/**
 * @generated from message stredono.TwitchUser
 */
export class TwitchUser extends Message<TwitchUser> {
  /**
   * @generated from field: string Id = 1;
   */
  Id = "";

  /**
   * @generated from field: string DisplayName = 2;
   */
  DisplayName = "";

  /**
   * @generated from field: string Login = 3;
   */
  Login = "";

  /**
   * @generated from field: string AvatarUrl = 4;
   */
  AvatarUrl = "";

  /**
   * @generated from field: string Description = 5;
   */
  Description = "";

  /**
   * @generated from field: int64 CreationTimestamp = 6;
   */
  CreationTimestamp = protoInt64.zero;

  constructor(data?: PartialMessage<TwitchUser>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "stredono.TwitchUser";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "Id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "DisplayName", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "Login", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "AvatarUrl", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "Description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "CreationTimestamp", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TwitchUser {
    return new TwitchUser().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TwitchUser {
    return new TwitchUser().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TwitchUser {
    return new TwitchUser().fromJsonString(jsonString, options);
  }

  static equals(a: TwitchUser | PlainMessage<TwitchUser> | undefined, b: TwitchUser | PlainMessage<TwitchUser> | undefined): boolean {
    return proto3.util.equals(TwitchUser, a, b);
  }
}

