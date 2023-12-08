//go:build linux && arm && !arm7

package sherpa_onnx

// #cgo LDFLAGS: -L ${SRCDIR}/lib/arm-unknown-linux-gnueabihf -lsherpa-onnx-c-api -lsherpa-onnx-core -lkaldi-native-fbank-core  -lkaldi-decoder-core -lsherpa-onnx-kaldifst-core -lsherpa-onnx-fst -lpiper_phonemize -lespeak-ng -lucd -lonnxruntime -Wl,-rpath,${SRCDIR}/lib/arm-unknown-linux-gnueabihf
import "C"
