#ifndef _DAPPLEY_NF_VM_V8_ERROR_H_
#define _DAPPLEY_NF_VM_V8_ERROR_H_
enum vmErrno {
    VM_SUCCESS = 0,
    VM_EXCEPTION_ERR = -1,
    VM_MEM_LIMIT_ERR = -2,
    VM_GAS_LIMIT_ERR = -3,
    VM_UNEXPECTED_ERR = -4,
    VM_EXE_TIMEOUT_ERR = -5,
    VM_INNER_EXE_ERR = -6,
};
#endif
