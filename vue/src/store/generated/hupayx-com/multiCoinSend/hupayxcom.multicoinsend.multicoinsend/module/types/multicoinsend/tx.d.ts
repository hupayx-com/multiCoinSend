import { Reader, Writer } from "protobufjs/minimal";
import { Coin } from "../cosmos/base/v1beta1/coin";
export declare const protobufPackage = "hupayxcom.multicoinsend.multicoinsend";
export interface MsgCoinSend {
    creator: string;
    receivers: Receiver[];
}
export interface Receiver {
    to: string;
    amount: Coin[];
}
export interface MsgCoinSendResponse {
    cntValue: string;
    totalAmount: string;
}
export declare const MsgCoinSend: {
    encode(message: MsgCoinSend, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCoinSend;
    fromJSON(object: any): MsgCoinSend;
    toJSON(message: MsgCoinSend): unknown;
    fromPartial(object: DeepPartial<MsgCoinSend>): MsgCoinSend;
};
export declare const Receiver: {
    encode(message: Receiver, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Receiver;
    fromJSON(object: any): Receiver;
    toJSON(message: Receiver): unknown;
    fromPartial(object: DeepPartial<Receiver>): Receiver;
};
export declare const MsgCoinSendResponse: {
    encode(message: MsgCoinSendResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCoinSendResponse;
    fromJSON(object: any): MsgCoinSendResponse;
    toJSON(message: MsgCoinSendResponse): unknown;
    fromPartial(object: DeepPartial<MsgCoinSendResponse>): MsgCoinSendResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    CoinSend(request: MsgCoinSend): Promise<MsgCoinSendResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CoinSend(request: MsgCoinSend): Promise<MsgCoinSendResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
