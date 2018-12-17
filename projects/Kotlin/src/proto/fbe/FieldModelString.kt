// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: fbe
// Version: 1.2.0.0

@file:Suppress("UnusedImport", "unused")

package fbe

import java.io.*
import java.lang.*
import java.lang.reflect.*
import java.math.*
import java.nio.charset.*
import java.time.*
import java.util.*

// Fast Binary Encoding string field model
class FieldModelString(buffer: Buffer, offset: Long) : FieldModel(buffer, offset)
{
    // Field size
    override val fbeSize: Long = 4

    // Field extra size
    override val fbeExtra: Long get()
    {
        if ((_buffer.offset + fbeOffset + fbeSize) > _buffer.size)
            return 0

        val fbeStringOffset = readUInt32(fbeOffset).toLong()
        if ((fbeStringOffset == 0L) || ((_buffer.offset + fbeStringOffset + 4) > _buffer.size))
            return 0

        val fbeStringSize = readUInt32(fbeStringOffset).toLong()
        return 4 + fbeStringSize
    }

    // Check if the string value is valid
    override fun verify(): Boolean
    {
        if ((_buffer.offset + fbeOffset + fbeSize) > _buffer.size)
            return true

        val fbeStringOffset = readUInt32(fbeOffset).toLong()
        if (fbeStringOffset == 0L)
            return true

        if ((_buffer.offset + fbeStringOffset + 4) > _buffer.size)
            return false

        val fbeStringSize = readUInt32(fbeStringOffset).toLong()
        if ((_buffer.offset + fbeStringOffset + 4 + fbeStringSize) > _buffer.size)
            return false

        return true
    }

    // Get the string value
    fun get(defaults: String = ""): String
    {
        if ((_buffer.offset + fbeOffset + fbeSize) > _buffer.size)
            return defaults

        val fbeStringOffset = readUInt32(fbeOffset).toLong()
        if (fbeStringOffset == 0L)
            return defaults

        assert((_buffer.offset + fbeStringOffset + 4) <= _buffer.size) { "Model is broken!" }
        if ((_buffer.offset + fbeStringOffset + 4) > _buffer.size)
            return defaults

        val fbeStringSize = readUInt32(fbeStringOffset).toLong()
        assert((_buffer.offset + fbeStringOffset + 4 + fbeStringSize) <= _buffer.size) { "Model is broken!" }
        if ((_buffer.offset + fbeStringOffset + 4 + fbeStringSize) > _buffer.size)
            return defaults

        return readString(fbeStringOffset + 4, fbeStringSize)
    }

    // Set the string value
    fun set(value: String)
    {
        assert((_buffer.offset + fbeOffset + fbeSize) <= _buffer.size) { "Model is broken!" }
        if ((_buffer.offset + fbeOffset + fbeSize) > _buffer.size)
            return

        val bytes = value.toByteArray(StandardCharsets.UTF_8)

        val fbeStringSize = bytes.size.toLong()
        val fbeStringOffset = _buffer.allocate(4 + fbeStringSize) - _buffer.offset
        assert((fbeStringOffset > 0) && ((_buffer.offset + fbeStringOffset + 4 + fbeStringSize) <= _buffer.size)) { "Model is broken!" }
        if ((fbeStringOffset <= 0) || ((_buffer.offset + fbeStringOffset + 4 + fbeStringSize) > _buffer.size))
            return

        write(fbeOffset, fbeStringOffset.toUInt())
        write(fbeStringOffset, fbeStringSize.toUInt())
        write(fbeStringOffset + 4, bytes)
    }
}