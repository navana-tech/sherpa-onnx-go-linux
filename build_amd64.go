//go:build !android && linux && amd64 && !musl

package sherpa_onnx

// #cgo LDFLAGS: -L ${SRCDIR}/lib/x86_64-unknown-linux-gnu -lsherpa-onnx-c-api -lsherpa-onnx-core -lkaldi-native-fbank-core -lonnxruntime -lkaldi-decoder-core -lsherpa-onnx-kaldifst-core -lsherpa-onnx-fst -Wl,-rpath,${SRCDIR}/lib/x86_64-unknown-linux-gnu
import "C"
