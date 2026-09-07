package actorindex

import "testing"

func TestCalculateIoU(t *testing.T) {
	b1 := BoundingBox{X: 0, Y: 0, W: 10, H: 10}
	b2 := BoundingBox{X: 0, Y: 0, W: 10, H: 10}
	iou := CalculateIoU(b1, b2)
	if iou != 1.0 {
		t.Errorf("expected IoU 1.0, got %f", iou)
	}
}
