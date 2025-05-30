package errors

type ErrorCode string

const IndexOutOfBoundsError ErrorCode = "INDEX_OUT_OF_BOUNDS_EXCEPTION"
const NoSuchElementError ErrorCode = "NO_SUCH_ELEMENT_EXCEPTION"
const NullPointerError ErrorCode = "NULL_POINTER_EXCEPTION"
const IllegalArgumentError ErrorCode = "ILLEGAL_ARGUMENT_EXCEPTION"
const EmptyStackError ErrorCode = "EMPTY_STACK_EXCEPTION"
const UnsupportedOperationError ErrorCode = "UNSUPPORTED_OPERATION_EXCEPTION"
const ConcurrentModificationError ErrorCode = "CONCURRENT_MODIFICATION_EXCEPTION"
const ClassCastError ErrorCode = "CLASS_CAST_EXCEPTION"
const NoSuchMethodError ErrorCode = "NO_SUCH_METHOD_EXCEPTION"
const IllegalStateError ErrorCode = "ILLEGAL_STATE_EXCEPTION"
const ArithmeticError ErrorCode = "ARITHMETIC_EXCEPTION"
const QueueIsEmptyError ErrorCode = "QueueIsEmptyError"
