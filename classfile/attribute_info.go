package classfile

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/
type AttributeInfo interface{}

type UnparsedAttribute struct {
	Name   string
	Length uint32
	Info   []byte
}

func readAttributes(reader *ClassReader) []AttributeInfo {
	attributesCount := reader.ReadUint16()
	attributes := make([]AttributeInfo, attributesCount)
	for i := range attributes {
		attributes[i] = readAttributeInfo(reader)
	}
	return attributes
}

func readAttributeInfo(reader *ClassReader) AttributeInfo {
	attrNameIndex := reader.ReadUint16()
	attrLen := reader.ReadUint32()
	attrName := reader.cp.getUtf8(attrNameIndex)

	switch attrName {
	// case "AnnotationDefault":
	case "BootstrapMethods":
		return readBootstrapMethodsAttribute(reader)
	case "Code":
		return readCodeAttribute(reader)
	case "ConstantValue":
		return readConstantValueAttribute(reader)
	case "Deprecated":
		return DeprecatedAttribute{}
	case "EnclosingMethod":
		return readEnclosingMethodAttribute(reader)
	case "Exceptions":
		return readExceptionsAttribute(reader)
	case "InnerClasses":
		return readInnerClassesAttribute(reader)
	case "LineNumberTable":
		return readLineNumberTableAttribute(reader)
	case "LocalVariableTable":
		return readLocalVariableTableAttribute(reader)
	case "LocalVariableTypeTable":
		return readLocalVariableTypeTableAttribute(reader)
	// case "MethodParameters":
	// case "RuntimeInvisibleAnnotations":
	// case "RuntimeInvisibleParameterAnnotations":
	// case "RuntimeInvisibleTypeAnnotations":
	// case "RuntimeVisibleAnnotations":
	// case "RuntimeVisibleParameterAnnotations":
	// case "RuntimeVisibleTypeAnnotations":
	case "Signature":
		return readSignatureAttribute(reader)
	case "SourceFile":
		return readSourceFileAttribute(reader)
	// case "SourceDebugExtension":
	// case "StackMapTable":
	case "Synthetic":
		return SyntheticAttribute{}
	default:
		// undefined attr
		return UnparsedAttribute{
			Name:   attrName,
			Length: attrLen,
			Info:   reader.ReadBytes(uint(attrLen)),
		}
	}
}
