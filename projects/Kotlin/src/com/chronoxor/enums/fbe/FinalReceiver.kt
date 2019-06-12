// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// Version: 1.3.0.0

@file:Suppress("UnusedImport", "unused")

package com.chronoxor.enums.fbe

// Fast Binary Encoding com.chronoxor.enums final receiver
@Suppress("MemberVisibilityCanBePrivate", "PrivatePropertyName", "UNUSED_PARAMETER")
open class FinalReceiver : com.chronoxor.fbe.Receiver, FinalReceiverListener
{
    // Receiver values accessors
    private val EnumsValue: com.chronoxor.enums.Enums

    // Receiver models accessors
    private val EnumsModel: EnumsFinalModel

    constructor() : super(true)
    {
        EnumsValue = com.chronoxor.enums.Enums()
        EnumsModel = EnumsFinalModel()
    }

    constructor(buffer: com.chronoxor.fbe.Buffer) : super(buffer, true)
    {
        EnumsValue = com.chronoxor.enums.Enums()
        EnumsModel = EnumsFinalModel()
    }

    override fun onReceive(type: Long, buffer: ByteArray, offset: Long, size: Long): Boolean
    {
        return onReceiveListener(this, type, buffer, offset, size)
    }

    open fun onReceiveListener(listener: FinalReceiverListener, type: Long, buffer: ByteArray, offset: Long, size: Long): Boolean
    {
        when (type)
        {
            com.chronoxor.enums.fbe.EnumsFinalModel.fbeTypeConst ->
            {
                // Deserialize the value from the FBE stream
                EnumsModel.attach(buffer, offset)
                assert(EnumsModel.verify()) { "com.chronoxor.enums.Enums validation failed!" }
                val deserialized = EnumsModel.deserialize(EnumsValue)
                assert(deserialized > 0) { "com.chronoxor.enums.Enums deserialization failed!" }

                // Log the value
                if (logging)
                {
                    val message = EnumsValue.toString()
                    onReceiveLog(message)
                }

                // Call receive handler with deserialized value
                listener.onReceive(EnumsValue)
                return true
            }
        }

        return false
    }
}
