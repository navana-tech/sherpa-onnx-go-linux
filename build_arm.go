//go:build linux && arm && !arm7

package sherpa_onnx

// #cgo LDFLAGS: -L ${SRCDIR}/lib/arm-unknown-linux-gnueabihf -lsherpa-onnx-c-api -lsherpa-onnx-core -lkaldi-native-fbank-core -lonnxruntime -Wl,-rpath,${SRCDIR}/lib/arm-unknown-linux-gnueabihf
import "C"
