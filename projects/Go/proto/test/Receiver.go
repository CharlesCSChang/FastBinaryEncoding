// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// Version: 1.1.0.0

package test

import "errors"
import "../fbe"
import "../proto"

// Workaround for Go unused imports issue
var _ = fbe.Version
var _ = proto.Version

// Receive StructSimple interface
type OnReceiveStructSimple interface {
    OnReceiveStructSimple(value *StructSimple)
}

// Receive StructSimple function
type OnReceiveStructSimpleFunc func(value *StructSimple)
func (f OnReceiveStructSimpleFunc) OnReceiveStructSimple(value *StructSimple) {
    f(value)
}

// Receive StructOptional interface
type OnReceiveStructOptional interface {
    OnReceiveStructOptional(value *StructOptional)
}

// Receive StructOptional function
type OnReceiveStructOptionalFunc func(value *StructOptional)
func (f OnReceiveStructOptionalFunc) OnReceiveStructOptional(value *StructOptional) {
    f(value)
}

// Receive StructNested interface
type OnReceiveStructNested interface {
    OnReceiveStructNested(value *StructNested)
}

// Receive StructNested function
type OnReceiveStructNestedFunc func(value *StructNested)
func (f OnReceiveStructNestedFunc) OnReceiveStructNested(value *StructNested) {
    f(value)
}

// Receive StructBytes interface
type OnReceiveStructBytes interface {
    OnReceiveStructBytes(value *StructBytes)
}

// Receive StructBytes function
type OnReceiveStructBytesFunc func(value *StructBytes)
func (f OnReceiveStructBytesFunc) OnReceiveStructBytes(value *StructBytes) {
    f(value)
}

// Receive StructArray interface
type OnReceiveStructArray interface {
    OnReceiveStructArray(value *StructArray)
}

// Receive StructArray function
type OnReceiveStructArrayFunc func(value *StructArray)
func (f OnReceiveStructArrayFunc) OnReceiveStructArray(value *StructArray) {
    f(value)
}

// Receive StructVector interface
type OnReceiveStructVector interface {
    OnReceiveStructVector(value *StructVector)
}

// Receive StructVector function
type OnReceiveStructVectorFunc func(value *StructVector)
func (f OnReceiveStructVectorFunc) OnReceiveStructVector(value *StructVector) {
    f(value)
}

// Receive StructList interface
type OnReceiveStructList interface {
    OnReceiveStructList(value *StructList)
}

// Receive StructList function
type OnReceiveStructListFunc func(value *StructList)
func (f OnReceiveStructListFunc) OnReceiveStructList(value *StructList) {
    f(value)
}

// Receive StructSet interface
type OnReceiveStructSet interface {
    OnReceiveStructSet(value *StructSet)
}

// Receive StructSet function
type OnReceiveStructSetFunc func(value *StructSet)
func (f OnReceiveStructSetFunc) OnReceiveStructSet(value *StructSet) {
    f(value)
}

// Receive StructMap interface
type OnReceiveStructMap interface {
    OnReceiveStructMap(value *StructMap)
}

// Receive StructMap function
type OnReceiveStructMapFunc func(value *StructMap)
func (f OnReceiveStructMapFunc) OnReceiveStructMap(value *StructMap) {
    f(value)
}

// Receive StructHash interface
type OnReceiveStructHash interface {
    OnReceiveStructHash(value *StructHash)
}

// Receive StructHash function
type OnReceiveStructHashFunc func(value *StructHash)
func (f OnReceiveStructHashFunc) OnReceiveStructHash(value *StructHash) {
    f(value)
}

// Receive StructHashEx interface
type OnReceiveStructHashEx interface {
    OnReceiveStructHashEx(value *StructHashEx)
}

// Receive StructHashEx function
type OnReceiveStructHashExFunc func(value *StructHashEx)
func (f OnReceiveStructHashExFunc) OnReceiveStructHashEx(value *StructHashEx) {
    f(value)
}

// Fast Binary Encoding test receiver
type Receiver struct {
    *fbe.Receiver
    protoReceiver *proto.Receiver
    structSimpleValue *StructSimple
    structSimpleModel *StructSimpleModel
    structOptionalValue *StructOptional
    structOptionalModel *StructOptionalModel
    structNestedValue *StructNested
    structNestedModel *StructNestedModel
    structBytesValue *StructBytes
    structBytesModel *StructBytesModel
    structArrayValue *StructArray
    structArrayModel *StructArrayModel
    structVectorValue *StructVector
    structVectorModel *StructVectorModel
    structListValue *StructList
    structListModel *StructListModel
    structSetValue *StructSet
    structSetModel *StructSetModel
    structMapValue *StructMap
    structMapModel *StructMapModel
    structHashValue *StructHash
    structHashModel *StructHashModel
    structHashExValue *StructHashEx
    structHashExModel *StructHashExModel

    // Receive StructSimple handler
    HandlerOnReceiveStructSimple OnReceiveStructSimple
    // Receive StructOptional handler
    HandlerOnReceiveStructOptional OnReceiveStructOptional
    // Receive StructNested handler
    HandlerOnReceiveStructNested OnReceiveStructNested
    // Receive StructBytes handler
    HandlerOnReceiveStructBytes OnReceiveStructBytes
    // Receive StructArray handler
    HandlerOnReceiveStructArray OnReceiveStructArray
    // Receive StructVector handler
    HandlerOnReceiveStructVector OnReceiveStructVector
    // Receive StructList handler
    HandlerOnReceiveStructList OnReceiveStructList
    // Receive StructSet handler
    HandlerOnReceiveStructSet OnReceiveStructSet
    // Receive StructMap handler
    HandlerOnReceiveStructMap OnReceiveStructMap
    // Receive StructHash handler
    HandlerOnReceiveStructHash OnReceiveStructHash
    // Receive StructHashEx handler
    HandlerOnReceiveStructHashEx OnReceiveStructHashEx
}

// Create a new test receiver with an empty buffer
func NewReceiver() *Receiver {
    return NewReceiverWithBuffer(fbe.NewEmptyBuffer())
}

// Create a new test receiver with the given buffer
func NewReceiverWithBuffer(buffer *fbe.Buffer) *Receiver {
    receiver := &Receiver{
        fbe.NewReceiver(buffer, false),
        proto.NewReceiverWithBuffer(buffer),
        NewStructSimple(),
        NewStructSimpleModel(buffer),
        NewStructOptional(),
        NewStructOptionalModel(buffer),
        NewStructNested(),
        NewStructNestedModel(buffer),
        NewStructBytes(),
        NewStructBytesModel(buffer),
        NewStructArray(),
        NewStructArrayModel(buffer),
        NewStructVector(),
        NewStructVectorModel(buffer),
        NewStructList(),
        NewStructListModel(buffer),
        NewStructSet(),
        NewStructSetModel(buffer),
        NewStructMap(),
        NewStructMapModel(buffer),
        NewStructHash(),
        NewStructHashModel(buffer),
        NewStructHashEx(),
        NewStructHashExModel(buffer),
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
    }
    receiver.SetupHandlerOnReceive(receiver)
    receiver.SetupHandlerOnReceiveStructSimpleFunc(func(value *StructSimple) {})
    receiver.SetupHandlerOnReceiveStructOptionalFunc(func(value *StructOptional) {})
    receiver.SetupHandlerOnReceiveStructNestedFunc(func(value *StructNested) {})
    receiver.SetupHandlerOnReceiveStructBytesFunc(func(value *StructBytes) {})
    receiver.SetupHandlerOnReceiveStructArrayFunc(func(value *StructArray) {})
    receiver.SetupHandlerOnReceiveStructVectorFunc(func(value *StructVector) {})
    receiver.SetupHandlerOnReceiveStructListFunc(func(value *StructList) {})
    receiver.SetupHandlerOnReceiveStructSetFunc(func(value *StructSet) {})
    receiver.SetupHandlerOnReceiveStructMapFunc(func(value *StructMap) {})
    receiver.SetupHandlerOnReceiveStructHashFunc(func(value *StructHash) {})
    receiver.SetupHandlerOnReceiveStructHashExFunc(func(value *StructHashEx) {})
    return receiver
}

// Imported receivers

// Get the proto receiver
func (r *Receiver) ProtoReceiver() *proto.Receiver { return r.protoReceiver }
// Set the proto receiver
func (r *Receiver) SetProtoReceiver(receiver *proto.Receiver) { r.protoReceiver = receiver }

// Setup handlers
func (r *Receiver) SetupHandlers(handlers interface{}) {
    r.Receiver.SetupHandlers(handlers)
    r.protoReceiver.SetupHandlers(handlers)
    if handler, ok := handlers.(OnReceiveStructSimple); ok {
        r.SetupHandlerOnReceiveStructSimple(handler)
    }
    if handler, ok := handlers.(OnReceiveStructOptional); ok {
        r.SetupHandlerOnReceiveStructOptional(handler)
    }
    if handler, ok := handlers.(OnReceiveStructNested); ok {
        r.SetupHandlerOnReceiveStructNested(handler)
    }
    if handler, ok := handlers.(OnReceiveStructBytes); ok {
        r.SetupHandlerOnReceiveStructBytes(handler)
    }
    if handler, ok := handlers.(OnReceiveStructArray); ok {
        r.SetupHandlerOnReceiveStructArray(handler)
    }
    if handler, ok := handlers.(OnReceiveStructVector); ok {
        r.SetupHandlerOnReceiveStructVector(handler)
    }
    if handler, ok := handlers.(OnReceiveStructList); ok {
        r.SetupHandlerOnReceiveStructList(handler)
    }
    if handler, ok := handlers.(OnReceiveStructSet); ok {
        r.SetupHandlerOnReceiveStructSet(handler)
    }
    if handler, ok := handlers.(OnReceiveStructMap); ok {
        r.SetupHandlerOnReceiveStructMap(handler)
    }
    if handler, ok := handlers.(OnReceiveStructHash); ok {
        r.SetupHandlerOnReceiveStructHash(handler)
    }
    if handler, ok := handlers.(OnReceiveStructHashEx); ok {
        r.SetupHandlerOnReceiveStructHashEx(handler)
    }
}

// Setup receive StructSimple handler
func (r *Receiver) SetupHandlerOnReceiveStructSimple(handler OnReceiveStructSimple) { r.HandlerOnReceiveStructSimple = handler }
// Setup receive StructSimple handler function
func (r *Receiver) SetupHandlerOnReceiveStructSimpleFunc(function func(value *StructSimple)) { r.HandlerOnReceiveStructSimple = OnReceiveStructSimpleFunc(function) }
// Setup receive StructOptional handler
func (r *Receiver) SetupHandlerOnReceiveStructOptional(handler OnReceiveStructOptional) { r.HandlerOnReceiveStructOptional = handler }
// Setup receive StructOptional handler function
func (r *Receiver) SetupHandlerOnReceiveStructOptionalFunc(function func(value *StructOptional)) { r.HandlerOnReceiveStructOptional = OnReceiveStructOptionalFunc(function) }
// Setup receive StructNested handler
func (r *Receiver) SetupHandlerOnReceiveStructNested(handler OnReceiveStructNested) { r.HandlerOnReceiveStructNested = handler }
// Setup receive StructNested handler function
func (r *Receiver) SetupHandlerOnReceiveStructNestedFunc(function func(value *StructNested)) { r.HandlerOnReceiveStructNested = OnReceiveStructNestedFunc(function) }
// Setup receive StructBytes handler
func (r *Receiver) SetupHandlerOnReceiveStructBytes(handler OnReceiveStructBytes) { r.HandlerOnReceiveStructBytes = handler }
// Setup receive StructBytes handler function
func (r *Receiver) SetupHandlerOnReceiveStructBytesFunc(function func(value *StructBytes)) { r.HandlerOnReceiveStructBytes = OnReceiveStructBytesFunc(function) }
// Setup receive StructArray handler
func (r *Receiver) SetupHandlerOnReceiveStructArray(handler OnReceiveStructArray) { r.HandlerOnReceiveStructArray = handler }
// Setup receive StructArray handler function
func (r *Receiver) SetupHandlerOnReceiveStructArrayFunc(function func(value *StructArray)) { r.HandlerOnReceiveStructArray = OnReceiveStructArrayFunc(function) }
// Setup receive StructVector handler
func (r *Receiver) SetupHandlerOnReceiveStructVector(handler OnReceiveStructVector) { r.HandlerOnReceiveStructVector = handler }
// Setup receive StructVector handler function
func (r *Receiver) SetupHandlerOnReceiveStructVectorFunc(function func(value *StructVector)) { r.HandlerOnReceiveStructVector = OnReceiveStructVectorFunc(function) }
// Setup receive StructList handler
func (r *Receiver) SetupHandlerOnReceiveStructList(handler OnReceiveStructList) { r.HandlerOnReceiveStructList = handler }
// Setup receive StructList handler function
func (r *Receiver) SetupHandlerOnReceiveStructListFunc(function func(value *StructList)) { r.HandlerOnReceiveStructList = OnReceiveStructListFunc(function) }
// Setup receive StructSet handler
func (r *Receiver) SetupHandlerOnReceiveStructSet(handler OnReceiveStructSet) { r.HandlerOnReceiveStructSet = handler }
// Setup receive StructSet handler function
func (r *Receiver) SetupHandlerOnReceiveStructSetFunc(function func(value *StructSet)) { r.HandlerOnReceiveStructSet = OnReceiveStructSetFunc(function) }
// Setup receive StructMap handler
func (r *Receiver) SetupHandlerOnReceiveStructMap(handler OnReceiveStructMap) { r.HandlerOnReceiveStructMap = handler }
// Setup receive StructMap handler function
func (r *Receiver) SetupHandlerOnReceiveStructMapFunc(function func(value *StructMap)) { r.HandlerOnReceiveStructMap = OnReceiveStructMapFunc(function) }
// Setup receive StructHash handler
func (r *Receiver) SetupHandlerOnReceiveStructHash(handler OnReceiveStructHash) { r.HandlerOnReceiveStructHash = handler }
// Setup receive StructHash handler function
func (r *Receiver) SetupHandlerOnReceiveStructHashFunc(function func(value *StructHash)) { r.HandlerOnReceiveStructHash = OnReceiveStructHashFunc(function) }
// Setup receive StructHashEx handler
func (r *Receiver) SetupHandlerOnReceiveStructHashEx(handler OnReceiveStructHashEx) { r.HandlerOnReceiveStructHashEx = handler }
// Setup receive StructHashEx handler function
func (r *Receiver) SetupHandlerOnReceiveStructHashExFunc(function func(value *StructHashEx)) { r.HandlerOnReceiveStructHashEx = OnReceiveStructHashExFunc(function) }

// Receive message handler
func (r *Receiver) OnReceive(fbeType int, buffer []byte) (bool, error) {
    switch fbeType {
    case r.structSimpleModel.FBEType():
        // Deserialize the value from the FBE stream
        r.structSimpleModel.Buffer().Attach(buffer)
        if !r.structSimpleModel.Verify() {
            return false, errors.New("test.StructSimple validation failed")
        }
        deserialized, err := r.structSimpleModel.DeserializeValue(r.structSimpleValue)
        if deserialized <= 0 {
            return false, errors.New("test.StructSimple deserialization failed")
        }
        if err != nil {
            return false, err
        }

        // Log the value
        if r.Logging() {
            message := r.structSimpleValue.String()
            r.HandlerOnReceiveLog.OnReceiveLog(message)
        }

        // Call receive handler with deserialized value
        r.HandlerOnReceiveStructSimple.OnReceiveStructSimple(r.structSimpleValue)
        return true, nil
    case r.structOptionalModel.FBEType():
        // Deserialize the value from the FBE stream
        r.structOptionalModel.Buffer().Attach(buffer)
        if !r.structOptionalModel.Verify() {
            return false, errors.New("test.StructOptional validation failed")
        }
        deserialized, err := r.structOptionalModel.DeserializeValue(r.structOptionalValue)
        if deserialized <= 0 {
            return false, errors.New("test.StructOptional deserialization failed")
        }
        if err != nil {
            return false, err
        }

        // Log the value
        if r.Logging() {
            message := r.structOptionalValue.String()
            r.HandlerOnReceiveLog.OnReceiveLog(message)
        }

        // Call receive handler with deserialized value
        r.HandlerOnReceiveStructOptional.OnReceiveStructOptional(r.structOptionalValue)
        return true, nil
    case r.structNestedModel.FBEType():
        // Deserialize the value from the FBE stream
        r.structNestedModel.Buffer().Attach(buffer)
        if !r.structNestedModel.Verify() {
            return false, errors.New("test.StructNested validation failed")
        }
        deserialized, err := r.structNestedModel.DeserializeValue(r.structNestedValue)
        if deserialized <= 0 {
            return false, errors.New("test.StructNested deserialization failed")
        }
        if err != nil {
            return false, err
        }

        // Log the value
        if r.Logging() {
            message := r.structNestedValue.String()
            r.HandlerOnReceiveLog.OnReceiveLog(message)
        }

        // Call receive handler with deserialized value
        r.HandlerOnReceiveStructNested.OnReceiveStructNested(r.structNestedValue)
        return true, nil
    case r.structBytesModel.FBEType():
        // Deserialize the value from the FBE stream
        r.structBytesModel.Buffer().Attach(buffer)
        if !r.structBytesModel.Verify() {
            return false, errors.New("test.StructBytes validation failed")
        }
        deserialized, err := r.structBytesModel.DeserializeValue(r.structBytesValue)
        if deserialized <= 0 {
            return false, errors.New("test.StructBytes deserialization failed")
        }
        if err != nil {
            return false, err
        }

        // Log the value
        if r.Logging() {
            message := r.structBytesValue.String()
            r.HandlerOnReceiveLog.OnReceiveLog(message)
        }

        // Call receive handler with deserialized value
        r.HandlerOnReceiveStructBytes.OnReceiveStructBytes(r.structBytesValue)
        return true, nil
    case r.structArrayModel.FBEType():
        // Deserialize the value from the FBE stream
        r.structArrayModel.Buffer().Attach(buffer)
        if !r.structArrayModel.Verify() {
            return false, errors.New("test.StructArray validation failed")
        }
        deserialized, err := r.structArrayModel.DeserializeValue(r.structArrayValue)
        if deserialized <= 0 {
            return false, errors.New("test.StructArray deserialization failed")
        }
        if err != nil {
            return false, err
        }

        // Log the value
        if r.Logging() {
            message := r.structArrayValue.String()
            r.HandlerOnReceiveLog.OnReceiveLog(message)
        }

        // Call receive handler with deserialized value
        r.HandlerOnReceiveStructArray.OnReceiveStructArray(r.structArrayValue)
        return true, nil
    case r.structVectorModel.FBEType():
        // Deserialize the value from the FBE stream
        r.structVectorModel.Buffer().Attach(buffer)
        if !r.structVectorModel.Verify() {
            return false, errors.New("test.StructVector validation failed")
        }
        deserialized, err := r.structVectorModel.DeserializeValue(r.structVectorValue)
        if deserialized <= 0 {
            return false, errors.New("test.StructVector deserialization failed")
        }
        if err != nil {
            return false, err
        }

        // Log the value
        if r.Logging() {
            message := r.structVectorValue.String()
            r.HandlerOnReceiveLog.OnReceiveLog(message)
        }

        // Call receive handler with deserialized value
        r.HandlerOnReceiveStructVector.OnReceiveStructVector(r.structVectorValue)
        return true, nil
    case r.structListModel.FBEType():
        // Deserialize the value from the FBE stream
        r.structListModel.Buffer().Attach(buffer)
        if !r.structListModel.Verify() {
            return false, errors.New("test.StructList validation failed")
        }
        deserialized, err := r.structListModel.DeserializeValue(r.structListValue)
        if deserialized <= 0 {
            return false, errors.New("test.StructList deserialization failed")
        }
        if err != nil {
            return false, err
        }

        // Log the value
        if r.Logging() {
            message := r.structListValue.String()
            r.HandlerOnReceiveLog.OnReceiveLog(message)
        }

        // Call receive handler with deserialized value
        r.HandlerOnReceiveStructList.OnReceiveStructList(r.structListValue)
        return true, nil
    case r.structSetModel.FBEType():
        // Deserialize the value from the FBE stream
        r.structSetModel.Buffer().Attach(buffer)
        if !r.structSetModel.Verify() {
            return false, errors.New("test.StructSet validation failed")
        }
        deserialized, err := r.structSetModel.DeserializeValue(r.structSetValue)
        if deserialized <= 0 {
            return false, errors.New("test.StructSet deserialization failed")
        }
        if err != nil {
            return false, err
        }

        // Log the value
        if r.Logging() {
            message := r.structSetValue.String()
            r.HandlerOnReceiveLog.OnReceiveLog(message)
        }

        // Call receive handler with deserialized value
        r.HandlerOnReceiveStructSet.OnReceiveStructSet(r.structSetValue)
        return true, nil
    case r.structMapModel.FBEType():
        // Deserialize the value from the FBE stream
        r.structMapModel.Buffer().Attach(buffer)
        if !r.structMapModel.Verify() {
            return false, errors.New("test.StructMap validation failed")
        }
        deserialized, err := r.structMapModel.DeserializeValue(r.structMapValue)
        if deserialized <= 0 {
            return false, errors.New("test.StructMap deserialization failed")
        }
        if err != nil {
            return false, err
        }

        // Log the value
        if r.Logging() {
            message := r.structMapValue.String()
            r.HandlerOnReceiveLog.OnReceiveLog(message)
        }

        // Call receive handler with deserialized value
        r.HandlerOnReceiveStructMap.OnReceiveStructMap(r.structMapValue)
        return true, nil
    case r.structHashModel.FBEType():
        // Deserialize the value from the FBE stream
        r.structHashModel.Buffer().Attach(buffer)
        if !r.structHashModel.Verify() {
            return false, errors.New("test.StructHash validation failed")
        }
        deserialized, err := r.structHashModel.DeserializeValue(r.structHashValue)
        if deserialized <= 0 {
            return false, errors.New("test.StructHash deserialization failed")
        }
        if err != nil {
            return false, err
        }

        // Log the value
        if r.Logging() {
            message := r.structHashValue.String()
            r.HandlerOnReceiveLog.OnReceiveLog(message)
        }

        // Call receive handler with deserialized value
        r.HandlerOnReceiveStructHash.OnReceiveStructHash(r.structHashValue)
        return true, nil
    case r.structHashExModel.FBEType():
        // Deserialize the value from the FBE stream
        r.structHashExModel.Buffer().Attach(buffer)
        if !r.structHashExModel.Verify() {
            return false, errors.New("test.StructHashEx validation failed")
        }
        deserialized, err := r.structHashExModel.DeserializeValue(r.structHashExValue)
        if deserialized <= 0 {
            return false, errors.New("test.StructHashEx deserialization failed")
        }
        if err != nil {
            return false, err
        }

        // Log the value
        if r.Logging() {
            message := r.structHashExValue.String()
            r.HandlerOnReceiveLog.OnReceiveLog(message)
        }

        // Call receive handler with deserialized value
        r.HandlerOnReceiveStructHashEx.OnReceiveStructHashEx(r.structHashExValue)
        return true, nil
    }

    if r.protoReceiver != nil {
        if ok, err := r.protoReceiver.OnReceive(fbeType, buffer); ok {
            return ok, err
        }
    }

    return false, nil
}
