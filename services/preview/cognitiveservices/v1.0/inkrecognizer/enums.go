package inkrecognizer

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// Application enumerates the values for application.
type Application string

const (
	// Drawing ...
	Drawing Application = "drawing"
	// Mixed ...
	Mixed Application = "mixed"
	// Writing ...
	Writing Application = "writing"
)

// PossibleApplicationValues returns an array of possible values for the Application const type.
func PossibleApplicationValues() []Application {
	return []Application{Drawing, Mixed, Writing}
}

// Category enumerates the values for category.
type Category string

const (
	// InkBullet ...
	InkBullet Category = "inkBullet"
	// InkDrawing ...
	InkDrawing Category = "inkDrawing"
	// InkWord ...
	InkWord Category = "inkWord"
	// Line ...
	Line Category = "line"
	// Paragraph ...
	Paragraph Category = "paragraph"
	// Root ...
	Root Category = "root"
	// Unknown ...
	Unknown Category = "unknown"
	// WritingRegion ...
	WritingRegion Category = "writingRegion"
)

// PossibleCategoryValues returns an array of possible values for the Category const type.
func PossibleCategoryValues() []Category {
	return []Category{InkBullet, InkDrawing, InkWord, Line, Paragraph, Root, Unknown, WritingRegion}
}

// Class enumerates the values for class.
type Class string

const (
	// ClassContainer ...
	ClassContainer Class = "container"
	// ClassLeaf ...
	ClassLeaf Class = "leaf"
)

// PossibleClassValues returns an array of possible values for the Class const type.
func PossibleClassValues() []Class {
	return []Class{ClassContainer, ClassLeaf}
}

// Container enumerates the values for container.
type Container string

const (
	// ContainerLine ...
	ContainerLine Container = "line"
	// ContainerParagraph ...
	ContainerParagraph Container = "paragraph"
	// ContainerRoot ...
	ContainerRoot Container = "root"
	// ContainerWritingRegion ...
	ContainerWritingRegion Container = "writingRegion"
)

// PossibleContainerValues returns an array of possible values for the Container const type.
func PossibleContainerValues() []Container {
	return []Container{ContainerLine, ContainerParagraph, ContainerRoot, ContainerWritingRegion}
}

// InputDevice enumerates the values for input device.
type InputDevice string

const (
	// Armature ...
	Armature InputDevice = "armature"
	// ArticulatedArm ...
	ArticulatedArm InputDevice = "articulatedArm"
	// Digitizer ...
	Digitizer InputDevice = "digitizer"
	// LightPen ...
	LightPen InputDevice = "lightPen"
	// Pen ...
	Pen InputDevice = "pen"
	// StereoPlotter ...
	StereoPlotter InputDevice = "stereoPlotter"
	// ThreedDigitizer ...
	ThreedDigitizer InputDevice = "3dDigitizer"
	// TouchPad ...
	TouchPad InputDevice = "touchPad"
	// TouchScreen ...
	TouchScreen InputDevice = "touchScreen"
	// WhiteBoard ...
	WhiteBoard InputDevice = "whiteBoard"
)

// PossibleInputDeviceValues returns an array of possible values for the InputDevice const type.
func PossibleInputDeviceValues() []InputDevice {
	return []InputDevice{Armature, ArticulatedArm, Digitizer, LightPen, Pen, StereoPlotter, ThreedDigitizer, TouchPad, TouchScreen, WhiteBoard}
}

// Kind enumerates the values for kind.
type Kind string

const (
	// KindInkDrawing ...
	KindInkDrawing Kind = "inkDrawing"
	// KindInkWriting ...
	KindInkWriting Kind = "inkWriting"
)

// PossibleKindValues returns an array of possible values for the Kind const type.
func PossibleKindValues() []Kind {
	return []Kind{KindInkDrawing, KindInkWriting}
}

// Leaf enumerates the values for leaf.
type Leaf string

const (
	// LeafInkBullet ...
	LeafInkBullet Leaf = "inkBullet"
	// LeafInkDrawing ...
	LeafInkDrawing Leaf = "inkDrawing"
	// LeafInkWord ...
	LeafInkWord Leaf = "inkWord"
	// LeafUnknown ...
	LeafUnknown Leaf = "unknown"
)

// PossibleLeafValues returns an array of possible values for the Leaf const type.
func PossibleLeafValues() []Leaf {
	return []Leaf{LeafInkBullet, LeafInkDrawing, LeafInkWord, LeafUnknown}
}

// RasterOp enumerates the values for raster op.
type RasterOp string

const (
	// CopyPen ...
	CopyPen RasterOp = "copyPen"
	// MaskPen ...
	MaskPen RasterOp = "maskPen"
	// NoOperation ...
	NoOperation RasterOp = "noOperation"
)

// PossibleRasterOpValues returns an array of possible values for the RasterOp const type.
func PossibleRasterOpValues() []RasterOp {
	return []RasterOp{CopyPen, MaskPen, NoOperation}
}

// Shape enumerates the values for shape.
type Shape string

const (
	// ShapeBlockArrow ...
	ShapeBlockArrow Shape = "blockArrow"
	// ShapeCircle ...
	ShapeCircle Shape = "circle"
	// ShapeCloud ...
	ShapeCloud Shape = "cloud"
	// ShapeCurve ...
	ShapeCurve Shape = "curve"
	// ShapeDiamond ...
	ShapeDiamond Shape = "diamond"
	// ShapeDrawing ...
	ShapeDrawing Shape = "drawing"
	// ShapeEllipse ...
	ShapeEllipse Shape = "ellipse"
	// ShapeEquilateralTriangle ...
	ShapeEquilateralTriangle Shape = "equilateralTriangle"
	// ShapeHeart ...
	ShapeHeart Shape = "heart"
	// ShapeHexagon ...
	ShapeHexagon Shape = "hexagon"
	// ShapeIsoscelesTriangle ...
	ShapeIsoscelesTriangle Shape = "isoscelesTriangle"
	// ShapeLine ...
	ShapeLine Shape = "line"
	// ShapeParallelogram ...
	ShapeParallelogram Shape = "parallelogram"
	// ShapePentagon ...
	ShapePentagon Shape = "pentagon"
	// ShapePolyLine ...
	ShapePolyLine Shape = "polyLine"
	// ShapeQuadrilateral ...
	ShapeQuadrilateral Shape = "quadrilateral"
	// ShapeRectangle ...
	ShapeRectangle Shape = "rectangle"
	// ShapeRightTriangle ...
	ShapeRightTriangle Shape = "rightTriangle"
	// ShapeSquare ...
	ShapeSquare Shape = "square"
	// ShapeStarCrossed ...
	ShapeStarCrossed Shape = "starCrossed"
	// ShapeStarSimple ...
	ShapeStarSimple Shape = "starSimple"
	// ShapeTrapezoid ...
	ShapeTrapezoid Shape = "trapezoid"
	// ShapeTriangle ...
	ShapeTriangle Shape = "triangle"
)

// PossibleShapeValues returns an array of possible values for the Shape const type.
func PossibleShapeValues() []Shape {
	return []Shape{ShapeBlockArrow, ShapeCircle, ShapeCloud, ShapeCurve, ShapeDiamond, ShapeDrawing, ShapeEllipse, ShapeEquilateralTriangle, ShapeHeart, ShapeHexagon, ShapeIsoscelesTriangle, ShapeLine, ShapeParallelogram, ShapePentagon, ShapePolyLine, ShapeQuadrilateral, ShapeRectangle, ShapeRightTriangle, ShapeSquare, ShapeStarCrossed, ShapeStarSimple, ShapeTrapezoid, ShapeTriangle}
}

// Tip enumerates the values for tip.
type Tip string

const (
	// Ellipse ...
	Ellipse Tip = "ellipse"
	// Rectangle ...
	Rectangle Tip = "rectangle"
)

// PossibleTipValues returns an array of possible values for the Tip const type.
func PossibleTipValues() []Tip {
	return []Tip{Ellipse, Rectangle}
}

// Unit enumerates the values for unit.
type Unit string

const (
	// Cm ...
	Cm Unit = "cm"
	// In ...
	In Unit = "in"
	// Mm ...
	Mm Unit = "mm"
)

// PossibleUnitValues returns an array of possible values for the Unit const type.
func PossibleUnitValues() []Unit {
	return []Unit{Cm, In, Mm}
}
