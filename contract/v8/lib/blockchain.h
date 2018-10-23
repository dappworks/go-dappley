#include <v8.h>

using namespace v8;

void NewBlockchainInstance(Isolate *isolate, Local<Context> context, void *handler);
void VerifyAddressCallback(const FunctionCallbackInfo<Value> &info);
