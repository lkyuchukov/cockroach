// Code generated by "stringer"; DO NOT EDIT.

package sqlproxyccl

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[codeAuthFailed-1]
	_ = x[codeBackendReadFailed-2]
	_ = x[codeBackendWriteFailed-3]
	_ = x[codeClientReadFailed-4]
	_ = x[codeClientWriteFailed-5]
	_ = x[codeUnexpectedInsecureStartupMessage-6]
	_ = x[codeSNIRoutingFailed-7]
	_ = x[codeUnexpectedStartupMessage-8]
	_ = x[codeParamsRoutingFailed-9]
	_ = x[codeBackendDown-10]
	_ = x[codeBackendRefusedTLS-11]
	_ = x[codeBackendDisconnected-12]
	_ = x[codeClientDisconnected-13]
	_ = x[codeProxyRefusedConnection-14]
	_ = x[codeExpiredClientConnection-15]
	_ = x[codeIdleDisconnect-16]
	_ = x[codeUnavailable-17]
}

const _errorCode_name = "codeAuthFailedcodeBackendReadFailedcodeBackendWriteFailedcodeClientReadFailedcodeClientWriteFailedcodeUnexpectedInsecureStartupMessagecodeSNIRoutingFailedcodeUnexpectedStartupMessagecodeParamsRoutingFailedcodeBackendDowncodeBackendRefusedTLScodeBackendDisconnectedcodeClientDisconnectedcodeProxyRefusedConnectioncodeExpiredClientConnectioncodeIdleDisconnectcodeUnavailable"

var _errorCode_index = [...]uint16{0, 14, 35, 57, 77, 98, 134, 154, 182, 205, 220, 241, 264, 286, 312, 339, 357, 372}

func (i errorCode) String() string {
	i -= 1
	if i < 0 || i >= errorCode(len(_errorCode_index)-1) {
		return "errorCode(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _errorCode_name[_errorCode_index[i]:_errorCode_index[i+1]]
}
