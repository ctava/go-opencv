// Copyright 2016 <chris1tava@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package opencv

//#include "opencv.h"
//#cgo linux  pkg-config: opencv
//#cgo darwin pkg-config: opencv
//#cgo freebsd pkg-config: opencv
//#cgo windows LDFLAGS: -lopencv_core242.dll -lopencv_imgproc242.dll -lopencv_photo242.dll -lopencv_highgui242.dll -lstdc++
import "C"

type Classifier struct {
	cascade *C.CvCascadeClassifer
}

/*
func LoadCascadeClassifier(lbp string) *CascadeClassifier {
	classifier := new(Classifier)
	classifier.cascade = C.cvLoadHaarClassifierCascade(C.CString(lbp), C.cvSize(1, 1))
	return classifier
}

func (this *Classifier) DetectObjects(image *IplImage) []*Rect {
	storage := C.cvCreateMemStorage(0)
	seq := C.cvHaarDetectObjects(unsafe.Pointer(image), this.cascade, storage, 1.1, 3, C.CV_HAAR_DO_CANNY_PRUNING, C.cvSize(0, 0), C.cvSize(0, 0))
	var faces []*Rect
	for i := 0; i < (int)(seq.total); i++ {
		rect := (*Rect)((*_Ctype_CvRect)(unsafe.Pointer(C.cvGetSeqElem(seq, C.int(i)))))
		faces = append(faces, rect)
	}

	storage_c := (*C.CvMemStorage)(storage)
	C.cvReleaseMemStorage(&storage_c)

	return faces
}

func (this *Classifier) Release() {
	cascade_c := (*C.CvLBPClassifierCascade)(this.cascade)
	C.cvReleaseLBPClassifierCascade(&cascade_c)
} */
