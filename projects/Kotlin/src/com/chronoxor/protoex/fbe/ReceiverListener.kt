// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: protoex.fbe
// Version: 1.3.0.0

@file:Suppress("UnusedImport", "unused")

package com.chronoxor.protoex.fbe

// Fast Binary Encoding com.chronoxor.protoex receiver listener
interface ReceiverListener : com.chronoxor.proto.fbe.ReceiverListener
{
    fun onReceive(value: com.chronoxor.protoex.Order) {}
    fun onReceive(value: com.chronoxor.protoex.Balance) {}
    fun onReceive(value: com.chronoxor.protoex.Account) {}
    fun onReceive(value: com.chronoxor.protoex.OrderMessage) {}
    fun onReceive(value: com.chronoxor.protoex.BalanceMessage) {}
    fun onReceive(value: com.chronoxor.protoex.AccountMessage) {}
}