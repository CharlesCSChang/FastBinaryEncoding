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

// Fast Binary Encoding test sender
type Sender struct {
    *fbe.Sender
    protoSender *proto.Sender
    structSimpleModel *StructSimpleModel
    structOptionalModel *StructOptionalModel
    structNestedModel *StructNestedModel
    structBytesModel *StructBytesModel
    structArrayModel *StructArrayModel
    structVectorModel *StructVectorModel
    structListModel *StructListModel
    structSetModel *StructSetModel
    structMapModel *StructMapModel
    structHashModel *StructHashModel
    structHashExModel *StructHashExModel
}

// Create a new test sender with an empty buffer
func NewSender() *Sender {
    return NewSenderWithBuffer(fbe.NewEmptyBuffer())
}

// Create a new test sender with the given buffer
func NewSenderWithBuffer(buffer *fbe.Buffer) *Sender {
    return &Sender{
        fbe.NewSender(buffer, false),
        proto.NewSenderWithBuffer(buffer),
        NewStructSimpleModel(buffer),
        NewStructOptionalModel(buffer),
        NewStructNestedModel(buffer),
        NewStructBytesModel(buffer),
        NewStructArrayModel(buffer),
        NewStructVectorModel(buffer),
        NewStructListModel(buffer),
        NewStructSetModel(buffer),
        NewStructMapModel(buffer),
        NewStructHashModel(buffer),
        NewStructHashExModel(buffer),
    }
}

// Imported senders

func (s *Sender) ProtoSender() *proto.Sender { return s.protoSender }

// Sender models accessors

func (s *Sender) StructSimpleModel() *StructSimpleModel { return s.structSimpleModel }
func (s *Sender) StructOptionalModel() *StructOptionalModel { return s.structOptionalModel }
func (s *Sender) StructNestedModel() *StructNestedModel { return s.structNestedModel }
func (s *Sender) StructBytesModel() *StructBytesModel { return s.structBytesModel }
func (s *Sender) StructArrayModel() *StructArrayModel { return s.structArrayModel }
func (s *Sender) StructVectorModel() *StructVectorModel { return s.structVectorModel }
func (s *Sender) StructListModel() *StructListModel { return s.structListModel }
func (s *Sender) StructSetModel() *StructSetModel { return s.structSetModel }
func (s *Sender) StructMapModel() *StructMapModel { return s.structMapModel }
func (s *Sender) StructHashModel() *StructHashModel { return s.structHashModel }
func (s *Sender) StructHashExModel() *StructHashExModel { return s.structHashExModel }

// Send methods

func (s *Sender) Send(value interface{}) (int, error) {
    switch value := value.(type) {
    case *StructSimple:
        return s.SendStructSimple(value)
    case *StructOptional:
        return s.SendStructOptional(value)
    case *StructNested:
        return s.SendStructNested(value)
    case *StructBytes:
        return s.SendStructBytes(value)
    case *StructArray:
        return s.SendStructArray(value)
    case *StructVector:
        return s.SendStructVector(value)
    case *StructList:
        return s.SendStructList(value)
    case *StructSet:
        return s.SendStructSet(value)
    case *StructMap:
        return s.SendStructMap(value)
    case *StructHash:
        return s.SendStructHash(value)
    case *StructHashEx:
        return s.SendStructHashEx(value)
    }
    if result, err := s.protoSender.Send(value); (result > 0) || (err != nil) {
        return result, err
    }
    return 0, nil
}

func (s *Sender) SendStructSimple(value *StructSimple) (int, error) {
    // Serialize the value into the FBE stream
    serialized, err := s.structSimpleModel.Serialize(value)
    if serialized <= 0 {
        return 0, errors.New("test.StructSimple serialization failed")
    }
    if err != nil {
        return 0, err
    }
    if !s.structSimpleModel.Verify() {
        return 0, errors.New("test.StructSimple validation failed")
    }

    // Log the value
    if s.Logging() {
        message := value.String()
        s.HandlerOnSendLog.OnSendLog(message)
    }

    // Send the serialized value
    return s.SendSerialized(serialized)
}

func (s *Sender) SendStructOptional(value *StructOptional) (int, error) {
    // Serialize the value into the FBE stream
    serialized, err := s.structOptionalModel.Serialize(value)
    if serialized <= 0 {
        return 0, errors.New("test.StructOptional serialization failed")
    }
    if err != nil {
        return 0, err
    }
    if !s.structOptionalModel.Verify() {
        return 0, errors.New("test.StructOptional validation failed")
    }

    // Log the value
    if s.Logging() {
        message := value.String()
        s.HandlerOnSendLog.OnSendLog(message)
    }

    // Send the serialized value
    return s.SendSerialized(serialized)
}

func (s *Sender) SendStructNested(value *StructNested) (int, error) {
    // Serialize the value into the FBE stream
    serialized, err := s.structNestedModel.Serialize(value)
    if serialized <= 0 {
        return 0, errors.New("test.StructNested serialization failed")
    }
    if err != nil {
        return 0, err
    }
    if !s.structNestedModel.Verify() {
        return 0, errors.New("test.StructNested validation failed")
    }

    // Log the value
    if s.Logging() {
        message := value.String()
        s.HandlerOnSendLog.OnSendLog(message)
    }

    // Send the serialized value
    return s.SendSerialized(serialized)
}

func (s *Sender) SendStructBytes(value *StructBytes) (int, error) {
    // Serialize the value into the FBE stream
    serialized, err := s.structBytesModel.Serialize(value)
    if serialized <= 0 {
        return 0, errors.New("test.StructBytes serialization failed")
    }
    if err != nil {
        return 0, err
    }
    if !s.structBytesModel.Verify() {
        return 0, errors.New("test.StructBytes validation failed")
    }

    // Log the value
    if s.Logging() {
        message := value.String()
        s.HandlerOnSendLog.OnSendLog(message)
    }

    // Send the serialized value
    return s.SendSerialized(serialized)
}

func (s *Sender) SendStructArray(value *StructArray) (int, error) {
    // Serialize the value into the FBE stream
    serialized, err := s.structArrayModel.Serialize(value)
    if serialized <= 0 {
        return 0, errors.New("test.StructArray serialization failed")
    }
    if err != nil {
        return 0, err
    }
    if !s.structArrayModel.Verify() {
        return 0, errors.New("test.StructArray validation failed")
    }

    // Log the value
    if s.Logging() {
        message := value.String()
        s.HandlerOnSendLog.OnSendLog(message)
    }

    // Send the serialized value
    return s.SendSerialized(serialized)
}

func (s *Sender) SendStructVector(value *StructVector) (int, error) {
    // Serialize the value into the FBE stream
    serialized, err := s.structVectorModel.Serialize(value)
    if serialized <= 0 {
        return 0, errors.New("test.StructVector serialization failed")
    }
    if err != nil {
        return 0, err
    }
    if !s.structVectorModel.Verify() {
        return 0, errors.New("test.StructVector validation failed")
    }

    // Log the value
    if s.Logging() {
        message := value.String()
        s.HandlerOnSendLog.OnSendLog(message)
    }

    // Send the serialized value
    return s.SendSerialized(serialized)
}

func (s *Sender) SendStructList(value *StructList) (int, error) {
    // Serialize the value into the FBE stream
    serialized, err := s.structListModel.Serialize(value)
    if serialized <= 0 {
        return 0, errors.New("test.StructList serialization failed")
    }
    if err != nil {
        return 0, err
    }
    if !s.structListModel.Verify() {
        return 0, errors.New("test.StructList validation failed")
    }

    // Log the value
    if s.Logging() {
        message := value.String()
        s.HandlerOnSendLog.OnSendLog(message)
    }

    // Send the serialized value
    return s.SendSerialized(serialized)
}

func (s *Sender) SendStructSet(value *StructSet) (int, error) {
    // Serialize the value into the FBE stream
    serialized, err := s.structSetModel.Serialize(value)
    if serialized <= 0 {
        return 0, errors.New("test.StructSet serialization failed")
    }
    if err != nil {
        return 0, err
    }
    if !s.structSetModel.Verify() {
        return 0, errors.New("test.StructSet validation failed")
    }

    // Log the value
    if s.Logging() {
        message := value.String()
        s.HandlerOnSendLog.OnSendLog(message)
    }

    // Send the serialized value
    return s.SendSerialized(serialized)
}

func (s *Sender) SendStructMap(value *StructMap) (int, error) {
    // Serialize the value into the FBE stream
    serialized, err := s.structMapModel.Serialize(value)
    if serialized <= 0 {
        return 0, errors.New("test.StructMap serialization failed")
    }
    if err != nil {
        return 0, err
    }
    if !s.structMapModel.Verify() {
        return 0, errors.New("test.StructMap validation failed")
    }

    // Log the value
    if s.Logging() {
        message := value.String()
        s.HandlerOnSendLog.OnSendLog(message)
    }

    // Send the serialized value
    return s.SendSerialized(serialized)
}

func (s *Sender) SendStructHash(value *StructHash) (int, error) {
    // Serialize the value into the FBE stream
    serialized, err := s.structHashModel.Serialize(value)
    if serialized <= 0 {
        return 0, errors.New("test.StructHash serialization failed")
    }
    if err != nil {
        return 0, err
    }
    if !s.structHashModel.Verify() {
        return 0, errors.New("test.StructHash validation failed")
    }

    // Log the value
    if s.Logging() {
        message := value.String()
        s.HandlerOnSendLog.OnSendLog(message)
    }

    // Send the serialized value
    return s.SendSerialized(serialized)
}

func (s *Sender) SendStructHashEx(value *StructHashEx) (int, error) {
    // Serialize the value into the FBE stream
    serialized, err := s.structHashExModel.Serialize(value)
    if serialized <= 0 {
        return 0, errors.New("test.StructHashEx serialization failed")
    }
    if err != nil {
        return 0, err
    }
    if !s.structHashExModel.Verify() {
        return 0, errors.New("test.StructHashEx validation failed")
    }

    // Log the value
    if s.Logging() {
        message := value.String()
        s.HandlerOnSendLog.OnSendLog(message)
    }

    // Send the serialized value
    return s.SendSerialized(serialized)
}
