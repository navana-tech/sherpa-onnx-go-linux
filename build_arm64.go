//go:build linux && arm64

package sherpa_onnx

// #cgo LDFLAGS: -L ${SRCDIR}/lib/aarch64-unknown-linux-gnu -lsherpa-onnx-c-api -lsherpa-onnx-core -lkaldi-native-fbank-core -lonnxruntime -lkaldi-decoder-core -lsherpa-onnx-kaldifst-core -lsherpa-onnx-fst -Wl,-rpath,${SRCDIR}/lib/aarch64-unknown-linux-gnu
import "C"
