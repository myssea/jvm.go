package misc

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
	_unsafe(arrayBaseOffset, "arrayBaseOffset", "(Ljava/lang/Class;)I")
	_unsafe(arrayIndexScale, "arrayIndexScale", "(Ljava/lang/Class;)I")
	_unsafe(objectFieldOffset, "objectFieldOffset", "(Ljava/lang/reflect/Field;)J")
	_unsafe(getBoolean, "getBoolean", "(Ljava/lang/Object;J)Z")
	_unsafe(putBoolean, "putBoolean", "(Ljava/lang/Object;JZ)V")
	_unsafe(getObject, "getObject", "(Ljava/lang/Object;J)Ljava/lang/Object;")
	_unsafe(putObject, "putObject", "(Ljava/lang/Object;JLjava/lang/Object;)V")
	_unsafe(getObjectVolatile, "getObjectVolatile", "(Ljava/lang/Object;J)Ljava/lang/Object;")
	_unsafe(putObjectVolatile, "putObjectVolatile", "(Ljava/lang/Object;JLjava/lang/Object;)V")
	_unsafe(getOrderedObject, "getOrderedObject", "(Ljava/lang/Object;J)Ljava/lang/Object;")
	_unsafe(putOrderedObject, "putOrderedObject", "(Ljava/lang/Object;JLjava/lang/Object;)V")
	_unsafe(getIntVolatile, "getIntVolatile", "(Ljava/lang/Object;J)I")
	_unsafe(getLongVolatile, "getLongVolatile", "(Ljava/lang/Object;J)J")
}

// public native int arrayBaseOffset(Class<?> type);
// (Ljava/lang/Class;)I
func arrayBaseOffset(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PushInt(0) // todo
}

// public native int arrayIndexScale(Class<?> type);
// (Ljava/lang/Class;)I
func arrayIndexScale(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PushInt(1) // todo
}

// public native long objectFieldOffset(Field field);
// (Ljava/lang/reflect/Field;)J
func objectFieldOffset(frame *rtda.Frame) {
	vars := frame.LocalVars()
	jField := vars.GetRef(1)

	offset := jField.GetFieldValue("slot", "I").(int32)

	stack := frame.OperandStack()
	stack.PushLong(int64(offset))
}

// public native boolean getBoolean(Object o, long offset);
// (Ljava/lang/Object;J)Z
func getBoolean(frame *rtda.Frame) {
	vars := frame.LocalVars()
	fields := vars.GetRef(1).Fields()
	offset := vars.GetLong(2)

	stack := frame.OperandStack()
	if anys, ok := fields.([]Any); ok {
		// object
		stack.PushBoolean(anys[offset].(int32) == 1)
	} else if bytes, ok := fields.([]int8); ok {
		// byte[]
		stack.PushBoolean(bytes[offset] == 1)
	} else {
		panic("getBoolean!")
	}
}

// public native void putBoolean(Object o, long offset, boolean x);
// (Ljava/lang/Object;JZ)V
func putBoolean(frame *rtda.Frame) {
	vars := frame.LocalVars()
	fields := vars.GetRef(1).Fields()
	offset := vars.GetLong(2)
	x := vars.GetInt(4)

	if anys, ok := fields.([]Any); ok {
		// object
		anys[offset] = x
	} else if bytes, ok := fields.([]int8); ok {
		// byte[]
		bytes[offset] = int8(x)
	} else {
		panic("putBoolean!")
	}
}

// public native void putObject(Object o, long offset, Object x);
// (Ljava/lang/Object;JLjava/lang/Object;)V
func putObject(frame *rtda.Frame) {
	vars := frame.LocalVars()
	fields := vars.GetRef(1).Fields()
	offset := vars.GetLong(2)
	x := vars.GetRef(4)

	if anys, ok := fields.([]Any); ok {
		// object
		anys[offset] = x
	} else if objs, ok := fields.([]*rtc.Obj); ok {
		// ref[]
		objs[offset] = x
	} else {
		panic("putObject!")
	}
}

// public native Object getObject(Object o, long offset);
// (Ljava/lang/Object;J)Ljava/lang/Object;
func getObject(frame *rtda.Frame) {
	vars := frame.LocalVars()
	fields := vars.GetRef(1).Fields()
	offset := vars.GetLong(2)

	if anys, ok := fields.([]Any); ok {
		// object
		x := _getObj(anys, offset)
		frame.OperandStack().PushRef(x)
	} else if objs, ok := fields.([]*rtc.Obj); ok {
		// ref[]
		x := objs[offset]
		frame.OperandStack().PushRef(x)
	} else {
		panic("getObject!")
	}
}
func _getObj(fields []Any, offset int64) *rtc.Obj {
	f := fields[offset]
	if f != nil {
		return f.(*rtc.Obj)
	} else {
		return nil
	}
}

// public native void putObjectVolatile(Object o, long offset, Object x);
// (Ljava/lang/Object;JLjava/lang/Object;)V
func putObjectVolatile(frame *rtda.Frame) {
	putObject(frame) // todo
}

// public native Object getObjectVolatile(Object o, long offset);
//(Ljava/lang/Object;J)Ljava/lang/Object;
func getObjectVolatile(frame *rtda.Frame) {
	getObject(frame) // todo
}

// public native void putOrderedObject(Object o, long offset, Object x);
// (Ljava/lang/Object;JLjava/lang/Object;)V
func putOrderedObject(frame *rtda.Frame) {
	putObjectVolatile(frame) // todo
}

// public native Object getOrderedObject(Object o, long offset);
//(Ljava/lang/Object;J)Ljava/lang/Object;
func getOrderedObject(frame *rtda.Frame) {
	getObjectVolatile(frame) // todo
}

// public native int getIntVolatile(Object o, long offset);
// (Ljava/lang/Object;J)I
func getIntVolatile(frame *rtda.Frame) {
	vars := frame.LocalVars()
	obj := vars.GetRef(1)
	offset := vars.GetLong(2)

	// todo
	value := obj.Fields().([]Any)[offset].(int32)
	frame.OperandStack().PushInt(value)
}

// public native long getLongVolatile(Object o, long offset);
// (Ljava/lang/Object;J)J
func getLongVolatile(frame *rtda.Frame) {
	vars := frame.LocalVars()
	obj := vars.GetRef(1)
	offset := vars.GetLong(2)

	// todo
	value := obj.Fields().([]Any)[offset].(int64)
	frame.OperandStack().PushLong(value)
}
