// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: proto.fbe
// Version: 1.3.0.0

package com.chronoxor.proto.fbe;

// Fast Binary Encoding com.chronoxor.proto sender
public class Sender extends com.chronoxor.fbe.Sender
{
    // Sender models accessors
    public final OrderModel OrderModel;
    public final BalanceModel BalanceModel;
    public final AccountModel AccountModel;

    public Sender()
    {
        super(false);
        OrderModel = new OrderModel(getBuffer());
        BalanceModel = new BalanceModel(getBuffer());
        AccountModel = new AccountModel(getBuffer());
    }
    public Sender(com.chronoxor.fbe.Buffer buffer)
    {
        super(buffer, false);
        OrderModel = new OrderModel(getBuffer());
        BalanceModel = new BalanceModel(getBuffer());
        AccountModel = new AccountModel(getBuffer());
    }

    public long send(Object obj)
    {
        if (obj instanceof com.chronoxor.proto.Order)
        {
            com.chronoxor.proto.Order value = (com.chronoxor.proto.Order)obj;
            return send(value);
        }
        if (obj instanceof com.chronoxor.proto.Balance)
        {
            com.chronoxor.proto.Balance value = (com.chronoxor.proto.Balance)obj;
            return send(value);
        }
        if (obj instanceof com.chronoxor.proto.Account)
        {
            com.chronoxor.proto.Account value = (com.chronoxor.proto.Account)obj;
            return send(value);
        }


        return 0;
    }

    public long send(com.chronoxor.proto.Order value)
    {
        // Serialize the value into the FBE stream
        long serialized = OrderModel.serialize(value);
        assert (serialized > 0) : "com.chronoxor.proto.Order serialization failed!";
        assert OrderModel.verify() : "com.chronoxor.proto.Order validation failed!";

        // Log the value
        if (getLogging())
        {
            String message = value.toString();
            onSendLog(message);
        }

        // Send the serialized value
        return sendSerialized(serialized);
    }
    public long send(com.chronoxor.proto.Balance value)
    {
        // Serialize the value into the FBE stream
        long serialized = BalanceModel.serialize(value);
        assert (serialized > 0) : "com.chronoxor.proto.Balance serialization failed!";
        assert BalanceModel.verify() : "com.chronoxor.proto.Balance validation failed!";

        // Log the value
        if (getLogging())
        {
            String message = value.toString();
            onSendLog(message);
        }

        // Send the serialized value
        return sendSerialized(serialized);
    }
    public long send(com.chronoxor.proto.Account value)
    {
        // Serialize the value into the FBE stream
        long serialized = AccountModel.serialize(value);
        assert (serialized > 0) : "com.chronoxor.proto.Account serialization failed!";
        assert AccountModel.verify() : "com.chronoxor.proto.Account validation failed!";

        // Log the value
        if (getLogging())
        {
            String message = value.toString();
            onSendLog(message);
        }

        // Send the serialized value
        return sendSerialized(serialized);
    }

    // Send message handler
    @Override
    protected long onSend(byte[] buffer, long offset, long size) { throw new UnsupportedOperationException("com.chronoxor.proto.fbe.Sender.onSend() not implemented!"); }
}
