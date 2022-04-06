/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Coin } from "../cosmos/base/v1beta1/coin";
export const protobufPackage = "hupayxcom.multicoinsend.multicoinsend";
const baseMsgCoinSend = { creator: "" };
export const MsgCoinSend = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        for (const v of message.receivers) {
            Receiver.encode(v, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCoinSend };
        message.receivers = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.receivers.push(Receiver.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgCoinSend };
        message.receivers = [];
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.receivers !== undefined && object.receivers !== null) {
            for (const e of object.receivers) {
                message.receivers.push(Receiver.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        if (message.receivers) {
            obj.receivers = message.receivers.map((e) => e ? Receiver.toJSON(e) : undefined);
        }
        else {
            obj.receivers = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgCoinSend };
        message.receivers = [];
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.receivers !== undefined && object.receivers !== null) {
            for (const e of object.receivers) {
                message.receivers.push(Receiver.fromPartial(e));
            }
        }
        return message;
    },
};
const baseReceiver = { to: "" };
export const Receiver = {
    encode(message, writer = Writer.create()) {
        if (message.to !== "") {
            writer.uint32(10).string(message.to);
        }
        for (const v of message.amount) {
            Coin.encode(v, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseReceiver };
        message.amount = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.to = reader.string();
                    break;
                case 2:
                    message.amount.push(Coin.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseReceiver };
        message.amount = [];
        if (object.to !== undefined && object.to !== null) {
            message.to = String(object.to);
        }
        else {
            message.to = "";
        }
        if (object.amount !== undefined && object.amount !== null) {
            for (const e of object.amount) {
                message.amount.push(Coin.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.to !== undefined && (obj.to = message.to);
        if (message.amount) {
            obj.amount = message.amount.map((e) => (e ? Coin.toJSON(e) : undefined));
        }
        else {
            obj.amount = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseReceiver };
        message.amount = [];
        if (object.to !== undefined && object.to !== null) {
            message.to = object.to;
        }
        else {
            message.to = "";
        }
        if (object.amount !== undefined && object.amount !== null) {
            for (const e of object.amount) {
                message.amount.push(Coin.fromPartial(e));
            }
        }
        return message;
    },
};
const baseMsgCoinSendResponse = { cntValue: "", totalAmount: "" };
export const MsgCoinSendResponse = {
    encode(message, writer = Writer.create()) {
        if (message.cntValue !== "") {
            writer.uint32(10).string(message.cntValue);
        }
        if (message.totalAmount !== "") {
            writer.uint32(18).string(message.totalAmount);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCoinSendResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.cntValue = reader.string();
                    break;
                case 2:
                    message.totalAmount = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgCoinSendResponse };
        if (object.cntValue !== undefined && object.cntValue !== null) {
            message.cntValue = String(object.cntValue);
        }
        else {
            message.cntValue = "";
        }
        if (object.totalAmount !== undefined && object.totalAmount !== null) {
            message.totalAmount = String(object.totalAmount);
        }
        else {
            message.totalAmount = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.cntValue !== undefined && (obj.cntValue = message.cntValue);
        message.totalAmount !== undefined &&
            (obj.totalAmount = message.totalAmount);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgCoinSendResponse };
        if (object.cntValue !== undefined && object.cntValue !== null) {
            message.cntValue = object.cntValue;
        }
        else {
            message.cntValue = "";
        }
        if (object.totalAmount !== undefined && object.totalAmount !== null) {
            message.totalAmount = object.totalAmount;
        }
        else {
            message.totalAmount = "";
        }
        return message;
    },
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    CoinSend(request) {
        const data = MsgCoinSend.encode(request).finish();
        const promise = this.rpc.request("hupayxcom.multicoinsend.multicoinsend.Msg", "CoinSend", data);
        return promise.then((data) => MsgCoinSendResponse.decode(new Reader(data)));
    }
}
